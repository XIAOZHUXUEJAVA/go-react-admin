package errors

// 业务错误码定义
// 采用5位数字分段: [模块(2位)][类别(1位)][序号(2位)]
const (
	// 成功
	CodeSuccess = 0

	// ========== 用户模块 (10xxx) ==========

	// 认证相关 (101xx)
	CodeInvalidCredentials = 10101 // 用户名或密码错误
	CodeAccountLocked      = 10102 // 账户已锁定
	CodeInvalidCaptcha     = 10103 // 验证码错误
	CodeInvalidToken       = 10104 // 无效的令牌
	CodeTokenExpired       = 10105 // 令牌已过期
	CodeUnauthorized       = 10106 // 未授权

	// 注册相关 (102xx)
	CodeUsernameExists = 10201 // 用户名已存在
	CodeEmailExists    = 10202 // 邮箱已存在
	CodeWeakPassword   = 10203 // 密码强度不足

	// 用户信息 (103xx)
	CodeUserNotFound = 10301 // 用户不存在
	CodeUserDisabled = 10302 // 用户已禁用

	// 用户操作内部错误 (104xx)
	CodePasswordHashFailed    = 10401 // 密码加密失败
	CodeUserCreateFailed      = 10402 // 用户创建失败
	CodeUserQueryFailed       = 10403 // 查询用户失败
	CodeUserUpdateFailed      = 10404 // 用户更新失败
	CodeUserDeleteFailed      = 10405 // 用户删除失败
	CodeUserListFailed        = 10406 // 查询用户列表失败
	CodeUsernameCheckFailed   = 10407 // 检查用户名失败
	CodeEmailCheckFailed      = 10408 // 检查邮箱失败
	CodeUserRoleSyncFailed    = 10409 // 同步用户角色失败
	CodeRoleFindFailed        = 10410 // 查找角色失败

	// 会话相关内部错误 (105xx)
	CodeTokenGenerateFailed   = 10501 // 生成令牌失败
	CodeSessionCreateFailed   = 10502 // 创建会话失败
	CodeSessionUpdateFailed   = 10503 // 更新会话失败
	CodeSessionDeleteFailed   = 10504 // 删除会话失败
	CodeSessionServiceUnavailable = 10505 // 会话服务不可用
	CodeTokenBlacklistFailed  = 10506 // 添加令牌到黑名单失败

	// ========== 角色权限模块 (20xxx) ==========

	// 角色相关 (201xx)
	CodeRoleNotFound = 20101 // 角色不存在
	CodeRoleExists   = 20102 // 角色已存在
	CodeRoleInUse    = 20103 // 角色正在使用中

	// 角色操作内部错误 (203xx)
	CodeRoleCheckFailed           = 20301 // 检查角色代码失败
	CodeRoleCreateFailed          = 20302 // 创建角色失败
	CodeRoleGetFailed             = 20303 // 获取角色失败
	CodeRoleUpdateFailed          = 20304 // 更新角色失败
	CodeRoleDeleteFailed          = 20305 // 删除角色失败
	CodeRoleListFailed            = 20306 // 获取角色列表失败
	CodeRoleCheckUsageFailed      = 20307 // 检查角色使用情况失败
	CodeRolePermissionDeleteFailed = 20308 // 删除角色权限失败
	CodeCasbinUpdateFailed        = 20309 // 更新Casbin角色权限失败
	CodeRolePermissionUpdateFailed = 20310 // 更新数据库角色权限失败
	CodeRolePermissionGetFailed   = 20311 // 获取角色权限ID失败
	CodeUserRoleRemoveFailed      = 20312 // 移除用户角色失败
	CodeUserCasbinRoleRemoveFailed = 20313 // 移除用户Casbin角色失败
	CodeUserRoleAssignFailed      = 20314 // 分配用户角色失败
	CodeUserCasbinRoleAddFailed   = 20315 // 添加用户Casbin角色失败
	CodeUserRoleGetFailed         = 20316 // 获取用户角色失败

	// 权限相关 (202xx)
	CodePermissionNotFound   = 20201 // 权限不存在
	CodePermissionDenied     = 20202 // 权限不足
	CodeNoAccess             = 20203 // 无访问权限

	// 权限操作内部错误 (204xx)
	CodePermissionCheckFailed     = 20401 // 检查权限代码失败
	CodePermissionCreateFailed    = 20402 // 创建权限失败
	CodePermissionGetFailed       = 20403 // 获取权限失败
	CodePermissionUpdateFailed    = 20404 // 更新权限失败
	CodePermissionDeleteFailed    = 20405 // 删除权限失败
	CodePermissionListFailed      = 20406 // 获取权限列表失败
	CodePermissionTreeFailed      = 20407 // 获取权限树失败
	CodePermissionByResourceFailed = 20408 // 根据资源获取权限失败
	CodePermissionByTypeFailed    = 20409 // 根据类型获取权限失败

	// ========== 系统模块 (30xxx) ==========

	// 限流相关 (301xx)
	CodeRateLimitExceeded = 30101 // 请求过于频繁
	CodeIPBlocked         = 30102 // IP被限制

	// 内部错误 (302xx)
	CodeInternalError  = 30201 // 内部服务错误
	CodeDatabaseError  = 30202 // 数据库错误
	CodeCacheError     = 30203 // 缓存错误
	CodeEncryptError   = 30204 // 加密错误
	CodeTokenGenError  = 30205 // 令牌生成错误
	CodeSessionError   = 30206 // 会话错误

	// 参数验证 (303xx)
	CodeInvalidParams = 30301 // 参数错误
	CodeMissingParams = 30302 // 缺少必需参数
	CodeBadRequest    = 30303 // 请求格式错误

	// 资源相关 (304xx)
	CodeResourceNotFound = 30401 // 资源不存在
	CodeResourceConflict = 30402 // 资源冲突

	// ========== 字典模块 (40xxx) ==========
	
	// 字典类型 (401xx)
	CodeDictTypeNotFound = 40101 // 字典类型不存在
	CodeDictTypeExists   = 40102 // 字典类型已存在
	CodeDictTypeInUse    = 40103 // 字典类型正在使用中
	
	// 字典项 (402xx)
	CodeDictItemNotFound = 40201 // 字典项不存在
	CodeDictItemExists   = 40202 // 字典项已存在

	// 字典类型操作内部错误 (403xx)
	CodeDictTypeCheckFailed  = 40301 // 检查字典类型代码失败
	CodeDictTypeCreateFailed = 40302 // 创建字典类型失败
	CodeDictTypeGetFailed    = 40303 // 查询字典类型失败
	CodeDictTypeUpdateFailed = 40304 // 更新字典类型失败
	CodeDictTypeDeleteFailed = 40305 // 删除字典类型失败
	CodeDictTypeListFailed   = 40306 // 查询字典类型列表失败
	CodeDictTypeCountFailed  = 40307 // 统计字典项数量失败

	// 字典项操作内部错误 (404xx)
	CodeDictItemTypeGetFailed     = 40401 // 查询字典类型失败
	CodeDictItemValueCheckFailed  = 40402 // 检查字典项值失败
	CodeDictItemExtraConvertFailed = 40403 // 转换Extra失败
	CodeDictItemDefaultClearFailed = 40404 // 清除默认值失败
	CodeDictItemCreateFailed      = 40405 // 创建字典项失败
	CodeDictItemGetFailed         = 40406 // 查询字典项失败
	CodeDictItemUpdateFailed      = 40407 // 更新字典项失败
	CodeDictItemDeleteFailed      = 40408 // 删除字典项失败
	CodeDictItemListFailed        = 40409 // 查询字典项列表失败
	CodeDictItemsByTypeGetFailed  = 40410 // 根据类型获取字典项失败

	// ========== 菜单模块 (50xxx) ==========
	// 菜单相关 (501xx)
	CodeMenuNotFound        = 50101 // 菜单不存在
	CodeMenuExists          = 50102 // 菜单已存在
	CodeMenuHasChildren     = 50103 // 菜单存在子菜单
	CodeMenuInvalidParent   = 50104 // 无效的父菜单（通用）
	CodeMenuCannotBeOwnChild = 50105 // 不能设置为自己的子菜单
	CodeMenuParentNotFound  = 50106 // 父菜单不存在

	// 菜单操作内部错误 (502xx)
	CodeMenuParentGetFailed    = 50201 // 查询父菜单失败
	CodeMenuPermissionGetFailed = 50202 // 查询权限失败
	CodeMenuCreateFailed       = 50203 // 创建菜单失败
	CodeMenuGetFailed          = 50204 // 获取菜单失败
	CodeMenuUpdateFailed       = 50205 // 更新菜单失败
	CodeMenuDeleteFailed       = 50206 // 删除菜单失败
	CodeMenuCheckChildrenFailed = 50207 // 检查子菜单失败
	CodeMenuListFailed         = 50208 // 获取菜单列表失败
	CodeMenuVisibleListFailed  = 50209 // 获取可见菜单失败
	CodeMenuOrderUpdateFailed  = 50210 // 更新菜单顺序失败

	// ========== 密码重置模块 (60xxx) ==========
	CodeInvalidResetToken = 60101 // 无效的重置令牌
	CodeResetTokenExpired = 60102 // 重置令牌已过期

	// 密码重置操作内部错误 (601xx)
	CodePasswordResetUserQueryFailed   = 60111 // 查询用户失败
	CodePasswordResetTokenCleanFailed  = 60112 // 清理旧Token失败
	CodePasswordResetTokenCreateFailed = 60113 // 创建重置Token失败
	CodePasswordResetEmailSendFailed   = 60114 // 发送重置邮件失败
	CodePasswordResetTokenQueryFailed  = 60115 // 查询Token失败
	CodePasswordResetHashFailed        = 60116 // 密码加密失败
	CodePasswordResetUpdateFailed      = 60117 // 更新密码失败
)

// GetBusinessCodeMessage 获取业务错误码对应的默认消息
func GetBusinessCodeMessage(code int) string {
	messages := map[int]string{
		CodeSuccess: "成功",

		// 用户模块
		CodeInvalidCredentials: "用户名或密码错误",
		CodeAccountLocked:      "账户已被锁定",
		CodeInvalidCaptcha:     "验证码错误",
		CodeInvalidToken:       "无效的令牌",
		CodeTokenExpired:       "令牌已过期",
		CodeUnauthorized:       "未授权访问",

		CodeUsernameExists: "用户名已存在",
		CodeEmailExists:    "邮箱已存在",
		CodeWeakPassword:   "密码强度不足",

		CodeUserNotFound: "用户不存在",
		CodeUserDisabled: "用户已禁用",

		// 用户操作内部错误
		CodePasswordHashFailed:   "密码加密失败",
		CodeUserCreateFailed:     "用户创建失败",
		CodeUserQueryFailed:      "查询用户失败",
		CodeUserUpdateFailed:     "用户更新失败",
		CodeUserDeleteFailed:     "用户删除失败",
		CodeUserListFailed:      "查询用户列表失败",
		CodeUsernameCheckFailed: "检查用户名失败",
		CodeEmailCheckFailed:    "检查邮箱失败",
		CodeUserRoleSyncFailed:  "同步用户角色失败",
		CodeRoleFindFailed:      "查找角色失败",

		// 会话相关内部错误
		CodeTokenGenerateFailed:       "生成令牌失败",
		CodeSessionCreateFailed:       "创建会话失败",
		CodeSessionUpdateFailed:       "更新会话失败",
		CodeSessionDeleteFailed:       "删除会话失败",
		CodeSessionServiceUnavailable: "会话服务不可用",
		CodeTokenBlacklistFailed:      "添加令牌到黑名单失败",

		// 角色权限模块
		CodeRoleNotFound: "角色不存在",
		CodeRoleExists:   "角色代码已存在",
		CodeRoleInUse:    "角色正在使用中",

		// 角色操作内部错误
		CodeRoleCheckFailed:            "检查角色代码失败",
		CodeRoleCreateFailed:           "创建角色失败",
		CodeRoleGetFailed:              "获取角色失败",
		CodeRoleUpdateFailed:           "更新角色失败",
		CodeRoleDeleteFailed:           "删除角色失败",
		CodeRoleListFailed:             "获取角色列表失败",
		CodeRoleCheckUsageFailed:       "检查角色使用情况失败",
		CodeRolePermissionDeleteFailed: "删除角色权限失败",
		CodeCasbinUpdateFailed:         "更新Casbin角色权限失败",
		CodeRolePermissionUpdateFailed: "更新数据库角色权限失败",
		CodeRolePermissionGetFailed:    "获取角色权限ID失败",
		CodeUserRoleRemoveFailed:       "移除用户角色失败",
		CodeUserCasbinRoleRemoveFailed: "移除用户Casbin角色失败",
		CodeUserRoleAssignFailed:       "分配用户角色失败",
		CodeUserCasbinRoleAddFailed:    "添加用户Casbin角色失败",
		CodeUserRoleGetFailed:          "获取用户角色失败",

		CodePermissionNotFound: "权限不存在",
		CodePermissionDenied:   "权限不足",
		CodeNoAccess:           "无访问权限",

		// 权限操作内部错误
		CodePermissionCheckFailed:      "检查权限代码失败",
		CodePermissionCreateFailed:     "创建权限失败",
		CodePermissionGetFailed:        "获取权限失败",
		CodePermissionUpdateFailed:     "更新权限失败",
		CodePermissionDeleteFailed:     "删除权限失败",
		CodePermissionListFailed:       "获取权限列表失败",
		CodePermissionTreeFailed:       "获取权限树失败",
		CodePermissionByResourceFailed: "根据资源获取权限失败",
		CodePermissionByTypeFailed:     "根据类型获取权限失败",

		// 系统模块
		CodeRateLimitExceeded: "请求过于频繁",
		CodeIPBlocked:         "IP被限制",

		CodeInternalError: "内部服务错误",
		CodeDatabaseError: "数据库错误",
		CodeCacheError:    "缓存错误",
		CodeEncryptError:  "加密错误",
		CodeTokenGenError: "令牌生成错误",
		CodeSessionError:  "会话错误",

		CodeInvalidParams: "参数错误",
		CodeMissingParams: "缺少必需参数",
		CodeBadRequest:    "请求格式错误",

		CodeResourceNotFound: "资源不存在",
		CodeResourceConflict: "资源冲突",

		// 字典模块
		CodeDictTypeNotFound: "字典类型不存在",
		CodeDictTypeExists:   "字典类型代码已存在",
		CodeDictTypeInUse:    "字典类型正在使用中",
		CodeDictItemNotFound: "字典项不存在",
		CodeDictItemExists:   "字典项值已存在",

		// 字典类型操作内部错误
		CodeDictTypeCheckFailed:  "检查字典类型代码失败",
		CodeDictTypeCreateFailed: "创建字典类型失败",
		CodeDictTypeGetFailed:    "查询字典类型失败",
		CodeDictTypeUpdateFailed: "更新字典类型失败",
		CodeDictTypeDeleteFailed: "删除字典类型失败",
		CodeDictTypeListFailed:   "查询字典类型列表失败",
		CodeDictTypeCountFailed:  "统计字典项数量失败",

		// 字典项操作内部错误
		CodeDictItemTypeGetFailed:     "查询字典类型失败",
		CodeDictItemValueCheckFailed:  "检查字典项值失败",
		CodeDictItemExtraConvertFailed: "转换Extra失败",
		CodeDictItemDefaultClearFailed: "清除默认值失败",
		CodeDictItemCreateFailed:      "创建字典项失败",
		CodeDictItemGetFailed:         "查询字典项失败",
		CodeDictItemUpdateFailed:      "更新字典项失败",
		CodeDictItemDeleteFailed:      "删除字典项失败",
		CodeDictItemListFailed:        "查询字典项列表失败",
		CodeDictItemsByTypeGetFailed:  "根据类型获取字典项失败",

		// 菜单模块
		CodeMenuNotFound:         "菜单不存在",
		CodeMenuExists:           "菜单已存在",
		CodeMenuHasChildren:      "菜单存在子菜单",
		CodeMenuInvalidParent:    "无效的父菜单",
		CodeMenuCannotBeOwnChild: "不能将菜单设置为自己的子菜单",
		CodeMenuParentNotFound:   "父菜单不存在",

		// 菜单操作内部错误
		CodeMenuParentGetFailed:     "查询父菜单失败",
		CodeMenuPermissionGetFailed: "查询权限失败",
		CodeMenuCreateFailed:        "创建菜单失败",
		CodeMenuGetFailed:           "获取菜单失败",
		CodeMenuUpdateFailed:        "更新菜单失败",
		CodeMenuDeleteFailed:        "删除菜单失败",
		CodeMenuCheckChildrenFailed: "检查子菜单失败",
		CodeMenuListFailed:          "获取菜单列表失败",
		CodeMenuVisibleListFailed:   "获取可见菜单失败",
		CodeMenuOrderUpdateFailed:   "更新菜单顺序失败",

		// 密码重置模块
		CodeInvalidResetToken: "无效的重置令牌",
		CodeResetTokenExpired: "重置令牌已过期",

		// 密码重置操作内部错误
		CodePasswordResetUserQueryFailed:   "查询用户失败",
		CodePasswordResetTokenCleanFailed:  "清理旧Token失败",
		CodePasswordResetTokenCreateFailed: "创建重置Token失败",
		CodePasswordResetEmailSendFailed:   "发送重置邮件失败",
		CodePasswordResetTokenQueryFailed:  "查询Token失败",
		CodePasswordResetHashFailed:        "密码加密失败",
		CodePasswordResetUpdateFailed:      "更新密码失败",
	}

	if msg, ok := messages[code]; ok {
		return msg
	}
	return "未知错误"
}
