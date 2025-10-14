package repository

import (
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"gorm.io/gorm"
)

// MenuRepository 菜单数据仓库
// 封装对 Menu 模型的所有数据库操作
type MenuRepository struct {
	db *gorm.DB
}

// NewMenuRepository 创建 MenuRepository 实例
func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

// Create 创建菜单
func (r *MenuRepository) Create(menu *model.Menu) error {
	return r.db.Create(menu).Error
}

// GetByID 根据 ID 获取菜单
func (r *MenuRepository) GetByID(id uint) (*model.Menu, error) {
	var menu model.Menu
	err := r.db.First(&menu, id).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

// Update 更新菜单信息
func (r *MenuRepository) Update(menu *model.Menu) error {
	return r.db.Save(menu).Error
}

// Delete 删除菜单
func (r *MenuRepository) Delete(id uint) error {
	return r.db.Delete(&model.Menu{}, id).Error
}

// GetAll 获取所有菜单
func (r *MenuRepository) GetAll() ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Order("order_num ASC, id ASC").Find(&menus).Error
	return menus, err
}

// GetByParentID 根据父菜单ID获取子菜单
func (r *MenuRepository) GetByParentID(parentID *uint) ([]model.Menu, error) {
	var menus []model.Menu
	query := r.db.Order("order_num ASC, id ASC")
	
	if parentID == nil {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", *parentID)
	}
	
	err := query.Find(&menus).Error
	return menus, err
}

// GetRootMenus 获取所有根菜单（没有父菜单的菜单）
func (r *MenuRepository) GetRootMenus() ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Where("parent_id IS NULL").Order("order_num ASC, id ASC").Find(&menus).Error
	return menus, err
}

// GetByType 根据类型获取菜单
func (r *MenuRepository) GetByType(menuType string) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Where("type = ?", menuType).Order("order_num ASC, id ASC").Find(&menus).Error
	return menus, err
}

// GetVisibleMenus 获取所有可见菜单
func (r *MenuRepository) GetVisibleMenus() ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Where("visible = ? AND status = ?", true, "active").
		Order("order_num ASC, id ASC").
		Find(&menus).Error
	return menus, err
}

// GetByPermissionCodes 根据权限代码列表获取菜单
func (r *MenuRepository) GetByPermissionCodes(codes []string) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Where("permission_code IN ? AND visible = ? AND status = ?", codes, true, "active").
		Order("order_num ASC, id ASC").
		Find(&menus).Error
	return menus, err
}

// GetMenusWithoutPermission 获取不需要权限的菜单
func (r *MenuRepository) GetMenusWithoutPermission() ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Where("(permission_code IS NULL OR permission_code = '') AND visible = ? AND status = ?", true, "active").
		Order("order_num ASC, id ASC").
		Find(&menus).Error
	return menus, err
}

// HasChildren 检查菜单是否有子菜单
func (r *MenuRepository) HasChildren(id uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.Menu{}).Where("parent_id = ?", id).Count(&count).Error
	return count > 0, err
}

