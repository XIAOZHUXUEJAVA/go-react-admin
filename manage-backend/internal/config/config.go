package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Environment   string                `mapstructure:"environment"`
	Port          string                `mapstructure:"port"`
	LogLevel      string                `mapstructure:"log_level"`
	Log           LogConfig             `mapstructure:"log"`
	Database      Database              `mapstructure:"database"`
	Redis         Redis                 `mapstructure:"redis"`
	JWT           JWT                   `mapstructure:"jwt"`
	Captcha       CaptchaConfig         `mapstructure:"captcha"`
	Email         EmailConfig           `mapstructure:"email"`
	PasswordReset PasswordResetConfig   `mapstructure:"password_reset"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Schema   string `mapstructure:"schema"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWT struct {
	Secret               string `mapstructure:"secret"`
	ExpireTime           int    `mapstructure:"expire_time"`           // Deprecated: use AccessTokenExpire
	AccessTokenExpire    int    `mapstructure:"access_token_expire"`   // Access token expiration in minutes
	RefreshTokenExpire   int    `mapstructure:"refresh_token_expire"`  // Refresh token expiration in hours
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `mapstructure:"level"`       // 日志级别: debug, info, warn, error
	Format     string `mapstructure:"format"`      // 日志格式: json, console
	OutputPath string `mapstructure:"output_path"` // 日志文件路径
	MaxSize    int    `mapstructure:"max_size"`    // 单个日志文件最大大小(MB)
	MaxBackups int    `mapstructure:"max_backups"` // 保留的旧日志文件最大数量
	MaxAge     int    `mapstructure:"max_age"`     // 保留旧日志文件的最大天数
	Compress   bool   `mapstructure:"compress"`    // 是否压缩旧日志文件
}

// EmailConfig 邮件配置
type EmailConfig struct {
	SMTPHost    string `mapstructure:"smtp_host"`
	SMTPPort    int    `mapstructure:"smtp_port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	FromName    string `mapstructure:"from_name"`
	FromAddress string `mapstructure:"from_address"`
}

// PasswordResetConfig 密码重置配置
type PasswordResetConfig struct {
	TokenExpireMinutes int    `mapstructure:"token_expire_minutes"`
	FrontendURL        string `mapstructure:"frontend_url"`
}

func Load() *Config {
	// 首先启用从环境变量读取配置
	viper.AutomaticEnv()
	
	// 从操作系统环境变量检查环境类型
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development"
	}

	// 设置配置文件类型
	viper.SetConfigType("yaml")
	
	// 添加配置文件搜索路径（所有环境通用）
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")        // 用于从子目录运行时
	viper.AddConfigPath("../../config")     // 用于更深层嵌套调用（如测试）
	viper.AddConfigPath(".")                // 回退到当前目录

	// 步骤1：加载基础配置文件 (config.yaml)
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		// 在日志器初始化之前，使用简单的输出
		fmt.Printf("警告: 未找到基础配置文件，使用默认值: %v\n", err)
	} else {
		fmt.Printf("已加载基础配置: %s\n", viper.ConfigFileUsed())
	}

	// 步骤2：合并环境特定配置
	envConfigName := fmt.Sprintf("config.%s", environment)
	viper.SetConfigName(envConfigName)
	
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("警告: 未找到环境 '%s' 的配置文件: %v\n", environment, err)
	} else {
		fmt.Printf("已合并环境配置: %s\n", viper.ConfigFileUsed())
	}

	// 步骤3：使用环境变量覆盖配置
	// 将环境变量映射到配置键
	viper.BindEnv("port", "PORT")
	viper.BindEnv("log_level", "LOG_LEVEL")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.user", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.name", "DB_NAME")
	viper.BindEnv("database.schema", "DB_SCHEMA")
	viper.BindEnv("redis.host", "REDIS_HOST")
	viper.BindEnv("redis.port", "REDIS_PORT")
	viper.BindEnv("redis.password", "REDIS_PASSWORD")
	viper.BindEnv("redis.db", "REDIS_DB")
	viper.BindEnv("jwt.secret", "JWT_SECRET")
	viper.BindEnv("jwt.expire_time", "JWT_EXPIRE_TIME")
	viper.BindEnv("jwt.access_token_expire", "JWT_ACCESS_TOKEN_EXPIRE")
	viper.BindEnv("jwt.refresh_token_expire", "JWT_REFRESH_TOKEN_EXPIRE")

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		// 配置解码失败是致命错误，使用 panic
		panic(fmt.Sprintf("无法解码配置: %v", err))
	}

	// 在配置中设置环境类型
	config.Environment = environment

	return &config
}