package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	Password  string         `json:"-" gorm:"not null"`
	Role      string         `json:"role" gorm:"default:user"`
	Status    string         `json:"status" gorm:"default:active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// UserResponse 用户响应结构体（不包含敏感信息）
type UserResponse struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"`
}

type UpdateUserRequest struct {
	Username string `json:"username" binding:"omitempty,min=3,max=50"`
	Email    string `json:"email" binding:"omitempty,email"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

type LoginRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CaptchaID   string `json:"captcha_id" binding:"required"`
	CaptchaCode string `json:"captcha_code" binding:"required"`
}

type LoginResponse struct {
	AccessToken      string       `json:"access_token"`
	RefreshToken     string       `json:"refresh_token"`
	ExpiresIn        int64        `json:"expires_in"`        // Access token expiration in seconds
	RefreshExpiresIn int64        `json:"refresh_expires_in"` // Refresh token expiration in seconds
	TokenType        string       `json:"token_type"`         // Always "Bearer"
	User             UserResponse `json:"user"`
}

// RefreshTokenRequest 刷新token请求
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// RefreshTokenResponse 刷新token响应
type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

// LogoutRequest 登出请求
type LogoutRequest struct {
	RefreshToken string `json:"refresh_token,omitempty"`
}

// CheckAvailabilityRequest 检查可用性请求
type CheckAvailabilityRequest struct {
	Username      string `json:"username,omitempty"`
	Email         string `json:"email,omitempty"`
	ExcludeUserID *uint  `json:"exclude_user_id,omitempty"`
}

// AvailabilityResult 可用性检查结果
type AvailabilityResult struct {
	Available bool   `json:"available"`
	Message   string `json:"message,omitempty"`
}

// CheckAvailabilityResponse 检查可用性响应
type CheckAvailabilityResponse struct {
	Username *AvailabilityResult `json:"username,omitempty"`
	Email    *AvailabilityResult `json:"email,omitempty"`
}

// SimpleAvailabilityResponse 简单可用性响应
type SimpleAvailabilityResponse struct {
	Available bool `json:"available"`
}

// CaptchaResponse 验证码响应
type CaptchaResponse struct {
	CaptchaID   string `json:"captcha_id"`
	CaptchaData string `json:"captcha_data"`
}