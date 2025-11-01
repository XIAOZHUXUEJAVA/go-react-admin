package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		Message: "bad request",
		Error:   message,
	})
}

// Unauthorized 401 错误响应
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, APIResponse{
		Code:    http.StatusUnauthorized,
		Message: "unauthorized",
		Error:   message,
	})
}

// Forbidden 403 错误响应
func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, APIResponse{
		Code:    http.StatusForbidden,
		Message: "forbidden",
		Error:   message,
	})
}

// NotFound 404 错误响应
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, APIResponse{
		Code:    http.StatusNotFound,
		Message: "not found",
		Error:   message,
	})
}

// Conflict 409 错误响应
func Conflict(c *gin.Context, message string) {
	c.JSON(http.StatusConflict, APIResponse{
		Code:    http.StatusConflict,
		Message: "conflict",
		Error:   message,
	})
}

// InternalServerError 500 错误响应
func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, APIResponse{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
		Error:   message,
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
		Message: "locked",
		Error:   message,
	})
}

// TooManyRequests 429 错误响应（限流）
func TooManyRequests(c *gin.Context, message string) {
	c.JSON(http.StatusTooManyRequests, APIResponse{
		Code:    http.StatusTooManyRequests,
		Message: "too many requests",
		Error:   message,
	})
}