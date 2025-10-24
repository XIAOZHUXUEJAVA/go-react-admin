package handler

import (
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/config"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/middleware"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/repository"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/service"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/auth"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/cache"
	casbinpkg "github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/casbin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.RouterGroup, db *gorm.DB, enforcer *casbin.Enforcer) {
	cfg := config.Load()
	
	// 使用配置初始化 JWT 管理器
	accessTokenExpire := cfg.JWT.AccessTokenExpire
	refreshTokenExpire := cfg.JWT.RefreshTokenExpire
	
	// 如果未配置则回退到默认值
	if accessTokenExpire == 0 {
		accessTokenExpire = 30 // 默认 30 分钟
	}
	if refreshTokenExpire == 0 {
		refreshTokenExpire = 720 // 默认 30 天（小时数）
	}
	
	jwtManager := auth.NewJWTManager(cfg.JWT.Secret, accessTokenExpire, refreshTokenExpire)

	// 初始化 Redis 客户端
	redisClient := cache.NewRedisClient(cfg.Redis)

	// 初始化仓储层
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	auditLogRepo := repository.NewAuditLogRepository(db)
	dictTypeRepo := repository.NewDictTypeRepository(db)
	dictItemRepo := repository.NewDictItemRepository(db)
	passwordResetRepo := repository.NewPasswordResetRepository(db)

	// 初始化服务层
	sessionService := service.NewSessionService(redisClient, jwtManager)
	casbinService := service.NewCasbinService(enforcer)
	roleService := service.NewRoleService(roleRepo, permissionRepo, casbinService)
	permissionService := service.NewPermissionService(permissionRepo)
	menuService := service.NewMenuService(menuRepo, permissionRepo)
	auditLogService := service.NewAuditLogService(auditLogRepo)
	dictTypeService := service.NewDictTypeService(dictTypeRepo, dictItemRepo)
	dictItemService := service.NewDictItemService(dictTypeRepo, dictItemRepo)
	emailService := service.NewEmailService(cfg)
	passwordResetRedisService := service.NewPasswordResetRedisService(redisClient)
	passwordResetService := service.NewPasswordResetService(cfg, userRepo, passwordResetRepo, emailService, auditLogService, passwordResetRedisService)
	
	// 验证码配置
	captchaConfig := service.CaptchaConfig{
		Type:            cfg.Captcha.Type,
		Length:          cfg.Captcha.Length,
		Width:           cfg.Captcha.Width,
		Height:          cfg.Captcha.Height,
		NoiseCount:      cfg.Captcha.NoiseCount,
		ShowLineOptions: cfg.Captcha.ShowLineOptions,
		Expiration:      cfg.Captcha.Expiration,
		Enabled:         cfg.Captcha.Enabled,
	}
	

	// 如果配置为空，使用默认配置
	if captchaConfig.Type == "" {
		captchaConfig = service.CaptchaConfig{
			Type:            "digit",
			Length:          5,
			Width:           240,
			Height:          80,
			NoiseCount:      0.7,
			ShowLineOptions: 80,
			Expiration:      5 * time.Minute,
			Enabled:         true,
		}
	}
	
	captchaService := service.NewCaptchaService(redisClient.GetClient(), captchaConfig)
	loginRateLimitService := service.NewLoginRateLimitService(redisClient)
	userService := service.NewUserService(userRepo, jwtManager, sessionService, captchaService, roleRepo, permissionService, loginRateLimitService)

	// 初始化处理器
	userHandler := NewUserHandler(userService)
	captchaHandler := NewCaptchaHandler(captchaService)
	roleHandler := NewRoleHandler(roleService)
	permissionHandler := NewPermissionHandler(permissionService)
	menuHandler := NewMenuHandler(menuService, roleRepo)
	auditLogHandler := NewAuditLogHandler(auditLogService)
	dictTypeHandler := NewDictTypeHandler(dictTypeService)
	dictItemHandler := NewDictItemHandler(dictItemService)
	passwordResetHandler := NewPasswordResetHandler(passwordResetService)

	// 审计日志中间件配置
	auditConfig := middleware.DefaultAuditLogConfig()
	router.Use(middleware.AuditLogger(db, auditConfig))


	// 用户可用性检查路由（无需认证）
	userCheck := router.Group("/users")
	{
		userCheck.GET("/check-username/:username", userHandler.CheckUsernameAvailable)
		userCheck.GET("/check-email/:email", userHandler.CheckEmailAvailable)
		userCheck.POST("/check-availability", userHandler.CheckUserDataAvailability)
	}

	// 认证路由（无需认证）
	authRoutes := router.Group("/auth")
	{
		authRoutes.GET("/captcha", captchaHandler.GenerateCaptcha)
		authRoutes.POST("/register", userHandler.Register)
		authRoutes.POST("/login", userHandler.Login)
		authRoutes.POST("/refresh", userHandler.RefreshToken)
		authRoutes.POST("/forgot-password", passwordResetHandler.ForgotPassword)
		authRoutes.POST("/verify-reset-token", passwordResetHandler.VerifyResetToken)
		authRoutes.POST("/reset-password", passwordResetHandler.ResetPassword)
	}

	// 受保护的路由（需要认证）
	protected := router.Group("/")
	protected.Use(middleware.JWTAuthWithSession(jwtManager, sessionService))
	{
		// 需要认证的认证路由
		authProtected := protected.Group("/auth")
		{
			authProtected.POST("/logout", userHandler.Logout)
		}

		// 用户路由
		users := protected.Group("/users")
		{
			users.GET("/profile", userHandler.GetProfile)
			users.PUT("/profile", userHandler.UpdateProfile)
			users.GET("/permissions", userHandler.GetUserPermissions)
			users.GET("", userHandler.ListUsers)
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
			
			// 用户角色管理
			users.GET("/:id/roles", roleHandler.GetUserRoles)
			users.PUT("/:id/roles", roleHandler.AssignRolesToUser)
		}

		// 角色路由
		roles := protected.Group("/roles")
		{
			roles.GET("", roleHandler.ListRoles)
			roles.GET("/all", roleHandler.GetAllRoles)
			roles.POST("", roleHandler.CreateRole)
			roles.GET("/:id", roleHandler.GetRole)
			roles.PUT("/:id", roleHandler.UpdateRole)
			roles.DELETE("/:id", roleHandler.DeleteRole)
			roles.GET("/:id/permissions", roleHandler.GetRolePermissions)
			roles.PUT("/:id/permissions", roleHandler.AssignPermissions)
		}

		// 权限路由
		permissions := protected.Group("/permissions")
		{
			permissions.GET("", permissionHandler.ListPermissions)
			permissions.GET("/all", permissionHandler.GetAllPermissions)
			permissions.GET("/tree", permissionHandler.GetPermissionTree)
			permissions.POST("", permissionHandler.CreatePermission)
			permissions.GET("/:id", permissionHandler.GetPermission)
			permissions.PUT("/:id", permissionHandler.UpdatePermission)
			permissions.DELETE("/:id", permissionHandler.DeletePermission)
			permissions.GET("/resource/:resource", permissionHandler.GetPermissionsByResource)
			permissions.GET("/type/:type", permissionHandler.GetPermissionsByType)
		}

		// 菜单路由
		menus := protected.Group("/menus")
		{
			menus.GET("/tree", menuHandler.GetMenuTree)
			menus.GET("/tree/visible", menuHandler.GetVisibleMenuTree)
			menus.GET("/user", menuHandler.GetUserMenuTree)
			menus.POST("", menuHandler.CreateMenu)
			menus.GET("/:id", menuHandler.GetMenu)
			menus.PUT("/:id", menuHandler.UpdateMenu)
			menus.PUT("/order", menuHandler.UpdateMenuOrder)
			menus.DELETE("/:id", menuHandler.DeleteMenu)
		}

		// 审计日志路由
		auditLogs := protected.Group("/audit-logs")
		{
			auditLogs.GET("", auditLogHandler.QueryAuditLogs)
			auditLogs.GET("/:id", auditLogHandler.GetAuditLog)
			auditLogs.POST("/clean", auditLogHandler.CleanOldAuditLogs)
		}

		// 字典类型路由
		dictTypes := protected.Group("/dict-types")
		{
			dictTypes.GET("", dictTypeHandler.ListDictTypes)
			dictTypes.GET("/all", dictTypeHandler.GetAllDictTypes)
			dictTypes.POST("", dictTypeHandler.CreateDictType)
			dictTypes.GET("/:id", dictTypeHandler.GetDictType)
			dictTypes.PUT("/:id", dictTypeHandler.UpdateDictType)
			dictTypes.DELETE("/:id", dictTypeHandler.DeleteDictType)
		}

		// 字典项路由
		dictItems := protected.Group("/dict-items")
		{
			dictItems.GET("", dictItemHandler.ListDictItems)
			dictItems.GET("/by-type/:code", dictItemHandler.GetDictItemsByType)
			dictItems.POST("", dictItemHandler.CreateDictItem)
			dictItems.GET("/:id", dictItemHandler.GetDictItem)
			dictItems.PUT("/:id", dictItemHandler.UpdateDictItem)
			dictItems.DELETE("/:id", dictItemHandler.DeleteDictItem)
		}
	}
}

// 未使用的变量，避免编译错误
var _ = casbinpkg.NewEnforcer
