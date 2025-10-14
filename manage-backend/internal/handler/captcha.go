package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/service"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
)

type CaptchaHandler struct {
	captchaService service.CaptchaServiceInterface
}

func NewCaptchaHandler(captchaService service.CaptchaServiceInterface) *CaptchaHandler {
	return &CaptchaHandler{
		captchaService: captchaService,
	}
}

// GenerateCaptcha godoc
// @Summary Generate captcha
// @Description Generate a new captcha image for login verification
// @Tags auth
// @Produce json
// @Success 200 {object} utils.APIResponse{data=service.CaptchaResponse} "验证码生成成功"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /auth/captcha [get]
func (h *CaptchaHandler) GenerateCaptcha(c *gin.Context) {
	captcha, err := h.captchaService.GenerateCaptcha()
	if err != nil {
		utils.InternalServerError(c, "Failed to generate captcha")
		return
	}

	utils.Success(c, captcha)
}