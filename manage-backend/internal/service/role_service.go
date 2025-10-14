package service

import (
	"errors"
	"fmt"

	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// RoleRepositoryInterface 定义角色仓库接口
type RoleRepositoryInterface interface {
	Create(role *model.Role) error
	GetByID(id uint) (*model.Role, error)
	GetByCode(code string) (*model.Role, error)
	Update(role *model.Role) error
	Delete(id uint) error
	List(offset, limit int) ([]model.Role, int64, error)
	GetAll() ([]model.Role, error)
	CheckCodeExists(code string) (bool, error)
	CheckCodeExistsExcludeID(code string, excludeID uint) (bool, error)
	GetUserRoles(userID uint) ([]model.Role, error)
	AssignRoleToUser(userID, roleID uint, assignedBy uint) error
	RemoveRoleFromUser(userID, roleID uint) error
	RemoveAllRolesFromUser(userID uint) error
	GetUsersByRole(roleID uint) ([]uint, error)
	UpdateRolePermissions(roleID uint, permissionIDs []uint) error
	GetRolePermissionIDs(roleID uint) ([]uint, error)
}

// PermissionRepositoryInterface 定义权限仓库接口
type PermissionRepositoryInterface interface {
	Create(permission *model.Permission) error
	GetByID(id uint) (*model.Permission, error)
	GetByCode(code string) (*model.Permission, error)
	Update(permission *model.Permission) error
	Delete(id uint) error
	List(offset, limit int) ([]model.Permission, int64, error)
	GetAll() ([]model.Permission, error)
	GetByResource(resource string) ([]model.Permission, error)
	GetByType(permType string) ([]model.Permission, error)
	CheckCodeExists(code string) (bool, error)
	CheckCodeExistsExcludeID(code string, excludeID uint) (bool, error)
	GetByIDs(ids []uint) ([]model.Permission, error)
	GetByCodes(codes []string) ([]model.Permission, error)
}

// CasbinServiceInterface 定义 Casbin 服务接口
type CasbinServiceInterface interface {
	AddRoleForUser(userID uint, roleCode string) error
	RemoveRoleForUser(userID uint, roleCode string) error
	RemoveAllRolesForUser(userID uint) error
	GetRolesForUser(userID uint) ([]string, error)
	AddPermissionForRole(roleCode, path, method string) error
	RemovePermissionForRole(roleCode, path, method string) error
	RemoveAllPermissionsForRole(roleCode string) error
	GetPermissionsForRole(roleCode string) ([][]string, error)
	UpdateRolePermissions(roleCode string, permissions [][]string) error
}

// RoleService 角色业务服务
type RoleService struct {
	roleRepo       RoleRepositoryInterface
	permissionRepo PermissionRepositoryInterface
	casbinService  CasbinServiceInterface
}

// NewRoleService 创建 RoleService 实例
func NewRoleService(
	roleRepo RoleRepositoryInterface,
	permissionRepo PermissionRepositoryInterface,
	casbinService CasbinServiceInterface,
) *RoleService {
	return &RoleService{
		roleRepo:       roleRepo,
		permissionRepo: permissionRepo,
		casbinService:  casbinService,
	}
}

// Create 创建角色
func (s *RoleService) Create(req *model.CreateRoleRequest) (*model.RoleResponse, error) {
	// 检查角色代码是否已存在
	exists, err := s.roleRepo.CheckCodeExists(req.Code)
	if err != nil {
		logger.Error("检查角色代码失败", zap.String("code", req.Code), zap.Error(err))
		return nil, err
	}
	if exists {
		return nil, errors.New("角色代码已存在")
	}

	// 创建角色
	role := &model.Role{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
		Status:      req.Status,
	}

	if role.Status == "" {
		role.Status = "active"
	}

	if err := s.roleRepo.Create(role); err != nil {
		logger.Error("创建角色失败", zap.String("code", req.Code), zap.Error(err))
		return nil, err
	}

	logger.Info("创建角色成功",
		zap.Uint("role_id", role.ID),
		zap.String("code", role.Code),
		zap.String("name", role.Name))

	return s.toRoleResponse(role), nil
}

// GetByID 根据ID获取角色
func (s *RoleService) GetByID(id uint) (*model.RoleResponse, error) {
	role, err := s.roleRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("角色不存在")
		}
		logger.Error("获取角色失败", zap.Uint("role_id", id), zap.Error(err))
		return nil, err
	}

	return s.toRoleResponse(role), nil
}

// Update 更新角色
func (s *RoleService) Update(id uint, req *model.UpdateRoleRequest) (*model.RoleResponse, error) {
	// 获取角色
	role, err := s.roleRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("角色不存在")
		}
		logger.Error("获取角色失败", zap.Uint("role_id", id), zap.Error(err))
		return nil, err
	}

	// 检查是否为系统角色
	if role.IsSystem {
		return nil, errors.New("系统角色不允许修改")
	}

	// 更新字段
	if req.Name != "" {
		role.Name = req.Name
	}
	if req.Description != "" {
		role.Description = req.Description
	}
	if req.Status != "" {
		role.Status = req.Status
	}

	if err := s.roleRepo.Update(role); err != nil {
		logger.Error("更新角色失败", zap.Uint("role_id", id), zap.Error(err))
		return nil, err
	}

	logger.Info("更新角色成功",
		zap.Uint("role_id", role.ID),
		zap.String("code", role.Code))

	return s.toRoleResponse(role), nil
}

// Delete 删除角色
func (s *RoleService) Delete(id uint) error {
	// 获取角色
	role, err := s.roleRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("角色不存在")
		}
		logger.Error("获取角色失败", zap.Uint("role_id", id), zap.Error(err))
		return err
	}

	// 检查是否为系统角色
	if role.IsSystem {
		return errors.New("系统角色不允许删除")
	}

	// 检查是否有用户使用该角色
	userIDs, err := s.roleRepo.GetUsersByRole(id)
	if err != nil {
		logger.Error("检查角色使用情况失败", zap.Uint("role_id", id), zap.Error(err))
		return err
	}
	if len(userIDs) > 0 {
		return fmt.Errorf("该角色正在被 %d 个用户使用，无法删除", len(userIDs))
	}

	// 删除 Casbin 中的角色权限
	if err := s.casbinService.RemoveAllPermissionsForRole(role.Code); err != nil {
		logger.Error("删除角色权限失败", zap.String("code", role.Code), zap.Error(err))
		return err
	}

	// 删除角色
	if err := s.roleRepo.Delete(id); err != nil {
		logger.Error("删除角色失败", zap.Uint("role_id", id), zap.Error(err))
		return err
	}

	logger.Info("删除角色成功",
		zap.Uint("role_id", id),
		zap.String("code", role.Code))

	return nil
}

// List 分页获取角色列表
func (s *RoleService) List(page, pageSize int) ([]model.RoleResponse, int64, error) {
	offset := (page - 1) * pageSize
	roles, total, err := s.roleRepo.List(offset, pageSize)
	if err != nil {
		logger.Error("获取角色列表失败", zap.Error(err))
		return nil, 0, err
	}

	responses := make([]model.RoleResponse, len(roles))
	for i, role := range roles {
		responses[i] = *s.toRoleResponse(&role)
	}

	return responses, total, nil
}

// GetAll 获取所有角色
func (s *RoleService) GetAll() ([]model.RoleResponse, error) {
	roles, err := s.roleRepo.GetAll()
	if err != nil {
		logger.Error("获取所有角色失败", zap.Error(err))
		return nil, err
	}

	responses := make([]model.RoleResponse, len(roles))
	for i, role := range roles {
		responses[i] = *s.toRoleResponse(&role)
	}

	return responses, nil
}

// AssignPermissions 为角色分配权限
func (s *RoleService) AssignPermissions(roleID uint, req *model.AssignRolePermissionsRequest) error {
	// 获取角色
	role, err := s.roleRepo.GetByID(roleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("角色不存在")
		}
		logger.Error("获取角色失败", zap.Uint("role_id", roleID), zap.Error(err))
		return err
	}

	// 获取权限列表
	permissions, err := s.permissionRepo.GetByIDs(req.PermissionIDs)
	if err != nil {
		logger.Error("获取权限列表失败", zap.Error(err))
		return err
	}

	logger.Info("获取到的权限数量",
		zap.Int("requested", len(req.PermissionIDs)),
		zap.Int("found", len(permissions)))

	if len(permissions) != len(req.PermissionIDs) {
		logger.Warn("部分权限不存在",
			zap.Uints("requested_ids", req.PermissionIDs),
			zap.Int("found_count", len(permissions)))
		return errors.New("部分权限不存在")
	}

	// 构建 Casbin 策略
	// 注意：只有 API 类型的权限才需要添加到 Casbin
	var casbinPolicies [][]string
	for _, perm := range permissions {
		// 只处理 API 类型的权限
		if perm.Type == "api" && perm.Path != "" && perm.Method != "" {
			casbinPolicies = append(casbinPolicies, []string{perm.Path, perm.Method})
			logger.Debug("添加API权限策略",
				zap.String("code", perm.Code),
				zap.String("path", perm.Path),
				zap.String("method", perm.Method))
		} else if perm.Type == "menu" {
			// 菜单类型权限不需要添加到 Casbin，只记录日志
			logger.Debug("跳过菜单类型权限",
				zap.String("code", perm.Code),
				zap.String("type", perm.Type))
		} else {
			logger.Warn("跳过无效权限",
				zap.String("code", perm.Code),
				zap.String("type", perm.Type),
				zap.String("path", perm.Path),
				zap.String("method", perm.Method))
		}
	}

	// 更新 Casbin 策略（只更新 API 类型的权限）
	if err := s.casbinService.UpdateRolePermissions(role.Code, casbinPolicies); err != nil {
		logger.Error("更新 Casbin 角色权限失败",
			zap.Uint("role_id", roleID),
			zap.String("code", role.Code),
			zap.Error(err))
		return err
	}

	// 更新数据库中的角色-权限关联（包括所有类型的权限）
	if err := s.roleRepo.UpdateRolePermissions(roleID, req.PermissionIDs); err != nil {
		logger.Error("更新数据库角色权限失败",
			zap.Uint("role_id", roleID),
			zap.Error(err))
		return err
	}

	logger.Info("分配角色权限成功",
		zap.Uint("role_id", roleID),
		zap.String("code", role.Code),
		zap.Int("permission_count", len(permissions)),
		zap.Int("casbin_policy_count", len(casbinPolicies)))

	return nil
}

// GetRolePermissions 获取角色的权限列表
func (s *RoleService) GetRolePermissions(roleID uint) (*model.RoleWithPermissions, error) {
	// 获取角色
	role, err := s.roleRepo.GetByID(roleID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("角色不存在")
		}
		logger.Error("获取角色失败", zap.Uint("role_id", roleID), zap.Error(err))
		return nil, err
	}

	// 获取所有权限用于匹配
	allPermissions, err := s.permissionRepo.GetAll()
	if err != nil {
		logger.Error("获取所有权限失败", zap.Error(err))
		return nil, err
	}

	// 如果是超级管理员，返回所有权限
	if role.Code == "admin" {
		permissions := make([]model.PermissionResponse, 0, len(allPermissions))
		for _, perm := range allPermissions {
			permissions = append(permissions, model.PermissionResponse{
				ID:          perm.ID,
				Name:        perm.Name,
				Code:        perm.Code,
				Resource:    perm.Resource,
				Action:      perm.Action,
				Path:        perm.Path,
				Method:      perm.Method,
				Description: perm.Description,
				Type:        perm.Type,
				Status:      perm.Status,
				CreatedAt:   perm.CreatedAt,
				UpdatedAt:   perm.UpdatedAt,
			})
		}
		return &model.RoleWithPermissions{
			Role:        *s.toRoleResponse(role),
			Permissions: permissions,
		}, nil
	}

	// 从数据库获取角色的权限ID列表
	permissionIDs, err := s.roleRepo.GetRolePermissionIDs(roleID)
	if err != nil {
		logger.Error("获取角色权限ID失败", zap.Uint("role_id", roleID), zap.Error(err))
		return nil, err
	}

	// 根据ID匹配权限
	permissions := make([]model.PermissionResponse, 0)
	permIDMap := make(map[uint]bool)
	for _, id := range permissionIDs {
		permIDMap[id] = true
	}

	for _, perm := range allPermissions {
		if permIDMap[perm.ID] {
			permissions = append(permissions, model.PermissionResponse{
				ID:          perm.ID,
				Name:        perm.Name,
				Code:        perm.Code,
				Resource:    perm.Resource,
				Action:      perm.Action,
				Path:        perm.Path,
				Method:      perm.Method,
				Description: perm.Description,
				Type:        perm.Type,
				Status:      perm.Status,
				CreatedAt:   perm.CreatedAt,
				UpdatedAt:   perm.UpdatedAt,
			})
		}
	}

	return &model.RoleWithPermissions{
		Role:        *s.toRoleResponse(role),
		Permissions: permissions,
	}, nil
}

// AssignRolesToUser 为用户分配角色
func (s *RoleService) AssignRolesToUser(userID uint, req *model.AssignUserRolesRequest, assignedBy uint) error {
	// 验证角色是否存在
	for _, roleID := range req.RoleIDs {
		_, err := s.roleRepo.GetByID(roleID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("角色 ID %d 不存在", roleID)
			}
			return err
		}
	}

	// 移除用户的所有现有角色
	if err := s.roleRepo.RemoveAllRolesFromUser(userID); err != nil {
		logger.Error("移除用户角色失败", zap.Uint("user_id", userID), zap.Error(err))
		return err
	}

	// 移除 Casbin 中的用户角色关系
	if err := s.casbinService.RemoveAllRolesForUser(userID); err != nil {
		logger.Error("移除用户Casbin角色失败", zap.Uint("user_id", userID), zap.Error(err))
		return err
	}

	// 分配新角色
	for _, roleID := range req.RoleIDs {
		role, _ := s.roleRepo.GetByID(roleID)

		// 添加到数据库
		if err := s.roleRepo.AssignRoleToUser(userID, roleID, assignedBy); err != nil {
			logger.Error("分配角色失败",
				zap.Uint("user_id", userID),
				zap.Uint("role_id", roleID),
				zap.Error(err))
			return err
		}

		// 添加到 Casbin
		if err := s.casbinService.AddRoleForUser(userID, role.Code); err != nil {
			logger.Error("添加Casbin角色失败",
				zap.Uint("user_id", userID),
				zap.String("role_code", role.Code),
				zap.Error(err))
			return err
		}
	}

	logger.Info("分配用户角色成功",
		zap.Uint("user_id", userID),
		zap.Int("role_count", len(req.RoleIDs)))

	return nil
}

// GetUserRoles 获取用户的角色列表
func (s *RoleService) GetUserRoles(userID uint) ([]model.RoleResponse, error) {
	roles, err := s.roleRepo.GetUserRoles(userID)
	if err != nil {
		logger.Error("获取用户角色失败", zap.Uint("user_id", userID), zap.Error(err))
		return nil, err
	}

	responses := make([]model.RoleResponse, len(roles))
	for i, role := range roles {
		responses[i] = *s.toRoleResponse(&role)
	}

	return responses, nil
}

// toRoleResponse 转换为响应结构
func (s *RoleService) toRoleResponse(role *model.Role) *model.RoleResponse {
	return &model.RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
		Status:      role.Status,
		IsSystem:    role.IsSystem,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}
}
