package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apperrors "github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/errors"
)

// APIResponse 统一的 API 响应结构
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PaginationMeta 分页元数据
type PaginationMeta struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// PaginatedResponse 分页响应结构
type PaginatedResponse struct {
	Code       int            `json:"code"`
	Message    string         `json:"message"`
	Data       interface{}    `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
	Error      string         `json:"error,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// SuccessWithMessage 带自定义消息的成功响应
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

// Created 创建成功响应
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, APIResponse{
		Code:    http.StatusCreated,
		Message: "created successfully",
		Data:    data,
	})
}

// BadRequest 400 错误响应
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, APIResponse{
		Code:    http.StatusBadRequest,
		Message: message,
		Error:   "bad request",
	})
}

// Unauthorized 401 错误响应
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, APIResponse{
		Code:    http.StatusUnauthorized,
		Message: message,
		Error:   "unauthorized",
	})
}

// Forbidden 403 错误响应
func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, APIResponse{
		Code:    http.StatusForbidden,
		Message: message,
		Error:   "forbidden",
	})
}

// NotFound 404 错误响应
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, APIResponse{
		Code:    http.StatusNotFound,
		Message: message,
		Error:   "not found",
	})
}

// Conflict 409 错误响应
func Conflict(c *gin.Context, message string) {
	c.JSON(http.StatusConflict, APIResponse{
		Code:    http.StatusConflict,
		Message: message,
		Error:   "conflict",
	})
}

// InternalServerError 500 错误响应
func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, APIResponse{
		Code:    http.StatusInternalServerError,
		Message: message,
		Error:   "internal server error",
	})
}

// PaginatedSuccess 分页成功响应
func PaginatedSuccess(c *gin.Context, data interface{}, pagination PaginationMeta) {
	c.JSON(http.StatusOK, PaginatedResponse{
		Code:       http.StatusOK,
		Message:    "success",
		Data:       data,
		Pagination: pagination,
	})
}

// SuccessWithPagination 带分页信息的成功响应
func SuccessWithPagination(c *gin.Context, data interface{}, page, pageSize, total int) {
	totalPages := (total + pageSize - 1) / pageSize
	if totalPages < 0 {
		totalPages = 0
	}

	c.JSON(http.StatusOK, PaginatedResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
		Pagination: PaginationMeta{
			Page:       page,
			PageSize:   pageSize,
			Total:      int64(total),
			TotalPages: totalPages,
		},
	})
}

// ValidationError 参数验证错误响应
func ValidationError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, APIResponse{
		Code:    http.StatusBadRequest,
		Message: "validation failed",
		Error:   err.Error(),
	})
}

// Locked 423 错误响应（账户锁定）
func Locked(c *gin.Context, message string) {
	c.JSON(http.StatusLocked, APIResponse{
		Code:    http.StatusLocked,
		Message: message,
		Error:   "locked",
	})
}

// TooManyRequests 429 错误响应（限流）
func TooManyRequests(c *gin.Context, message string) {
	c.JSON(http.StatusTooManyRequests, APIResponse{
		Code:    http.StatusTooManyRequests,
		Message: message,
		Error:   "too many requests",
	})
}

// HandleError 根据自定义错误类型自动返回对应的HTTP响应
// 这是一个通用的错误处理函数，会根据错误类型返回正确的HTTP状态码
func HandleError(c *gin.Context, err error) {
	// 尝试提取自定义错误
	if appErr, ok := apperrors.GetAppError(err); ok {
		// 根据错误的HTTP状态码返回对应的响应
		switch appErr.Code {
		case 400:
			BadRequest(c, appErr.Message)
		case 401:
			Unauthorized(c, appErr.Message)
		case 403:
			Forbidden(c, appErr.Message)
		case 404:
			NotFound(c, appErr.Message)
		case 409:
			Conflict(c, appErr.Message)
		case 423:
			Locked(c, appErr.Message)
		case 429:
			TooManyRequests(c, appErr.Message)
		case 500:
			InternalServerError(c, appErr.Message)
		default:
			// 未知状态码，默认返回500
			InternalServerError(c, appErr.Message)
		}
		return
	}
	
	// 如果不是自定义错误，默认返回500
	InternalServerError(c, err.Error())
}