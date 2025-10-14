package casbin

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// NewEnforcer 创建并初始化 Casbin Enforcer
// 参数：
//   - db: GORM 数据库连接
//   - modelPath: Casbin 模型配置文件路径
// 返回：
//   - *casbin.Enforcer: Casbin 执行器实例
//   - error: 错误信息
func NewEnforcer(db *gorm.DB, modelPath string) (*casbin.Enforcer, error) {
	// 1. 初始化 GORM Adapter
	// 使用现有的数据库连接，策略将存储在 casbin_rule 表中
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create casbin adapter: %w", err)
	}

	// 2. 创建 Enforcer
	// 使用模型配置文件和适配器创建执行器
	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		return nil, fmt.Errorf("failed to create casbin enforcer: %w", err)
	}

	// 3. 加载策略
	// 从数据库加载所有策略到内存
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, fmt.Errorf("failed to load casbin policy: %w", err)
	}

	// 4. 启用自动保存
	// 当策略发生变更时，自动保存到数据库
	enforcer.EnableAutoSave(true)

	// 5. 启用日志（可选，用于调试）
	// enforcer.EnableLog(true)

	return enforcer, nil
}

// NewEnforcerWithAutoMigrate 创建 Enforcer 并自动迁移表结构
// 如果 casbin_rule 表不存在，会自动创建
func NewEnforcerWithAutoMigrate(db *gorm.DB, modelPath string) (*casbin.Enforcer, error) {
	// 自动迁移 casbin_rule 表
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create casbin adapter: %w", err)
	}

	enforcer, err := casbin.NewEnforcer(modelPath, adapter)
	if err != nil {
		return nil, fmt.Errorf("failed to create casbin enforcer: %w", err)
	}

	if err := enforcer.LoadPolicy(); err != nil {
		return nil, fmt.Errorf("failed to load casbin policy: %w", err)
	}

	enforcer.EnableAutoSave(true)

	return enforcer, nil
}

// ReloadPolicy 重新加载策略
// 用于在策略发生变更后手动刷新
func ReloadPolicy(enforcer *casbin.Enforcer) error {
	return enforcer.LoadPolicy()
}

// ClearPolicy 清除所有策略
// 谨慎使用，会删除所有权限配置
func ClearPolicy(enforcer *casbin.Enforcer) error {
	enforcer.ClearPolicy()
	return enforcer.SavePolicy()
}
