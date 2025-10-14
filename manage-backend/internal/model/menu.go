package model

import (
	"time"

	"gorm.io/gorm"
)

// Menu 菜单模型
type Menu struct {
	ID             uint           `json:"id" gorm:"primarykey"`
	ParentID       *uint          `json:"parent_id" gorm:"index"`
	Name           string         `json:"name" gorm:"not null"`
	Title          string         `json:"title" gorm:"not null"`
	Path           string         `json:"path"`
	Component      string         `json:"component"`
	Icon           string         `json:"icon"`
	OrderNum       int            `json:"order_num" gorm:"default:0"`
	Type           string         `json:"type" gorm:"not null"`
	PermissionCode string         `json:"permission_code"`
	Visible        bool           `json:"visible" gorm:"default:true"`
	Status         string         `json:"status" gorm:"default:active"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
}

// MenuResponse 菜单响应结构体
type MenuResponse struct {
	ID             uint           `json:"id"`
	ParentID       *uint          `json:"parent_id"`
	Name           string         `json:"name"`
	Title          string         `json:"title"`
	Path           string         `json:"path"`
	Component      string         `json:"component"`
	Icon           string         `json:"icon"`
	OrderNum       int            `json:"order_num"`
	Type           string         `json:"type"`
	PermissionCode string         `json:"permission_code"`
	Visible        bool           `json:"visible"`
	Status         string         `json:"status"`
	Children       []MenuResponse `json:"children,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

// CreateMenuRequest 创建菜单请求
type CreateMenuRequest struct {
	ParentID       *uint  `json:"parent_id"`
	Name           string `json:"name" binding:"required,min=2,max=50"`
	Title          string `json:"title" binding:"required,min=2,max=100"`
	Path           string `json:"path" binding:"omitempty,max=255"`
	Component      string `json:"component" binding:"omitempty,max=255"`
	Icon           string `json:"icon" binding:"omitempty,max=50"`
	OrderNum       int    `json:"order_num"`
	Type           string `json:"type" binding:"required,oneof=menu button"`
	PermissionCode string `json:"permission_code" binding:"omitempty,max=100"`
	Visible        bool   `json:"visible"`
}

// UpdateMenuRequest 更新菜单请求
type UpdateMenuRequest struct {
	ParentID       *uint  `json:"parent_id"`
	Name           string `json:"name" binding:"omitempty,min=2,max=50"`
	Title          string `json:"title" binding:"omitempty,min=2,max=100"`
	Path           string `json:"path" binding:"omitempty,max=255"`
	Component      string `json:"component" binding:"omitempty,max=255"`
	Icon           string `json:"icon" binding:"omitempty,max=50"`
	OrderNum       int    `json:"order_num"`
	PermissionCode string `json:"permission_code" binding:"omitempty,max=100"`
	Visible        bool   `json:"visible"`
	Status         string `json:"status" binding:"omitempty,oneof=active inactive"`
}

// MenuOrderUpdate 菜单顺序更新
type MenuOrderUpdate struct {
	ID       uint  `json:"id" binding:"required"`
	OrderNum int   `json:"order_num" binding:"required"`
	ParentID *uint `json:"parent_id"`
}
