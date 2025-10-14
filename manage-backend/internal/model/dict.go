package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// DictType 字典类型模型
type DictType struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Code        string         `json:"code" gorm:"uniqueIndex;not null"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Status      string         `json:"status" gorm:"default:active"`
	SortOrder   int            `json:"sort_order" gorm:"default:0"`
	IsSystem    bool           `json:"is_system" gorm:"default:false"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// DictTypeResponse 字典类型响应结构体
type DictTypeResponse struct {
	ID          uint      `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	SortOrder   int       `json:"sort_order"`
	IsSystem    bool      `json:"is_system"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateDictTypeRequest 创建字典类型请求
type CreateDictTypeRequest struct {
	Code        string `json:"code" binding:"required,min=2,max=50"`
	Name        string `json:"name" binding:"required,min=2,max=100"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Status      string `json:"status" binding:"omitempty,oneof=active inactive"`
	SortOrder   int    `json:"sort_order" binding:"omitempty,min=0"`
}

// UpdateDictTypeRequest 更新字典类型请求
type UpdateDictTypeRequest struct {
	Name        string `json:"name" binding:"omitempty,min=2,max=100"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Status      string `json:"status" binding:"omitempty,oneof=active inactive"`
	SortOrder   int    `json:"sort_order" binding:"omitempty,min=0"`
}

// DictItem 字典项模型
type DictItem struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	DictTypeCode string         `json:"dict_type_code" gorm:"not null;index"`
	Label        string         `json:"label" gorm:"not null"`
	Value        string         `json:"value" gorm:"not null"`
	Extra        datatypes.JSON `json:"extra" gorm:"type:jsonb"`
	Description  string         `json:"description"`
	Status       string         `json:"status" gorm:"default:active"`
	SortOrder    int            `json:"sort_order" gorm:"default:0"`
	IsDefault    bool           `json:"is_default" gorm:"default:false"`
	IsSystem     bool           `json:"is_system" gorm:"default:false"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

// DictItemResponse 字典项响应结构体
type DictItemResponse struct {
	ID           uint                   `json:"id"`
	DictTypeCode string                 `json:"dict_type_code"`
	Label        string                 `json:"label"`
	Value        string                 `json:"value"`
	Extra        map[string]interface{} `json:"extra,omitempty"`
	Description  string                 `json:"description"`
	Status       string                 `json:"status"`
	SortOrder    int                    `json:"sort_order"`
	IsDefault    bool                   `json:"is_default"`
	IsSystem     bool                   `json:"is_system"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
}

// CreateDictItemRequest 创建字典项请求
type CreateDictItemRequest struct {
	DictTypeCode string                 `json:"dict_type_code" binding:"required,min=2,max=50"`
	Label        string                 `json:"label" binding:"required,min=1,max=100"`
	Value        string                 `json:"value" binding:"required,min=1,max=100"`
	Extra        map[string]interface{} `json:"extra" binding:"omitempty"`
	Description  string                 `json:"description" binding:"omitempty,max=255"`
	Status       string                 `json:"status" binding:"omitempty,oneof=active inactive"`
	SortOrder    int                    `json:"sort_order" binding:"omitempty,min=0"`
	IsDefault    bool                   `json:"is_default" binding:"omitempty"`
}

// UpdateDictItemRequest 更新字典项请求
type UpdateDictItemRequest struct {
	Label       string                 `json:"label" binding:"omitempty,min=1,max=100"`
	Extra       map[string]interface{} `json:"extra" binding:"omitempty"`
	Description string                 `json:"description" binding:"omitempty,max=255"`
	Status      string                 `json:"status" binding:"omitempty,oneof=active inactive"`
	SortOrder   int                    `json:"sort_order" binding:"omitempty,min=0"`
	IsDefault   bool                   `json:"is_default" binding:"omitempty"`
}

// DictTypeListRequest 字典类型列表查询请求
type DictTypeListRequest struct {
	Page     int    `form:"page" binding:"omitempty,min=1"`
	PageSize int    `form:"page_size" binding:"omitempty,min=1,max=100"`
	Status   string `form:"status" binding:"omitempty,oneof=active inactive"`
	Keyword  string `form:"keyword" binding:"omitempty,max=50"`
}

// DictItemListRequest 字典项列表查询请求
type DictItemListRequest struct {
	Page         int    `form:"page" binding:"omitempty,min=1"`
	PageSize     int    `form:"page_size" binding:"omitempty,min=1,max=100"`
	DictTypeCode string `form:"dict_type_code" binding:"omitempty,max=50"`
	Status       string `form:"status" binding:"omitempty,oneof=active inactive"`
}
