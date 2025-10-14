package repository

import (
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"gorm.io/gorm"
)

// RoleRepository 角色数据仓库
// 封装对 Role 模型的所有数据库操作
type RoleRepository struct {
	db *gorm.DB
}

// NewRoleRepository 创建 RoleRepository 实例
func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

// Create 创建角色
func (r *RoleRepository) Create(role *model.Role) error {
	return r.db.Create(role).Error
}

// GetByID 根据 ID 获取角色
func (r *RoleRepository) GetByID(id uint) (*model.Role, error) {
	var role model.Role
	err := r.db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// GetByCode 根据角色代码获取角色
func (r *RoleRepository) GetByCode(code string) (*model.Role, error) {
	var role model.Role
	err := r.db.Where("code = ?", code).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// Update 更新角色信息
func (r *RoleRepository) Update(role *model.Role) error {
	return r.db.Save(role).Error
}

// Delete 删除角色
func (r *RoleRepository) Delete(id uint) error {
	return r.db.Delete(&model.Role{}, id).Error
}

// List 分页获取角色列表
func (r *RoleRepository) List(offset, limit int) ([]model.Role, int64, error) {
	var roles []model.Role
	var total int64

	// 获取总数
	err := r.db.Model(&model.Role{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = r.db.Offset(offset).Limit(limit).Order("created_at DESC").Find(&roles).Error
	return roles, total, err
}

// GetAll 获取所有角色
func (r *RoleRepository) GetAll() ([]model.Role, error) {
	var roles []model.Role
	err := r.db.Order("created_at DESC").Find(&roles).Error
	return roles, err
}

// CheckCodeExists 检查角色代码是否已存在
func (r *RoleRepository) CheckCodeExists(code string) (bool, error) {
	var count int64
	err := r.db.Model(&model.Role{}).Where("code = ?", code).Count(&count).Error
	return count > 0, err
}

// CheckCodeExistsExcludeID 检查角色代码是否已存在（排除指定ID）
func (r *RoleRepository) CheckCodeExistsExcludeID(code string, excludeID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.Role{}).Where("code = ? AND id != ?", code, excludeID).Count(&count).Error
	return count > 0, err
}

// GetUserRoles 获取用户的所有角色
func (r *RoleRepository) GetUserRoles(userID uint) ([]model.Role, error) {
	var roles []model.Role
	err := r.db.
		Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}

// AssignRoleToUser 为用户分配角色
func (r *RoleRepository) AssignRoleToUser(userID, roleID uint, assignedBy uint) error {
	userRole := model.UserRole{
		UserID:     userID,
		RoleID:     roleID,
		AssignedBy: assignedBy,
	}
	return r.db.Create(&userRole).Error
}

// RemoveRoleFromUser 移除用户的角色
func (r *RoleRepository) RemoveRoleFromUser(userID, roleID uint) error {
	return r.db.Where("user_id = ? AND role_id = ?", userID, roleID).Delete(&model.UserRole{}).Error
}

// RemoveAllRolesFromUser 移除用户的所有角色
func (r *RoleRepository) RemoveAllRolesFromUser(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.UserRole{}).Error
}

// GetUsersByRole 获取拥有指定角色的所有用户ID
func (r *RoleRepository) GetUsersByRole(roleID uint) ([]uint, error) {
	var userIDs []uint
	err := r.db.Model(&model.UserRole{}).
		Where("role_id = ?", roleID).
		Pluck("user_id", &userIDs).Error
	return userIDs, err
}

// UpdateRolePermissions 更新角色的权限关联
func (r *RoleRepository) UpdateRolePermissions(roleID uint, permissionIDs []uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 删除该角色的所有权限关联
		if err := tx.Where("role_id = ?", roleID).Delete(&model.RolePermission{}).Error; err != nil {
			return err
		}

		// 如果没有新权限，直接返回
		if len(permissionIDs) == 0 {
			return nil
		}

		// 批量插入新的权限关联
		rolePermissions := make([]model.RolePermission, 0, len(permissionIDs))
		for _, permID := range permissionIDs {
			rolePermissions = append(rolePermissions, model.RolePermission{
				RoleID:       roleID,
				PermissionID: permID,
			})
		}

		return tx.Create(&rolePermissions).Error
	})
}

// GetRolePermissionIDs 获取角色的所有权限ID
func (r *RoleRepository) GetRolePermissionIDs(roleID uint) ([]uint, error) {
	var permissionIDs []uint
	err := r.db.Model(&model.RolePermission{}).
		Where("role_id = ?", roleID).
		Pluck("permission_id", &permissionIDs).Error
	return permissionIDs, err
}
