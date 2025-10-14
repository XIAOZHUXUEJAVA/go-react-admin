package service

import (
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// AuditLogRepositoryInterface 定义审计日志仓库接口
type AuditLogRepositoryInterface interface {
	Create(log *model.AuditLog) error
	GetByID(id uint) (*model.AuditLog, error)
	Query(query *model.AuditLogQuery) ([]model.AuditLog, int64, error)
	DeleteOldLogs(days int) error
}

// AuditLogService 审计日志业务服务
type AuditLogService struct {
	auditLogRepo AuditLogRepositoryInterface
}

// NewAuditLogService 创建 AuditLogService 实例
func NewAuditLogService(auditLogRepo AuditLogRepositoryInterface) *AuditLogService {
	return &AuditLogService{
		auditLogRepo: auditLogRepo,
	}
}

// GetByID 根据ID获取审计日志
func (s *AuditLogService) GetByID(id uint) (*model.AuditLogResponse, error) {
	log, err := s.auditLogRepo.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		logger.Error("获取审计日志失败", zap.Uint("id", id), zap.Error(err))
		return nil, err
	}

	return s.toAuditLogResponse(log), nil
}

// Query 查询审计日志
func (s *AuditLogService) Query(query *model.AuditLogQuery) ([]model.AuditLogResponse, int64, error) {
	// 设置默认值
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 {
		query.PageSize = 10
	}
	if query.PageSize > 100 {
		query.PageSize = 100
	}

	logs, total, err := s.auditLogRepo.Query(query)
	if err != nil {
		logger.Error("查询审计日志失败", zap.Error(err))
		return nil, 0, err
	}

	responses := make([]model.AuditLogResponse, len(logs))
	for i, log := range logs {
		responses[i] = *s.toAuditLogResponse(&log)
	}

	return responses, total, nil
}

// CleanOldLogs 清理旧的审计日志
func (s *AuditLogService) CleanOldLogs(days int) error {
	if days < 1 {
		days = 90 // 默认保留90天
	}

	err := s.auditLogRepo.DeleteOldLogs(days)
	if err != nil {
		logger.Error("清理旧审计日志失败", zap.Int("days", days), zap.Error(err))
		return err
	}

	logger.Info("清理旧审计日志成功", zap.Int("days", days))
	return nil
}

// toAuditLogResponse 转换为响应结构
func (s *AuditLogService) toAuditLogResponse(log *model.AuditLog) *model.AuditLogResponse {
	return &model.AuditLogResponse{
		ID:          log.ID,
		UserID:      log.UserID,
		Username:    log.Username,
		Action:      log.Action,
		Resource:    log.Resource,
		ResourceID:  log.ResourceID,
		Method:      log.Method,
		Path:        log.Path,
		IP:          log.IP,
		UserAgent:   log.UserAgent,
		Status:      log.Status,
		ErrorMsg:    log.ErrorMsg,
		RequestBody: log.RequestBody,
		Duration:    log.Duration,
		CreatedAt:   log.CreatedAt,
	}
}
