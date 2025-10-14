package model

import (
	"time"

	"gorm.io/gorm"
)

// Permission 权限模型
type Permission struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Name        string         `json:"name" gorm:"not null"`
	Code        string         `json:"code" gorm:"uniqueIndex;not null"`
	Resource    string         `json:"resource" gorm:"not null"`
	Action      string         `json:"action" gorm:"not null"`
	Path        string         `json:"path"`
	Method      string         `json:"method"`
	Description string         `json:"description"`
	Type        string         `json:"type" gorm:"default:api"`
	Status      string         `json:"status" gorm:"default:active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// PermissionResponse 权限响应结构体
type PermissionResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Resource    string    `json:"resource"`
	Action      string    `json:"action"`
	Path        string    `json:"path"`
	Method      string    `json:"method"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreatePermissionRequest 创建权限请求
type CreatePermissionRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=100"`
	Code        string `json:"code" binding:"required,min=2,max=100"`
	Resource    string `json:"resource" binding:"required,max=50"`
	Action      string `json:"action" binding:"required,max=50"`
	Path        string `json:"path" binding:"omitempty,max=255"`
	Method      string `json:"method" binding:"omitempty,oneof=GET POST PUT DELETE PATCH"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Type        string `json:"type" binding:"omitempty,oneof=api menu button"`
}

// UpdatePermissionRequest 更新权限请求
type UpdatePermissionRequest struct {
	Name        string `json:"name" binding:"omitempty,min=2,max=100"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Path        string `json:"path" binding:"omitempty,max=255"`
	Method      string `json:"method" binding:"omitempty,oneof=GET POST PUT DELETE PATCH"`
	Status      string `json:"status" binding:"omitempty,oneof=active inactive"`
}

// PermissionTree 权限树结构（按资源分组）
type PermissionTree struct {
	Resource    string               `json:"resource"`
	Permissions []PermissionResponse `json:"permissions"`
}

// UserPermissionsResponse 用户权限响应
type UserPermissionsResponse struct {
	Roles       []string             `json:"roles"`
	Permissions []string             `json:"permissions"`
	Menus       []MenuResponse       `json:"menus"`
}
