package repository

import (
	"encoding/json"

	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// DictTypeRepository 字典类型数据仓库
type DictTypeRepository struct {
	db *gorm.DB
}

// NewDictTypeRepository 创建 DictTypeRepository 实例
func NewDictTypeRepository(db *gorm.DB) *DictTypeRepository {
	return &DictTypeRepository{db: db}
}

// Create 创建字典类型
func (r *DictTypeRepository) Create(dictType *model.DictType) error {
	return r.db.Create(dictType).Error
}

// GetByID 根据 ID 获取字典类型
func (r *DictTypeRepository) GetByID(id uint) (*model.DictType, error) {
	var dictType model.DictType
	err := r.db.First(&dictType, id).Error
	if err != nil {
		return nil, err
	}
	return &dictType, nil
}

// GetByCode 根据代码获取字典类型
func (r *DictTypeRepository) GetByCode(code string) (*model.DictType, error) {
	var dictType model.DictType
	err := r.db.Where("code = ?", code).First(&dictType).Error
	if err != nil {
		return nil, err
	}
	return &dictType, nil
}

// Update 更新字典类型
func (r *DictTypeRepository) Update(dictType *model.DictType) error {
	return r.db.Save(dictType).Error
}

// Delete 删除字典类型
func (r *DictTypeRepository) Delete(id uint) error {
	return r.db.Delete(&model.DictType{}, id).Error
}

// List 获取字典类型列表（分页）
func (r *DictTypeRepository) List(offset, limit int, status, keyword string) ([]model.DictType, int64, error) {
	var dictTypes []model.DictType
	var total int64

	query := r.db.Model(&model.DictType{})

	// 状态筛选
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 关键字搜索
	if keyword != "" {
		query = query.Where("code LIKE ? OR name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	err := query.Order("sort_order ASC, created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&dictTypes).Error

	return dictTypes, total, err
}

// GetAll 获取所有字典类型（不分页）
func (r *DictTypeRepository) GetAll() ([]model.DictType, error) {
	var dictTypes []model.DictType
	err := r.db.Where("status = ?", "active").
		Order("sort_order ASC, created_at DESC").
		Find(&dictTypes).Error
	return dictTypes, err
}

// CheckCodeExists 检查代码是否存在
func (r *DictTypeRepository) CheckCodeExists(code string) (bool, error) {
	var count int64
	err := r.db.Model(&model.DictType{}).Where("code = ?", code).Count(&count).Error
	return count > 0, err
}

// CheckCodeExistsExcludeID 检查代码是否存在（排除指定ID）
func (r *DictTypeRepository) CheckCodeExistsExcludeID(code string, excludeID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.DictType{}).
		Where("code = ? AND id != ?", code, excludeID).
		Count(&count).Error
	return count > 0, err
}

// ==================== DictItemRepository ====================

// DictItemRepository 字典项数据仓库
type DictItemRepository struct {
	db *gorm.DB
}

// NewDictItemRepository 创建 DictItemRepository 实例
func NewDictItemRepository(db *gorm.DB) *DictItemRepository {
	return &DictItemRepository{db: db}
}

// Create 创建字典项
func (r *DictItemRepository) Create(dictItem *model.DictItem) error {
	return r.db.Create(dictItem).Error
}

// GetByID 根据 ID 获取字典项
func (r *DictItemRepository) GetByID(id uint) (*model.DictItem, error) {
	var dictItem model.DictItem
	err := r.db.First(&dictItem, id).Error
	if err != nil {
		return nil, err
	}
	return &dictItem, nil
}

// GetByTypeCodeAndValue 根据类型代码和值获取字典项
func (r *DictItemRepository) GetByTypeCodeAndValue(typeCode, value string) (*model.DictItem, error) {
	var dictItem model.DictItem
	err := r.db.Where("dict_type_code = ? AND value = ?", typeCode, value).First(&dictItem).Error
	if err != nil {
		return nil, err
	}
	return &dictItem, nil
}

// Update 更新字典项
func (r *DictItemRepository) Update(dictItem *model.DictItem) error {
	return r.db.Save(dictItem).Error
}

// Delete 删除字典项
func (r *DictItemRepository) Delete(id uint) error {
	return r.db.Delete(&model.DictItem{}, id).Error
}

// List 获取字典项列表（分页）
func (r *DictItemRepository) List(offset, limit int, typeCode, status string) ([]model.DictItem, int64, error) {
	var dictItems []model.DictItem
	var total int64

	query := r.db.Model(&model.DictItem{})

	// 类型代码筛选
	if typeCode != "" {
		query = query.Where("dict_type_code = ?", typeCode)
	}

	// 状态筛选
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	err := query.Order("sort_order ASC, created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&dictItems).Error

	return dictItems, total, err
}

// GetByTypeCode 根据类型代码获取所有字典项
func (r *DictItemRepository) GetByTypeCode(typeCode string, activeOnly bool) ([]model.DictItem, error) {
	var dictItems []model.DictItem
	query := r.db.Where("dict_type_code = ?", typeCode)

	if activeOnly {
		query = query.Where("status = ?", "active")
	}

	err := query.Order("sort_order ASC, created_at DESC").Find(&dictItems).Error
	return dictItems, err
}

// CheckValueExists 检查值是否存在
func (r *DictItemRepository) CheckValueExists(typeCode, value string) (bool, error) {
	var count int64
	err := r.db.Model(&model.DictItem{}).
		Where("dict_type_code = ? AND value = ?", typeCode, value).
		Count(&count).Error
	return count > 0, err
}

// CheckValueExistsExcludeID 检查值是否存在（排除指定ID）
func (r *DictItemRepository) CheckValueExistsExcludeID(typeCode, value string, excludeID uint) (bool, error) {
	var count int64
	err := r.db.Model(&model.DictItem{}).
		Where("dict_type_code = ? AND value = ? AND id != ?", typeCode, value, excludeID).
		Count(&count).Error
	return count > 0, err
}

// ClearDefaultByType 清除指定类型下的所有默认值
func (r *DictItemRepository) ClearDefaultByType(typeCode string) error {
	return r.db.Model(&model.DictItem{}).
		Where("dict_type_code = ? AND is_default = ?", typeCode, true).
		Update("is_default", false).Error
}

// GetDefaultByType 获取指定类型的默认字典项
func (r *DictItemRepository) GetDefaultByType(typeCode string) (*model.DictItem, error) {
	var dictItem model.DictItem
	err := r.db.Where("dict_type_code = ? AND is_default = ? AND status = ?", typeCode, true, "active").
		First(&dictItem).Error
	if err != nil {
		return nil, err
	}
	return &dictItem, nil
}

// CountByTypeCode 统计指定类型的字典项数量
func (r *DictItemRepository) CountByTypeCode(typeCode string) (int64, error) {
	var count int64
	err := r.db.Model(&model.DictItem{}).
		Where("dict_type_code = ?", typeCode).
		Count(&count).Error
	return count, err
}

// ConvertDictItemToResponse 转换字典项为响应格式
func ConvertDictItemToResponse(item *model.DictItem) *model.DictItemResponse {
	response := &model.DictItemResponse{
		ID:           item.ID,
		DictTypeCode: item.DictTypeCode,
		Label:        item.Label,
		Value:        item.Value,
		Description:  item.Description,
		Status:       item.Status,
		SortOrder:    item.SortOrder,
		IsDefault:    item.IsDefault,
		IsSystem:     item.IsSystem,
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
	}

	// 解析 Extra JSON
	if item.Extra != nil {
		var extra map[string]interface{}
		if err := json.Unmarshal(item.Extra, &extra); err == nil {
			response.Extra = extra
		}
	}

	return response
}

// ConvertMapToJSON 转换 map 为 JSON
func ConvertMapToJSON(data map[string]interface{}) (datatypes.JSON, error) {
	if data == nil {
		return nil, nil
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return datatypes.JSON(jsonData), nil
}
