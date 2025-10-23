package repository

import (
	"context"
	"time"

	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"gorm.io/gorm"
)

// PasswordResetRepository 密码重置Token数据仓库
// 封装对 PasswordResetToken 模型的所有数据库操作
type PasswordResetRepository struct {
	db *gorm.DB
}

// NewPasswordResetRepository 创建 PasswordResetRepository 实例
func NewPasswordResetRepository(db *gorm.DB) *PasswordResetRepository {
	return &PasswordResetRepository{db: db}
}

// Create 创建密码重置Token
// 参数: ctx - 上下文, token - 密码重置Token对象
// 返回: error - 操作是否成功
func (r *PasswordResetRepository) Create(ctx context.Context, token *model.PasswordResetToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

// FindByToken 根据Token查找密码重置记录
// 参数: ctx - 上下文, token - Token字符串
// 返回: *model.PasswordResetToken - 密码重置Token对象, error - 查询是否成功
func (r *PasswordResetRepository) FindByToken(ctx context.Context, token string) (*model.PasswordResetToken, error) {
	var resetToken model.PasswordResetToken
	err := r.db.WithContext(ctx).Where("token = ?", token).First(&resetToken).Error
	if err != nil {
		return nil, err
	}
	return &resetToken, nil
}

// MarkAsUsed 标记Token为已使用
// 参数: ctx - 上下文, tokenID - Token ID
// 返回: error - 操作是否成功
func (r *PasswordResetRepository) MarkAsUsed(ctx context.Context, tokenID uint) error {
	now := time.Now()
	return r.db.WithContext(ctx).Model(&model.PasswordResetToken{}).
		Where("id = ?", tokenID).
		Update("used_at", now).Error
}

// DeleteExpiredTokens 删除过期的Token
// 参数: ctx - 上下文
// 返回: error - 操作是否成功
func (r *PasswordResetRepository) DeleteExpiredTokens(ctx context.Context) error {
	return r.db.WithContext(ctx).
		Where("expires_at < ?", time.Now()).
		Delete(&model.PasswordResetToken{}).Error
}

// DeleteByUserID 删除指定用户的未使用Token
// 参数: ctx - 上下文, userID - 用户ID
// 返回: error - 操作是否成功
func (r *PasswordResetRepository) DeleteByUserID(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND used_at IS NULL", userID).
		Delete(&model.PasswordResetToken{}).Error
}
