package errors

import "errors"

// ErrorType 错误类型枚举
type ErrorType int

const (
	ErrorTypeUnknown ErrorType = iota
	ErrorTypeRateLimit          // 限流错误
	ErrorTypeAccountLocked      // 账户锁定
	ErrorTypeInvalidCredentials // 凭证错误（用户名或密码错误）
	ErrorTypeInvalidCaptcha     // 验证码错误
	ErrorTypeUserNotFound       // 用户不存在
	ErrorTypeAccountDisabled    // 账户已禁用
)

// AppError 应用错误
type AppError struct {
	Type    ErrorType // 错误类型
	Message string    // 错误消息
	Code    int       // HTTP状态码
}

// Error 实现error接口
func (e *AppError) Error() string {
	return e.Message
}

// ==================== 构造函数 ====================

// NewRateLimitError 创建限流错误
func NewRateLimitError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeRateLimit,
		Message: message,
		Code:    429, // Too Many Requests
	}
}

// NewAccountLockedError 创建账户锁定错误
func NewAccountLockedError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeAccountLocked,
		Message: message,
		Code:    423, // Locked
	}
}

// NewInvalidCredentialsError 创建凭证错误
func NewInvalidCredentialsError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeInvalidCredentials,
		Message: message,
		Code:    401, // Unauthorized
	}
}

// NewInvalidCaptchaError 创建验证码错误
func NewInvalidCaptchaError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeInvalidCaptcha,
		Message: message,
		Code:    401, // Unauthorized
	}
}

// NewUserNotFoundError 创建用户不存在错误
func NewUserNotFoundError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeUserNotFound,
		Message: message,
		Code:    401, // Unauthorized (安全实践：不暴露用户是否存在)
	}
}

// NewAccountDisabledError 创建账户禁用错误
func NewAccountDisabledError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeAccountDisabled,
		Message: message,
		Code:    403, // Forbidden
	}
}

// ==================== 类型判断函数 ====================

// IsRateLimitError 判断是否是限流错误
func IsRateLimitError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypeRateLimit
}

// IsAccountLockedError 判断是否是账户锁定错误
func IsAccountLockedError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypeAccountLocked
}

// IsInvalidCredentialsError 判断是否是凭证错误
func IsInvalidCredentialsError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypeInvalidCredentials
}

// IsInvalidCaptchaError 判断是否是验证码错误
func IsInvalidCaptchaError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypeInvalidCaptcha
}

// IsUserNotFoundError 判断是否是用户不存在错误
func IsUserNotFoundError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypeUserNotFound
}

// IsAccountDisabledError 判断是否是账户禁用错误
func IsAccountDisabledError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypeAccountDisabled
}

// ==================== 辅助函数 ====================

// GetAppError 从error中提取AppError
func GetAppError(err error) (*AppError, bool) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}

// GetHTTPCode 获取错误对应的HTTP状态码
func GetHTTPCode(err error) int {
	if appErr, ok := GetAppError(err); ok {
		return appErr.Code
	}
	return 500 // Internal Server Error
}

// GetErrorType 获取错误类型
func GetErrorType(err error) ErrorType {
	if appErr, ok := GetAppError(err); ok {
		return appErr.Type
	}
	return ErrorTypeUnknown
}
