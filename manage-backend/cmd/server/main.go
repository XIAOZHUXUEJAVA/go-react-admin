package main

import (
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/config"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/handler"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/middleware"
	casbinpkg "github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/casbin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/database"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	_ "github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/docs" // 导入生成的 docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title Go 管理系统启动器 API
// @version 1.0
// @description 基于 Go 和 Gin 构建的管理系统 API
// @termsOfService http://swagger.io/terms/

// @contact.name API 支持
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// 加载配置
	cfg := config.Load()

	// 初始化日志器
	logger.Init(logger.LogConfig{
		Level:      cfg.Log.Level,
		Format:     cfg.Log.Format,
		OutputPath: cfg.Log.OutputPath,
		MaxSize:    cfg.Log.MaxSize,
		MaxBackups: cfg.Log.MaxBackups,
		MaxAge:     cfg.Log.MaxAge,
		Compress:   cfg.Log.Compress,
	})

	// 记录配置详情
	config.LogConfigDetails(cfg)

	// 初始化数据库
	db, err := database.Init(cfg.Database)
	if err != nil {
		logger.Fatal("数据库初始化失败", zap.Error(err))
	}

	// 运行数据库迁移
	if err := database.RunMigrations(db, cfg); err != nil {
		logger.Fatal("数据库迁移失败", zap.Error(err))
	}

	// 如需种子数据，可以手动调用: database.SeedDatabase(db, cfg.Environment)
	logger.Info("✅ 数据库连接成功")

	// 初始化 Casbin Enforcer
	modelPath := "pkg/casbin/model.conf"
	enforcer, err := casbinpkg.NewEnforcer(db, modelPath)
	if err != nil {
		logger.Fatal("Casbin 初始化失败", zap.Error(err))
	}
	logger.Info("✅ Casbin 权限引擎初始化成功")
	
	// TODO: enforcer 将在后续的权限中间件中使用
	_ = enforcer

	// 根据环境初始化 Gin 路由器设置
	switch cfg.Environment {
	case "production":
		gin.SetMode(gin.ReleaseMode)
		logger.Info("🏭 运行在生产模式", zap.String("environment", cfg.Environment))
	case "test":
		gin.SetMode(gin.TestMode)
		logger.Info("🧪 运行在测试模式", zap.String("environment", cfg.Environment))
	default:
		gin.SetMode(gin.DebugMode)
		logger.Info("🔧 运行在开发模式", zap.String("environment", cfg.Environment))
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// API 路由
	api := router.Group("/api/v1")
	handler.SetupRoutes(api, db, enforcer)

	// Swagger 文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 启动服务器
	logger.Info("服务器正在启动", zap.String("port", cfg.Port))
	if err := router.Run(":" + cfg.Port); err != nil {
		logger.Fatal("服务器启动失败", zap.Error(err))
	}
}