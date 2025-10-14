package repository

import (
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"gorm.io/gorm"
)

// AuditLogRepository 审计日志数据仓库
type AuditLogRepository struct {
	db *gorm.DB
}

// NewAuditLogRepository 创建 AuditLogRepository 实例
func NewAuditLogRepository(db *gorm.DB) *AuditLogRepository {
	return &AuditLogRepository{db: db}
}

// Create 创建审计日志
func (r *AuditLogRepository) Create(log *model.AuditLog) error {
	return r.db.Create(log).Error
}

// GetByID 根据ID获取审计日志
func (r *AuditLogRepository) GetByID(id uint) (*model.AuditLog, error) {
	var log model.AuditLog
	err := r.db.First(&log, id).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// Query 根据条件查询审计日志
func (r *AuditLogRepository) Query(query *model.AuditLogQuery) ([]model.AuditLog, int64, error) {
	var logs []model.AuditLog
	var total int64

	db := r.db.Model(&model.AuditLog{})

	// 构建查询条件
	if query.UserID != nil {
		db = db.Where("user_id = ?", *query.UserID)
	}
	if query.Username != "" {
		db = db.Where("username LIKE ?", "%"+query.Username+"%")
	}
	if query.Action != "" {
		db = db.Where("action LIKE ?", "%"+query.Action+"%")
	}
	if query.Resource != "" {
		db = db.Where("resource = ?", query.Resource)
	}
	if query.Method != "" {
		db = db.Where("method = ?", query.Method)
	}
	if query.Status != nil {
		db = db.Where("status = ?", *query.Status)
	}
	if !query.StartTime.IsZero() {
		db = db.Where("created_at >= ?", query.StartTime)
	}
	if !query.EndTime.IsZero() {
		db = db.Where("created_at <= ?", query.EndTime)
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (query.Page - 1) * query.PageSize
	err := db.Order("created_at DESC").
		Offset(offset).
		Limit(query.PageSize).
		Find(&logs).Error

	return logs, total, err
}

// DeleteOldLogs 删除旧的审计日志
func (r *AuditLogRepository) DeleteOldLogs(days int) error {
	return r.db.Where("created_at < NOW() - INTERVAL ? DAY", days).
		Delete(&model.AuditLog{}).Error
}
