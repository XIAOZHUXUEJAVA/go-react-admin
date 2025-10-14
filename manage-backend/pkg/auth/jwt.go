package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims 自定义的 JWT 声明 (Payload)
// 包含用户的基本信息和标准声明
type Claims struct {
	UserID   uint   `json:"user_id"` // 用户ID
	Username string `json:"username"` // 用户名
	Role     string `json:"role"`     // 用户角色
	JTI      string `json:"jti"`      // JWT唯一ID (用于黑名单支持)
	jwt.RegisteredClaims              // 标准 JWT 声明 (exp, iat, nbf, iss, sub 等)
}

// JWTManager 用于生成和验证 JWT
type JWTManager struct {
	secretKey          string        // 密钥，用于签名和验证
	accessTokenExpire  time.Duration // Access Token 过期时间
	refreshTokenExpire time.Duration // Refresh Token 过期时间
}

// TokenPair 封装访问令牌和刷新令牌
type TokenPair struct {
	AccessToken      string `json:"access_token"`       // 访问令牌
	RefreshToken     string `json:"refresh_token"`      // 刷新令牌
	ExpiresIn        int64  `json:"expires_in"`         // Access Token 有效时长（秒）
	RefreshExpiresIn int64  `json:"refresh_expires_in"` // Refresh Token 有效时长（秒）
}

// NewJWTManager 创建一个新的 JWT 管理器
// secretKey: 签名密钥
// accessExpireMinutes: access token 有效期（分钟）
// refreshExpireHours: refresh token 有效期（小时）
func NewJWTManager(secretKey string, accessExpireMinutes, refreshExpireHours int) *JWTManager {
	return &JWTManager{
		secretKey:          secretKey,
		accessTokenExpire:  time.Duration(accessExpireMinutes) * time.Minute,
		refreshTokenExpire: time.Duration(refreshExpireHours) * time.Hour,
	}
}

// GenerateTokenPair 生成 Access Token 和 Refresh Token
func (j *JWTManager) GenerateTokenPair(userID uint, username, role string) (*TokenPair, error) {
	// 生成 Access Token 的 JTI
	accessJTI, err := j.generateJTI()
	if err != nil {
		return nil, err
	}

	// 设置 Access Token 的声明
	accessClaims := &Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		JTI:      accessJTI,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.accessTokenExpire)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                          // 生效时间
			Issuer:    "go-manage-starter",                                     // 签发者
			Subject:   "access",                                                // Token 类型
		},
	}

	// 签名生成 Access Token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(j.secretKey))
	if err != nil {
		return nil, err
	}

	// 生成 Refresh Token 的 JTI
	refreshJTI, err := j.generateJTI()
	if err != nil {
		return nil, err
	}

	// 设置 Refresh Token 的声明
	refreshClaims := &Claims{
		UserID:   userID,
		Username: username,
		Role:     role,
		JTI:      refreshJTI,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshTokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "go-manage-starter",
			Subject:   "refresh", // 标记为刷新令牌
		},
	}

	// 签名生成 Refresh Token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(j.secretKey))
	if err != nil {
		return nil, err
	}

	// 返回 Token 对
	return &TokenPair{
		AccessToken:      accessTokenString,
		RefreshToken:     refreshTokenString,
		ExpiresIn:        int64(j.accessTokenExpire.Seconds()),
		RefreshExpiresIn: int64(j.refreshTokenExpire.Seconds()),
	}, nil
}

// GenerateToken 仅生成 Access Token（兼容旧版本使用）
func (j *JWTManager) GenerateToken(userID uint, username, role string) (string, error) {
	tokenPair, err := j.GenerateTokenPair(userID, username, role)
	if err != nil {
		return "", err
	}
	return tokenPair.AccessToken, nil
}

// ValidateToken 验证 JWT Token 并解析为 Claims
func (j *JWTManager) ValidateToken(tokenString string) (*Claims, error) {
	// 使用密钥验证 Token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("不支持的签名方法")
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证 Token 是否有效并返回 Claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的 Token")
}

// ValidateRefreshToken 专门验证 Refresh Token
func (j *JWTManager) ValidateRefreshToken(tokenString string) (*Claims, error) {
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.Subject != "refresh" {
		return nil, errors.New("不是有效的刷新令牌")
	}

	return claims, nil
}

// generateJTI 生成唯一的 JWT ID
func (j *JWTManager) generateJTI() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// GetTokenExpiration 获取 Token 剩余的过期时间
func (j *JWTManager) GetTokenExpiration(claims *Claims) time.Duration {
	if claims.ExpiresAt == nil {
		return 0
	}
	return time.Until(claims.ExpiresAt.Time)
}
