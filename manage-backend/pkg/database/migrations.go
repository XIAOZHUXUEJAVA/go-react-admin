package database

import (
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/gorm"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
)


// Migration 定义单个迁移的结构
// 每个迁移需要一个唯一的 ID，以及对应的 Up（执行迁移）和 Down（回滚迁移）方法
type Migration struct {
	ID   string
	Up   func(*gorm.DB) error
	Down func(*gorm.DB) error
}

// executeSQLFile 读取并执行 SQL 文件
func executeSQLFile(db *gorm.DB, filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file %s: %w", filePath, err)
	}

	// 执行 SQL 内容
	if err := db.Exec(string(content)).Error; err != nil {
		return fmt.Errorf("failed to execute SQL file %s: %w", filePath, err)
	}

	return nil
}

// migrations 迁移列表
// 在这里定义所有需要的迁移，按顺序执行
var migrations = []Migration{
	{
		ID: "001_create_users_table",
		Up: func(db *gorm.DB) error {
			// 执行 SQL 文件
			sqlFile := filepath.Join("migrations", "000001_create_users_table.up.sql")
			return executeSQLFile(db, sqlFile)
		},
		Down: func(db *gorm.DB) error {
			// 执行回滚 SQL 文件
			sqlFile := filepath.Join("migrations", "000001_create_users_table.down.sql")
			return executeSQLFile(db, sqlFile)
		},
	},
	{
		ID: "002_create_rbac_tables",
		Up: func(db *gorm.DB) error {
			// 执行 RBAC 表创建和初始化
			sqlFile := filepath.Join("migrations", "000002_create_rbac_tables.up.sql")
			return executeSQLFile(db, sqlFile)
		},
		Down: func(db *gorm.DB) error {
			// 执行 RBAC 表回滚
			sqlFile := filepath.Join("migrations", "000002_create_rbac_tables.down.sql")
			return executeSQLFile(db, sqlFile)
		},
	},
	{
		ID: "003_create_audit_logs_table",
		Up: func(db *gorm.DB) error {
			// 执行审计日志表创建
			sqlFile := filepath.Join("migrations", "000003_create_audit_logs_table.up.sql")
			return executeSQLFile(db, sqlFile)
		},
		Down: func(db *gorm.DB) error {
			// 执行审计日志表回滚
			sqlFile := filepath.Join("migrations", "000003_create_audit_logs_table.down.sql")
			return executeSQLFile(db, sqlFile)
		},
	},
	{
		ID: "004_create_dict_tables",
		Up: func(db *gorm.DB) error {
			// 执行字典管理表创建和初始化
			sqlFile := filepath.Join("migrations", "000004_create_dict_tables.up.sql")
			return executeSQLFile(db, sqlFile)
		},
		Down: func(db *gorm.DB) error {
			// 执行字典管理表回滚
			sqlFile := filepath.Join("migrations", "000004_create_dict_tables.down.sql")
			return executeSQLFile(db, sqlFile)
		},
	},
	// 在这里继续追加其他迁移
}


// RollbackMigration 回滚指定的迁移
// - 按照 migrationID 找到对应迁移
// - 执行 Down 回滚逻辑
// - 从 migration_records 表中删除迁移记录
// - 输出日志
func RollbackMigration(db *gorm.DB, migrationID string) error {
	for _, migration := range migrations {
		if migration.ID == migrationID {
			// 执行 Down 回滚
			if err := migration.Down(db); err != nil {
				return fmt.Errorf("rollback migration %s failed: %w", migrationID, err)
			}

			// 删除迁移记录
			if err := db.Where("migration_id = ?", migrationID).Delete(&MigrationRecord{}).Error; err != nil {
				return fmt.Errorf("failed to remove migration record %s: %w", migrationID, err)
			}

			// 记录日志
			logger.Info("迁移回滚成功", zap.String("migration_id", migrationID))
			return nil
		}
	}
	// 如果找不到对应 ID
	return fmt.Errorf("migration %s not found", migrationID)
}


// GetMigrationStatus 获取所有迁移的执行状态
// 逻辑：
// 1. 检查 migration_records 表是否存在：
//    - 如果不存在：认为是开发模式，尝试用实际表的存在情况来判断
//    - 如果存在：从 migration_records 读取已执行的迁移
// 2. 返回每个迁移的状态，包括是否执行过、执行时间
func GetMigrationStatus(db *gorm.DB) ([]MigrationStatus, error) {
	// 检查 migration_records 表是否存在
	if !db.Migrator().HasTable(&MigrationRecord{}) {
		// 表不存在，使用实际表情况来推测状态
		var status []MigrationStatus
		for _, migration := range migrations {
			s := MigrationStatus{
				ID:       migration.ID,
				Executed: false,
			}

			// 特殊处理：001_create_users_table
			if migration.ID == "001_create_users_table" {
				// 判断 users 表是否存在
				s.Executed = db.Migrator().HasTable("users")
				if s.Executed {
					// 没有 migration_records 表，只能标记为自动迁移
					s.ExecutedAt = "Auto-migrated (development mode)"
				}
			}

			status = append(status, s)
		}
		return status, nil
	}

	// migration_records 存在，直接读取
	var records []MigrationRecord
	if err := db.Find(&records).Error; err != nil {
		return nil, fmt.Errorf("failed to get migration records: %w", err)
	}

	// 转换成 map，方便查找
	recordMap := make(map[string]MigrationRecord)
	for _, record := range records {
		recordMap[record.MigrationID] = record
	}

	// 遍历所有迁移，生成状态
	var status []MigrationStatus
	for _, migration := range migrations {
		s := MigrationStatus{
			ID:       migration.ID,
			Executed: false,
		}

		// 如果有记录，则迁移已执行
		if record, exists := recordMap[migration.ID]; exists {
			s.Executed = true
			s.ExecutedAt = record.ExecutedAt
		}

		status = append(status, s)
	}

	return status, nil
}


// MigrationStatus 表示迁移状态
// - ID: 迁移 ID
// - Executed: 是否已执行
// - ExecutedAt: 执行时间（可能是时间戳、字符串说明）
type MigrationStatus struct {
	ID         string
	Executed   bool
	ExecutedAt interface{}
}
