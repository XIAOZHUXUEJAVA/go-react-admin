package model

import (
	"time"

	"gorm.io/gorm"
)

// AuditLog 审计日志模型
type AuditLog struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	UserID      uint           `json:"user_id" gorm:"index"`
	Username    string         `json:"username"`
	Action      string         `json:"action" gorm:"not null"`
	Resource    string         `json:"resource"`
	ResourceID  string         `json:"resource_id"`
	Method      string         `json:"method"`
	Path        string         `json:"path"`
	IP          string         `json:"ip"`
	UserAgent   string         `json:"user_agent"`
	Status      int            `json:"status"`
	ErrorMsg    string         `json:"error_msg"`
	RequestBody string         `json:"request_body" gorm:"type:text"`
	Duration    int64          `json:"duration"` // 请求耗时（毫秒）
	CreatedAt   time.Time      `json:"created_at" gorm:"index"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

// AuditLogResponse 审计日志响应结构体
type AuditLogResponse struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Username    string    `json:"username"`
	Action      string    `json:"action"`
	Resource    string    `json:"resource"`
	ResourceID  string    `json:"resource_id"`
	Method      string    `json:"method"`
	Path        string    `json:"path"`
	IP          string    `json:"ip"`
	UserAgent   string    `json:"user_agent"`
	Status      int       `json:"status"`
	ErrorMsg    string    `json:"error_msg"`
	RequestBody string    `json:"request_body"`
	Duration    int64     `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
}

// AuditLogQuery 审计日志查询参数
type AuditLogQuery struct {
	UserID     *uint     `json:"user_id" form:"user_id"`
	Username   string    `json:"username" form:"username"`
	Action     string    `json:"action" form:"action"`
	Resource   string    `json:"resource" form:"resource"`
	Method     string    `json:"method" form:"method"`
	Status     *int      `json:"status" form:"status"`
	StartTime  time.Time `json:"start_time" form:"start_time"`
	EndTime    time.Time `json:"end_time" form:"end_time"`
	Page       int       `json:"page" form:"page"`
	PageSize   int       `json:"page_size" form:"page_size"`
}
