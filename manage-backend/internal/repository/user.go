package repository

import (
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
	return r.db.Create(user).Error
}

// GetByID 根据 ID 获取用户
// 参数: id - 用户ID
// 返回: *model.User - 用户对象, error - 查询是否成功
func (r *UserRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername 根据用户名获取用户
// 参数: username - 用户名
// 返回: *model.User - 用户对象, error - 查询是否成功
func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail 根据邮箱获取用户
// 参数: email - 邮箱地址
// 返回: *model.User - 用户对象, error - 查询是否成功
func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update 更新用户信息
// 参数: user - 用户对象（需包含ID）
// 返回: error - 操作是否成功
func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete 删除用户
// 参数: id - 用户ID
// 返回: error - 操作是否成功
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

// List 分页获取用户列表
// 参数: offset - 偏移量, limit - 每页数量
// 返回: []model.User - 用户列表, int64 - 总记录数, error - 查询是否成功
func (r *UserRepository) List(offset, limit int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	// 获取总数
	err := r.db.Model(&model.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = r.db.Offset(offset).Limit(limit).Find(&users).Error
	return users, total, err
}

// CheckUsernameExists 检查用户名是否已存在
// 参数: username - 用户名
// 返回: bool - 是否存在, error - 查询是否成功
func (r *UserRepository) CheckUsernameExists(username string) (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

// CheckEmailExists 检查邮箱是否已存在
// 参数: email - 邮箱地址
// 返回: bool - 是否存在, error - 查询是否成功
func (r *UserRepository) CheckEmailExists(email string) (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

// CheckUsernameExistsExcludeID 检查用户名是否已存在（排除指定ID）
// 参数: username - 用户名, excludeID - 排除的用户ID（用于更新时排除自己）
// 返回: bool - 是否存在, error - 查询是否成功
func (r *UserRepository) CheckUsernameExistsExcludeID(username string, excludeID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).Where("username = ? AND id != ?", username, excludeID).Count(&count).Error
	return count > 0, err
}

// CheckEmailExistsExcludeID 检查邮箱是否已存在（排除指定ID）
// 参数: email - 邮箱地址, excludeID - 排除的用户ID（用于更新时排除自己）
// 返回: bool - 是否存在, error - 查询是否成功
func (r *UserRepository) CheckEmailExistsExcludeID(email string, excludeID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).Where("email = ? AND id != ?", email, excludeID).Count(&count).Error
	return count > 0, err
}
