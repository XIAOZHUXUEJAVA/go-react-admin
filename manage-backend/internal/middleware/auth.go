package middleware

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/auth"
)

// SessionServiceInterface 会话服务接口（用于中间件扩展）
// 提供了 Token 黑名单检测和用户活跃状态更新的方法
type SessionServiceInterface interface {
	IsTokenBlacklisted(ctx context.Context, jti string) bool // 判断 Token 是否在黑名单中
	UpdateLastActivity(ctx context.Context, userID uint) error // 更新用户最后活跃时间
	SetUserActive(ctx context.Context, userID uint) error      // 设置用户为活跃状态
}

// JWTAuth 基础 JWT 鉴权中间件（不包含会话服务）
func JWTAuth(jwtManager *auth.JWTManager) gin.HandlerFunc {
	return JWTAuthWithSession(jwtManager, nil)
}

// JWTAuthWithSession 带会话管理的 JWT 鉴权中间件
// jwtManager: JWT 管理器，用于验证 Token
// sessionService: 可选，会话服务接口（支持 Token 黑名单和用户活跃状态更新）
func JWTAuthWithSession(jwtManager *auth.JWTManager, sessionService SessionServiceInterface) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// 从请求头中获取 Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Unauthorized(c, "需要提供 Authorization 头")
			c.Abort()
			return
		}

		// 提取 Bearer Token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			utils.Unauthorized(c, "Authorization 格式错误，需要 Bearer Token")
			c.Abort()
			return
		}

		// 验证 Token
		claims, err := jwtManager.ValidateToken(tokenString)
		if err != nil {
			utils.Unauthorized(c, "无效的 Token")
			c.Abort()
			return
		}

		// 如果启用了会话服务，则检查黑名单和更新用户状态
		if sessionService != nil {
			ctx := context.Background()

			// 检查 Token 是否已被拉黑
			if sessionService.IsTokenBlacklisted(ctx, claims.JTI) {
				utils.Unauthorized(c, "Token 已被吊销")
				c.Abort()
				return
			}

			// 更新用户活跃状态
			sessionService.UpdateLastActivity(ctx, claims.UserID)
			sessionService.SetUserActive(ctx, claims.UserID)
		}

		// 将用户信息保存到 Gin Context 中，后续 Handler 可以直接使用
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("jti", claims.JTI)
		c.Set("access_token", tokenString)

		c.Next()
	})
}

// RefreshTokenAuth 刷新 Token 鉴权中间件
// 用于刷新 Access Token 的接口，要求客户端传入 refresh_token
func RefreshTokenAuth(jwtManager *auth.JWTManager) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		// 解析请求体，必须包含 refresh_token
		var req struct {
			RefreshToken string `json:"refresh_token" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			utils.BadRequest(c, "请求格式错误，缺少 refresh_token")
			c.Abort()
			return
		}

		// 验证 refresh token
		claims, err := jwtManager.ValidateRefreshToken(req.RefreshToken)
		if err != nil {
			utils.Unauthorized(c, "无效的 Refresh Token")
			c.Abort()
			return
		}

		// 将用户信息保存到 Gin Context，供后续 Handler 使用
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("refresh_token", req.RefreshToken)

		c.Next()
	})
}
