package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/auth"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/cache"
)

// SessionInfo 表示用户会话信息
// - UserID: 用户ID
// - Username: 用户名
// - RefreshToken: 刷新令牌（用于续期）
// - DeviceInfo: 设备信息
// - IPAddress: 登录时的IP地址
// - UserAgent: 浏览器/客户端标识
// - LoginTime: 登录时间
// - LastActivity: 最后活跃时间
type SessionInfo struct {
	UserID       uint      `json:"user_id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	DeviceInfo   string    `json:"device_info"`
	IPAddress    string    `json:"ip_address"`
	UserAgent    string    `json:"user_agent"`
	LoginTime    time.Time `json:"login_time"`
	LastActivity time.Time `json:"last_activity"`
}

// SessionService 会话服务
// 负责管理用户会话、令牌黑名单、用户活跃状态以及权限缓存
type SessionService struct {
	redisClient *cache.RedisClient
	jwtManager  *auth.JWTManager
}

// NewSessionService 创建会话服务实例
func NewSessionService(redisClient *cache.RedisClient, jwtManager *auth.JWTManager) *SessionService {
	return &SessionService{
		redisClient: redisClient,
		jwtManager:  jwtManager,
	}
}

// CreateSession 创建一个新的用户会话并存储到 Redis
// 会话有效期为 30 天（与刷新令牌一致）
func (s *SessionService) CreateSession(ctx context.Context, userID uint, username, refreshToken, deviceInfo, ipAddress, userAgent string) error {
	sessionInfo := &SessionInfo{
		UserID:       userID,
		Username:     username,
		RefreshToken: refreshToken,
		DeviceInfo:   deviceInfo,
		IPAddress:    ipAddress,
		UserAgent:    userAgent,
		LoginTime:    time.Now(),
		LastActivity: time.Now(),
	}

	sessionData, err := json.Marshal(sessionInfo)
	if err != nil {
		return fmt.Errorf("序列化会话信息失败: %w", err)
	}

	sessionKey := fmt.Sprintf("user:session:%d", userID)
	return s.redisClient.Set(ctx, sessionKey, sessionData, 30*24*time.Hour)
}

// GetSession 从 Redis 获取用户会话信息
func (s *SessionService) GetSession(ctx context.Context, userID uint) (*SessionInfo, error) {
	sessionKey := fmt.Sprintf("user:session:%d", userID)
	sessionData, err := s.redisClient.Get(ctx, sessionKey)
	if err != nil {
		return nil, fmt.Errorf("未找到会话: %w", err)
	}

	var sessionInfo SessionInfo
	if err := json.Unmarshal([]byte(sessionData), &sessionInfo); err != nil {
		return nil, fmt.Errorf("反序列化会话信息失败: %w", err)
	}

	return &sessionInfo, nil
}

// UpdateLastActivity 更新用户会话的最后活跃时间
// 每次用户有请求时调用，用于刷新 TTL
func (s *SessionService) UpdateLastActivity(ctx context.Context, userID uint) error {
	sessionInfo, err := s.GetSession(ctx, userID)
	if err != nil {
		return err
	}

	sessionInfo.LastActivity = time.Now()

	sessionData, err := json.Marshal(sessionInfo)
	if err != nil {
		return fmt.Errorf("序列化会话信息失败: %w", err)
	}

	sessionKey := fmt.Sprintf("user:session:%d", userID)
	return s.redisClient.Set(ctx, sessionKey, sessionData, 30*24*time.Hour)
}

// DeleteSession 删除 Redis 中的用户会话
func (s *SessionService) DeleteSession(ctx context.Context, userID uint) error {
	sessionKey := fmt.Sprintf("user:session:%d", userID)
	return s.redisClient.Del(ctx, sessionKey)
}

// ValidateRefreshToken 校验刷新令牌并验证 Redis 中的会话
// 步骤：
// 1. 校验刷新令牌的有效性（JWT 格式）
// 2. 检查是否在黑名单
// 3. 从 Redis 获取会话并校验是否匹配
func (s *SessionService) ValidateRefreshToken(ctx context.Context, refreshToken string) (*SessionInfo, error) {
	claims, err := s.jwtManager.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("刷新令牌无效: %w", err)
	}

	if s.IsTokenBlacklisted(ctx, claims.JTI) {
		return nil, fmt.Errorf("刷新令牌已被加入黑名单")
	}

	sessionInfo, err := s.GetSession(ctx, claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("未找到会话: %w", err)
	}

	if sessionInfo.RefreshToken != refreshToken {
		return nil, fmt.Errorf("刷新令牌不匹配")
	}

	return sessionInfo, nil
}

// AddTokenToBlacklist 将指定 JTI 的令牌加入黑名单（设置过期时间）
func (s *SessionService) AddTokenToBlacklist(ctx context.Context, jti string, expiration time.Duration) error {
	blacklistKey := fmt.Sprintf("token:blacklist:%s", jti)
	return s.redisClient.Set(ctx, blacklistKey, "blacklisted", expiration)
}

// IsTokenBlacklisted 检查指定 JTI 的令牌是否在黑名单中
func (s *SessionService) IsTokenBlacklisted(ctx context.Context, jti string) bool {
	blacklistKey := fmt.Sprintf("token:blacklist:%s", jti)
	exists, err := s.redisClient.Exists(ctx, blacklistKey)
	if err != nil {
		return false
	}
	return exists > 0
}

// SetUserActive 设置用户为活跃状态（TTL 30 分钟）
// 一般在用户请求时调用，用于标记在线状态
func (s *SessionService) SetUserActive(ctx context.Context, userID uint) error {
	activeKey := fmt.Sprintf("user:active:%d", userID)
	return s.redisClient.Set(ctx, activeKey, time.Now().Unix(), 30*time.Minute)
}

// IsUserActive 判断用户当前是否处于活跃状态
func (s *SessionService) IsUserActive(ctx context.Context, userID uint) bool {
	activeKey := fmt.Sprintf("user:active:%d", userID)
	exists, err := s.redisClient.Exists(ctx, activeKey)
	if err != nil {
		return false
	}
	return exists > 0
}

// CacheUserPermissions 缓存用户角色和权限到 Redis
// 缓存内容包括：角色、权限列表、缓存时间
// TTL 默认 1 小时
func (s *SessionService) CacheUserPermissions(ctx context.Context, userID uint, role string, permissions []string) error {
	permissionData := map[string]interface{}{
		"role":        role,
		"permissions": permissions,
		"cached_at":   time.Now().Unix(),
	}

	data, err := json.Marshal(permissionData)
	if err != nil {
		return fmt.Errorf("序列化权限数据失败: %w", err)
	}

	permissionKey := fmt.Sprintf("user:permissions:%d", userID)
	return s.redisClient.Set(ctx, permissionKey, data, time.Hour)
}

// GetCachedUserPermissions 获取缓存的用户权限
// 返回用户的角色和权限列表
func (s *SessionService) GetCachedUserPermissions(ctx context.Context, userID uint) (string, []string, error) {
	permissionKey := fmt.Sprintf("user:permissions:%d", userID)
	data, err := s.redisClient.Get(ctx, permissionKey)
	if err != nil {
		return "", nil, fmt.Errorf("权限未缓存: %w", err)
	}

	var permissionData map[string]interface{}
	if err := json.Unmarshal([]byte(data), &permissionData); err != nil {
		return "", nil, fmt.Errorf("反序列化权限数据失败: %w", err)
	}

	role, _ := permissionData["role"].(string)
	permissionsInterface, _ := permissionData["permissions"].([]interface{})

	permissions := make([]string, len(permissionsInterface))
	for i, p := range permissionsInterface {
		permissions[i], _ = p.(string)
	}

	return role, permissions, nil
}

// CleanupExpiredSessions 清理过期会话（占位方法）
// 实际上 Redis TTL 已经能自动清理大部分过期会话
// 可以作为定时任务扩展，用于更复杂的清理逻辑
func (s *SessionService) CleanupExpiredSessions(ctx context.Context) error {
	return nil
}
