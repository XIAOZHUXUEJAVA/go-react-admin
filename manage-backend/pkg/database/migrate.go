package database

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/config"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
)

// RunMigrations 根据环境运行不同的迁移策略
func RunMigrations(db *gorm.DB, cfg *config.Config) error {
	logger.Info("开始运行数据库迁移", zap.String("environment", cfg.Environment))
	
	switch cfg.Environment {
	case "development", "test":
		// 开发和测试环境使用 AutoMigrate
		return autoMigrate(db, cfg)
	case "production":
		// 生产环境使用版本化迁移
		return runVersionedMigrations(db)
	default:
		return fmt.Errorf("unknown environment: %s", cfg.Environment)
	}
}

// autoMigrate 自动迁移所有模型
func autoMigrate(db *gorm.DB, cfg *config.Config) error {
	logger.Info("开始自动迁移")
	
	// 如果配置了非 public 模式，先创建模式
	if cfg.Database.Schema != "" && cfg.Database.Schema != "public" {
		createSchemaSQL := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", cfg.Database.Schema)
		if err := db.Exec(createSchemaSQL).Error; err != nil {
			logger.Warn("创建数据库模式失败", zap.String("schema", cfg.Database.Schema), zap.Error(err))
		} else {
			logger.Info("数据库模式创建成功或已存在", zap.String("schema", cfg.Database.Schema))
		}
	}
	
	err := db.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Permission{},
		&model.Menu{},
		&model.UserRole{},
		&model.RolePermission{},
		// 在这里添加其他模型
	)
	
	if err != nil {
		return fmt.Errorf("auto migration failed: %w", err)
	}
	
	logger.Info("自动迁移完成")
	
	// 开发环境也运行版本化迁移来执行初始数据插入
	logger.Info("执行初始数据迁移")
	return runVersionedMigrations(db)
}

// runVersionedMigrations 运行版本化迁移（生产环境）
func runVersionedMigrations(db *gorm.DB) error {
	logger.Info("开始运行版本化迁移")
	
	// 创建迁移记录表
	if err := db.AutoMigrate(&MigrationRecord{}); err != nil {
		return fmt.Errorf("failed to create migration table: %w", err)
	}
	
	// 运行所有未执行的迁移
	for _, migration := range migrations {
		var record MigrationRecord
		result := db.Where("migration_id = ?", migration.ID).First(&record)
		
		if result.Error == gorm.ErrRecordNotFound {
			// 执行迁移
			if err := migration.Up(db); err != nil {
				return fmt.Errorf("migration %s failed: %w", migration.ID, err)
			}
			
			// 记录迁移
			record = MigrationRecord{
				MigrationID: migration.ID,
				ExecutedAt:  time.Now(),
			}
			if err := db.Create(&record).Error; err != nil {
				return fmt.Errorf("failed to record migration %s: %w", migration.ID, err)
			}
			
			logger.Info("迁移执行成功", zap.String("migration_id", migration.ID))
		}
	}
	
	logger.Info("版本化迁移完成")
	return nil
}

// MigrationRecord 迁移记录模型
type MigrationRecord struct {
	ID          uint      `gorm:"primaryKey"`
	MigrationID string    `gorm:"uniqueIndex;not null"`
	ExecutedAt  time.Time `gorm:"not null"`
}