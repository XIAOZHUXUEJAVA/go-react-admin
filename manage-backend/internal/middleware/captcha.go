package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/service"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
)

// CaptchaMiddleware 验证码中间件
func CaptchaMiddleware(captchaService service.CaptchaServiceInterface, enabled bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果验证码未启用，直接跳过验证
		if !enabled {
			c.Next()
			return
		}

		// 只对 POST 请求进行验证码检查
		if c.Request.Method != http.MethodPost {
			c.Next()
			return
		}

		// 检查是否是需要验证码的路由
		path := c.FullPath()
		needsCaptcha := path == "/api/v1/auth/login" || path == "/api/v1/auth/register"
		
		if !needsCaptcha {
			c.Next()
			return
		}

		// 从请求中获取验证码信息
		var captchaData struct {
			CaptchaID   string `json:"captcha_id"`
			CaptchaCode string `json:"captcha_code"`
		}

		// 尝试绑定 JSON 数据
		if err := c.ShouldBindJSON(&captchaData); err != nil {
			utils.BadRequest(c, "Invalid request format")
			c.Abort()
			return
		}

		// 验证验证码
		if !captchaService.VerifyCaptcha(captchaData.CaptchaID, captchaData.CaptchaCode) {
			utils.BadRequest(c, "Invalid captcha")
			c.Abort()
			return
		}

		c.Next()
	}
}