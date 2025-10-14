package handler

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/service"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
)

type AuditLogHandler struct {
	auditLogService *service.AuditLogService
}

func NewAuditLogHandler(auditLogService *service.AuditLogService) *AuditLogHandler {
	return &AuditLogHandler{auditLogService: auditLogService}
}

// GetAuditLog godoc
// @Summary 获取审计日志详情
// @Description 根据ID获取审计日志详情
// @Tags audit-logs
// @Produce json
// @Security BearerAuth
// @Param id path int true "审计日志ID"
// @Success 200 {object} utils.APIResponse{data=model.AuditLogResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /audit-logs/{id} [get]
func (h *AuditLogHandler) GetAuditLog(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的审计日志ID")
		return
	}

	log, err := h.auditLogService.GetByID(uint(id))
	if err != nil {
		utils.NotFound(c, "审计日志不存在")
		return
	}

	utils.Success(c, log)
}

// QueryAuditLogs godoc
// @Summary 查询审计日志
// @Description 根据条件查询审计日志列表
// @Tags audit-logs
// @Produce json
// @Security BearerAuth
// @Param user_id query int false "用户ID"
// @Param username query string false "用户名"
// @Param action query string false "操作"
// @Param resource query string false "资源"
// @Param method query string false "HTTP方法"
// @Param status query int false "状态码"
// @Param start_time query string false "开始时间" format(date-time)
// @Param end_time query string false "结束时间" format(date-time)
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} utils.PaginatedResponse{data=[]model.AuditLogResponse}
// @Failure 400 {object} utils.APIResponse
// @Router /audit-logs [get]
func (h *AuditLogHandler) QueryAuditLogs(c *gin.Context) {
	var query model.AuditLogQuery

	// 解析查询参数
	if userIDStr := c.Query("user_id"); userIDStr != "" {
		userID, err := strconv.ParseUint(userIDStr, 10, 32)
		if err == nil {
			uid := uint(userID)
			query.UserID = &uid
		}
	}

	query.Username = c.Query("username")
	query.Action = c.Query("action")
	query.Resource = c.Query("resource")
	query.Method = c.Query("method")

	if statusStr := c.Query("status"); statusStr != "" {
		status, err := strconv.Atoi(statusStr)
		if err == nil {
			query.Status = &status
		}
	}

	// 解析时间
	if startTimeStr := c.Query("start_time"); startTimeStr != "" {
		if t, err := time.Parse(time.RFC3339, startTimeStr); err == nil {
			query.StartTime = t
		}
	}

	if endTimeStr := c.Query("end_time"); endTimeStr != "" {
		if t, err := time.Parse(time.RFC3339, endTimeStr); err == nil {
			query.EndTime = t
		}
	}

	// 分页参数
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	query.Page = page
	query.PageSize = pageSize

	// 查询
	logs, total, err := h.auditLogService.Query(&query)
	if err != nil {
		utils.InternalServerError(c, "查询审计日志失败")
		return
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))
	pagination := utils.PaginationMeta{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}

	utils.PaginatedSuccess(c, logs, pagination)
}

// CleanOldAuditLogs godoc
// @Summary 清理旧审计日志
// @Description 清理指定天数之前的审计日志
// @Tags audit-logs
// @Produce json
// @Security BearerAuth
// @Param days query int false "保留天数" default(90)
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /audit-logs/clean [post]
func (h *AuditLogHandler) CleanOldAuditLogs(c *gin.Context) {
	daysStr := c.DefaultQuery("days", "90")
	days, err := strconv.Atoi(daysStr)
	if err != nil || days < 1 {
		utils.BadRequest(c, "无效的天数参数")
		return
	}

	if err := h.auditLogService.CleanOldLogs(days); err != nil {
		utils.InternalServerError(c, "清理审计日志失败")
		return
	}

	utils.Success(c, gin.H{"message": "审计日志清理成功"})
}
