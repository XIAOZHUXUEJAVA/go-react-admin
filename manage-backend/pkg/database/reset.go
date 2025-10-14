package database

import (
	"fmt"

	"gorm.io/gorm"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
)

// ResetDatabase 重置数据库（删除所有表）
func ResetDatabase(db *gorm.DB) error {
	logger.Info("开始重置数据库")
	
	// 获取所有表名
	tables := []string{
		"migration_records",
		"users",
		// 在这里添加其他表名
	}
	
	// 删除所有表
	for _, table := range tables {
		if db.Migrator().HasTable(table) {
			if err := db.Migrator().DropTable(table); err != nil {
				logger.Warn("删除表失败", zap.String("table", table), zap.Error(err))
			} else {
				logger.Info("删除表成功", zap.String("table", table))
			}
		}
	}
	
	// 或者使用模型来删除表（更安全的方式）
	models := []interface{}{
		&MigrationRecord{},
		&model.User{},
		// 在这里添加其他模型
	}
	
	for _, model := range models {
		if err := db.Migrator().DropTable(model); err != nil {
			logger.Warn("删除模型表失败", zap.String("model", fmt.Sprintf("%T", model)), zap.Error(err))
		}
	}
	
	logger.Info("数据库重置完成")
	return nil
}