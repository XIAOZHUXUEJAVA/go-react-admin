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
	ErrorTypeValidation         // 参数验证错误
	ErrorTypeNotFound           // 资源不存在
	ErrorTypeConflict           // 资源冲突（如用户名已存在）
	ErrorTypePermissionDenied   // 权限不足
	ErrorTypeInvalidToken       // Token无效
	ErrorTypeTokenExpired       // Token已过期
	ErrorTypeTokenUsed          // Token已使用
	ErrorTypeUnauthorized       // 未授权
	ErrorTypeInternal           // 内部错误
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

// NewValidationError 创建参数验证错误
func NewValidationError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeValidation,
		Message: message,
		Code:    400, // Bad Request
	}
}

// NewNotFoundError 创建资源不存在错误
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeNotFound,
		Message: message,
		Code:    404, // Not Found
	}
}

// NewConflictError 创建资源冲突错误
func NewConflictError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeConflict,
		Message: message,
		Code:    409, // Conflict
	}
}

// NewPermissionDeniedError 创建权限不足错误
func NewPermissionDeniedError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypePermissionDenied,
		Message: message,
		Code:    403, // Forbidden
	}
}

// NewInvalidTokenError 创建Token无效错误
func NewInvalidTokenError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeInvalidToken,
		Message: message,
		Code:    400, // Bad Request
	}
}

// NewTokenExpiredError 创建Token过期错误
func NewTokenExpiredError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeTokenExpired,
		Message: message,
		Code:    400, // Bad Request
	}
}

// NewTokenUsedError 创建Token已使用错误
func NewTokenUsedError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeTokenUsed,
		Message: message,
		Code:    400, // Bad Request
	}
}

// NewUnauthorizedError 创建未授权错误
func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeUnauthorized,
		Message: message,
		Code:    401, // Unauthorized
	}
}

// NewInternalError 创建内部错误
func NewInternalError(message string) *AppError {
	return &AppError{
		Type:    ErrorTypeInternal,
		Message: message,
		Code:    500, // Internal Server Error
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

// IsValidationError 判断是否是参数验证错误
func IsValidationError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypeValidation
}

// IsNotFoundError 判断是否是资源不存在错误
func IsNotFoundError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypeNotFound
}

// IsConflictError 判断是否是资源冲突错误
func IsConflictError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypeConflict
}

// IsPermissionDeniedError 判断是否是权限不足错误
func IsPermissionDeniedError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypePermissionDenied
}

// IsInvalidTokenError 判断是否是Token无效错误
func IsInvalidTokenError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypeInvalidToken
}

// IsTokenExpiredError 判断是否是Token过期错误
func IsTokenExpiredError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypeTokenExpired
}

// IsTokenUsedError 判断是否是Token已使用错误
func IsTokenUsedError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ErrorTypeTokenUsed
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
