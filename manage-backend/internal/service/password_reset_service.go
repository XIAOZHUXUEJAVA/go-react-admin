package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/config"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/repository"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
	apperrors "github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/errors"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// PasswordResetService 密码重置服务接口
type PasswordResetService interface {
	RequestPasswordReset(ctx context.Context, email, ipAddress, userAgent string) error
	VerifyResetToken(ctx context.Context, token string) (*model.User, error)
	ResetPassword(ctx context.Context, token, newPassword, ipAddress, userAgent string) error
}

type passwordResetService struct {
	config          *config.Config
	userRepo        *repository.UserRepository
	resetTokenRepo  *repository.PasswordResetRepository
	emailService    EmailService
	auditLogService *AuditLogService
	redisService    *PasswordResetRedisService
}

// NewPasswordResetService 创建密码重置服务实例
func NewPasswordResetService(
	cfg *config.Config,
	userRepo *repository.UserRepository,
	resetTokenRepo *repository.PasswordResetRepository,
	emailService EmailService,
	auditLogService *AuditLogService,
	redisService *PasswordResetRedisService,
) PasswordResetService {
	logger.Info("密码重置服务初始化成功（已启用Redis加速）")
	return &passwordResetService{
		config:          cfg,
		userRepo:        userRepo,
		resetTokenRepo:  resetTokenRepo,
		emailService:    emailService,
		auditLogService: auditLogService,
		redisService:    redisService,
	}
}

// RequestPasswordReset 请求密码重置
func (s *passwordResetService) RequestPasswordReset(ctx context.Context, email, ipAddress, userAgent string) error {
	logger.Info("开始处理密码重置请求",
		zap.String("email", email),
		zap.String("ip_address", ipAddress),
		zap.String("operation", "request_password_reset"))

	// 1. Redis限流检查 - IP级别
	if s.redisService != nil {
		allowed, remaining, err := s.redisService.CheckIPRateLimit(ctx, ipAddress)
		if err != nil {
			logger.Warn("IP限流检查失败，继续处理", zap.Error(err))
		} else if !allowed {
			logger.Warn("IP请求频率超限，拒绝请求",
				zap.String("ip", ipAddress),
				zap.Int("remaining", remaining))
			return apperrors.NewRateLimitErrorWithCode(fmt.Sprintf("请求过于频繁，请1小时后再试（剩余次数：%d）", remaining))
		} else {
			logger.Debug("IP限流检查通过",
				zap.String("ip", ipAddress),
				zap.Int("remaining", remaining))
		}
	}

	// 2. Redis限流检查 - 邮箱级别
	if s.redisService != nil {
		allowed, remaining, err := s.redisService.CheckEmailRateLimit(ctx, email)
		if err != nil {
			logger.Warn("邮箱限流检查失败，继续处理", zap.Error(err))
		} else if !allowed {
			logger.Warn("邮箱请求频率超限，拒绝请求",
				zap.String("email", email),
				zap.Int("remaining", remaining))
			return apperrors.NewRateLimitErrorWithCode(fmt.Sprintf("该邮箱请求过于频繁，请1小时后再试（剩余次数：%d）", remaining))
		} else {
			logger.Debug("邮箱限流检查通过",
				zap.String("email", email),
				zap.Int("remaining", remaining))
		}
	}

	// 3. 查找用户
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 安全实践：即使用户不存在也返回成功，防止邮箱枚举攻击
			logger.Warn("密码重置请求：用户不存在（返回成功以防止枚举）",
				zap.String("email", email),
				zap.String("ip_address", ipAddress),
				zap.String("operation", "request_password_reset"))
			return nil
		}
		logger.Error("查询用户失败",
			zap.String("email", email),
			zap.Error(err),
			zap.String("operation", "request_password_reset"))
		return apperrors.NewPasswordResetUserQueryFailedError()
	}

	// 2. 检查用户状态
	if user.Status != "active" {
		logger.Warn("密码重置请求失败：账户已被禁用",
			zap.Uint("user_id", user.ID),
			zap.String("status", user.Status),
			zap.String("operation", "request_password_reset"))
		return apperrors.NewAccountDisabledError("账户已被禁用，无法重置密码")
	}

	// 3. 删除该用户之前未使用的Token（防止重复请求）
	if err := s.resetTokenRepo.DeleteByUserID(ctx, user.ID); err != nil {
		logger.Error("清理旧Token失败",
			zap.Uint("user_id", user.ID),
			zap.Error(err),
			zap.String("operation", "request_password_reset"))
		return apperrors.NewPasswordResetTokenCleanFailedError()
	}

	logger.Debug("已清理用户旧的重置Token",
		zap.Uint("user_id", user.ID),
		zap.String("email", email))

	// 4. 生成新Token
	token := uuid.New().String()
	expiresAt := time.Now().Add(time.Duration(s.config.PasswordReset.TokenExpireMinutes) * time.Minute)

	resetToken := &model.PasswordResetToken{
		UserID:    user.ID,
		Email:     email,
		Token:     token,
		ExpiresAt: expiresAt,
		IPAddress: ipAddress,
		UserAgent: userAgent,
	}

	if err := s.resetTokenRepo.Create(ctx, resetToken); err != nil {
		logger.Error("创建重置Token失败",
			zap.Uint("user_id", user.ID),
			zap.String("email", email),
			zap.Error(err),
			zap.String("operation", "request_password_reset"))
		return apperrors.NewPasswordResetTokenCreateFailedError()
	}

	logger.Debug("重置Token创建成功",
		zap.Uint("user_id", user.ID),
		zap.String("email", email),
		zap.Time("expires_at", expiresAt))

	// 4.5. 同时保存Token到Redis（加速验证）
	if s.redisService != nil {
		expiration := time.Duration(s.config.PasswordReset.TokenExpireMinutes) * time.Minute
		if err := s.redisService.SaveToken(ctx, token, user.ID, email, expiration); err != nil {
			logger.Warn("保存Token到Redis失败，但不影响主流程", zap.Error(err))
		} else {
			logger.Debug("Token已同步到Redis")
		}
	}

	// 5. 发送邮件
	if err := s.emailService.SendPasswordResetEmail(email, token, user.Username); err != nil {
		logger.Error("发送重置邮件失败",
			zap.Uint("user_id", user.ID),
			zap.String("email", email),
			zap.Error(err),
			zap.String("operation", "request_password_reset"))
		return apperrors.NewPasswordResetEmailSendFailedError()
	}

	// 6. 记录审计日志
	s.auditLogService.auditLogRepo.Create(&model.AuditLog{
		UserID:     user.ID,
		Username:   user.Username,
		Action:     "请求密码重置",
		Resource:   "password_reset",
		ResourceID: fmt.Sprintf("%d", resetToken.ID),
		Method:     "POST",
		IP:         ipAddress,
		UserAgent:  userAgent,
		Status:     200,
	})

	logger.Info("密码重置请求处理成功",
		zap.String("email", email),
		zap.Uint("user_id", user.ID),
		zap.String("username", user.Username),
		zap.String("ip_address", ipAddress),
		zap.String("operation", "request_password_reset"))

	return nil
}

// VerifyResetToken 验证重置Token（混合方案：优先Redis，降级PostgreSQL）
func (s *passwordResetService) VerifyResetToken(ctx context.Context, token string) (*model.User, error) {
	logger.Debug("开始验证重置Token",
		zap.String("operation", "verify_reset_token"))

	var userID uint
	var email string

	// 1. 优先从Redis获取Token（快速路径）
	if s.redisService != nil {
		redisToken, err := s.redisService.GetToken(ctx, token)
		if err != nil {
			logger.Warn("从Redis获取Token失败，降级到数据库", zap.Error(err))
		} else if redisToken != nil {
			// Redis命中
			logger.Debug("Token在Redis中命中（快速验证）",
				zap.Uint("user_id", redisToken.UserID),
				zap.String("email", redisToken.Email))
			userID = redisToken.UserID
			email = redisToken.Email
		}
	}

	// 2. Redis未命中，从PostgreSQL查找（降级路径）
	if userID == 0 {
		logger.Debug("Redis未命中，从数据库查询Token")
		resetToken, err := s.resetTokenRepo.FindByToken(ctx, token)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				logger.Warn("验证失败：Token不存在",
					zap.String("operation", "verify_reset_token"))
				return nil, apperrors.NewInvalidResetTokenError()
			}
			logger.Error("查询Token失败",
				zap.Error(err),
				zap.String("operation", "verify_reset_token"))
			return nil, apperrors.NewPasswordResetTokenQueryFailedError()
		}

		// 验证Token有效性
		if !resetToken.IsValid() {
			if resetToken.IsExpired() {
				logger.Warn("验证失败：Token已过期",
					zap.Uint("token_id", resetToken.ID),
					zap.Time("expires_at", resetToken.ExpiresAt),
					zap.String("operation", "verify_reset_token"))
				return nil, apperrors.NewResetTokenExpiredError()
			}
			if resetToken.IsUsed() {
				logger.Warn("验证失败：Token已使用",
					zap.Uint("token_id", resetToken.ID),
					zap.Time("used_at", *resetToken.UsedAt),
					zap.String("operation", "verify_reset_token"))
				return nil, apperrors.NewInvalidResetTokenError()
			}
		}

		userID = resetToken.UserID
		email = resetToken.Email

		// 回写Redis（缓存预热）
		if s.redisService != nil {
			ttl := time.Until(resetToken.ExpiresAt)
			if ttl > 0 {
				if err := s.redisService.SaveToken(ctx, token, userID, email, ttl); err != nil {
					logger.Warn("回写Token到Redis失败", zap.Error(err))
				} else {
					logger.Debug("Token已回写到Redis")
				}
			}
		}

		logger.Debug("Token验证通过（数据库）",
			zap.Uint("token_id", resetToken.ID),
			zap.Uint("user_id", userID))
	}

	// 3. 查找用户
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		logger.Error("查询用户失败",
			zap.Uint("user_id", userID),
			zap.Error(err),
			zap.String("operation", "verify_reset_token"))
		return nil, apperrors.NewPasswordResetUserQueryFailedError()
	}

	logger.Debug("Token验证成功",
		zap.Uint("user_id", user.ID),
		zap.String("username", user.Username),
		zap.String("email", user.Email),
		zap.String("operation", "verify_reset_token"))

	return user, nil
}

// ResetPassword 重置密码
func (s *passwordResetService) ResetPassword(ctx context.Context, token, newPassword, ipAddress, userAgent string) error {
	logger.Info("开始重置密码",
		zap.String("ip_address", ipAddress),
		zap.String("operation", "reset_password"))

	// 1. 验证Token并获取用户
	user, err := s.VerifyResetToken(ctx, token)
	if err != nil {
		logger.Warn("密码重置失败：Token验证失败",
			zap.Error(err),
			zap.String("ip_address", ipAddress),
			zap.String("operation", "reset_password"))
		return err
	}

	// 2. 加密新密码
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		logger.Error("密码加密失败",
			zap.Uint("user_id", user.ID),
			zap.Error(err),
			zap.String("operation", "reset_password"))
		return apperrors.NewPasswordResetHashFailedError()
	}

	logger.Debug("新密码加密成功",
		zap.Uint("user_id", user.ID))

	// 3. 更新密码
	if err := s.userRepo.UpdatePassword(user.ID, hashedPassword); err != nil {
		logger.Error("更新密码失败",
			zap.Uint("user_id", user.ID),
			zap.String("username", user.Username),
			zap.Error(err),
			zap.String("operation", "reset_password"))
		return apperrors.NewPasswordResetUpdateFailedError()
	}

	logger.Debug("密码更新成功",
		zap.Uint("user_id", user.ID),
		zap.String("username", user.Username))

	// 3.5. 删除Redis中的Token（立即失效）
	if s.redisService != nil {
		if err := s.redisService.DeleteToken(ctx, token); err != nil {
			logger.Warn("删除Redis Token失败", zap.Error(err))
		} else {
			logger.Debug("已删除Redis中的Token")
		}
	}

	// 4. 标记Token为已使用
	resetToken, _ := s.resetTokenRepo.FindByToken(ctx, token)
	if resetToken != nil {
		if err := s.resetTokenRepo.MarkAsUsed(ctx, resetToken.ID); err != nil {
			logger.Error("标记Token为已使用失败",
				zap.Uint("token_id", resetToken.ID),
				zap.Error(err),
				zap.String("operation", "reset_password"))
			// 不返回错误，因为密码已经更新成功
		} else {
			logger.Debug("Token已标记为已使用",
				zap.Uint("token_id", resetToken.ID))
		}
	}

	// 5. 记录审计日志
	s.auditLogService.auditLogRepo.Create(&model.AuditLog{
		UserID:     user.ID,
		Username:   user.Username,
		Action:     "密码重置成功",
		Resource:   "password_reset",
		ResourceID: fmt.Sprintf("%d", user.ID),
		Method:     "POST",
		IP:         ipAddress,
		UserAgent:  userAgent,
		Status:     200,
	})

	logger.Info("密码重置成功",
		zap.Uint("user_id", user.ID),
		zap.String("username", user.Username),
		zap.String("email", user.Email),
		zap.String("ip_address", ipAddress),
		zap.String("operation", "reset_password"))

	return nil
}
