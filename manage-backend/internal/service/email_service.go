package service

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/config"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

// EmailService 邮件服务接口
type EmailService interface {
	SendPasswordResetEmail(to, token, username string) error
}

type emailService struct {
	config *config.Config
	dialer *gomail.Dialer
}

// NewEmailService 创建邮件服务实例
func NewEmailService(cfg *config.Config) EmailService {
	// 创建SMTP拨号器（可复用）
	dialer := gomail.NewDialer(
		cfg.Email.SMTPHost,
		cfg.Email.SMTPPort,
		cfg.Email.Username,
		cfg.Email.Password,
	)

	// 如果使用465端口（SSL），需要设置
	if cfg.Email.SMTPPort == 465 {
		dialer.SSL = true
	}

	logger.Info("邮件服务初始化成功",
		zap.String("smtp_host", cfg.Email.SMTPHost),
		zap.Int("smtp_port", cfg.Email.SMTPPort),
		zap.String("from_address", cfg.Email.FromAddress))

	return &emailService{
		config: cfg,
		dialer: dialer,
	}
}

// SendPasswordResetEmail 发送密码重置邮件
func (s *emailService) SendPasswordResetEmail(to, token, username string) error {
	logger.Info("开始发送密码重置邮件",
		zap.String("to", to),
		zap.String("username", username),
		zap.String("operation", "send_reset_email"))

	resetURL := fmt.Sprintf("%s/reset-password?token=%s",
		s.config.PasswordReset.FrontendURL, token)

	// HTML邮件模板
	htmlTemplate := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body { 
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
            line-height: 1.6; 
            color: #333; 
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }
        .container { 
            max-width: 600px; 
            margin: 20px auto; 
            background: white;
            border-radius: 8px;
            overflow: hidden;
            box-shadow: 0 2px 8px rgba(0,0,0,0.1);
        }
        .header { 
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white; 
            padding: 30px 20px; 
            text-align: center; 
        }
        .header h1 {
            margin: 0;
            font-size: 24px;
            font-weight: 600;
        }
        .content { 
            padding: 40px 30px; 
        }
        .content p {
            margin: 0 0 15px 0;
            color: #555;
        }
        .button { 
            display: inline-block; 
            padding: 14px 40px; 
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white !important; 
            text-decoration: none; 
            border-radius: 6px; 
            margin: 25px 0;
            font-weight: 600;
            box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
        }
        .link-box {
            word-break: break-all; 
            background: #f8f9fa; 
            padding: 15px; 
            border-radius: 6px;
            border-left: 4px solid #667eea;
            font-size: 13px;
            color: #666;
            margin: 20px 0;
        }
        .warning { 
            background: #fff3cd;
            border-left: 4px solid #ffc107;
            padding: 15px;
            border-radius: 6px;
            margin: 20px 0;
        }
        .warning-icon {
            color: #ff9800;
            font-size: 18px;
            font-weight: bold;
        }
        .footer { 
            text-align: center; 
            color: #999; 
            font-size: 12px; 
            padding: 20px;
            background: #f8f9fa;
            border-top: 1px solid #e9ecef;
        }
        .footer p {
            margin: 5px 0;
        }
        @media only screen and (max-width: 600px) {
            .container {
                margin: 0;
                border-radius: 0;
            }
            .content {
                padding: 30px 20px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🔐 密码重置请求</h1>
        </div>
        <div class="content">
            <p style="font-size: 16px;"><strong>您好，{{.Username}}！</strong></p>
            <p>我们收到了您的密码重置请求。请点击下面的按钮来重置您的密码：</p>
            
            <div style="text-align: center;">
                <a href="{{.ResetURL}}" class="button">重置密码</a>
            </div>
            
            <p style="margin-top: 30px;">如果按钮无法点击，请复制以下链接到浏览器地址栏：</p>
            <div class="link-box">
                {{.ResetURL}}
            </div>
            
            <div class="warning">
                <span class="warning-icon">⚠️</span>
                <strong>重要提示：</strong>
                <ul style="margin: 10px 0 0 0; padding-left: 20px;">
                    <li>此链接将在 <strong>1小时</strong> 后过期</li>
                    <li>链接仅可使用 <strong>一次</strong></li>
                    <li>如果您没有请求重置密码，请忽略此邮件</li>
                </ul>
            </div>
            
            <p style="margin-top: 30px; color: #999; font-size: 14px;">
                为了您的账户安全，请勿将此链接分享给他人。
            </p>
        </div>
        <div class="footer">
            <p>此邮件由系统自动发送，请勿回复</p>
            <p>&copy; 2025 Go Manage System. All rights reserved.</p>
            <p style="margin-top: 10px;">
                <a href="{{.FrontendURL}}" style="color: #667eea; text-decoration: none;">访问网站</a>
            </p>
        </div>
    </div>
</body>
</html>
`

	// 解析模板
	tmpl, err := template.New("reset").Parse(htmlTemplate)
	if err != nil {
		logger.Error("解析邮件模板失败",
			zap.String("to", to),
			zap.Error(err),
			zap.String("operation", "send_reset_email"))
		return fmt.Errorf("解析邮件模板失败: %w", err)
	}

	// 渲染模板
	var body bytes.Buffer
	data := struct {
		Username    string
		ResetURL    string
		FrontendURL string
	}{
		Username:    username,
		ResetURL:    resetURL,
		FrontendURL: s.config.PasswordReset.FrontendURL,
	}

	if err := tmpl.Execute(&body, data); err != nil {
		logger.Error("渲染邮件模板失败",
			zap.String("to", to),
			zap.Error(err),
			zap.String("operation", "send_reset_email"))
		return fmt.Errorf("渲染邮件模板失败: %w", err)
	}

	// 创建邮件消息
	m := gomail.NewMessage()

	// 设置发件人
	m.SetHeader("From", m.FormatAddress(s.config.Email.FromAddress, s.config.Email.FromName))

	// 设置收件人
	m.SetHeader("To", to)

	// 设置主题
	m.SetHeader("Subject", "密码重置请求 - Go Manage System")

	// 设置邮件正文（HTML格式）
	m.SetBody("text/html", body.String())

	// 可选：设置纯文本备用内容（某些邮件客户端不支持HTML）
	plainText := fmt.Sprintf(`
您好，%s！

我们收到了您的密码重置请求。

请访问以下链接重置密码：
%s

此链接将在1小时后过期，且仅可使用一次。

如果您没有请求重置密码，请忽略此邮件。

---
Go Manage System
此邮件由系统自动发送，请勿回复
`, username, resetURL)

	m.AddAlternative("text/plain", plainText)

	// 发送邮件
	if err := s.dialer.DialAndSend(m); err != nil {
		logger.Error("发送邮件失败",
			zap.String("to", to),
			zap.String("smtp_host", s.config.Email.SMTPHost),
			zap.Error(err),
			zap.String("operation", "send_reset_email"))
		return fmt.Errorf("发送邮件失败: %w", err)
	}

	logger.Info("密码重置邮件发送成功",
		zap.String("to", to),
		zap.String("username", username),
		zap.String("operation", "send_reset_email"))

	return nil
}
