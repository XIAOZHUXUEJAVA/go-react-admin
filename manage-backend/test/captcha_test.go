package test

import (
	"context"
	"testing"
	"time"

	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/service"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCaptchaService(t *testing.T) {
	// 创建 Redis 客户端（测试环境）
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1, // 使用测试数据库
	})

	// 测试 Redis 连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		t.Skip("Redis not available, skipping captcha tests")
	}

	// 清理测试数据
	defer func() {
		keys, _ := rdb.Keys(ctx, "captcha:*").Result()
		if len(keys) > 0 {
			rdb.Del(ctx, keys...)
		}
		rdb.Close()
	}()

	// 创建验证码配置
	config := service.CaptchaConfig{
		Type:            "digit",
		Length:          5,
		Width:           240,
		Height:          80,
		NoiseCount:      0.7,
		ShowLineOptions: 80,
		Expiration:      1 * time.Minute,
		Enabled:         true,
	}

	// 创建验证码服务
	captchaService := service.NewCaptchaService(rdb, config)

	t.Run("Generate Captcha", func(t *testing.T) {
		// 生成验证码
		captcha, err := captchaService.GenerateCaptcha()
		require.NoError(t, err)
		require.NotNil(t, captcha)

		// 验证返回数据
		assert.NotEmpty(t, captcha.CaptchaID)
		assert.NotEmpty(t, captcha.CaptchaData)
		assert.Contains(t, captcha.CaptchaData, "data:image/png;base64,")

		t.Logf("Generated captcha ID: %s", captcha.CaptchaID)
		t.Logf("Captcha data length: %d", len(captcha.CaptchaData))
	})

	t.Run("Verify Captcha - Invalid Cases", func(t *testing.T) {
		// 测试空值
		result := captchaService.VerifyCaptcha("", "")
		assert.False(t, result)

		// 测试不存在的ID
		result = captchaService.VerifyCaptcha("nonexistent", "12345")
		assert.False(t, result)

		// 测试错误的验证码
		captcha, err := captchaService.GenerateCaptcha()
		require.NoError(t, err)
		
		result = captchaService.VerifyCaptcha(captcha.CaptchaID, "wrong")
		assert.False(t, result)
	})

	t.Run("Captcha Expiration", func(t *testing.T) {
		// 创建短过期时间的配置
		shortConfig := config
		shortConfig.Expiration = 100 * time.Millisecond
		
		shortCaptchaService := service.NewCaptchaService(rdb, shortConfig)
		
		// 生成验证码
		captcha, err := shortCaptchaService.GenerateCaptcha()
		require.NoError(t, err)

		// 等待过期
		time.Sleep(200 * time.Millisecond)

		// 验证应该失败
		result := shortCaptchaService.VerifyCaptcha(captcha.CaptchaID, "12345")
		assert.False(t, result)
	})

	t.Run("Different Captcha Types", func(t *testing.T) {
		types := []string{"digit", "string", "math", "chinese"}
		
		for _, captchaType := range types {
			t.Run(captchaType, func(t *testing.T) {
				typeConfig := config
				typeConfig.Type = captchaType
				
				typeCaptchaService := service.NewCaptchaService(rdb, typeConfig)
				
				captcha, err := typeCaptchaService.GenerateCaptcha()
				require.NoError(t, err)
				require.NotNil(t, captcha)
				
				assert.NotEmpty(t, captcha.CaptchaID)
				assert.NotEmpty(t, captcha.CaptchaData)
				assert.Contains(t, captcha.CaptchaData, "data:image/png;base64,")
				
				t.Logf("Generated %s captcha ID: %s", captchaType, captcha.CaptchaID)
			})
		}
	})
}

func TestRedisCaptchaStore(t *testing.T) {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})

	// 测试连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		t.Skip("Redis not available, skipping store tests")
	}

	// 清理测试数据
	defer func() {
		keys, _ := rdb.Keys(ctx, "captcha:*").Result()
		if len(keys) > 0 {
			rdb.Del(ctx, keys...)
		}
		rdb.Close()
	}()

	store := service.NewRedisCaptchaStore(rdb, 1*time.Minute)

	t.Run("Set and Get", func(t *testing.T) {
		id := "test123"
		value := "12345"

		// 存储
		err := store.Set(id, value)
		require.NoError(t, err)

		// 获取（不清除）
		result := store.Get(id, false)
		assert.Equal(t, value, result)

		// 再次获取应该仍然存在
		result = store.Get(id, false)
		assert.Equal(t, value, result)

		// 获取并清除
		result = store.Get(id, true)
		assert.Equal(t, value, result)

		// 再次获取应该为空
		result = store.Get(id, false)
		assert.Empty(t, result)
	})

	t.Run("Verify", func(t *testing.T) {
		id := "verify123"
		value := "54321"

		// 存储
		err := store.Set(id, value)
		require.NoError(t, err)

		// 验证正确值
		result := store.Verify(id, value, false)
		assert.True(t, result)

		// 验证错误值
		result = store.Verify(id, "wrong", false)
		assert.False(t, result)

		// 验证并清除
		result = store.Verify(id, value, true)
		assert.True(t, result)

		// 再次验证应该失败
		result = store.Verify(id, value, false)
		assert.False(t, result)
	})

	t.Run("Nonexistent Key", func(t *testing.T) {
		// 获取不存在的键
		result := store.Get("nonexistent", false)
		assert.Empty(t, result)

		// 验证不存在的键
		result2 := store.Verify("nonexistent", "value", false)
		assert.False(t, result2)
	})
}