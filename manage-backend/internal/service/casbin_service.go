package service

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
)

// CasbinService Casbin 权限服务
// 封装 Casbin Enforcer 的所有权限操作
type CasbinService struct {
	enforcer *casbin.Enforcer
}

// NewCasbinService 创建 CasbinService 实例
func NewCasbinService(enforcer *casbin.Enforcer) *CasbinService {
	return &CasbinService{
		enforcer: enforcer,
	}
}

// CheckPermission 检查用户权限
// 参数：
//   - userID: 用户ID
//   - path: 请求路径
//   - method: HTTP方法
// 返回：bool - 是否有权限
func (s *CasbinService) CheckPermission(userID uint, path, method string) (bool, error) {
	sub := fmt.Sprintf("user:%d", userID)
	ok, err := s.enforcer.Enforce(sub, path, method)
	if err != nil {
		logger.Error("Casbin权限检查失败", 
			zap.Uint("user_id", userID),
			zap.String("path", path),
			zap.String("method", method),
			zap.Error(err))
		return false, err
	}
	return ok, nil
}

// AddRoleForUser 为用户分配角色
func (s *CasbinService) AddRoleForUser(userID uint, roleCode string) error {
	sub := fmt.Sprintf("user:%d", userID)
	role := fmt.Sprintf("role:%s", roleCode)
	
	_, err := s.enforcer.AddGroupingPolicy(sub, role)
	if err != nil {
		logger.Error("添加用户角色失败",
			zap.Uint("user_id", userID),
			zap.String("role", roleCode),
			zap.Error(err))
		return err
	}
	
	logger.Info("添加用户角色成功",
		zap.Uint("user_id", userID),
		zap.String("role", roleCode))
	return nil
}

// RemoveRoleForUser 移除用户的角色
func (s *CasbinService) RemoveRoleForUser(userID uint, roleCode string) error {
	sub := fmt.Sprintf("user:%d", userID)
	role := fmt.Sprintf("role:%s", roleCode)
	
	_, err := s.enforcer.RemoveGroupingPolicy(sub, role)
	if err != nil {
		logger.Error("移除用户角色失败",
			zap.Uint("user_id", userID),
			zap.String("role", roleCode),
			zap.Error(err))
		return err
	}
	
	logger.Info("移除用户角色成功",
		zap.Uint("user_id", userID),
		zap.String("role", roleCode))
	return nil
}

// RemoveAllRolesForUser 移除用户的所有角色
func (s *CasbinService) RemoveAllRolesForUser(userID uint) error {
	sub := fmt.Sprintf("user:%d", userID)
	
	_, err := s.enforcer.DeleteRolesForUser(sub)
	if err != nil {
		logger.Error("移除用户所有角色失败",
			zap.Uint("user_id", userID),
			zap.Error(err))
		return err
	}
	
	logger.Info("移除用户所有角色成功", zap.Uint("user_id", userID))
	return nil
}

// GetRolesForUser 获取用户的所有角色
func (s *CasbinService) GetRolesForUser(userID uint) ([]string, error) {
	sub := fmt.Sprintf("user:%d", userID)
	roles, err := s.enforcer.GetRolesForUser(sub)
	if err != nil {
		logger.Error("获取用户角色失败",
			zap.Uint("user_id", userID),
			zap.Error(err))
		return nil, err
	}
	return roles, nil
}

// AddPermissionForRole 为角色添加权限
func (s *CasbinService) AddPermissionForRole(roleCode, path, method string) error {
	sub := fmt.Sprintf("role:%s", roleCode)
	
	_, err := s.enforcer.AddPolicy(sub, path, method)
	if err != nil {
		logger.Error("添加角色权限失败",
			zap.String("role", roleCode),
			zap.String("path", path),
			zap.String("method", method),
			zap.Error(err))
		return err
	}
	
	logger.Info("添加角色权限成功",
		zap.String("role", roleCode),
		zap.String("path", path),
		zap.String("method", method))
	return nil
}

// RemovePermissionForRole 移除角色的权限
func (s *CasbinService) RemovePermissionForRole(roleCode, path, method string) error {
	sub := fmt.Sprintf("role:%s", roleCode)
	
	_, err := s.enforcer.RemovePolicy(sub, path, method)
	if err != nil {
		logger.Error("移除角色权限失败",
			zap.String("role", roleCode),
			zap.String("path", path),
			zap.String("method", method),
			zap.Error(err))
		return err
	}
	
	logger.Info("移除角色权限成功",
		zap.String("role", roleCode),
		zap.String("path", path),
		zap.String("method", method))
	return nil
}

// RemoveAllPermissionsForRole 移除角色的所有权限
func (s *CasbinService) RemoveAllPermissionsForRole(roleCode string) error {
	sub := fmt.Sprintf("role:%s", roleCode)
	
	_, err := s.enforcer.DeletePermissionsForUser(sub)
	if err != nil {
		logger.Error("移除角色所有权限失败",
			zap.String("role", roleCode),
			zap.Error(err))
		return err
	}
	
	logger.Info("移除角色所有权限成功", zap.String("role", roleCode))
	return nil
}

// GetPermissionsForRole 获取角色的所有权限
func (s *CasbinService) GetPermissionsForRole(roleCode string) ([][]string, error) {
	sub := fmt.Sprintf("role:%s", roleCode)
	permissions, err := s.enforcer.GetPermissionsForUser(sub)
	if err != nil {
		logger.Error("获取角色权限失败",
			zap.String("role", roleCode),
			zap.Error(err))
		return nil, err
	}
	return permissions, nil
}

// UpdateRolePermissions 更新角色的权限（先删除所有权限，再添加新权限）
func (s *CasbinService) UpdateRolePermissions(roleCode string, permissions [][]string) error {
	// 删除角色的所有权限
	if err := s.RemoveAllPermissionsForRole(roleCode); err != nil {
		return err
	}
	
	// 添加新权限
	sub := fmt.Sprintf("role:%s", roleCode)
	for _, perm := range permissions {
		if len(perm) >= 2 {
			path := perm[0]
			method := perm[1]
			if _, err := s.enforcer.AddPolicy(sub, path, method); err != nil {
				logger.Error("添加权限失败",
					zap.String("role", roleCode),
					zap.String("path", path),
					zap.String("method", method),
					zap.Error(err))
				return err
			}
		}
	}
	
	logger.Info("更新角色权限成功",
		zap.String("role", roleCode),
		zap.Int("permission_count", len(permissions)))
	return nil
}

// GetUsersForRole 获取拥有指定角色的所有用户
func (s *CasbinService) GetUsersForRole(roleCode string) ([]string, error) {
	role := fmt.Sprintf("role:%s", roleCode)
	users, err := s.enforcer.GetUsersForRole(role)
	if err != nil {
		logger.Error("获取角色用户失败",
			zap.String("role", roleCode),
			zap.Error(err))
		return nil, err
	}
	return users, nil
}

// ReloadPolicy 重新加载策略
func (s *CasbinService) ReloadPolicy() error {
	err := s.enforcer.LoadPolicy()
	if err != nil {
		logger.Error("重新加载Casbin策略失败", zap.Error(err))
		return err
	}
	logger.Info("重新加载Casbin策略成功")
	return nil
}

// GetAllRoles 获取所有角色
func (s *CasbinService) GetAllRoles() ([]string, error) {
	roles, err := s.enforcer.GetAllRoles()
	if err != nil {
		logger.Error("获取所有角色失败", zap.Error(err))
		return nil, err
	}
	return roles, nil
}

// GetAllSubjects 获取所有主体（用户）
func (s *CasbinService) GetAllSubjects() ([]string, error) {
	subjects, err := s.enforcer.GetAllSubjects()
	if err != nil {
		logger.Error("获取所有主体失败", zap.Error(err))
		return nil, err
	}
	return subjects, nil
}
