package middleware

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
)

// CasbinEnforcer Casbin 权限中间件
// 使用 Casbin 进行权限检查，需要先通过 JWT 认证中间件
func CasbinEnforcer(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户ID（由 JWT 中间件设置）
		userID, exists := c.Get("user_id")
		if !exists {
			logger.Warn("权限检查失败：未找到用户ID",
				zap.String("path", c.Request.URL.Path),
				zap.String("method", c.Request.Method))
			utils.Unauthorized(c, "未授权访问")
			c.Abort()
			return
		}

		// 构建 Casbin 主体
		sub := fmt.Sprintf("user:%d", userID.(uint))
		obj := c.Request.URL.Path
		act := c.Request.Method

		// 执行权限检查
		ok, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			logger.Error("权限检查失败",
				zap.Uint("user_id", userID.(uint)),
				zap.String("path", obj),
				zap.String("method", act),
				zap.Error(err))
			utils.InternalServerError(c, "权限检查失败")
			c.Abort()
			return
		}

		if !ok {
			logger.Warn("权限不足",
				zap.Uint("user_id", userID.(uint)),
				zap.String("path", obj),
				zap.String("method", act))
			utils.Forbidden(c, "您没有权限访问此资源")
			c.Abort()
			return
		}

		// 权限检查通过，继续处理请求
		logger.Debug("权限检查通过",
			zap.Uint("user_id", userID.(uint)),
			zap.String("path", obj),
			zap.String("method", act))

		c.Next()
	}
}

// RequirePermission 要求特定权限的中间件（用于单个路由）
// 使用方式：router.GET("/api/users", middleware.RequirePermission(enforcer, "user:read"), handler)
func RequirePermission(enforcer *casbin.Enforcer, permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户ID
		userID, exists := c.Get("user_id")
		if !exists {
			logger.Warn("权限检查失败：未找到用户ID",
				zap.String("permission", permission))
			utils.Unauthorized(c, "未授权访问")
			c.Abort()
			return
		}

		// 构建 Casbin 主体
		sub := fmt.Sprintf("user:%d", userID.(uint))
		obj := c.Request.URL.Path
		act := c.Request.Method

		// 执行权限检查
		ok, err := enforcer.Enforce(sub, obj, act)
		if err != nil {
			logger.Error("权限检查失败",
				zap.Uint("user_id", userID.(uint)),
				zap.String("permission", permission),
				zap.Error(err))
			utils.InternalServerError(c, "权限检查失败")
			c.Abort()
			return
		}

		if !ok {
			logger.Warn("权限不足",
				zap.Uint("user_id", userID.(uint)),
				zap.String("permission", permission),
				zap.String("path", obj))
			utils.Forbidden(c, fmt.Sprintf("需要 %s 权限", permission))
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireRole 要求特定角色的中间件
// 使用方式：router.GET("/admin/users", middleware.RequireRole(enforcer, "admin"), handler)
func RequireRole(enforcer *casbin.Enforcer, roleCode string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户ID
		userID, exists := c.Get("user_id")
		if !exists {
			logger.Warn("角色检查失败：未找到用户ID",
				zap.String("required_role", roleCode))
			utils.Unauthorized(c, "未授权访问")
			c.Abort()
			return
		}

		// 构建 Casbin 主体
		sub := fmt.Sprintf("user:%d", userID.(uint))
		role := fmt.Sprintf("role:%s", roleCode)

		// 检查用户是否拥有该角色
		ok, err := enforcer.HasRoleForUser(sub, role)
		if err != nil {
			logger.Error("角色检查失败",
				zap.Uint("user_id", userID.(uint)),
				zap.String("required_role", roleCode),
				zap.Error(err))
			utils.InternalServerError(c, "角色检查失败")
			c.Abort()
			return
		}

		if !ok {
			logger.Warn("角色不足",
				zap.Uint("user_id", userID.(uint)),
				zap.String("required_role", roleCode))
			utils.Forbidden(c, fmt.Sprintf("需要 %s 角色", roleCode))
			c.Abort()
			return
		}

		logger.Debug("角色检查通过",
			zap.Uint("user_id", userID.(uint)),
			zap.String("role", roleCode))

		c.Next()
	}
}

// RequireAnyRole 要求任意一个角色的中间件
// 使用方式：router.GET("/api/data", middleware.RequireAnyRole(enforcer, []string{"admin", "manager"}), handler)
func RequireAnyRole(enforcer *casbin.Enforcer, roleCodes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户ID
		userID, exists := c.Get("user_id")
		if !exists {
			logger.Warn("角色检查失败：未找到用户ID")
			utils.Unauthorized(c, "未授权访问")
			c.Abort()
			return
		}

		// 构建 Casbin 主体
		sub := fmt.Sprintf("user:%d", userID.(uint))

		// 检查用户是否拥有任意一个角色
		hasRole := false
		for _, roleCode := range roleCodes {
			role := fmt.Sprintf("role:%s", roleCode)
			ok, err := enforcer.HasRoleForUser(sub, role)
			if err != nil {
				logger.Error("角色检查失败",
					zap.Uint("user_id", userID.(uint)),
					zap.String("role", roleCode),
					zap.Error(err))
				continue
			}
			if ok {
				hasRole = true
				break
			}
		}

		if !hasRole {
			logger.Warn("角色不足",
				zap.Uint("user_id", userID.(uint)),
				zap.Strings("required_roles", roleCodes))
			utils.Forbidden(c, "您没有权限访问此资源")
			c.Abort()
			return
		}

		logger.Debug("角色检查通过",
			zap.Uint("user_id", userID.(uint)))

		c.Next()
	}
}

// SkipPermissionCheck 跳过权限检查的中间件（用于公开接口）
// 使用方式：在路由组中排除某些路径
func SkipPermissionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置标记，表示跳过权限检查
		c.Set("skip_permission_check", true)
		c.Next()
	}
}
