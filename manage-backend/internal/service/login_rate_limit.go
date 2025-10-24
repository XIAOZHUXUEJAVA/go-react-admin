package service

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/cache"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
)

// Redis Key 常量
const (
	// IP限流 - login:limit:ip:{ip}
	RedisKeyLoginLimitIP = "login:limit:ip:%s"
	
	// 账户失败次数 - login:fail:account:{username}
	RedisKeyLoginFailAccount = "login:fail:account:%s"
	
	// 账户锁定 - login:locked:account:{username}
	RedisKeyLoginLockedAccount = "login:locked:account:%s"
)

// 限流配置
const (
	// IP级别：每小时最多10次登录尝试
	MaxLoginAttemptsPerIP = 10
	IPLimitWindow         = 1 * time.Hour
	
	// 账户级别：连续5次失败后锁定
	MaxLoginFailsPerAccount = 5
	AccountLockDuration     = 15 * time.Minute
	FailCountWindow         = 15 * time.Minute
)

// LoginRateLimitService 登录限流服务
type LoginRateLimitService struct {
	redisClient *cache.RedisClient
}

// NewLoginRateLimitService 创建登录限流服务
func NewLoginRateLimitService(redisClient *cache.RedisClient) *LoginRateLimitService {
	return &LoginRateLimitService{
		redisClient: redisClient,
	}
}

// CheckIPRateLimit 检查IP限流
// 返回：是否允许，剩余次数，错误
func (s *LoginRateLimitService) CheckIPRateLimit(ctx context.Context, ip string) (bool, int, error) {
	key := fmt.Sprintf(RedisKeyLoginLimitIP, ip)
	
	// 递增计数
	count, err := s.redisClient.Incr(ctx, key)
	if err != nil {
		logger.Error("IP登录限流计数失败", zap.String("ip", ip), zap.Error(err))
		return false, 0, err
	}
	
	// 第一次请求，设置过期时间
	if count == 1 {
		if err := s.redisClient.Expire(ctx, key, IPLimitWindow); err != nil {
			logger.Error("设置IP登录限流过期时间失败", zap.String("ip", ip), zap.Error(err))
		}
	}
	
	remaining := int(MaxLoginAttemptsPerIP - count)
	if remaining < 0 {
		remaining = 0
	}
	
	allowed := count <= MaxLoginAttemptsPerIP
	
	if !allowed {
		logger.Warn("IP登录请求频率超限",
			zap.String("ip", ip),
			zap.Int64("count", count),
			zap.Int("max", MaxLoginAttemptsPerIP))
	}
	
	return allowed, remaining, nil
}

// CheckAccountLocked 检查账户是否被锁定
// 返回：是否锁定，剩余锁定时间，错误
func (s *LoginRateLimitService) CheckAccountLocked(ctx context.Context, username string) (bool, time.Duration, error) {
	key := fmt.Sprintf(RedisKeyLoginLockedAccount, username)
	
	exists, err := s.redisClient.Exists(ctx, key)
	if err != nil {
		logger.Error("检查账户锁定状态失败", zap.String("username", username), zap.Error(err))
		return false, 0, err
	}
	
	if exists > 0 {
		// 账户已锁定，获取剩余时间
		ttl, err := s.redisClient.TTL(ctx, key)
		if err != nil {
			logger.Error("获取账户锁定TTL失败", zap.String("username", username), zap.Error(err))
			return true, 0, err
		}
		
		logger.Warn("账户处于锁定状态",
			zap.String("username", username),
			zap.Duration("remaining", ttl))
		
		return true, ttl, nil
	}
	
	return false, 0, nil
}

// RecordLoginFailure 记录登录失败
// 返回：当前失败次数，是否应该锁定账户，错误
func (s *LoginRateLimitService) RecordLoginFailure(ctx context.Context, username string) (int, bool, error) {
	key := fmt.Sprintf(RedisKeyLoginFailAccount, username)
	
	// 递增失败计数
	count, err := s.redisClient.Incr(ctx, key)
	if err != nil {
		logger.Error("记录登录失败次数失败", zap.String("username", username), zap.Error(err))
		return 0, false, err
	}
	
	// 第一次失败，设置过期时间
	if count == 1 {
		if err := s.redisClient.Expire(ctx, key, FailCountWindow); err != nil {
			logger.Error("设置失败计数过期时间失败", zap.String("username", username), zap.Error(err))
		}
	}
	
	logger.Info("记录登录失败",
		zap.String("username", username),
		zap.Int64("fail_count", count),
		zap.Int("max", MaxLoginFailsPerAccount))
	
	// 检查是否达到锁定阈值
	shouldLock := count >= int64(MaxLoginFailsPerAccount)
	
	if shouldLock {
		// 锁定账户
		if err := s.LockAccount(ctx, username); err != nil {
			logger.Error("锁定账户失败", zap.String("username", username), zap.Error(err))
			return int(count), true, err
		}
		
		logger.Warn("账户因连续失败被锁定",
			zap.String("username", username),
			zap.Int64("fail_count", count),
			zap.Duration("lock_duration", AccountLockDuration))
	}
	
	return int(count), shouldLock, nil
}

// LockAccount 锁定账户
func (s *LoginRateLimitService) LockAccount(ctx context.Context, username string) error {
	key := fmt.Sprintf(RedisKeyLoginLockedAccount, username)
	
	lockReason := fmt.Sprintf("连续登录失败%d次", MaxLoginFailsPerAccount)
	
	if err := s.redisClient.Set(ctx, key, lockReason, AccountLockDuration); err != nil {
		logger.Error("锁定账户失败", zap.String("username", username), zap.Error(err))
		return fmt.Errorf("锁定账户失败: %w", err)
	}
	
	logger.Info("账户已锁定",
		zap.String("username", username),
		zap.String("reason", lockReason),
		zap.Duration("duration", AccountLockDuration))
	
	return nil
}

// ClearLoginFailures 清除登录失败记录（登录成功时调用）
func (s *LoginRateLimitService) ClearLoginFailures(ctx context.Context, username string) error {
	key := fmt.Sprintf(RedisKeyLoginFailAccount, username)
	
	if err := s.redisClient.Del(ctx, key); err != nil {
		logger.Error("清除登录失败记录失败", zap.String("username", username), zap.Error(err))
		return err
	}
	
	logger.Debug("已清除登录失败记录", zap.String("username", username))
	return nil
}

// GetFailureCount 获取当前失败次数
func (s *LoginRateLimitService) GetFailureCount(ctx context.Context, username string) (int, error) {
	key := fmt.Sprintf(RedisKeyLoginFailAccount, username)
	
	countStr, err := s.redisClient.Get(ctx, key)
	if err != nil {
		if err == redis.Nil {
			return 0, nil // 没有失败记录
		}
		logger.Error("获取失败次数失败", zap.String("username", username), zap.Error(err))
		return 0, err
	}
	
	var count int
	fmt.Sscanf(countStr, "%d", &count)
	return count, nil
}

// GetRemainingAttempts 获取剩余尝试次数
func (s *LoginRateLimitService) GetRemainingAttempts(ctx context.Context, username string) int {
	count, err := s.GetFailureCount(ctx, username)
	if err != nil {
		return MaxLoginFailsPerAccount // 出错时返回最大值
	}
	
	remaining := MaxLoginFailsPerAccount - count
	if remaining < 0 {
		remaining = 0
	}
	
	return remaining
}
