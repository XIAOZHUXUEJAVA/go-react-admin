package config

import "time"

// CaptchaConfig 验证码配置
type CaptchaConfig struct {
	// 验证码类型: digit(数字), string(字符串), math(数学), chinese(中文), audio(音频)
	Type string `mapstructure:"type" yaml:"type"`
	
	// 验证码长度
	Length int `mapstructure:"length" yaml:"length"`
	
	// 图片宽度
	Width int `mapstructure:"width" yaml:"width"`
	
	// 图片高度
	Height int `mapstructure:"height" yaml:"height"`
	
	// 噪点强度 (0.0-1.0)
	NoiseCount float64 `mapstructure:"noise_count" yaml:"noise_count"`
	
	// 干扰线数量
	ShowLineOptions int `mapstructure:"show_line_options" yaml:"show_line_options"`
	
	// 过期时间
	Expiration time.Duration `mapstructure:"expiration" yaml:"expiration"`
	
	// 是否启用验证码
	Enabled bool `mapstructure:"enabled" yaml:"enabled"`
}

// GetDefaultCaptchaConfig 获取默认验证码配置
func GetDefaultCaptchaConfig() CaptchaConfig {
	return CaptchaConfig{
		Type:            "digit",
		Length:          5,
		Width:           240,
		Height:          80,
		NoiseCount:      0.7,
		ShowLineOptions: 80,
		Expiration:      5 * time.Minute,
		Enabled:         true,
	}
}