package model

import (
	"time"

	"gorm.io/gorm"
)

// Role 角色模型
type Role struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Name        string         `json:"name" gorm:"uniqueIndex;not null"`
	Code        string         `json:"code" gorm:"uniqueIndex;not null"`
	Description string         `json:"description"`
	Status      string         `json:"status" gorm:"default:active"`
	IsSystem    bool           `json:"is_system" gorm:"default:false"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// RoleResponse 角色响应结构体
type RoleResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	IsSystem    bool      `json:"is_system"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateRoleRequest 创建角色请求
type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=50"`
	Code        string `json:"code" binding:"required,min=2,max=50"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Status      string `json:"status" binding:"omitempty,oneof=active inactive"`
}

// UpdateRoleRequest 更新角色请求
type UpdateRoleRequest struct {
	Name        string `json:"name" binding:"omitempty,min=2,max=50"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Status      string `json:"status" binding:"omitempty,oneof=active inactive"`
}

// AssignRolePermissionsRequest 分配角色权限请求
type AssignRolePermissionsRequest struct {
	PermissionIDs []uint `json:"permission_ids" binding:"required"`
}

// AssignUserRolesRequest 分配用户角色请求
type AssignUserRolesRequest struct {
	RoleIDs []uint `json:"role_ids" binding:"required"`
}

// RoleWithPermissions 角色及其权限
type RoleWithPermissions struct {
	Role        RoleResponse       `json:"role"`
	Permissions []PermissionResponse `json:"permissions"`
}

// RolePermission 角色-权限关联模型
type RolePermission struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	RoleID       uint      `json:"role_id" gorm:"not null;index;uniqueIndex:idx_role_permission"`
	PermissionID uint      `json:"permission_id" gorm:"not null;index;uniqueIndex:idx_role_permission"`
	AssignedAt   time.Time `json:"assigned_at" gorm:"autoCreateTime"`
}

// UserRole 用户-角色关联模型
type UserRole struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	UserID     uint      `json:"user_id" gorm:"not null;index;uniqueIndex:idx_user_role"`
	RoleID     uint      `json:"role_id" gorm:"not null;index;uniqueIndex:idx_user_role"`
	AssignedAt time.Time `json:"assigned_at" gorm:"default:CURRENT_TIMESTAMP"`
	AssignedBy uint      `json:"assigned_by"`
}
