package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/cache"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
)

// Redis Key 常量
const (
	// Token验证 - password_reset:token:{token}
	RedisKeyResetToken = "password_reset:token:%s"
	
	// IP限流 - password_reset:limit:ip:{ip}
	RedisKeyLimitIP = "password_reset:limit:ip:%s"
	
	// 邮箱限流 - password_reset:limit:email:{email}
	RedisKeyLimitEmail = "password_reset:limit:email:%s"
)

// 限流配置
const (
	// IP级别：每小时最多5次请求
	MaxRequestsPerIP = 5
	
	// 邮箱级别：每小时最多3次请求
	MaxRequestsPerEmail = 3
	
	// 限流时间窗口
	RateLimitWindow = 1 * time.Hour
)

// RedisResetToken Redis中存储的Token数据
type RedisResetToken struct {
	UserID    uint      `json:"user_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// PasswordResetRedisService Redis密码重置服务
type PasswordResetRedisService struct {
	redisClient *cache.RedisClient
}

// NewPasswordResetRedisService 创建Redis密码重置服务
func NewPasswordResetRedisService(redisClient *cache.RedisClient) *PasswordResetRedisService {
	return &PasswordResetRedisService{
		redisClient: redisClient,
	}
}

// CheckIPRateLimit 检查IP限流
// 返回：是否允许请求，剩余次数，错误
func (s *PasswordResetRedisService) CheckIPRateLimit(ctx context.Context, ip string) (bool, int, error) {
	key := fmt.Sprintf(RedisKeyLimitIP, ip)
	
	// 递增计数
	count, err := s.redisClient.Incr(ctx, key)
	if err != nil {
		logger.Error("IP限流计数失败", zap.String("ip", ip), zap.Error(err))
		return false, 0, err
	}
	
	// 第一次请求，设置过期时间
	if count == 1 {
		if err := s.redisClient.Expire(ctx, key, RateLimitWindow); err != nil {
			logger.Error("设置IP限流过期时间失败", zap.String("ip", ip), zap.Error(err))
		}
	}
	
	remaining := int(MaxRequestsPerIP - count)
	if remaining < 0 {
		remaining = 0
	}
	
	allowed := count <= MaxRequestsPerIP
	
	if !allowed {
		logger.Warn("IP请求频率超限",
			zap.String("ip", ip),
			zap.Int64("count", count),
			zap.Int("max", MaxRequestsPerIP))
	}
	
	return allowed, remaining, nil
}

// CheckEmailRateLimit 检查邮箱限流
// 返回：是否允许请求，剩余次数，错误
func (s *PasswordResetRedisService) CheckEmailRateLimit(ctx context.Context, email string) (bool, int, error) {
	key := fmt.Sprintf(RedisKeyLimitEmail, email)
	
	// 递增计数
	count, err := s.redisClient.Incr(ctx, key)
	if err != nil {
		logger.Error("邮箱限流计数失败", zap.String("email", email), zap.Error(err))
		return false, 0, err
	}
	
	// 第一次请求，设置过期时间
	if count == 1 {
		if err := s.redisClient.Expire(ctx, key, RateLimitWindow); err != nil {
			logger.Error("设置邮箱限流过期时间失败", zap.String("email", email), zap.Error(err))
		}
	}
	
	remaining := int(MaxRequestsPerEmail - count)
	if remaining < 0 {
		remaining = 0
	}
	
	allowed := count <= MaxRequestsPerEmail
	
	if !allowed {
		logger.Warn("邮箱请求频率超限",
			zap.String("email", email),
			zap.Int64("count", count),
			zap.Int("max", MaxRequestsPerEmail))
	}
	
	return allowed, remaining, nil
}

// SaveToken 保存Token到Redis
func (s *PasswordResetRedisService) SaveToken(ctx context.Context, token string, userID uint, email string, expiration time.Duration) error {
	key := fmt.Sprintf(RedisKeyResetToken, token)
	
	data := RedisResetToken{
		UserID:    userID,
		Email:     email,
		CreatedAt: time.Now(),
	}
	
	// 序列化为JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Error("序列化Token数据失败", zap.Error(err))
		return fmt.Errorf("序列化Token数据失败: %w", err)
	}
	
	// 保存到Redis，设置过期时间
	if err := s.redisClient.Set(ctx, key, jsonData, expiration); err != nil {
		logger.Error("保存Token到Redis失败",
			zap.String("token", token),
			zap.Uint("user_id", userID),
			zap.Error(err))
		return fmt.Errorf("保存Token到Redis失败: %w", err)
	}
	
	logger.Debug("Token已保存到Redis",
		zap.Uint("user_id", userID),
		zap.String("email", email),
		zap.Duration("expiration", expiration))
	
	return nil
}

// GetToken 从Redis获取Token
func (s *PasswordResetRedisService) GetToken(ctx context.Context, token string) (*RedisResetToken, error) {
	key := fmt.Sprintf(RedisKeyResetToken, token)
	
	// 从Redis获取
	jsonData, err := s.redisClient.Get(ctx, key)
	if err != nil {
		if err == redis.Nil {
			logger.Debug("Token在Redis中不存在", zap.String("token", token))
			return nil, nil // Token不存在
		}
		logger.Error("从Redis获取Token失败", zap.String("token", token), zap.Error(err))
		return nil, fmt.Errorf("从Redis获取Token失败: %w", err)
	}
	
	// 反序列化
	var data RedisResetToken
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		logger.Error("反序列化Token数据失败", zap.Error(err))
		return nil, fmt.Errorf("反序列化Token数据失败: %w", err)
	}
	
	logger.Debug("从Redis获取Token成功",
		zap.Uint("user_id", data.UserID),
		zap.String("email", data.Email))
	
	return &data, nil
}

// DeleteToken 删除Redis中的Token
func (s *PasswordResetRedisService) DeleteToken(ctx context.Context, token string) error {
	key := fmt.Sprintf(RedisKeyResetToken, token)
	
	if err := s.redisClient.Del(ctx, key); err != nil {
		logger.Error("删除Redis Token失败", zap.String("token", token), zap.Error(err))
		return fmt.Errorf("删除Redis Token失败: %w", err)
	}
	
	logger.Debug("已删除Redis Token")
	return nil
}

// GetTokenTTL 获取Token剩余有效时间
func (s *PasswordResetRedisService) GetTokenTTL(ctx context.Context, token string) (time.Duration, error) {
	key := fmt.Sprintf(RedisKeyResetToken, token)
	
	ttl, err := s.redisClient.TTL(ctx, key)
	if err != nil {
		logger.Error("获取Token TTL失败", zap.String("token", token), zap.Error(err))
		return 0, fmt.Errorf("获取Token TTL失败: %w", err)
	}
	
	return ttl, nil
}
