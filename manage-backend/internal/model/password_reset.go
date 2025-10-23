package model

import (
	"time"
)

// PasswordResetToken 密码重置Token模型
type PasswordResetToken struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	UserID    uint       `json:"user_id" gorm:"not null;index"`
	Email     string     `json:"email" gorm:"not null;index;size:255"`
	Token     string     `json:"token" gorm:"not null;uniqueIndex;size:255"`
	ExpiresAt time.Time  `json:"expires_at" gorm:"not null;index"`
	UsedAt    *time.Time `json:"used_at" gorm:"index"`
	IPAddress string     `json:"ip_address" gorm:"size:50"`
	UserAgent string     `json:"user_agent" gorm:"type:text"`
	CreatedAt time.Time  `json:"created_at"`
}

// TableName 指定表名
func (PasswordResetToken) TableName() string {
	return "password_reset_tokens"
}

// IsExpired 检查Token是否过期
func (t *PasswordResetToken) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}

// IsUsed 检查Token是否已使用
func (t *PasswordResetToken) IsUsed() bool {
	return t.UsedAt != nil
}

// IsValid 检查Token是否有效（未过期且未使用）
func (t *PasswordResetToken) IsValid() bool {
	return !t.IsExpired() && !t.IsUsed()
}

// ForgotPasswordRequest 忘记密码请求
type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// ResetPasswordRequest 重置密码请求
type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// VerifyResetTokenRequest 验证重置Token请求
type VerifyResetTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

// VerifyResetTokenResponse 验证重置Token响应
type VerifyResetTokenResponse struct {
	Valid bool   `json:"valid"`
	Email string `json:"email,omitempty"`
}
