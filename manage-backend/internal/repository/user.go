package repository

import (
	"fmt"

	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"gorm.io/gorm"
)

// UserRepository 用户数据仓库
// 封装对 User 模型的所有数据库操作
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建 UserRepository 实例
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create 新增用户
// 参数: user - 用户对象
// 返回: error - 操作是否成功
func (r *UserRepository) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return fmt.Errorf("创建用户失败 [username=%s]: %w", user.Username, err)
	}
	return nil
}

// GetByID 根据 ID 获取用户
// 参数: id - 用户ID
// 返回: *model.User - 用户对象, error - 查询是否成功
func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("根据ID查询用户失败 [id=%d]: %w", id, err)
	}
	return &user, nil
}

// GetByUsername 根据用户名获取用户
// 参数: username - 用户名
// 返回: *model.User - 用户对象, error - 查询是否成功
func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, fmt.Errorf("根据用户名查询用户失败 [username=%s]: %w", username, err)
	}
	return &user, nil
}

// GetByEmail 根据邮箱获取用户
// 参数: email - 邮箱地址
// 返回: *model.User - 用户对象, error - 查询是否成功
func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("根据邮箱查询用户失败 [email=%s]: %w", email, err)
	}
	return &user, nil
}

// Update 更新用户信息
// 参数: user - 用户对象（需包含ID）
// 返回: error - 操作是否成功
func (r *UserRepository) Update(user *model.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return fmt.Errorf("更新用户失败 [id=%d, username=%s]: %w", user.ID, user.Username, err)
	}
	return nil
}

// Delete 删除用户
// 参数: id - 用户ID
// 返回: error - 操作是否成功
func (r *UserRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.User{}, id).Error; err != nil {
		return fmt.Errorf("删除用户失败 [id=%d]: %w", id, err)
	}
	return nil
}

// List 分页获取用户列表
// 参数: offset - 偏移量, limit - 每页数量
// 返回: []model.User - 用户列表, int64 - 总记录数, error - 查询是否成功
func (r *UserRepository) List(offset, limit int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	// 获取总数
	if err := r.db.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("统计用户总数失败: %w", err)
	}

	// 分页查询
	if err := r.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("分页查询用户列表失败 [offset=%d, limit=%d]: %w", offset, limit, err)
	}
	return users, total, nil
}

// ListByVisibility 根据用户角色和可见性规则分页获取用户列表
// 参数: currentUserID - 当前用户ID, currentUserRole - 当前用户角色, offset - 偏移量, limit - 每页数量
// 返回: []model.User - 用户列表, int64 - 总记录数, error - 查询是否成功
func (r *UserRepository) ListByVisibility(currentUserID uint, currentUserRole string, offset, limit int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := r.db.Model(&model.User{})

	// 超级管理员可以看到所有用户
	if currentUserRole == "admin" {
		// 不添加任何过滤条件
	} else if currentUserRole == "manager" {
		// 管理员只能看到：
		// 1. 自己 (id = currentUserID)
		// 2. 自己创建的用户 (created_by = currentUserID)
		// 3. 排除超级管理员和其他管理员（但不排除自己）
		query = query.Where(
			r.db.Where("id = ?", currentUserID).
				Or(r.db.Where("created_by = ?", currentUserID).Where("role NOT IN (?)", []string{"manager", "admin"})),
		)
	} else {
		// 普通用户只能看到自己
		query = query.Where("id = ?", currentUserID)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("统计可见用户总数失败 [userID=%d, role=%s]: %w", currentUserID, currentUserRole, err)
	}

	// 分页查询
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("分页查询可见用户列表失败 [userID=%d, role=%s, offset=%d, limit=%d]: %w", currentUserID, currentUserRole, offset, limit, err)
	}
	return users, total, nil
}

// CheckUsernameExists 检查用户名是否已存在
// 参数: username - 用户名
// 返回: bool - 是否存在, error - 查询是否成功
func (r *UserRepository) CheckUsernameExists(username string) (bool, error) {
	var count int64
	if err := r.db.Model(&model.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, fmt.Errorf("检查用户名是否存在失败 [username=%s]: %w", username, err)
	}
	return count > 0, nil
}

// CheckEmailExists 检查邮箱是否已存在
// 参数: email - 邮箱地址
// 返回: bool - 是否存在, error - 查询是否成功
func (r *UserRepository) CheckEmailExists(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, fmt.Errorf("检查邮箱是否存在失败 [email=%s]: %w", email, err)
	}
	return count > 0, nil
}

// CheckUsernameExistsExcludeID 检查用户名是否已存在（排除指定ID）
// 参数: username - 用户名, excludeID - 排除的用户ID（用于更新时排除自己）
// 返回: bool - 是否存在, error - 查询是否成功
func (r *UserRepository) CheckUsernameExistsExcludeID(username string, excludeID uint) (bool, error) {
	var count int64
	if err := r.db.Model(&model.User{}).Where("username = ? AND id != ?", username, excludeID).Count(&count).Error; err != nil {
		return false, fmt.Errorf("检查用户名是否存在失败 [username=%s, excludeID=%d]: %w", username, excludeID, err)
	}
	return count > 0, nil
}

// CheckEmailExistsExcludeID 检查邮箱是否已存在（排除指定ID）
// 参数: email - 邮箱地址, excludeID - 排除的用户ID（用于更新时排除自己）
// 返回: bool - 是否存在, error - 查询是否成功
func (r *UserRepository) CheckEmailExistsExcludeID(email string, excludeID uint) (bool, error) {
	var count int64
	if err := r.db.Model(&model.User{}).Where("email = ? AND id != ?", email, excludeID).Count(&count).Error; err != nil {
		return false, fmt.Errorf("检查邮箱是否存在失败 [email=%s, excludeID=%d]: %w", email, excludeID, err)
	}
	return count > 0, nil
}

// UpdatePassword 更新用户密码
// 参数: userID - 用户ID, hashedPassword - 加密后的密码
// 返回: error - 操作是否成功
func (r *UserRepository) UpdatePassword(userID uint, hashedPassword string) error {
	if err := r.db.Model(&model.User{}).
		Where("id = ?", userID).
		Update("password", hashedPassword).Error; err != nil {
		return fmt.Errorf("更新用户密码失败 [userID=%d]: %w", userID, err)
	}
	return nil
}
