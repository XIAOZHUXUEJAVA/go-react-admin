package service

import (
	"errors"

	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// PermissionService 权限业务服务
type PermissionService struct {
	permissionRepo PermissionRepositoryInterface
}

// NewPermissionService 创建 PermissionService 实例
func NewPermissionService(permissionRepo PermissionRepositoryInterface) *PermissionService {
	return &PermissionService{
		permissionRepo: permissionRepo,
	}
}

// Create 创建权限
func (s *PermissionService) Create(req *model.CreatePermissionRequest) (*model.PermissionResponse, error) {
	// 检查权限代码是否已存在
	exists, err := s.permissionRepo.CheckCodeExists(req.Code)
	if err != nil {
		logger.Error("检查权限代码失败", zap.String("code", req.Code), zap.Error(err))
		return nil, err
	}
	if exists {
		return nil, errors.New("权限代码已存在")
	}

	// 创建权限
	permission := &model.Permission{
		Name:        req.Name,
		Code:        req.Code,
		Resource:    req.Resource,
		Action:      req.Action,
		Path:        req.Path,
		Method:      req.Method,
		Description: req.Description,
		Type:        req.Type,
		Status:      "active",
	}

	if permission.Type == "" {
		permission.Type = "api"
	}

	if err := s.permissionRepo.Create(permission); err != nil {
		logger.Error("创建权限失败", zap.String("code", req.Code), zap.Error(err))
		return nil, err
	}

	logger.Info("创建权限成功",
		zap.Uint("permission_id", permission.ID),
		zap.String("code", permission.Code),
		zap.String("name", permission.Name))

	return s.toPermissionResponse(permission), nil
}

// GetByID 根据ID获取权限
func (s *PermissionService) GetByID(id uint) (*model.PermissionResponse, error) {
	permission, err := s.permissionRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("权限不存在")
		}
		logger.Error("获取权限失败", zap.Uint("permission_id", id), zap.Error(err))
		return nil, err
	}

	return s.toPermissionResponse(permission), nil
}

// Update 更新权限
func (s *PermissionService) Update(id uint, req *model.UpdatePermissionRequest) (*model.PermissionResponse, error) {
	// 获取权限
	permission, err := s.permissionRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("权限不存在")
		}
		logger.Error("获取权限失败", zap.Uint("permission_id", id), zap.Error(err))
		return nil, err
	}

	// 更新字段
	if req.Name != "" {
		permission.Name = req.Name
	}
	if req.Description != "" {
		permission.Description = req.Description
	}
	if req.Path != "" {
		permission.Path = req.Path
	}
	if req.Method != "" {
		permission.Method = req.Method
	}
	if req.Status != "" {
		permission.Status = req.Status
	}

	if err := s.permissionRepo.Update(permission); err != nil {
		logger.Error("更新权限失败", zap.Uint("permission_id", id), zap.Error(err))
		return nil, err
	}

	logger.Info("更新权限成功",
		zap.Uint("permission_id", permission.ID),
		zap.String("code", permission.Code))

	return s.toPermissionResponse(permission), nil
}

// Delete 删除权限
func (s *PermissionService) Delete(id uint) error {
	// 获取权限
	permission, err := s.permissionRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("权限不存在")
		}
		logger.Error("获取权限失败", zap.Uint("permission_id", id), zap.Error(err))
		return err
	}

	// 删除权限
	if err := s.permissionRepo.Delete(id); err != nil {
		logger.Error("删除权限失败", zap.Uint("permission_id", id), zap.Error(err))
		return err
	}

	logger.Info("删除权限成功",
		zap.Uint("permission_id", id),
		zap.String("code", permission.Code))

	return nil
}

// List 分页获取权限列表
func (s *PermissionService) List(page, pageSize int) ([]model.PermissionResponse, int64, error) {
	offset := (page - 1) * pageSize
	permissions, total, err := s.permissionRepo.List(offset, pageSize)
	if err != nil {
		logger.Error("获取权限列表失败", zap.Error(err))
		return nil, 0, err
	}

	responses := make([]model.PermissionResponse, len(permissions))
	for i, permission := range permissions {
		responses[i] = *s.toPermissionResponse(&permission)
	}

	return responses, total, nil
}

// GetAll 获取所有权限
func (s *PermissionService) GetAll() ([]model.PermissionResponse, error) {
	permissions, err := s.permissionRepo.GetAll()
	if err != nil {
		logger.Error("获取所有权限失败", zap.Error(err))
		return nil, err
	}

	responses := make([]model.PermissionResponse, len(permissions))
	for i, permission := range permissions {
		responses[i] = *s.toPermissionResponse(&permission)
	}

	return responses, nil
}

// GetByResource 根据资源类型获取权限
func (s *PermissionService) GetByResource(resource string) ([]model.PermissionResponse, error) {
	permissions, err := s.permissionRepo.GetByResource(resource)
	if err != nil {
		logger.Error("根据资源获取权限失败", zap.String("resource", resource), zap.Error(err))
		return nil, err
	}

	responses := make([]model.PermissionResponse, len(permissions))
	for i, permission := range permissions {
		responses[i] = *s.toPermissionResponse(&permission)
	}

	return responses, nil
}

// GetByType 根据类型获取权限
func (s *PermissionService) GetByType(permType string) ([]model.PermissionResponse, error) {
	permissions, err := s.permissionRepo.GetByType(permType)
	if err != nil {
		logger.Error("根据类型获取权限失败", zap.String("type", permType), zap.Error(err))
		return nil, err
	}

	responses := make([]model.PermissionResponse, len(permissions))
	for i, permission := range permissions {
		responses[i] = *s.toPermissionResponse(&permission)
	}

	return responses, nil
}

// GetPermissionTree 获取权限树（按资源分组）
func (s *PermissionService) GetPermissionTree() ([]model.PermissionTree, error) {
	permissions, err := s.permissionRepo.GetAll()
	if err != nil {
		logger.Error("获取权限树失败", zap.Error(err))
		return nil, err
	}

	// 按资源分组
	resourceMap := make(map[string][]model.PermissionResponse)
	for _, perm := range permissions {
		response := s.toPermissionResponse(&perm)
		resourceMap[perm.Resource] = append(resourceMap[perm.Resource], *response)
	}

	// 转换为树形结构
	var tree []model.PermissionTree
	for resource, perms := range resourceMap {
		tree = append(tree, model.PermissionTree{
			Resource:    resource,
			Permissions: perms,
		})
	}

	return tree, nil
}

// GetUserPermissions 获取用户的所有权限（通过角色）
func (s *PermissionService) GetUserPermissions(userID uint, roleRepo RoleRepositoryInterface) ([]string, error) {
	// 获取用户角色
	roles, err := roleRepo.GetUserRoles(userID)
	if err != nil {
		logger.Error("获取用户角色失败", zap.Uint("user_id", userID), zap.Error(err))
		return nil, err
	}

	// 收集所有权限代码
	permissionCodes := make(map[string]bool)
	for range roles {
		// 这里可以扩展：从 Casbin 或数据库获取角色的权限
		// 暂时返回空，后续可以通过 Casbin 获取
	}

	// 转换为列表
	var codes []string
	for code := range permissionCodes {
		codes = append(codes, code)
	}

	return codes, nil
}

// toPermissionResponse 转换为响应结构
func (s *PermissionService) toPermissionResponse(permission *model.Permission) *model.PermissionResponse {
	return &model.PermissionResponse{
		ID:          permission.ID,
		Name:        permission.Name,
		Code:        permission.Code,
		Resource:    permission.Resource,
		Action:      permission.Action,
		Path:        permission.Path,
		Method:      permission.Method,
		Description: permission.Description,
		Type:        permission.Type,
		Status:      permission.Status,
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}
}
