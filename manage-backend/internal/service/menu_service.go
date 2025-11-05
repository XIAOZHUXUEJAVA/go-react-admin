package service

import (
	"errors"

	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"

	apperrors "github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/errors"
)

// MenuRepositoryInterface 定义菜单仓库接口
type MenuRepositoryInterface interface {
	Create(menu *model.Menu) error
	GetByID(id uint) (*model.Menu, error)
	Update(menu *model.Menu) error
	Delete(id uint) error
	GetAll() ([]model.Menu, error)
	GetByParentID(parentID *uint) ([]model.Menu, error)
	GetRootMenus() ([]model.Menu, error)
	GetByType(menuType string) ([]model.Menu, error)
	GetVisibleMenus() ([]model.Menu, error)
	GetByPermissionCodes(codes []string) ([]model.Menu, error)
	GetMenusWithoutPermission() ([]model.Menu, error)
	HasChildren(id uint) (bool, error)
}

// MenuService 菜单业务服务
type MenuService struct {
	menuRepo       MenuRepositoryInterface
	permissionRepo PermissionRepositoryInterface
}

// NewMenuService 创建 MenuService 实例
func NewMenuService(
	menuRepo MenuRepositoryInterface,
	permissionRepo PermissionRepositoryInterface,
) *MenuService {
	return &MenuService{
		menuRepo:       menuRepo,
		permissionRepo: permissionRepo,
	}
}

// Create 创建菜单
func (s *MenuService) Create(req *model.CreateMenuRequest) (*model.MenuResponse, error) {
	// 如果有父菜单，验证父菜单是否存在
	if req.ParentID != nil {
		_, err := s.menuRepo.GetByID(*req.ParentID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, apperrors.NewMenuParentNotFoundError()
			}
			return nil, apperrors.NewMenuParentGetFailedError()
		}
	}

	// 如果有权限代码，验证权限是否存在
	if req.PermissionCode != "" {
		_, err := s.permissionRepo.GetByCode(req.PermissionCode)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, apperrors.NewPermissionNotFoundError("权限代码不存在")
			}
			return nil, apperrors.NewMenuPermissionGetFailedError()
		}
	}

	// 创建菜单
	menu := &model.Menu{
		ParentID:       req.ParentID,
		Name:           req.Name,
		Title:          req.Title,
		Path:           req.Path,
		Component:      req.Component,
		Icon:           req.Icon,
		OrderNum:       req.OrderNum,
		Type:           req.Type,
		PermissionCode: req.PermissionCode,
		Visible:        req.Visible,
		Status:         "active",
	}

	if err := s.menuRepo.Create(menu); err != nil {
		logger.Error("创建菜单失败", zap.String("name", req.Name), zap.Error(err))
		return nil, apperrors.NewMenuCreateFailedError()
	}

	logger.Info("创建菜单成功",
		zap.Uint("menu_id", menu.ID),
		zap.String("name", menu.Name),
		zap.String("title", menu.Title))

	return s.toMenuResponse(menu), nil
}

// GetByID 根据ID获取菜单
func (s *MenuService) GetByID(id uint) (*model.MenuResponse, error) {
	menu, err := s.menuRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewMenuNotFoundError()
		}
		logger.Error("获取菜单失败", zap.Uint("menu_id", id), zap.Error(err))
		return nil, apperrors.NewMenuGetFailedError()
	}

	return s.toMenuResponse(menu), nil
}

// Update 更新菜单
func (s *MenuService) Update(id uint, req *model.UpdateMenuRequest) (*model.MenuResponse, error) {
	// 获取菜单
	menu, err := s.menuRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.NewMenuNotFoundError()
		}
		logger.Error("获取菜单失败", zap.Uint("menu_id", id), zap.Error(err))
		return nil, apperrors.NewMenuGetFailedError()
	}

	// 如果修改父菜单，验证父菜单是否存在
	if req.ParentID != nil {
		// 不能将菜单设置为自己的子菜单
		if *req.ParentID == id {
			return nil, apperrors.NewMenuCannotBeOwnChildError()
		}

		_, err := s.menuRepo.GetByID(*req.ParentID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, apperrors.NewMenuParentNotFoundError()
			}
			return nil, apperrors.NewMenuParentGetFailedError()
		}
		menu.ParentID = req.ParentID
	}

	// 如果修改权限代码，验证权限是否存在
	if req.PermissionCode != "" {
		_, err := s.permissionRepo.GetByCode(req.PermissionCode)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, apperrors.NewPermissionNotFoundError("权限代码不存在")
			}
			return nil, apperrors.NewMenuPermissionGetFailedError()
		}
		menu.PermissionCode = req.PermissionCode
	}

	// 更新字段
	if req.Name != "" {
		menu.Name = req.Name
	}
	if req.Title != "" {
		menu.Title = req.Title
	}
	if req.Path != "" {
		menu.Path = req.Path
	}
	if req.Component != "" {
		menu.Component = req.Component
	}
	if req.Icon != "" {
		menu.Icon = req.Icon
	}
	if req.OrderNum != 0 {
		menu.OrderNum = req.OrderNum
	}
	menu.Visible = req.Visible
	if req.Status != "" {
		menu.Status = req.Status
	}

	if err := s.menuRepo.Update(menu); err != nil {
		logger.Error("更新菜单失败", zap.Uint("menu_id", id), zap.Error(err))
		return nil, apperrors.NewMenuUpdateFailedError()
	}

	logger.Info("更新菜单成功",
		zap.Uint("menu_id", menu.ID),
		zap.String("name", menu.Name))

	return s.toMenuResponse(menu), nil
}

// Delete 删除菜单
func (s *MenuService) Delete(id uint) error {
	// 获取菜单
	menu, err := s.menuRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperrors.NewMenuNotFoundError()
		}
		logger.Error("获取菜单失败", zap.Uint("menu_id", id), zap.Error(err))
		return apperrors.NewMenuGetFailedError()
	}

	// 检查是否有子菜单
	hasChildren, err := s.menuRepo.HasChildren(id)
	if err != nil {
		logger.Error("检查子菜单失败", zap.Uint("menu_id", id), zap.Error(err))
		return apperrors.NewMenuCheckChildrenFailedError()
	}
	if hasChildren {
		return apperrors.NewMenuHasChildrenError()
	}

	// 删除菜单
	if err := s.menuRepo.Delete(id); err != nil {
		logger.Error("删除菜单失败", zap.Uint("menu_id", id), zap.Error(err))
		return apperrors.NewMenuDeleteFailedError()
	}

	logger.Info("删除菜单成功",
		zap.Uint("menu_id", id),
		zap.String("name", menu.Name))

	return nil
}

// GetMenuTree 获取菜单树（所有菜单）
func (s *MenuService) GetMenuTree() ([]model.MenuResponse, error) {
	menus, err := s.menuRepo.GetAll()
	if err != nil {
		logger.Error("获取菜单列表失败", zap.Error(err))
		return nil, apperrors.NewMenuListFailedError()
	}

	return s.buildMenuTree(menus, nil), nil
}

// GetVisibleMenuTree 获取可见菜单树
func (s *MenuService) GetVisibleMenuTree() ([]model.MenuResponse, error) {
	menus, err := s.menuRepo.GetVisibleMenus()
	if err != nil {
		logger.Error("获取可见菜单失败", zap.Error(err))
		return nil, apperrors.NewMenuVisibleListFailedError()
	}

	return s.buildMenuTree(menus, nil), nil
}

// GetUserMenuTree 获取用户的菜单树（根据角色菜单关联）
func (s *MenuService) GetUserMenuTree(userID uint, roleRepo RoleRepositoryInterface) ([]model.MenuResponse, error) {
	// 获取用户角色
	roles, err := roleRepo.GetUserRoles(userID)
	if err != nil {
		logger.Error("获取用户角色失败", zap.Uint("user_id", userID), zap.Error(err))
		return nil, apperrors.NewUserRoleGetFailedError()
	}

	if len(roles) == 0 {
		logger.Warn("用户没有分配角色", zap.Uint("user_id", userID))
		return []model.MenuResponse{}, nil
	}

	// 如果用户是超级管理员，返回所有可见菜单
	for _, role := range roles {
		if role.Code == "admin" {
			return s.GetVisibleMenuTree()
		}
	}

	// 获取所有可见菜单
	allVisibleMenus, err := s.menuRepo.GetVisibleMenus()
	if err != nil {
		logger.Error("获取可见菜单失败", zap.Error(err))
		return nil, apperrors.NewMenuVisibleListFailedError()
	}

	// 获取用户的所有权限代码
	userPermissions := make(map[string]bool)
	for _, role := range roles {
		// 从 role_permissions 表获取权限ID
		permissionIDs, err := roleRepo.GetRolePermissionIDs(role.ID)
		if err == nil {
			// 根据权限ID获取权限详情
			permissions, err := s.permissionRepo.GetByIDs(permissionIDs)
			if err == nil {
				for _, perm := range permissions {
					userPermissions[perm.Code] = true
				}
			}
		}
	}

	// 根据权限过滤菜单
	userMenus := make([]model.Menu, 0)
	for _, menu := range allVisibleMenus {
		// 如果菜单没有权限要求（公共菜单），或者用户有该权限
		if menu.PermissionCode == "" || userPermissions[menu.PermissionCode] {
			userMenus = append(userMenus, menu)
		}
	}

	logger.Info("获取用户菜单成功",
		zap.Uint("user_id", userID),
		zap.Int("menu_count", len(userMenus)),
		zap.Int("permission_count", len(userPermissions)))

	return s.buildMenuTree(userMenus, nil), nil
}

// buildMenuTree 构建菜单树
func (s *MenuService) buildMenuTree(menus []model.Menu, parentID *uint) []model.MenuResponse {
	var tree []model.MenuResponse

	for _, menu := range menus {
		// 匹配父菜单
		if (parentID == nil && menu.ParentID == nil) ||
			(parentID != nil && menu.ParentID != nil && *menu.ParentID == *parentID) {

			menuResponse := s.toMenuResponse(&menu)

			// 递归构建子菜单
			children := s.buildMenuTree(menus, &menu.ID)
			if len(children) > 0 {
				menuResponse.Children = children
			}

			tree = append(tree, *menuResponse)
		}
	}

	return tree
}

// toMenuResponse 转换为响应结构
func (s *MenuService) toMenuResponse(menu *model.Menu) *model.MenuResponse {
	return &model.MenuResponse{
		ID:             menu.ID,
		ParentID:       menu.ParentID,
		Name:           menu.Name,
		Title:          menu.Title,
		Path:           menu.Path,
		Component:      menu.Component,
		Icon:           menu.Icon,
		OrderNum:       menu.OrderNum,
		Type:           menu.Type,
		PermissionCode: menu.PermissionCode,
		Visible:        menu.Visible,
		Status:         menu.Status,
		CreatedAt:      menu.CreatedAt,
		UpdatedAt:      menu.UpdatedAt,
	}
}

// UpdateMenuOrder 批量更新菜单顺序
func (s *MenuService) UpdateMenuOrder(updates []model.MenuOrderUpdate) error {
	for _, update := range updates {
		menu, err := s.menuRepo.GetByID(update.ID)
		if err != nil {
			logger.Error("获取菜单失败", zap.Uint("menu_id", update.ID), zap.Error(err))
			return apperrors.NewMenuNotFoundError()
		}

		menu.OrderNum = update.OrderNum
		if update.ParentID != nil {
			menu.ParentID = update.ParentID
		}

		if err := s.menuRepo.Update(menu); err != nil {
			logger.Error("更新菜单顺序失败", zap.Uint("menu_id", update.ID), zap.Error(err))
			return apperrors.NewMenuOrderUpdateFailedError()
		}
	}

	logger.Info("批量更新菜单顺序成功", zap.Int("count", len(updates)))
	return nil
}
