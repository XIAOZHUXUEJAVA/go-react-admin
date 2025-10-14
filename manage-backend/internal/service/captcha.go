package service

import (
	"context"
	"fmt"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
)

// CaptchaConfig 验证码配置
type CaptchaConfig struct {
	Type            string        `json:"type"`
	Length          int           `json:"length"`
	Width           int           `json:"width"`
	Height          int           `json:"height"`
	NoiseCount      float64       `json:"noise_count"`
	ShowLineOptions int           `json:"show_line_options"`
	Expiration      time.Duration `json:"expiration"`
	Enabled         bool          `json:"enabled"`
}

// CaptchaService 验证码服务
type CaptchaService struct {
	store  base64Captcha.Store
	driver base64Captcha.Driver
	rdb    *redis.Client
}

// CaptchaResponse 验证码响应
type CaptchaResponse struct {
	CaptchaID   string `json:"captcha_id"`
	CaptchaData string `json:"captcha_data"`
}

// NewCaptchaService 创建验证码服务实例
func NewCaptchaService(rdb *redis.Client, config CaptchaConfig) *CaptchaService {
	var driver base64Captcha.Driver
	
	// 根据配置创建不同类型的验证码驱动
	switch config.Type {
	case "string":
		driver = base64Captcha.NewDriverString(
			config.Height, config.Width, int(config.NoiseCount*100), 
			config.ShowLineOptions, config.Length, 
			"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 
			nil, nil, nil,
		)
	case "math":
		driver = base64Captcha.NewDriverMath(
			config.Height, config.Width, int(config.NoiseCount*100), 
			config.ShowLineOptions, nil, nil, nil,
		)
	case "chinese":
		driver = base64Captcha.NewDriverChinese(
			config.Height, config.Width, int(config.NoiseCount*100), 
			config.ShowLineOptions, config.Length, 
			"设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,假设,我们,需要,使用,一个,不同,的音,频输,出路径,来处,理每个,不同,的声道,输出,音频,输出的,最后,一个,概念,是音,频输,出设,备的,概念,在这,里每,个音,频输,出设,备代,表着,一个,硬件,设备,例如,扬声器,或者,头戴式,耳机,音频,输出,设备,通常,是物,理设,备和,驱动,程序,的抽,象化,表示,驱动,程序,控制,着硬,件设,备并,向计,算机,提供,一致,的软,件接,口", 
			nil, nil, nil,
		)
	default: // digit
		driver = base64Captcha.NewDriverDigit(
			config.Height, config.Width, config.Length, 
			config.NoiseCount, config.ShowLineOptions,
		)
	}
	
	// 使用 Redis 作为存储
	store := NewRedisCaptchaStore(rdb, config.Expiration)
	
	return &CaptchaService{
		store:  store,
		driver: driver,
		rdb:    rdb,
	}
}

// GenerateCaptcha 生成验证码
func (s *CaptchaService) GenerateCaptcha() (*CaptchaResponse, error) {
	captcha := base64Captcha.NewCaptcha(s.driver, s.store)
	
	id, b64s, answer, err := captcha.Generate()
	if err != nil {
		return nil, fmt.Errorf("failed to generate captcha: %w", err)
	}
	
	// 记录生成的验证码答案（用于调试，生产环境应移除）
	// log.Printf("Generated captcha ID: %s, Answer: %s", id, answer)
	_ = answer // 避免未使用变量警告
	
	return &CaptchaResponse{
		CaptchaID:   id,
		CaptchaData: b64s,
	}, nil
}

// VerifyCaptcha 验证验证码
func (s *CaptchaService) VerifyCaptcha(captchaID, captchaValue string) bool {
	if captchaID == "" || captchaValue == "" {
		return false
	}
	
	return s.store.Verify(captchaID, captchaValue, true) // true 表示验证后清除
}

// RedisCaptchaStore Redis 验证码存储实现
type RedisCaptchaStore struct {
	rdb        *redis.Client
	expiration time.Duration
}

// NewRedisCaptchaStore 创建 Redis 验证码存储
func NewRedisCaptchaStore(rdb *redis.Client, expiration time.Duration) *RedisCaptchaStore {
	return &RedisCaptchaStore{
		rdb:        rdb,
		expiration: expiration,
	}
}

// Set 存储验证码
func (r *RedisCaptchaStore) Set(id string, value string) error {
	key := r.getCaptchaKey(id)
	return r.rdb.Set(context.Background(), key, value, r.expiration).Err()
}

// Get 获取验证码
func (r *RedisCaptchaStore) Get(id string, clear bool) string {
	key := r.getCaptchaKey(id)
	ctx := context.Background()
	
	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	
	if clear {
		r.rdb.Del(ctx, key)
	}
	
	return val
}

// Verify 验证验证码
func (r *RedisCaptchaStore) Verify(id, answer string, clear bool) bool {
	storedAnswer := r.Get(id, clear)
	return storedAnswer != "" && storedAnswer == answer
}

// getCaptchaKey 获取验证码在 Redis 中的键
func (r *RedisCaptchaStore) getCaptchaKey(id string) string {
	return fmt.Sprintf("captcha:%s", id)
}

// CaptchaServiceInterface 验证码服务接口
type CaptchaServiceInterface interface {
	GenerateCaptcha() (*CaptchaResponse, error)
	VerifyCaptcha(captchaID, captchaValue string) bool
}

// 确保 CaptchaService 实现了接口
var _ CaptchaServiceInterface = (*CaptchaService)(nil)