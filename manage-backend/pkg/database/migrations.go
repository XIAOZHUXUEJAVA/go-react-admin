package database

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
// 
// ⚠️ 重要提示：数据库本身（go_manage_starter）必须在运行迁移前已经存在
// 
// 数据库初始化方式：
// 
// 【方式一】Docker Compose 启动（推荐，自动化）
//   - 执行: docker-start.sh 或者 docker-start.bat脚本
//   - 自动流程:
//     1. 创建数据库 go_manage_starter
//     2. 执行 scripts/01-init-db.sh 创建 schema
//     3. 执行 scripts/manage_dev.sql 初始化表和数据
//   - 优点: 无需手动操作，一键完成所有初始化
//   - 注意: 使用此方式时，main.go 中的 RunMigrations 应保持注释状态
// 
// 【方式二】本地开发环境（手动）
//   步骤 1: 创建数据库
//     - 方式 A: psql -U postgres -h localhost -f scripts/setup-dev-db.sql
//     - 方式 B: 手动执行 SQL: CREATE DATABASE go_manage_starter;
//   步骤 2: 运行迁移
//     - 方式 A: 取消 main.go 中 RunMigrations 的注释并运行程序
//     - 方式 B: 使用 Makefile: make migrate
//   - 优点: 适合不使用 Docker 的本地开发
//   - 注意: 需要确保 PostgreSQL 服务已启动
var migrations = []Migration{
	{
		ID: "001_create_schemas",
		Up: func(db *gorm.DB) error {
			// 创建 schemas（假设数据库已存在）
			schemas := []string{"manage_dev", "manage_test", "manage_prod"}
			for _, schema := range schemas {
				if err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema)).Error; err != nil {
					return fmt.Errorf("failed to create schema %s: %w", schema, err)
				}
			}
			
			// 设置默认搜索路径
			if err := db.Exec("SET search_path TO manage_dev, public").Error; err != nil {
				return fmt.Errorf("failed to set search_path: %w", err)
			}
			
			logger.Info("Schema 创建成功")
			return nil
		},
		Down: func(db *gorm.DB) error {
			// 删除所有 schema（CASCADE 会自动删除 schema 下的所有对象）
			schemas := []string{"manage_dev", "manage_test", "manage_prod"}
			for _, schema := range schemas {
				if err := db.Exec(fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE", schema)).Error; err != nil {
					return fmt.Errorf("failed to drop schema %s: %w", schema, err)
				}
			}
			
			logger.Info("Schema 删除成功")
			return nil
		},
	},
	{
		ID: "002_create_tables",
		Up: func(db *gorm.DB) error {
			// 执行完整的数据库表初始化 SQL 文件
			sqlFile := filepath.Join("scripts", "manage_dev.sql")
			if err := executeSQLFile(db, sqlFile); err != nil {
				return err
			}
			
			logger.Info("数据库表创建成功")
			return nil
		},
		Down: func(db *gorm.DB) error {
			// 回滚：删除所有表
			tables := []string{
				"audit_logs",
				"casbin_rule",
				"dict_items",
				"dict_types",
				"menus",
				"migration_records",
				"password_reset_tokens",
				"permissions",
				"role_permissions",
				"roles",
				"user_roles",
				"users",
			}

			for _, table := range tables {
				if err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS \"manage_dev\".\"%s\" CASCADE", table)).Error; err != nil {
					return fmt.Errorf("failed to drop table %s: %w", table, err)
				}
			}

			// 删除序列
			sequences := []string{
				"audit_logs_id_seq",
				"casbin_rule_id_seq",
				"dict_items_id_seq",
				"dict_types_id_seq",
				"menus_id_seq",
				"migration_records_id_seq",
				"password_reset_tokens_id_seq",
				"permissions_id_seq",
				"role_permissions_id_seq",
				"roles_id_seq",
				"user_roles_id_seq",
				"users_id_seq",
			}

			for _, seq := range sequences {
				if err := db.Exec(fmt.Sprintf("DROP SEQUENCE IF EXISTS \"manage_dev\".\"%s\"", seq)).Error; err != nil {
					return fmt.Errorf("failed to drop sequence %s: %w", seq, err)
				}
			}

			logger.Info("数据库表删除成功")
			return nil
		},
	},
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

			// 特殊处理：检查迁移是否已执行
			switch migration.ID {
			case "001_create_schemas":
				// 检查 schema 是否存在
				var schemaExists bool
				err := db.Raw("SELECT EXISTS(SELECT 1 FROM information_schema.schemata WHERE schema_name = 'manage_dev')").Scan(&schemaExists).Error
				if err == nil && schemaExists {
					s.Executed = true
					s.ExecutedAt = "Auto-migrated (development mode)"
				}
			case "002_create_tables":
				// 检查 users 表是否存在
				s.Executed = db.Migrator().HasTable("users")
				if s.Executed {
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
