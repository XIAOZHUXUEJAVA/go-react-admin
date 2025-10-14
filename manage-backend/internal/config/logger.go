package config

import (
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
)

// LogConfigDetails 使用结构化日志记录配置详情
func LogConfigDetails(cfg *Config) {
	logger.Info("🔧 运行时配置详情")
	logger.Info("配置信息", 
		zap.String("环境", cfg.Environment),
		zap.String("端口", cfg.Port),
		zap.String("日志级别", cfg.LogLevel))
	
	logger.Info("数据库配置",
		zap.String("用户", cfg.Database.User),
		zap.String("主机", cfg.Database.Host),
		zap.String("端口", cfg.Database.Port),
		zap.String("数据库名", cfg.Database.Name),
		zap.String("模式", cfg.Database.Schema))
	
	redisPassword := "无"
	if cfg.Redis.Password != "" {
		redisPassword = "***已设置***"
	}
	logger.Info("Redis配置",
		zap.String("主机", cfg.Redis.Host),
		zap.String("端口", cfg.Redis.Port),
		zap.Int("数据库", cfg.Redis.DB),
		zap.String("密码", redisPassword))
	
	jwtSecret := "未设置"
	if cfg.JWT.Secret != "" {
		jwtSecret = "***已设置***"
	}
	logger.Info("JWT配置",
		zap.String("密钥", jwtSecret),
		zap.Int("过期时间", cfg.JWT.ExpireTime))
}