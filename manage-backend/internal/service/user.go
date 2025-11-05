package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/auth"
	apperrors "github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/errors"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// UserRepositoryInterface 定义用户仓库接口
type UserRepositoryInterface interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
	List(offset, limit int) ([]model.User, int64, error)
	ListByVisibility(currentUserID uint, currentUserRole string, offset, limit int) ([]model.User, int64, error)
	CheckUsernameExists(username string) (bool, error)
	CheckEmailExists(email string) (bool, error)
	CheckUsernameExistsExcludeID(username string, excludeID uint) (bool, error)
	CheckEmailExistsExcludeID(email string, excludeID uint) (bool, error)
}

// JWTManagerInterface 定义 JWT 管理器接口
type JWTManagerInterface interface {
	GenerateToken(userID uint, username, role string) (string, error)
	GenerateTokenPair(userID uint, username, role string) (*auth.TokenPair, error)
	ValidateToken(tokenString string) (*auth.Claims, error)
	ValidateRefreshToken(tokenString string) (*auth.Claims, error)
	GetTokenExpiration(claims *auth.Claims) time.Duration
}

// SessionServiceInterface 定义会话服务接口
type SessionServiceInterface interface {
	CreateSession(ctx context.Context, userID uint, username, refreshToken, deviceInfo, ipAddress, userAgent string) error
	GetSession(ctx context.Context, userID uint) (*SessionInfo, error)
	UpdateLastActivity(ctx context.Context, userID uint) error
	DeleteSession(ctx context.Context, userID uint) error
	ValidateRefreshToken(ctx context.Context, refreshToken string) (*SessionInfo, error)
	AddTokenToBlacklist(ctx context.Context, jti string, expiration time.Duration) error
	IsTokenBlacklisted(ctx context.Context, jti string) bool
	SetUserActive(ctx context.Context, userID uint) error
	CacheUserPermissions(ctx context.Context, userID uint, role string, permissions []string) error
}

type UserService struct {
	userRepo             UserRepositoryInterface
	jwtManager           JWTManagerInterface
	sessionService       SessionServiceInterface
	captchaService       CaptchaServiceInterface
	roleRepo             RoleRepositoryInterface
	permissionService    *PermissionService
	loginRateLimitService *LoginRateLimitService
}

func NewUserService(
	userRepo UserRepositoryInterface,
	jwtManager JWTManagerInterface,
	sessionService SessionServiceInterface,
	captchaService CaptchaServiceInterface,
	roleRepo RoleRepositoryInterface,
	permissionService *PermissionService,
	loginRateLimitService *LoginRateLimitService,
) *UserService {
	return &UserService{
		userRepo:             userRepo,
		jwtManager:           jwtManager,
		sessionService:       sessionService,
		captchaService:       captchaService,
		roleRepo:             roleRepo,
		permissionService:    permissionService,
		loginRateLimitService: loginRateLimitService,
	}
}

func (s *UserService) Register(req *model.CreateUserRequest) (*model.User, error) {
	return s.RegisterWithCreator(req, nil)
}

// RegisterWithCreator 创建用户（带创建者ID）
func (s *UserService) RegisterWithCreator(req *model.CreateUserRequest, creatorID *uint) (*model.User, error) {
	logger.Info("开始用户注册流程",
		zap.String("username", req.Username),
		zap.String("email", req.Email),
		zap.String("role", req.Role),
		zap.Any("creator_id", creatorID))

	// 检查用户名是否已存在
	_, err := s.userRepo.GetByUsername(req.Username)
	if err == nil {
		logger.Warn("用户注册失败：用户名已存在", 
			zap.String("username", req.Username),
			zap.String("operation", "register"))
		return nil, apperrors.NewConflictError("用户名已存在")
	}

	// 检查邮箱是否已存在
	_, err = s.userRepo.GetByEmail(req.Email)
	if err == nil {
		logger.Warn("用户注册失败：邮箱已存在", 
			zap.String("username", req.Username),
			zap.String("email", req.Email),
			zap.String("operation", "register"))
		return nil, apperrors.NewConflictError("邮箱已存在")
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		logger.Error("密码加密失败", 
			zap.String("username", req.Username),
			zap.Error(err),
			zap.String("operation", "register"))
		return nil, apperrors.NewInternalError("密码加密失败")
	}

	user := &model.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		Role:      req.Role,
		CreatedBy: creatorID,
	}

	if user.Role == "" {
		user.Role = "user"
		logger.Debug("用户角色为空，设置为默认角色", 
			zap.String("username", req.Username),
			zap.String("default_role", "user"))
	}

	err = s.userRepo.Create(user)
	if err != nil {
		logger.Error("用户创建失败", 
			zap.String("username", req.Username),
			zap.String("email", req.Email),
			zap.Error(err),
			zap.String("operation", "register"))
		return nil, apperrors.NewInternalError("用户创建失败")
	}

	// 同步到 user_roles 表
	if err := s.syncUserRole(user.ID, user.Role); err != nil {
		logger.Error("同步用户角色失败", 
			zap.Uint("user_id", user.ID),
			zap.String("role", user.Role),
			zap.Error(err))
		// 不返回错误，因为用户已创建成功
	}

	logger.Info("用户注册成功", 
		zap.String("username", user.Username),
		zap.Uint("user_id", user.ID),
		zap.String("role", user.Role),
		zap.Any("created_by", creatorID),
		zap.String("operation", "register"))

	return user, nil
}

func (s *UserService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	return s.LoginWithContext(context.Background(), req, "", "", "")
}

// LoginWithContext 带会话上下文信息的登录
func (s *UserService) LoginWithContext(ctx context.Context, req *model.LoginRequest, deviceInfo, ipAddress, userAgent string) (*model.LoginResponse, error) {
	logger.Info("开始用户登录流程", 
		zap.String("username", req.Username),
		zap.String("ip_address", ipAddress),
		zap.String("user_agent", userAgent),
		zap.String("device_info", deviceInfo))

	// 1. IP限流检查
	if s.loginRateLimitService != nil {
		allowed, remaining, err := s.loginRateLimitService.CheckIPRateLimit(ctx, ipAddress)
		if err != nil {
			logger.Warn("IP限流检查失败，继续处理", zap.Error(err))
		} else if !allowed {
			logger.Warn("IP登录请求频率超限",
				zap.String("ip", ipAddress),
				zap.Int("remaining", remaining))
			return nil, apperrors.NewRateLimitError("登录请求过于频繁，请1小时后再试")
		}
	}

	// 2. 检查账户是否被锁定
	if s.loginRateLimitService != nil {
		locked, ttl, err := s.loginRateLimitService.CheckAccountLocked(ctx, req.Username)
		if err != nil {
			logger.Warn("账户锁定检查失败，继续处理", zap.Error(err))
		} else if locked {
			minutes := int(ttl.Minutes())
			if minutes < 1 {
				minutes = 1
			}
			logger.Warn("账户处于锁定状态",
				zap.String("username", req.Username),
				zap.Duration("remaining", ttl))
			return nil, apperrors.NewAccountLockedError(fmt.Sprintf("账户已被锁定，请%d分钟后再试", minutes))
		}
	}

	// 3. 验证验证码
	if s.captchaService != nil {
		if !s.captchaService.VerifyCaptcha(req.CaptchaID, req.CaptchaCode) {
			logger.Warn("登录失败：验证码错误", 
				zap.String("username", req.Username),
				zap.String("captcha_id", req.CaptchaID),
				zap.String("ip_address", ipAddress),
				zap.String("operation", "login"))
			return nil, apperrors.NewInvalidCaptchaError("验证码错误")
		}
		logger.Debug("验证码验证通过", 
			zap.String("username", req.Username),
			zap.String("captcha_id", req.CaptchaID))
	}

	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("登录失败：用户不存在", 
				zap.String("username", req.Username),
				zap.String("ip_address", ipAddress),
				zap.String("operation", "login"))
			return nil, apperrors.NewInvalidCredentialsError("用户名或密码错误")
		}
		logger.Error("登录失败：查询用户时发生错误", 
			zap.String("username", req.Username),
			zap.Error(err),
			zap.String("operation", "login"))
		return nil, apperrors.NewInternalError("查询用户失败")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		logger.Warn("登录失败：密码错误", 
			zap.String("username", req.Username),
			zap.Uint("user_id", user.ID),
			zap.String("ip_address", ipAddress),
			zap.String("operation", "login"))
		
		// 记录登录失败
		if s.loginRateLimitService != nil {
			failCount, shouldLock, err := s.loginRateLimitService.RecordLoginFailure(ctx, req.Username)
			if err != nil {
				logger.Error("记录登录失败次数失败", zap.Error(err))
			} else if shouldLock {
				// 账户已被锁定
				return nil, apperrors.NewAccountLockedError("连续登录失败5次，账户已被锁定15分钟")
			} else {
				// 返回剩余尝试次数
				remaining := MaxLoginFailsPerAccount - failCount
				logger.Info("记录登录失败",
					zap.String("username", req.Username),
					zap.Int("fail_count", failCount),
					zap.Int("remaining", remaining))
				return nil, apperrors.NewInvalidCredentialsError(fmt.Sprintf("用户名或密码错误（剩余%d次机会）", remaining))
			}
		}
		
		return nil, apperrors.NewInvalidCredentialsError("用户名或密码错误")
	}

	logger.Debug("用户认证成功", 
		zap.String("username", user.Username),
		zap.Uint("user_id", user.ID),
		zap.String("role", user.Role))

	// 清除登录失败记录（登录成功）
	if s.loginRateLimitService != nil {
		if err := s.loginRateLimitService.ClearLoginFailures(ctx, req.Username); err != nil {
			logger.Warn("清除登录失败记录失败", zap.Error(err))
		}
	}

	// 生成令牌对
	tokenPair, err := s.jwtManager.GenerateTokenPair(user.ID, user.Username, user.Role)
	if err != nil {
		logger.Error("生成令牌失败", 
			zap.String("username", user.Username),
			zap.Uint("user_id", user.ID),
			zap.Error(err),
			zap.String("operation", "login"))
		return nil, apperrors.NewInternalError("生成令牌失败")
	}

	// 在 Redis 中创建会话
	if s.sessionService != nil {
		err = s.sessionService.CreateSession(ctx, user.ID, user.Username, tokenPair.RefreshToken, deviceInfo, ipAddress, userAgent)
		if err != nil {
			logger.Error("创建会话失败", 
				zap.String("username", user.Username),
				zap.Uint("user_id", user.ID),
				zap.Error(err),
				zap.String("operation", "login"))
			return nil, apperrors.NewInternalError("创建会话失败")
		}

		// 设置用户为活跃状态
		s.sessionService.SetUserActive(ctx, user.ID)

		// 缓存用户权限
		permissions := []string{} // 可根据权限系统扩展
		s.sessionService.CacheUserPermissions(ctx, user.ID, user.Role, permissions)
		
		logger.Debug("会话创建成功", 
			zap.String("username", user.Username),
			zap.Uint("user_id", user.ID))
	}

	// 创建安全的用户响应（不包含密码）
	safeUser := model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	logger.Info("用户登录成功", 
		zap.String("username", user.Username),
		zap.Uint("user_id", user.ID),
		zap.String("role", user.Role),
		zap.String("ip_address", ipAddress),
		zap.String("operation", "login"))

	return &model.LoginResponse{
		AccessToken:      tokenPair.AccessToken,
		RefreshToken:     tokenPair.RefreshToken,
		ExpiresIn:        tokenPair.ExpiresIn,
		RefreshExpiresIn: tokenPair.RefreshExpiresIn,
		TokenType:        "Bearer",
		User:             safeUser,
	}, nil
}

// RefreshToken 使用刷新令牌更新访问令牌
func (s *UserService) RefreshToken(ctx context.Context, req *model.RefreshTokenRequest) (*model.RefreshTokenResponse, error) {
	logger.Debug("开始刷新令牌流程")

	if s.sessionService == nil {
		logger.Error("刷新令牌失败：会话服务不可用", 
			zap.String("operation", "refresh_token"))
		return nil, apperrors.NewInternalError("会话服务不可用")
	}

	// 验证刷新令牌并获取会话
	sessionInfo, err := s.sessionService.ValidateRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		logger.Warn("刷新令牌失败：无效的刷新令牌", 
			zap.Error(err),
			zap.String("operation", "refresh_token"))
		return nil, apperrors.NewUnauthorizedError("无效的刷新令牌")
	}

	logger.Debug("刷新令牌验证成功", 
		zap.String("username", sessionInfo.Username),
		zap.Uint("user_id", sessionInfo.UserID))

	// 生成新的令牌对
	tokenPair, err := s.jwtManager.GenerateTokenPair(sessionInfo.UserID, sessionInfo.Username, "user") // 角色可从会话中获取
	if err != nil {
		logger.Error("生成新令牌失败", 
			zap.String("username", sessionInfo.Username),
			zap.Uint("user_id", sessionInfo.UserID),
			zap.Error(err),
			zap.String("operation", "refresh_token"))
		return nil, apperrors.NewInternalError("生成新令牌失败")
	}

	// 用新的刷新令牌更新会话
	err = s.sessionService.CreateSession(ctx, sessionInfo.UserID, sessionInfo.Username, tokenPair.RefreshToken, sessionInfo.DeviceInfo, sessionInfo.IPAddress, sessionInfo.UserAgent)
	if err != nil {
		logger.Error("更新会话失败", 
			zap.String("username", sessionInfo.Username),
			zap.Uint("user_id", sessionInfo.UserID),
			zap.Error(err),
			zap.String("operation", "refresh_token"))
		return nil, apperrors.NewInternalError("更新会话失败")
	}

	// 更新最后活跃时间
	s.sessionService.UpdateLastActivity(ctx, sessionInfo.UserID)

	logger.Info("令牌刷新成功", 
		zap.String("username", sessionInfo.Username),
		zap.Uint("user_id", sessionInfo.UserID),
		zap.String("operation", "refresh_token"))

	return &model.RefreshTokenResponse{
		AccessToken: tokenPair.AccessToken,
		ExpiresIn:   tokenPair.ExpiresIn,
		TokenType:   "Bearer",
	}, nil
}

// Logout 用户登出
func (s *UserService) Logout(ctx context.Context, userID uint, accessToken string, req *model.LogoutRequest) error {
	logger.Info("开始用户登出流程", 
		zap.Uint("user_id", userID),
		zap.String("operation", "logout"))

	if s.sessionService == nil {
		logger.Error("登出失败：会话服务不可用", 
			zap.Uint("user_id", userID),
			zap.String("operation", "logout"))
		return apperrors.NewInternalError("会话服务不可用")
	}

	// 验证并获取访问令牌声明
	claims, err := s.jwtManager.ValidateToken(accessToken)
	if err != nil {
		logger.Warn("登出失败：无效的访问令牌", 
			zap.Uint("user_id", userID),
			zap.Error(err),
			zap.String("operation", "logout"))
		return apperrors.NewUnauthorizedError("无效的访问令牌")
	}

	logger.Debug("访问令牌验证成功", 
		zap.Uint("user_id", userID),
		zap.String("jti", claims.JTI))

	// 将访问令牌加入黑名单
	expiration := s.jwtManager.GetTokenExpiration(claims)
	if expiration > 0 {
		err = s.sessionService.AddTokenToBlacklist(ctx, claims.JTI, expiration)
		if err != nil {
			logger.Error("添加访问令牌到黑名单失败", 
				zap.Uint("user_id", userID),
				zap.String("jti", claims.JTI),
				zap.Error(err),
				zap.String("operation", "logout"))
			return apperrors.NewInternalError("添加令牌到黑名单失败")
		}
		logger.Debug("访问令牌已加入黑名单", 
			zap.Uint("user_id", userID),
			zap.String("jti", claims.JTI))
	}

	// 如果提供了刷新令牌，也验证并拉黑
	if req.RefreshToken != "" {
		refreshClaims, err := s.jwtManager.ValidateRefreshToken(req.RefreshToken)
		if err == nil {
			refreshExpiration := s.jwtManager.GetTokenExpiration(refreshClaims)
			if refreshExpiration > 0 {
				s.sessionService.AddTokenToBlacklist(ctx, refreshClaims.JTI, refreshExpiration)
				logger.Debug("刷新令牌已加入黑名单", 
					zap.Uint("user_id", userID),
					zap.String("refresh_jti", refreshClaims.JTI))
			}
		} else {
			logger.Warn("刷新令牌验证失败", 
				zap.Uint("user_id", userID),
				zap.Error(err))
		}
	}

	// 删除会话
	err = s.sessionService.DeleteSession(ctx, userID)
	if err != nil {
		logger.Error("删除会话失败", 
			zap.Uint("user_id", userID),
			zap.Error(err),
			zap.String("operation", "logout"))
		return apperrors.NewInternalError("删除会话失败")
	}

	logger.Info("用户登出成功", 
		zap.Uint("user_id", userID),
		zap.String("operation", "logout"))

	return nil
}

func (s *UserService) GetByID(id uint) (*model.User, error) {
	logger.Debug("查询用户信息", 
		zap.Uint("user_id", id),
		zap.String("operation", "get_user"))

	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("用户不存在", 
				zap.Uint("user_id", id),
				zap.String("operation", "get_user"))
			return nil, apperrors.NewNotFoundError("用户不存在")
		}
		logger.Error("查询用户失败", 
			zap.Uint("user_id", id),
			zap.Error(err),
			zap.String("operation", "get_user"))
		return nil, apperrors.NewInternalError("查询用户失败")
	}

	logger.Debug("用户查询成功", 
		zap.Uint("user_id", id),
		zap.String("username", user.Username),
		zap.String("operation", "get_user"))

	return user, nil
}

func (s *UserService) Update(id uint, req *model.UpdateUserRequest) (*model.User, error) {
	logger.Info("开始更新用户信息", 
		zap.Uint("user_id", id),
		zap.String("operation", "update_user"))

	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("更新失败：用户不存在", 
				zap.Uint("user_id", id),
				zap.String("operation", "update_user"))
			return nil, apperrors.NewNotFoundError("用户不存在")
		}
		logger.Error("查询用户失败", 
			zap.Uint("user_id", id),
			zap.Error(err),
			zap.String("operation", "update_user"))
		return nil, apperrors.NewInternalError("查询用户失败")
	}

	// 记录更新的字段
	updatedFields := []string{}
	if req.Username != "" {
		logger.Debug("更新用户名", 
			zap.Uint("user_id", id),
			zap.String("old_username", user.Username),
			zap.String("new_username", req.Username))
		user.Username = req.Username
		updatedFields = append(updatedFields, "username")
	}
	if req.Email != "" {
		logger.Debug("更新邮箱", 
			zap.Uint("user_id", id),
			zap.String("old_email", user.Email),
			zap.String("new_email", req.Email))
		user.Email = req.Email
		updatedFields = append(updatedFields, "email")
	}
	roleChanged := false
	if req.Role != "" && req.Role != user.Role {
		logger.Debug("更新角色", 
			zap.Uint("user_id", id),
			zap.String("old_role", user.Role),
			zap.String("new_role", req.Role))
		user.Role = req.Role
		updatedFields = append(updatedFields, "role")
		roleChanged = true
	}
	if req.Status != "" {
		logger.Debug("更新状态", 
			zap.Uint("user_id", id),
			zap.String("old_status", user.Status),
			zap.String("new_status", req.Status))
		user.Status = req.Status
		updatedFields = append(updatedFields, "status")
	}

	err = s.userRepo.Update(user)
	if err != nil {
		logger.Error("用户更新失败", 
			zap.Uint("user_id", id),
			zap.Strings("updated_fields", updatedFields),
			zap.Error(err),
			zap.String("operation", "update_user"))
		return nil, apperrors.NewInternalError("用户更新失败")
	}

	// 如果角色发生变化，同步到 user_roles 表
	if roleChanged {
		if err := s.syncUserRole(user.ID, user.Role); err != nil {
			logger.Error("同步用户角色失败", 
				zap.Uint("user_id", user.ID),
				zap.String("role", user.Role),
				zap.Error(err))
			// 不返回错误，因为用户已更新成功
		}
	}

	logger.Info("用户更新成功", 
		zap.Uint("user_id", id),
		zap.String("username", user.Username),
		zap.Strings("updated_fields", updatedFields),
		zap.String("operation", "update_user"))

	return user, nil
}

func (s *UserService) Delete(id uint) error {
	logger.Info("开始删除用户", 
		zap.Uint("user_id", id),
		zap.String("operation", "delete_user"))

	// 先查询用户信息用于日志记录
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("删除失败：用户不存在", 
				zap.Uint("user_id", id),
				zap.String("operation", "delete_user"))
			return apperrors.NewNotFoundError("用户不存在")
		}
		logger.Error("查询用户失败", 
			zap.Uint("user_id", id),
			zap.Error(err),
			zap.String("operation", "delete_user"))
		return apperrors.NewInternalError("查询用户失败")
	}

	err = s.userRepo.Delete(id)
	if err != nil {
		logger.Error("用户删除失败", 
			zap.Uint("user_id", id),
			zap.String("username", user.Username),
			zap.Error(err),
			zap.String("operation", "delete_user"))
		return apperrors.NewInternalError("用户删除失败")
	}

	logger.Info("用户删除成功", 
		zap.Uint("user_id", id),
		zap.String("username", user.Username),
		zap.String("operation", "delete_user"))

	return nil
}

func (s *UserService) List(page, pageSize int) ([]model.User, int64, error) {
	logger.Debug("查询用户列表", 
		zap.Int("page", page),
		zap.Int("page_size", pageSize),
		zap.String("operation", "list_users"))

	offset := (page - 1) * pageSize
	users, total, err := s.userRepo.List(offset, pageSize)
	if err != nil {
		logger.Error("查询用户列表失败", 
			zap.Int("page", page),
			zap.Int("page_size", pageSize),
			zap.Error(err),
			zap.String("operation", "list_users"))
		return nil, 0, apperrors.NewInternalError("查询用户列表失败")
	}

	logger.Debug("用户列表查询成功", 
		zap.Int("page", page),
		zap.Int("page_size", pageSize),
		zap.Int64("total", total),
		zap.Int("returned_count", len(users)),
		zap.String("operation", "list_users"))

	return users, total, nil
}

// ListWithVisibility 根据用户角色和可见性规则获取用户列表
func (s *UserService) ListWithVisibility(currentUserID uint, currentUserRole string, page, pageSize int) ([]model.User, int64, error) {
	logger.Debug("查询用户列表（带可见性控制）", 
		zap.Uint("current_user_id", currentUserID),
		zap.String("current_user_role", currentUserRole),
		zap.Int("page", page),
		zap.Int("page_size", pageSize),
		zap.String("operation", "list_users_with_visibility"))

	offset := (page - 1) * pageSize
	users, total, err := s.userRepo.ListByVisibility(currentUserID, currentUserRole, offset, pageSize)
	if err != nil {
		logger.Error("查询用户列表失败", 
			zap.Uint("current_user_id", currentUserID),
			zap.String("current_user_role", currentUserRole),
			zap.Int("page", page),
			zap.Int("page_size", pageSize),
			zap.Error(err),
			zap.String("operation", "list_users_with_visibility"))
		return nil, 0, apperrors.NewInternalError("查询用户列表失败")
	}

	logger.Debug("用户列表查询成功", 
		zap.Uint("current_user_id", currentUserID),
		zap.String("current_user_role", currentUserRole),
		zap.Int("page", page),
		zap.Int("page_size", pageSize),
		zap.Int64("total", total),
		zap.Int("returned_count", len(users)),
		zap.String("operation", "list_users_with_visibility"))

	return users, total, nil
}

// CheckUsernameAvailable 检查用户名是否可用
func (s *UserService) CheckUsernameAvailable(username string) (bool, error) {
	logger.Debug("检查用户名可用性", 
		zap.String("username", username),
		zap.String("operation", "check_username"))

	exists, err := s.userRepo.CheckUsernameExists(username)
	if err != nil {
		logger.Error("检查用户名可用性失败", 
			zap.String("username", username),
			zap.Error(err),
			zap.String("operation", "check_username"))
		return false, apperrors.NewInternalError("检查用户名失败")
	}

	available := !exists
	logger.Debug("用户名可用性检查完成", 
		zap.String("username", username),
		zap.Bool("available", available),
		zap.String("operation", "check_username"))

	return available, nil // 不存在则可用
}

// CheckEmailAvailable 检查邮箱是否可用
func (s *UserService) CheckEmailAvailable(email string) (bool, error) {
	logger.Debug("检查邮箱可用性", 
		zap.String("email", email),
		zap.String("operation", "check_email"))

	exists, err := s.userRepo.CheckEmailExists(email)
	if err != nil {
		logger.Error("检查邮箱可用性失败", 
			zap.String("email", email),
			zap.Error(err),
			zap.String("operation", "check_email"))
		return false, apperrors.NewInternalError("检查邮箱失败")
	}

	available := !exists
	logger.Debug("邮箱可用性检查完成", 
		zap.String("email", email),
		zap.Bool("available", available),
		zap.String("operation", "check_email"))

	return available, nil // 不存在则可用
}

// CheckUserDataAvailability 批量检查用户数据可用性
func (s *UserService) CheckUserDataAvailability(req *model.CheckAvailabilityRequest) (*model.CheckAvailabilityResponse, error) {
	logger.Debug("开始批量检查用户数据可用性", 
		zap.String("username", req.Username),
		zap.String("email", req.Email),
		zap.Any("exclude_user_id", req.ExcludeUserID),
		zap.String("operation", "check_availability"))

	response := &model.CheckAvailabilityResponse{}

	// 检查用户名
	if req.Username != "" {
		var available bool
		var err error
		
		if req.ExcludeUserID != nil && *req.ExcludeUserID > 0 {
			logger.Debug("检查用户名可用性（排除指定用户）", 
				zap.String("username", req.Username),
				zap.Uint("exclude_user_id", *req.ExcludeUserID))
			exists, err := s.userRepo.CheckUsernameExistsExcludeID(req.Username, *req.ExcludeUserID)
			if err != nil {
				logger.Error("检查用户名可用性失败", 
					zap.String("username", req.Username),
					zap.Uint("exclude_user_id", *req.ExcludeUserID),
					zap.Error(err),
					zap.String("operation", "check_availability"))
				return nil, apperrors.NewInternalError("检查用户名失败")
			}
			available = !exists
		} else {
			available, err = s.CheckUsernameAvailable(req.Username)
			if err != nil {
				return nil, err
			}
		}

		message := "用户名可用"
		if !available {
			message = "用户名已被使用"
		}

		response.Username = &model.AvailabilityResult{
			Available: available,
			Message:   message,
		}

		logger.Debug("用户名可用性检查结果", 
			zap.String("username", req.Username),
			zap.Bool("available", available))
	}

	// 检查邮箱
	if req.Email != "" {
		var available bool
		var err error
		
		if req.ExcludeUserID != nil && *req.ExcludeUserID > 0 {
			logger.Debug("检查邮箱可用性（排除指定用户）", 
				zap.String("email", req.Email),
				zap.Uint("exclude_user_id", *req.ExcludeUserID))
			exists, err := s.userRepo.CheckEmailExistsExcludeID(req.Email, *req.ExcludeUserID)
			if err != nil {
				logger.Error("检查邮箱可用性失败", 
					zap.String("email", req.Email),
					zap.Uint("exclude_user_id", *req.ExcludeUserID),
					zap.Error(err),
					zap.String("operation", "check_availability"))
				return nil, apperrors.NewInternalError("检查邮箱失败")
			}
			available = !exists
		} else {
			available, err = s.CheckEmailAvailable(req.Email)
			if err != nil {
				return nil, err
			}
		}

		message := "邮箱可用"
		if !available {
			message = "邮箱已被使用"
		}

		response.Email = &model.AvailabilityResult{
			Available: available,
			Message:   message,
		}

		logger.Debug("邮箱可用性检查结果", 
			zap.String("email", req.Email),
			zap.Bool("available", available))
	}

	logger.Debug("批量检查用户数据可用性完成", 
		zap.String("operation", "check_availability"))

	return response, nil
}

// GetUserPermissions 获取用户的权限信息
func (s *UserService) GetUserPermissions(userID uint) (*model.UserPermissionsResponse, error) {
	logger.Info("获取用户权限", zap.Uint("user_id", userID))

	// 获取用户角色
	roles, err := s.roleRepo.GetUserRoles(userID)
	if err != nil {
		logger.Error("获取用户角色失败", 
			zap.Uint("user_id", userID),
			zap.Error(err))
		return nil, apperrors.NewInternalError("获取用户角色失败")
	}

	// 提取角色代码
	roleCodes := make([]string, 0, len(roles))
	for _, role := range roles {
		roleCodes = append(roleCodes, role.Code)
	}

	// 获取所有权限代码
	permissionCodes := make(map[string]bool)
	
	// 遍历每个角色，获取其权限
	for _, role := range roles {
		// 获取角色的权限ID列表
		permissionIDs, err := s.roleRepo.GetRolePermissionIDs(role.ID)
		if err != nil {
			logger.Error("获取角色权限ID失败", 
				zap.Uint("role_id", role.ID),
				zap.Error(err))
			continue
		}

		// 获取权限详情
		permissions, err := s.permissionService.permissionRepo.GetByIDs(permissionIDs)
		if err != nil {
			logger.Error("获取权限详情失败", 
				zap.Uints("permission_ids", permissionIDs),
				zap.Error(err))
			continue
		}

		// 收集权限代码
		for _, perm := range permissions {
			if perm.Status == "active" {
				permissionCodes[perm.Code] = true
			}
		}
	}

	// 转换为列表
	permissions := make([]string, 0, len(permissionCodes))
	for code := range permissionCodes {
		permissions = append(permissions, code)
	}

	logger.Info("获取用户权限成功", 
		zap.Uint("user_id", userID),
		zap.Int("role_count", len(roles)),
		zap.Int("permission_count", len(permissions)))

	return &model.UserPermissionsResponse{
		Roles:       roleCodes,
		Permissions: permissions,
		Menus:       []model.MenuResponse{}, // 菜单由前端单独获取
	}, nil
}

// syncUserRole 同步用户角色到 user_roles 表和 Casbin
func (s *UserService) syncUserRole(userID uint, roleCode string) error {
	if roleCode == "" {
		return nil
	}

	// 根据角色代码查找角色
	role, err := s.roleRepo.GetByCode(roleCode)
	if err != nil {
		logger.Error("查找角色失败", 
			zap.Uint("user_id", userID),
			zap.String("role_code", roleCode),
			zap.Error(err))
		return apperrors.NewInternalError("查找角色失败")
	}

	// 移除用户的所有现有角色
	if err := s.roleRepo.RemoveAllRolesFromUser(userID); err != nil {
		logger.Error("移除用户角色失败", 
			zap.Uint("user_id", userID),
			zap.Error(err))
		return apperrors.NewInternalError("移除用户角色失败")
	}

	// 分配新角色到 user_roles 表
	if err := s.roleRepo.AssignRoleToUser(userID, role.ID, 0); err != nil {
		logger.Error("分配角色失败", 
			zap.Uint("user_id", userID),
			zap.Uint("role_id", role.ID),
			zap.Error(err))
		return apperrors.NewInternalError("分配角色失败")
	}

	logger.Info("同步用户角色成功", 
		zap.Uint("user_id", userID),
		zap.String("role_code", roleCode),
		zap.Uint("role_id", role.ID))

	return nil
}