package errors

import (
	"errors"
	"fmt"
)

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
	Type         ErrorType // 错误类型
	Message      string    // 错误消息
	Code         int       // HTTP状态码
	BusinessCode int       // 业务错误码
}

// Error 实现error接口
func (e *AppError) Error() string {
	if e.BusinessCode > 0 {
		return fmt.Sprintf("[%d] %s", e.BusinessCode, e.Message)
	}
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

// ==================== 专用错误构造函数（带业务错误码） ====================

// ========== 用户模块 ==========

// NewUsernameExistsError 用户名已存在
func NewUsernameExistsError() *AppError {
	code := CodeUsernameExists
	return &AppError{
		Type:         ErrorTypeConflict,
		Message:      GetBusinessCodeMessage(code),
		Code:         409,
		BusinessCode: code,
	}
}

// NewEmailExistsError 邮箱已存在
func NewEmailExistsError() *AppError {
	code := CodeEmailExists
	return &AppError{
		Type:         ErrorTypeConflict,
		Message:      GetBusinessCodeMessage(code),
		Code:         409,
		BusinessCode: code,
	}
}

// NewUserNotFoundErrorWithCode 用户不存在（带业务码）
func NewUserNotFoundErrorWithCode() *AppError {
	code := CodeUserNotFound
	return &AppError{
		Type:         ErrorTypeUserNotFound,
		Message:      GetBusinessCodeMessage(code),
		Code:         401,
		BusinessCode: code,
	}
}

// ========== 用户操作内部错误 ==========

// NewPasswordHashFailedError 密码加密失败
func NewPasswordHashFailedError() *AppError {
	code := CodePasswordHashFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewUserCreateFailedError 用户创建失败
func NewUserCreateFailedError() *AppError {
	code := CodeUserCreateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewUserQueryFailedError 查询用户失败
func NewUserQueryFailedError() *AppError {
	code := CodeUserQueryFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewUserUpdateFailedError 用户更新失败
func NewUserUpdateFailedError() *AppError {
	code := CodeUserUpdateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewUserDeleteFailedError 用户删除失败
func NewUserDeleteFailedError() *AppError {
	code := CodeUserDeleteFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewUserListFailedError 查询用户列表失败
func NewUserListFailedError() *AppError {
	code := CodeUserListFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewUsernameCheckFailedError 检查用户名失败
func NewUsernameCheckFailedError() *AppError {
	code := CodeUsernameCheckFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewEmailCheckFailedError 检查邮箱失败
func NewEmailCheckFailedError() *AppError {
	code := CodeEmailCheckFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewUserRoleSyncFailedError 同步用户角色失败
func NewUserRoleSyncFailedError() *AppError {
	code := CodeUserRoleSyncFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewRoleFindFailedError 查找角色失败
func NewRoleFindFailedError() *AppError {
	code := CodeRoleFindFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// ========== 会话相关内部错误 ==========

// NewTokenGenerateFailedError 生成令牌失败
func NewTokenGenerateFailedError() *AppError {
	code := CodeTokenGenerateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewSessionCreateFailedError 创建会话失败
func NewSessionCreateFailedError() *AppError {
	code := CodeSessionCreateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewSessionUpdateFailedError 更新会话失败
func NewSessionUpdateFailedError() *AppError {
	code := CodeSessionUpdateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewSessionDeleteFailedError 删除会话失败
func NewSessionDeleteFailedError() *AppError {
	code := CodeSessionDeleteFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewSessionServiceUnavailableError 会话服务不可用
func NewSessionServiceUnavailableError() *AppError {
	code := CodeSessionServiceUnavailable
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         503,
		BusinessCode: code,
	}
}

// NewTokenBlacklistFailedError 添加令牌到黑名单失败
func NewTokenBlacklistFailedError() *AppError {
	code := CodeTokenBlacklistFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewInvalidCredentialsErrorWithCode 用户名或密码错误（带业务码）
func NewInvalidCredentialsErrorWithCode(message string) *AppError {
	code := CodeInvalidCredentials
	if message == "" {
		message = GetBusinessCodeMessage(code)
	}
	return &AppError{
		Type:         ErrorTypeInvalidCredentials,
		Message:      message,
		Code:         401,
		BusinessCode: code,
	}
}

// NewAccountLockedErrorWithCode 账户已锁定（带业务码）
func NewAccountLockedErrorWithCode(message string) *AppError {
	code := CodeAccountLocked
	if message == "" {
		message = GetBusinessCodeMessage(code)
	}
	return &AppError{
		Type:         ErrorTypeAccountLocked,
		Message:      message,
		Code:         423,
		BusinessCode: code,
	}
}

// NewInvalidCaptchaErrorWithCode 验证码错误（带业务码）
func NewInvalidCaptchaErrorWithCode(message string) *AppError {
	code := CodeInvalidCaptcha
	if message == "" {
		message = GetBusinessCodeMessage(code)
	}
	return &AppError{
		Type:         ErrorTypeInvalidCaptcha,
		Message:      message,
		Code:         401,
		BusinessCode: code,
	}
}

// NewInvalidTokenErrorWithCode 无效的令牌（带业务码）
func NewInvalidTokenErrorWithCode(message string) *AppError {
	code := CodeInvalidToken
	if message == "" {
		message = GetBusinessCodeMessage(code)
	}
	return &AppError{
		Type:         ErrorTypeInvalidToken,
		Message:      message,
		Code:         400,
		BusinessCode: code,
	}
}

// NewUnauthorizedErrorWithCode 未授权（带业务码）
func NewUnauthorizedErrorWithCode(message string) *AppError {
	code := CodeUnauthorized
	if message == "" {
		message = GetBusinessCodeMessage(code)
	}
	return &AppError{
		Type:         ErrorTypeUnauthorized,
		Message:      message,
		Code:         401,
		BusinessCode: code,
	}
}

// ========== 系统模块 ==========

// NewRateLimitErrorWithCode 请求过于频繁（带业务码）
func NewRateLimitErrorWithCode(message string) *AppError {
	code := CodeRateLimitExceeded
	if message == "" {
		message = GetBusinessCodeMessage(code)
	}
	return &AppError{
		Type:         ErrorTypeRateLimit,
		Message:      message,
		Code:         429,
		BusinessCode: code,
	}
}

// NewInternalErrorWithCode 内部服务错误（带业务码）
func NewInternalErrorWithCode(message string) *AppError {
	code := CodeInternalError
	if message == "" {
		message = GetBusinessCodeMessage(code)
	}
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      message,
		Code:         500,
		BusinessCode: code,
	}
}

// NewBadRequestError 请求格式错误
func NewBadRequestError(message string) *AppError {
	code := CodeBadRequest
	if message == "" {
		message = GetBusinessCodeMessage(code)
	}
	return &AppError{
		Type:         ErrorTypeValidation,
		Message:      message,
		Code:         400,
		BusinessCode: code,
	}
}

// NewResourceNotFoundError 资源不存在
func NewResourceNotFoundError(message string) *AppError {
	code := CodeResourceNotFound
	if message == "" {
		message = GetBusinessCodeMessage(code)
	}
	return &AppError{
		Type:         ErrorTypeNotFound,
		Message:      message,
		Code:         404,
		BusinessCode: code,
	}
}

// NewPermissionDeniedErrorWithCode 权限不足（带业务码）
func NewPermissionDeniedErrorWithCode(message string) *AppError {
	code := CodePermissionDenied
	if message == "" {
		message = GetBusinessCodeMessage(code)
	}
	return &AppError{
		Type:         ErrorTypePermissionDenied,
		Message:      message,
		Code:         403,
		BusinessCode: code,
	}
}

// ========== 角色操作内部错误 ==========

// NewRoleCheckFailedError 检查角色代码失败
func NewRoleCheckFailedError() *AppError {
	code := CodeRoleCheckFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewRoleCreateFailedError 创建角色失败
func NewRoleCreateFailedError() *AppError {
	code := CodeRoleCreateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewRoleGetFailedError 获取角色失败
func NewRoleGetFailedError() *AppError {
	code := CodeRoleGetFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewRoleUpdateFailedError 更新角色失败
func NewRoleUpdateFailedError() *AppError {
	code := CodeRoleUpdateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewRoleDeleteFailedError 删除角色失败
func NewRoleDeleteFailedError() *AppError {
	code := CodeRoleDeleteFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewRoleListFailedError 获取角色列表失败
func NewRoleListFailedError() *AppError {
	code := CodeRoleListFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewRoleCheckUsageFailedError 检查角色使用情况失败
func NewRoleCheckUsageFailedError() *AppError {
	code := CodeRoleCheckUsageFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewRolePermissionDeleteFailedError 删除角色权限失败
func NewRolePermissionDeleteFailedError() *AppError {
	code := CodeRolePermissionDeleteFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewCasbinUpdateFailedError 更新Casbin角色权限失败
func NewCasbinUpdateFailedError() *AppError {
	code := CodeCasbinUpdateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewRolePermissionUpdateFailedError 更新数据库角色权限失败
func NewRolePermissionUpdateFailedError() *AppError {
	code := CodeRolePermissionUpdateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewRolePermissionGetFailedError 获取角色权限ID失败
func NewRolePermissionGetFailedError() *AppError {
	code := CodeRolePermissionGetFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewUserRoleRemoveFailedError 移除用户角色失败
func NewUserRoleRemoveFailedError() *AppError {
	code := CodeUserRoleRemoveFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewUserCasbinRoleRemoveFailedError 移除用户Casbin角色失败
func NewUserCasbinRoleRemoveFailedError() *AppError {
	code := CodeUserCasbinRoleRemoveFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewUserRoleAssignFailedError 分配用户角色失败
func NewUserRoleAssignFailedError() *AppError {
	code := CodeUserRoleAssignFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewUserCasbinRoleAddFailedError 添加用户Casbin角色失败
func NewUserCasbinRoleAddFailedError() *AppError {
	code := CodeUserCasbinRoleAddFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewUserRoleGetFailedError 获取用户角色失败
func NewUserRoleGetFailedError() *AppError {
	code := CodeUserRoleGetFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// ========== 权限操作内部错误 ==========

// NewPermissionCheckFailedError 检查权限代码失败
func NewPermissionCheckFailedError() *AppError {
	code := CodePermissionCheckFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPermissionCreateFailedError 创建权限失败
func NewPermissionCreateFailedError() *AppError {
	code := CodePermissionCreateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPermissionGetFailedError 获取权限失败
func NewPermissionGetFailedError() *AppError {
	code := CodePermissionGetFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPermissionUpdateFailedError 更新权限失败
func NewPermissionUpdateFailedError() *AppError {
	code := CodePermissionUpdateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPermissionDeleteFailedError 删除权限失败
func NewPermissionDeleteFailedError() *AppError {
	code := CodePermissionDeleteFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPermissionListFailedError 获取权限列表失败
func NewPermissionListFailedError() *AppError {
	code := CodePermissionListFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPermissionTreeFailedError 获取权限树失败
func NewPermissionTreeFailedError() *AppError {
	code := CodePermissionTreeFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPermissionByResourceFailedError 根据资源获取权限失败
func NewPermissionByResourceFailedError() *AppError {
	code := CodePermissionByResourceFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPermissionByTypeFailedError 根据类型获取权限失败
func NewPermissionByTypeFailedError() *AppError {
	code := CodePermissionByTypeFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// ========== 密码重置模块 ==========

// NewInvalidResetTokenError 无效的重置令牌
func NewInvalidResetTokenError() *AppError {
	code := CodeInvalidResetToken
	return &AppError{
		Type:         ErrorTypeInvalidToken,
		Message:      GetBusinessCodeMessage(code),
		Code:         400,
		BusinessCode: code,
	}
}

// NewResetTokenExpiredError 重置令牌已过期
func NewResetTokenExpiredError() *AppError {
	code := CodeResetTokenExpired
	return &AppError{
		Type:         ErrorTypeTokenExpired,
		Message:      GetBusinessCodeMessage(code),
		Code:         400,
		BusinessCode: code,
	}
}

// ========== 密码重置操作内部错误 ==========

// NewPasswordResetUserQueryFailedError 查询用户失败
func NewPasswordResetUserQueryFailedError() *AppError {
	code := CodePasswordResetUserQueryFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPasswordResetTokenCleanFailedError 清理旧Token失败
func NewPasswordResetTokenCleanFailedError() *AppError {
	code := CodePasswordResetTokenCleanFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPasswordResetTokenCreateFailedError 创建重置Token失败
func NewPasswordResetTokenCreateFailedError() *AppError {
	code := CodePasswordResetTokenCreateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPasswordResetEmailSendFailedError 发送重置邮件失败
func NewPasswordResetEmailSendFailedError() *AppError {
	code := CodePasswordResetEmailSendFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPasswordResetTokenQueryFailedError 查询Token失败
func NewPasswordResetTokenQueryFailedError() *AppError {
	code := CodePasswordResetTokenQueryFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPasswordResetHashFailedError 密码加密失败
func NewPasswordResetHashFailedError() *AppError {
	code := CodePasswordResetHashFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPasswordResetUpdateFailedError 更新密码失败
func NewPasswordResetUpdateFailedError() *AppError {
	code := CodePasswordResetUpdateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// ========== 菜单操作内部错误 ==========

// NewMenuParentGetFailedError 查询父菜单失败
func NewMenuParentGetFailedError() *AppError {
	code := CodeMenuParentGetFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewMenuPermissionGetFailedError 查询权限失败
func NewMenuPermissionGetFailedError() *AppError {
	code := CodeMenuPermissionGetFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewMenuCreateFailedError 创建菜单失败
func NewMenuCreateFailedError() *AppError {
	code := CodeMenuCreateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewMenuGetFailedError 获取菜单失败
func NewMenuGetFailedError() *AppError {
	code := CodeMenuGetFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewMenuUpdateFailedError 更新菜单失败
func NewMenuUpdateFailedError() *AppError {
	code := CodeMenuUpdateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewMenuDeleteFailedError 删除菜单失败
func NewMenuDeleteFailedError() *AppError {
	code := CodeMenuDeleteFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewMenuCheckChildrenFailedError 检查子菜单失败
func NewMenuCheckChildrenFailedError() *AppError {
	code := CodeMenuCheckChildrenFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewMenuListFailedError 获取菜单列表失败
func NewMenuListFailedError() *AppError {
	code := CodeMenuListFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewMenuVisibleListFailedError 获取可见菜单失败
func NewMenuVisibleListFailedError() *AppError {
	code := CodeMenuVisibleListFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewMenuOrderUpdateFailedError 更新菜单顺序失败
func NewMenuOrderUpdateFailedError() *AppError {
	code := CodeMenuOrderUpdateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// ========== 角色权限模块 ==========

// NewRoleNotFoundError 角色不存在
func NewRoleNotFoundError() *AppError {
	code := CodeRoleNotFound
	return &AppError{
		Type:         ErrorTypeNotFound,
		Message:      GetBusinessCodeMessage(code),
		Code:         404,
		BusinessCode: code,
	}
}

// NewRoleExistsError 角色已存在
func NewRoleExistsError() *AppError {
	code := CodeRoleExists
	return &AppError{
		Type:         ErrorTypeConflict,
		Message:      GetBusinessCodeMessage(code),
		Code:         409,
		BusinessCode: code,
	}
}

// NewRoleInUseError 角色正在使用中
func NewRoleInUseError(userCount int) *AppError {
	return &AppError{
		Type:         ErrorTypeConflict,
		Message:      fmt.Sprintf("该角色正在被 %d 个用户使用，无法删除", userCount),
		Code:         409,
		BusinessCode: CodeRoleInUse,
	}
}

// ========== 字典类型操作内部错误 ==========

// NewDictTypeCheckFailedError 检查字典类型代码失败
func NewDictTypeCheckFailedError() *AppError {
	code := CodeDictTypeCheckFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictTypeCreateFailedError 创建字典类型失败
func NewDictTypeCreateFailedError() *AppError {
	code := CodeDictTypeCreateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictTypeGetFailedError 查询字典类型失败
func NewDictTypeGetFailedError() *AppError {
	code := CodeDictTypeGetFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictTypeUpdateFailedError 更新字典类型失败
func NewDictTypeUpdateFailedError() *AppError {
	code := CodeDictTypeUpdateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictTypeDeleteFailedError 删除字典类型失败
func NewDictTypeDeleteFailedError() *AppError {
	code := CodeDictTypeDeleteFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictTypeListFailedError 查询字典类型列表失败
func NewDictTypeListFailedError() *AppError {
	code := CodeDictTypeListFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictTypeCountFailedError 统计字典项数量失败
func NewDictTypeCountFailedError() *AppError {
	code := CodeDictTypeCountFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// ========== 字典项操作内部错误 ==========

// NewDictItemTypeGetFailedError 查询字典类型失败
func NewDictItemTypeGetFailedError() *AppError {
	code := CodeDictItemTypeGetFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictItemValueCheckFailedError 检查字典项值失败
func NewDictItemValueCheckFailedError() *AppError {
	code := CodeDictItemValueCheckFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictItemDefaultClearFailedError 清除默认值失败
func NewDictItemDefaultClearFailedError() *AppError {
	code := CodeDictItemDefaultClearFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictItemCreateFailedError 创建字典项失败
func NewDictItemCreateFailedError() *AppError {
	code := CodeDictItemCreateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictItemGetFailedError 查询字典项失败
func NewDictItemGetFailedError() *AppError {
	code := CodeDictItemGetFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictItemUpdateFailedError 更新字典项失败
func NewDictItemUpdateFailedError() *AppError {
	code := CodeDictItemUpdateFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictItemDeleteFailedError 删除字典项失败
func NewDictItemDeleteFailedError() *AppError {
	code := CodeDictItemDeleteFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictItemListFailedError 查询字典项列表失败
func NewDictItemListFailedError() *AppError {
	code := CodeDictItemListFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewDictItemsByTypeGetFailedError 根据类型获取字典项失败
func NewDictItemsByTypeGetFailedError() *AppError {
	code := CodeDictItemsByTypeGetFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewPermissionNotFoundError 权限不存在
func NewPermissionNotFoundError(message string) *AppError {
	code := CodePermissionNotFound
	if message == "" {
		message = GetBusinessCodeMessage(code)
	}
	return &AppError{
		Type:         ErrorTypeNotFound,
		Message:      message,
		Code:         404,
		BusinessCode: code,
	}
}

// ========== 菜单模块 ==========

// NewMenuNotFoundError 菜单不存在
func NewMenuNotFoundError() *AppError {
	code := CodeMenuNotFound
	return &AppError{
		Type:         ErrorTypeNotFound,
		Message:      GetBusinessCodeMessage(code),
		Code:         404,
		BusinessCode: code,
	}
}

// NewMenuHasChildrenError 菜单存在子菜单
func NewMenuHasChildrenError() *AppError {
	code := CodeMenuHasChildren
	return &AppError{
		Type:         ErrorTypeConflict,
		Message:      GetBusinessCodeMessage(code),
		Code:         409,
		BusinessCode: code,
	}
}

// NewMenuInvalidParentError 无效的父菜单（通用）
func NewMenuInvalidParentError(message string) *AppError {
	code := CodeMenuInvalidParent
	if message == "" {
		message = GetBusinessCodeMessage(code)
	}
	return &AppError{
		Type:         ErrorTypeValidation,
		Message:      message,
		Code:         400,
		BusinessCode: code,
	}
}

// NewMenuCannotBeOwnChildError 不能将菜单设置为自己的子菜单
func NewMenuCannotBeOwnChildError() *AppError {
	code := CodeMenuCannotBeOwnChild
	return &AppError{
		Type:         ErrorTypeValidation,
		Message:      GetBusinessCodeMessage(code),
		Code:         400,
		BusinessCode: code,
	}
}

// NewMenuParentNotFoundError 父菜单不存在
func NewMenuParentNotFoundError() *AppError {
	code := CodeMenuParentNotFound
	return &AppError{
		Type:         ErrorTypeNotFound,
		Message:      GetBusinessCodeMessage(code),
		Code:         404,
		BusinessCode: code,
	}
}

// ========== 字典模块 ==========

// NewDictTypeNotFoundError 字典类型不存在
func NewDictTypeNotFoundError() *AppError {
	code := CodeDictTypeNotFound
	return &AppError{
		Type:         ErrorTypeNotFound,
		Message:      GetBusinessCodeMessage(code),
		Code:         404,
		BusinessCode: code,
	}
}

// NewDictTypeExistsError 字典类型已存在
func NewDictTypeExistsError() *AppError {
	code := CodeDictTypeExists
	return &AppError{
		Type:         ErrorTypeConflict,
		Message:      GetBusinessCodeMessage(code),
		Code:         409,
		BusinessCode: code,
	}
}

// NewDictTypeInUseError 字典类型正在使用中
func NewDictTypeInUseError(count int64) *AppError {
	return &AppError{
		Type:         ErrorTypeConflict,
		Message:      fmt.Sprintf("字典类型下存在 %d 个字典项，请先删除字典项", count),
		Code:         409,
		BusinessCode: CodeDictTypeInUse,
	}
}

// NewDictItemNotFoundError 字典项不存在
func NewDictItemNotFoundError() *AppError {
	code := CodeDictItemNotFound
	return &AppError{
		Type:         ErrorTypeNotFound,
		Message:      GetBusinessCodeMessage(code),
		Code:         404,
		BusinessCode: code,
	}
}

// NewDictItemExistsError 字典项已存在
func NewDictItemExistsError() *AppError {
	code := CodeDictItemExists
	return &AppError{
		Type:         ErrorTypeConflict,
		Message:      GetBusinessCodeMessage(code),
		Code:         409,
		BusinessCode: code,
	}
}

// ========== 审计日志模块错误 ==========

// NewAuditLogNotFoundError 审计日志不存在
func NewAuditLogNotFoundError() *AppError {
	code := CodeAuditLogNotFound
	return &AppError{
		Type:         ErrorTypeNotFound,
		Message:      GetBusinessCodeMessage(code),
		Code:         404,
		BusinessCode: code,
	}
}

// NewAuditLogGetFailedError 获取审计日志失败
func NewAuditLogGetFailedError() *AppError {
	code := CodeAuditLogGetFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewAuditLogQueryFailedError 查询审计日志失败
func NewAuditLogQueryFailedError() *AppError {
	code := CodeAuditLogQueryFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}

// NewAuditLogCleanFailedError 清理审计日志失败
func NewAuditLogCleanFailedError() *AppError {
	code := CodeAuditLogCleanFailed
	return &AppError{
		Type:         ErrorTypeInternal,
		Message:      GetBusinessCodeMessage(code),
		Code:         500,
		BusinessCode: code,
	}
}
