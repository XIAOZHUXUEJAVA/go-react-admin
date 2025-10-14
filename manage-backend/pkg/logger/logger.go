package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

// LogConfig 日志配置
type LogConfig struct {
	Level      string
	Format     string
	OutputPath string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

// Init 初始化日志器，支持文件输出和日志轮转
func Init(cfg LogConfig) {
	// 解析日志级别
	var zapLevel zapcore.Level
	switch cfg.Level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	// 配置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		 EncodeTime:     zapcore.ISO8601TimeEncoder,
		 EncodeDuration: zapcore.SecondsDurationEncoder,
		 EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 根据格式选择编码器
	var encoder zapcore.Encoder
	if cfg.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		// console 格式使用彩色输出
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 配置输出目标
	var writeSyncer zapcore.WriteSyncer
	
	if cfg.OutputPath != "" && cfg.OutputPath != "stdout" {
		// 文件输出 + 日志轮转
		fileWriter := &lumberjack.Logger{
			Filename:   cfg.OutputPath,
			MaxSize:    cfg.MaxSize,    // MB
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,     // days
			Compress:   cfg.Compress,
		}
		
		// 同时输出到控制台和文件
		writeSyncer = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(fileWriter),
		)
	} else {
		// 仅输出到控制台
		writeSyncer = zapcore.AddSync(os.Stdout)
	}

	// 创建 core
	core := zapcore.NewCore(encoder, writeSyncer, zapLevel)

	// 创建 logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}