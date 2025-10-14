package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// AuditLogConfig 审计日志配置
type AuditLogConfig struct {
	Enabled         bool     // 是否启用审计日志
	LogRequestBody  bool     // 是否记录请求体
	MaxBodySize     int      // 最大请求体大小（字节）
	SkipPaths       []string // 跳过的路径
	SensitivePaths  []string // 敏感路径（强制记录）
}

// DefaultAuditLogConfig 默认审计日志配置
func DefaultAuditLogConfig() AuditLogConfig {
	return AuditLogConfig{
		Enabled:        true,
		LogRequestBody: true,
		MaxBodySize:    10240, // 10KB
		SkipPaths: []string{
			"/health",
			"/api/v1/auth/captcha",
		},
		SensitivePaths: []string{
			"/api/v1/users",
			"/api/v1/roles",
			"/api/v1/permissions",
		},
	}
}

// AuditLogger 审计日志中间件
func AuditLogger(db *gorm.DB, config AuditLogConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查是否启用
		if !config.Enabled {
			c.Next()
			return
		}

		// 检查是否跳过该路径
		path := c.Request.URL.Path
		for _, skipPath := range config.SkipPaths {
			if path == skipPath {
				c.Next()
				return
			}
		}

		// 记录开始时间
		startTime := time.Now()

		// 读取请求体（如果需要）
		var requestBody string
		if config.LogRequestBody && shouldLogBody(c.Request.Method) {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				// 限制大小
				if len(bodyBytes) <= config.MaxBodySize {
					requestBody = string(bodyBytes)
				} else {
					requestBody = string(bodyBytes[:config.MaxBodySize]) + "...(truncated)"
				}
				// 恢复请求体供后续使用
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// 创建自定义 ResponseWriter 以捕获状态码
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 处理请求
		c.Next()

		// 计算耗时
		duration := time.Since(startTime).Milliseconds()

		// 获取用户信息
		var userID uint
		var username string
		if uid, exists := c.Get("user_id"); exists {
			userID = uid.(uint)
		}
		if uname, exists := c.Get("username"); exists {
			username = uname.(string)
		}

		// 获取错误信息
		var errorMsg string
		if len(c.Errors) > 0 {
			errorMsg = c.Errors.String()
		}

		// 构建审计日志
		auditLog := &model.AuditLog{
			UserID:      userID,
			Username:    username,
			Action:      buildAction(c.Request.Method, path),
			Resource:    extractResource(path),
			ResourceID:  c.Param("id"),
			Method:      c.Request.Method,
			Path:        path,
			IP:          c.ClientIP(),
			UserAgent:   c.Request.UserAgent(),
			Status:      blw.status,
			ErrorMsg:    errorMsg,
			RequestBody: requestBody,
			Duration:    duration,
		}

		// 异步保存审计日志
		go saveAuditLog(db, auditLog)

		// 记录到应用日志（敏感操作或失败请求）
		if isSensitiveOperation(path, config.SensitivePaths) || blw.status >= 400 {
			logger.Info("审计日志",
				zap.Uint("user_id", userID),
				zap.String("username", username),
				zap.String("action", auditLog.Action),
				zap.String("path", path),
				zap.String("method", c.Request.Method),
				zap.Int("status", blw.status),
				zap.Int64("duration_ms", duration),
				zap.String("ip", auditLog.IP))
		}
	}
}

// bodyLogWriter 自定义 ResponseWriter 以捕获状态码
type bodyLogWriter struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *bodyLogWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// shouldLogBody 判断是否应该记录请求体
func shouldLogBody(method string) bool {
	return method == "POST" || method == "PUT" || method == "PATCH"
}

// buildAction 构建操作描述
func buildAction(method, path string) string {
	action := method + " " + path
	
	// 可以根据路径和方法生成更友好的描述
	switch method {
	case "POST":
		if contains(path, "/login") {
			return "用户登录"
		} else if contains(path, "/logout") {
			return "用户登出"
		} else if contains(path, "/register") {
			return "用户注册"
		}
		return "创建资源: " + path
	case "PUT", "PATCH":
		return "更新资源: " + path
	case "DELETE":
		return "删除资源: " + path
	case "GET":
		return "查询资源: " + path
	default:
		return action
	}
}

// extractResource 从路径中提取资源类型
func extractResource(path string) string {
	// 简单实现：提取路径的第一个部分
	// 例如：/api/v1/users/123 -> users
	parts := splitPath(path)
	if len(parts) >= 3 {
		return parts[2] // 跳过 /api/v1
	}
	return path
}

// splitPath 分割路径
func splitPath(path string) []string {
	var parts []string
	start := 0
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			if i > start {
				parts = append(parts, path[start:i])
			}
			start = i + 1
		}
	}
	if start < len(path) {
		parts = append(parts, path[start:])
	}
	return parts
}

// contains 检查字符串是否包含子串
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// isSensitiveOperation 判断是否为敏感操作
func isSensitiveOperation(path string, sensitivePaths []string) bool {
	for _, sensitivePath := range sensitivePaths {
		if contains(path, sensitivePath) {
			return true
		}
	}
	return false
}

// saveAuditLog 保存审计日志到数据库
func saveAuditLog(db *gorm.DB, auditLog *model.AuditLog) {
	if err := db.Create(auditLog).Error; err != nil {
		logger.Error("保存审计日志失败",
			zap.Uint("user_id", auditLog.UserID),
			zap.String("action", auditLog.Action),
			zap.Error(err))
	}
}

// AuditAction 手动记录审计日志（用于非 HTTP 操作）
func AuditAction(db *gorm.DB, userID uint, username, action, resource string) {
	auditLog := &model.AuditLog{
		UserID:   userID,
		Username: username,
		Action:   action,
		Resource: resource,
		Status:   200,
	}
	
	go saveAuditLog(db, auditLog)
	
	logger.Info("审计日志",
		zap.Uint("user_id", userID),
		zap.String("username", username),
		zap.String("action", action),
		zap.String("resource", resource))
}
