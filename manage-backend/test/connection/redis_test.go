package connection

import (
	"context"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

// TestRedisConnection 测试Redis连接
func TestRedisConnection(t *testing.T) {
	// 硬编码的Redis配置 - 用于快速连通性测试
	config := struct {
		Host     string
		Port     string
		Password string
		DB       int
	}{
		Host:     "localhost",
		Port:     "6379",
		Password: "", // 开发环境通常无密码
		DB:       0,  // 默认数据库
	}

	t.Logf("🔌 测试Redis连接: %s:%s (DB: %d)", config.Host, config.Port, config.DB)

	// 创建Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})
	defer rdb.Close()

	ctx := context.Background()

	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		t.Fatalf("❌ Redis连接失败: %v", err)
	}

	if pong != "PONG" {
		t.Fatalf("❌ Redis Ping 响应异常: %s", pong)
	}

	// 测试基本操作
	testKey := "test:connection:key"
	testValue := "test_value_" + time.Now().Format("20060102150405")

	// 设置值
	err = rdb.Set(ctx, testKey, testValue, time.Minute).Err()
	if err != nil {
		t.Fatalf("❌ Redis SET 操作失败: %v", err)
	}

	// 获取值
	val, err := rdb.Get(ctx, testKey).Result()
	if err != nil {
		t.Fatalf("❌ Redis GET 操作失败: %v", err)
	}

	if val != testValue {
		t.Fatalf("❌ Redis 值不匹配: 期望 %s, 实际 %s", testValue, val)
	}

	// 删除测试键
	err = rdb.Del(ctx, testKey).Err()
	if err != nil {
		t.Fatalf("❌ Redis DEL 操作失败: %v", err)
	}

	// 获取Redis信息
	info, err := rdb.Info(ctx, "server").Result()
	if err != nil {
		t.Logf("⚠️ 无法获取Redis服务器信息: %v", err)
	}

	t.Logf("✅ Redis连接成功!")
	t.Logf("📊 Ping响应: %s", pong)
	t.Logf("🔧 基本操作测试通过")
	if info != "" {
		t.Logf("ℹ️ Redis服务器信息获取成功")
	}

	// 测试连接池状态
	poolStats := rdb.PoolStats()
	t.Logf("🔗 连接池统计:")
	t.Logf("   - 总连接数: %d", poolStats.TotalConns)
	t.Logf("   - 空闲连接数: %d", poolStats.IdleConns)
	t.Logf("   - 过期连接数: %d", poolStats.StaleConns)
}

// TestRedisConnectionWithPassword 测试带密码的Redis连接
func TestRedisConnectionWithPassword(t *testing.T) {
	t.Logf("🔌 测试带密码的Redis连接...")

	// 使用密码的配置（如果你的Redis设置了密码，修改这里）
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "your_redis_password", // 如果有密码，在这里设置
		DB:       1,                     // 使用不同的数据库
	})
	defer rdb.Close()

	ctx := context.Background()

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		// 如果没有设置密码或密码错误，这是预期的
		t.Logf("⚠️ 带密码的Redis连接失败（可能是预期的）: %v", err)
		return
	}

	t.Logf("✅ 带密码的Redis连接成功!")
}

// TestRedisConnectionWithWrongPort 测试错误端口的情况
func TestRedisConnectionWithWrongPort(t *testing.T) {
	t.Logf("🔌 测试错误的Redis端口...")

	// 故意使用错误的端口
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:9999", // 错误的端口
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err == nil {
		t.Fatalf("❌ 预期连接失败，但连接成功了")
	}

	t.Logf("✅ 错误端口测试通过: %v", err)
}

// TestRedisMultipleOperations 测试Redis多种操作
func TestRedisMultipleOperations(t *testing.T) {
	t.Logf("🔌 测试Redis多种操作...")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       2, // 使用DB 2 避免冲突
	})
	defer rdb.Close()

	ctx := context.Background()

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		t.Fatalf("❌ Redis连接失败: %v", err)
	}

	// 测试字符串操作
	t.Logf("📝 测试字符串操作...")
	err = rdb.Set(ctx, "test:string", "hello", time.Minute).Err()
	if err != nil {
		t.Fatalf("❌ 字符串SET失败: %v", err)
	}

	// 测试哈希操作
	t.Logf("📝 测试哈希操作...")
	err = rdb.HSet(ctx, "test:hash", "field1", "value1", "field2", "value2").Err()
	if err != nil {
		t.Fatalf("❌ 哈希HSET失败: %v", err)
	}

	// 测试列表操作
	t.Logf("📝 测试列表操作...")
	err = rdb.LPush(ctx, "test:list", "item1", "item2", "item3").Err()
	if err != nil {
		t.Fatalf("❌ 列表LPUSH失败: %v", err)
	}

	// 测试集合操作
	t.Logf("📝 测试集合操作...")
	err = rdb.SAdd(ctx, "test:set", "member1", "member2", "member3").Err()
	if err != nil {
		t.Fatalf("❌ 集合SADD失败: %v", err)
	}

	// 清理测试数据
	keys := []string{"test:string", "test:hash", "test:list", "test:set"}
	err = rdb.Del(ctx, keys...).Err()
	if err != nil {
		t.Logf("⚠️ 清理测试数据失败: %v", err)
	}

	t.Logf("✅ Redis多种操作测试通过!")
}