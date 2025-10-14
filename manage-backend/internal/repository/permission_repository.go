package repository

import (
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"gorm.io/gorm"
)

// PermissionRepository 权限数据仓库
// 封装对 Permission 模型的所有数据库操作
type PermissionRepository struct {
	db *gorm.DB
}

// NewPermissionRepository 创建 PermissionRepository 实例
func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{db: db}
}

// Create 创建权限
func (r *PermissionRepository) Create(permission *model.Permission) error {
	return r.db.Create(permission).Error
}

// GetByID 根据 ID 获取权限
func (r *PermissionRepository) GetByID(id uint) (*model.Permission, error) {
	var permission model.Permission
	err := r.db.First(&permission, id).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

// GetByCode 根据权限代码获取权限
func (r *PermissionRepository) GetByCode(code string) (*model.Permission, error) {
	var permission model.Permission
	err := r.db.Where("code = ?", code).First(&permission).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

// Update 更新权限信息
func (r *PermissionRepository) Update(permission *model.Permission) error {
	return r.db.Save(permission).Error
}

// Delete 删除权限
func (r *PermissionRepository) Delete(id uint) error {
	return r.db.Delete(&model.Permission{}, id).Error
}

// List 分页获取权限列表
func (r *PermissionRepository) List(offset, limit int) ([]model.Permission, int64, error) {
	var permissions []model.Permission
	var total int64

	// 获取总数
	err := r.db.Model(&model.Permission{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = r.db.Offset(offset).Limit(limit).Order("resource, action").Find(&permissions).Error
	return permissions, total, err
}

// GetAll 获取所有权限
func (r *PermissionRepository) GetAll() ([]model.Permission, error) {
	var permissions []model.Permission
	err := r.db.Order("resource, action").Find(&permissions).Error
	return permissions, err
}

// GetByResource 根据资源类型获取权限
func (r *PermissionRepository) GetByResource(resource string) ([]model.Permission, error) {
	var permissions []model.Permission
	err := r.db.Where("resource = ?", resource).Order("action").Find(&permissions).Error
	return permissions, err
}

// GetByType 根据类型获取权限
func (r *PermissionRepository) GetByType(permType string) ([]model.Permission, error) {
	var permissions []model.Permission
	err := r.db.Where("type = ?", permType).Order("resource, action").Find(&permissions).Error
	return permissions, err
}

// CheckCodeExists 检查权限代码是否已存在
func (r *PermissionRepository) CheckCodeExists(code string) (bool, error) {
	var count int64
	err := r.db.Model(&model.Permission{}).Where("code = ?", code).Count(&count).Error
	return count > 0, err
}

// CheckCodeExistsExcludeID 检查权限代码是否已存在（排除指定ID）
func (r *PermissionRepository) CheckCodeExistsExcludeID(code string, excludeID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.Permission{}).Where("code = ? AND id != ?", code, excludeID).Count(&count).Error
	return count > 0, err
}

// GetByIDs 根据 ID 列表批量获取权限
func (r *PermissionRepository) GetByIDs(ids []uint) ([]model.Permission, error) {
	var permissions []model.Permission
	err := r.db.Where("id IN ?", ids).Find(&permissions).Error
	return permissions, err
}

// GetByCodes 根据权限代码列表批量获取权限
func (r *PermissionRepository) GetByCodes(codes []string) ([]model.Permission, error) {
	var permissions []model.Permission
	err := r.db.Where("code IN ?", codes).Find(&permissions).Error
	return permissions, err
}
