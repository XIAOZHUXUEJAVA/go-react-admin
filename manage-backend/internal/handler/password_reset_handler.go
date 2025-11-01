package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/service"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
	apperrors "github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/errors"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
)

// PasswordResetHandler 密码重置处理器
type PasswordResetHandler struct {
	service service.PasswordResetService
}

// NewPasswordResetHandler 创建密码重置处理器实例
func NewPasswordResetHandler(service service.PasswordResetService) *PasswordResetHandler {
	return &PasswordResetHandler{service: service}
}

// ForgotPassword godoc
// @Summary 请求密码重置
// @Description 用户通过邮箱请求密码重置，系统将发送重置链接到邮箱
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.ForgotPasswordRequest true "邮箱"
// @Success 200 {object} utils.APIResponse "请求成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /auth/forgot-password [post]
func (h *PasswordResetHandler) ForgotPassword(c *gin.Context) {
	var req model.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("忘记密码请求参数错误",
			zap.Error(err),
			zap.String("operation", "forgot_password"))
		utils.BadRequest(c, "请求参数错误")
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	logger.Info("收到忘记密码请求",
		zap.String("email", req.Email),
		zap.String("ip_address", ipAddress),
		zap.String("operation", "forgot_password"))

	if err := h.service.RequestPasswordReset(c.Request.Context(), req.Email, ipAddress, userAgent); err != nil {
		logger.Error("处理忘记密码请求失败",
			zap.String("email", req.Email),
			zap.Error(err),
			zap.String("operation", "forgot_password"))
		
		// 使用自定义错误类型处理
		if appErr, ok := apperrors.GetAppError(err); ok {
			// 根据错误类型返回相应的HTTP状态码
			switch appErr.Code {
			case 400:
				utils.BadRequest(c, appErr.Message)
			case 401:
				utils.Unauthorized(c, appErr.Message)
			case 403:
				utils.Forbidden(c, appErr.Message)
			case 423:
				utils.Locked(c, appErr.Message)
			case 429:
				utils.TooManyRequests(c, appErr.Message)
			default:
				utils.InternalServerError(c, appErr.Message)
			}
			return
		}
		
		// 未知错误
		utils.InternalServerError(c, "处理失败，请稍后重试")
		return
	}

	// 安全实践：无论邮箱是否存在，都返回相同的成功消息
	utils.Success(c, gin.H{
		"message": "如果该邮箱存在，重置链接已发送，请查收邮件",
	})
}

// VerifyResetToken godoc
// @Summary 验证重置Token
// @Description 验证密码重置Token的有效性
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.VerifyResetTokenRequest true "Token"
// @Success 200 {object} utils.APIResponse{data=model.VerifyResetTokenResponse} "验证成功"
// @Failure 400 {object} utils.APIResponse "Token无效或已过期"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /auth/verify-reset-token [post]
func (h *PasswordResetHandler) VerifyResetToken(c *gin.Context) {
	var req model.VerifyResetTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("验证Token请求参数错误",
			zap.Error(err),
			zap.String("operation", "verify_reset_token"))
		utils.BadRequest(c, "请求参数错误")
		return
	}

	logger.Debug("收到验证Token请求",
		zap.String("operation", "verify_reset_token"))

	user, err := h.service.VerifyResetToken(c.Request.Context(), req.Token)
	if err != nil {
		logger.Warn("Token验证失败",
			zap.Error(err),
			zap.String("operation", "verify_reset_token"))
		utils.BadRequest(c, err.Error())
		return
	}

	logger.Info("Token验证成功",
		zap.Uint("user_id", user.ID),
		zap.String("email", user.Email),
		zap.String("operation", "verify_reset_token"))

	utils.Success(c, model.VerifyResetTokenResponse{
		Valid: true,
		Email: user.Email,
	})
}

// ResetPassword godoc
// @Summary 重置密码
// @Description 使用有效的Token重置用户密码
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.ResetPasswordRequest true "重置信息"
// @Success 200 {object} utils.APIResponse "重置成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误或Token无效"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /auth/reset-password [post]
func (h *PasswordResetHandler) ResetPassword(c *gin.Context) {
	var req model.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("重置密码请求参数错误",
			zap.Error(err),
			zap.String("operation", "reset_password"))
		utils.BadRequest(c, "请求参数错误")
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	logger.Info("收到重置密码请求",
		zap.String("ip_address", ipAddress),
		zap.String("operation", "reset_password"))

	if err := h.service.ResetPassword(c.Request.Context(), req.Token, req.NewPassword, ipAddress, userAgent); err != nil {
		logger.Warn("密码重置失败",
			zap.Error(err),
			zap.String("ip_address", ipAddress),
			zap.String("operation", "reset_password"))
		
		// 根据错误类型返回不同的状态码
		if err.Error() == "无效的重置链接" || 
		   err.Error() == "重置链接已过期，请重新申请" || 
		   err.Error() == "重置链接已使用，请重新申请" {
			utils.BadRequest(c, err.Error())
		} else {
			utils.InternalServerError(c, "密码重置失败，请稍后重试")
		}
		return
	}

	logger.Info("密码重置成功",
		zap.String("ip_address", ipAddress),
		zap.String("operation", "reset_password"))

	utils.SuccessWithMessage(c, "密码重置成功，请使用新密码登录", nil)
}
