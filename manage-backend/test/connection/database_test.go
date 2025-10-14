package connection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TestDatabaseConnection 测试数据库连接
func TestDatabaseConnection(t *testing.T) {
	config := struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		Schema   string
		SSLMode  string
	}{
		Host:     "localhost",
		Port:     5432,
		User:     "xiaozhu",
		Password: "12345679",
		DBName:   "go_manage_starter",
		Schema:   "manage_dev",
		SSLMode:  "disable",
	}

	t.Logf("🔌 测试数据库连接: %s@%s:%d/%s", config.User, config.Host, config.Port, config.DBName)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.SSLMode,
		config.Schema,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err, "❌ 数据库连接失败")

	sqlDB, err := db.DB()
	require.NoError(t, err, "❌ 获取数据库实例失败")
	defer sqlDB.Close()

	require.NoError(t, sqlDB.Ping(), "❌ 数据库 Ping 失败")

	var version string
	require.NoError(t, db.Raw("SELECT version()").Scan(&version).Error, "❌ 查询 PostgreSQL 版本失败")

	var schemaExists bool
	query := "SELECT EXISTS(SELECT 1 FROM information_schema.schemata WHERE schema_name = ?)"
	require.NoError(t, db.Raw(query, config.Schema).Scan(&schemaExists).Error, "❌ Schema 检查失败")

	assert.True(t, schemaExists, "❌ 期望 Schema %s 存在", config.Schema)

	t.Logf("✅ 数据库连接成功!")
	t.Logf("📊 PostgreSQL 版本: %s", version)
	t.Logf("📁 Schema '%s' 存在: %v", config.Schema, schemaExists)

	stats := sqlDB.Stats()
	t.Logf("🔗 连接统计:")
	t.Logf("   - 打开连接数: %d", stats.OpenConnections)
	t.Logf("   - 使用中连接数: %d", stats.InUse)
	t.Logf("   - 空闲连接数: %d", stats.Idle)
}

// TestDatabaseConnectionWithWrongCredentials 测试错误凭据
func TestDatabaseConnectionWithWrongCredentials(t *testing.T) {
	t.Log("🔌 测试错误的数据库凭据...")

	dsn := "host=localhost user=xiaozhu password=wrong_password dbname=go_manage_starter_dev port=5432 sslmode=disable"

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.Error(t, err, "❌ 预期连接失败，但连接成功了")

	t.Logf("✅ 错误凭据测试通过: %v", err)
}

// TestDatabaseConnectionWithWrongHost 测试错误主机
func TestDatabaseConnectionWithWrongHost(t *testing.T) {
	t.Log("🔌 测试错误的数据库主机...")

	dsn := "host=nonexistent-host user=xiaozhu password=12345679 dbname=go_manage_starter_dev port=5432 sslmode=disable"

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.Error(t, err, "❌ 预期连接失败，但连接成功了")

	t.Logf("✅ 错误主机测试通过: %v", err)
}
