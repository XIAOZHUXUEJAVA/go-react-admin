package service

import (
	"errors"

	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/repository"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/logger"
	apperrors "github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// DictTypeRepositoryInterface 定义字典类型仓库接口
type DictTypeRepositoryInterface interface {
	Create(dictType *model.DictType) error
	GetByID(id uint) (*model.DictType, error)
	GetByCode(code string) (*model.DictType, error)
	Update(dictType *model.DictType) error
	Delete(id uint) error
	List(offset, limit int, status, keyword string) ([]model.DictType, int64, error)
	GetAll() ([]model.DictType, error)
	CheckCodeExists(code string) (bool, error)
	CheckCodeExistsExcludeID(code string, excludeID uint) (bool, error)
}

// DictItemRepositoryInterface 定义字典项仓库接口
type DictItemRepositoryInterface interface {
	Create(dictItem *model.DictItem) error
	GetByID(id uint) (*model.DictItem, error)
	GetByTypeCodeAndValue(typeCode, value string) (*model.DictItem, error)
	Update(dictItem *model.DictItem) error
	Delete(id uint) error
	List(offset, limit int, typeCode, status string) ([]model.DictItem, int64, error)
	GetByTypeCode(typeCode string, activeOnly bool) ([]model.DictItem, error)
	CheckValueExists(typeCode, value string) (bool, error)
	CheckValueExistsExcludeID(typeCode, value string, excludeID uint) (bool, error)
	ClearDefaultByType(typeCode string) error
	GetDefaultByType(typeCode string) (*model.DictItem, error)
	CountByTypeCode(typeCode string) (int64, error)
}

// DictTypeService 字典类型业务服务
type DictTypeService struct {
	dictTypeRepo DictTypeRepositoryInterface
	dictItemRepo DictItemRepositoryInterface
}

// NewDictTypeService 创建 DictTypeService 实例
func NewDictTypeService(
	dictTypeRepo DictTypeRepositoryInterface,
	dictItemRepo DictItemRepositoryInterface,
) *DictTypeService {
	return &DictTypeService{
		dictTypeRepo: dictTypeRepo,
		dictItemRepo: dictItemRepo,
	}
}

// Create 创建字典类型
func (s *DictTypeService) Create(req *model.CreateDictTypeRequest) (*model.DictType, error) {
	logger.Info("开始创建字典类型",
		zap.String("code", req.Code),
		zap.String("name", req.Name),
		zap.String("operation", "create_dict_type"))

	// 检查代码是否已存在
	exists, err := s.dictTypeRepo.CheckCodeExists(req.Code)
	if err != nil {
		logger.Error("检查字典类型代码是否存在失败",
			zap.String("code", req.Code),
			zap.Error(err),
			zap.String("operation", "create_dict_type"))
		return nil, apperrors.NewDictTypeCheckFailedError()
	}
	if exists {
		logger.Warn("字典类型代码已存在",
			zap.String("code", req.Code),
			zap.String("operation", "create_dict_type"))
		return nil, apperrors.NewDictTypeExistsError()
	}

	// 创建字典类型
	dictType := &model.DictType{
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		SortOrder:   req.SortOrder,
		IsSystem:    false, // 用户创建的都不是系统内置
	}

	if dictType.Status == "" {
		dictType.Status = "active"
	}

	err = s.dictTypeRepo.Create(dictType)
	if err != nil {
		logger.Error("创建字典类型失败",
			zap.String("code", req.Code),
			zap.Error(err),
			zap.String("operation", "create_dict_type"))
		return nil, apperrors.NewDictTypeCreateFailedError()
	}

	logger.Info("字典类型创建成功",
		zap.Uint("id", dictType.ID),
		zap.String("code", dictType.Code),
		zap.String("operation", "create_dict_type"))

	return dictType, nil
}

// GetByID 根据ID获取字典类型
func (s *DictTypeService) GetByID(id uint) (*model.DictType, error) {
	logger.Debug("查询字典类型",
		zap.Uint("id", id),
		zap.String("operation", "get_dict_type"))

	dictType, err := s.dictTypeRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("字典类型不存在",
				zap.Uint("id", id),
				zap.String("operation", "get_dict_type"))
		} else {
			logger.Error("查询字典类型失败",
				zap.Uint("id", id),
				zap.Error(err),
				zap.String("operation", "get_dict_type"))
			return nil, apperrors.NewDictTypeNotFoundError()
		}
		return nil, apperrors.NewDictTypeGetFailedError()
	}

	return dictType, nil
}

// GetByCode 根据代码获取字典类型
func (s *DictTypeService) GetByCode(code string) (*model.DictType, error) {
	logger.Debug("根据代码查询字典类型",
		zap.String("code", code),
		zap.String("operation", "get_dict_type_by_code"))

	dictType, err := s.dictTypeRepo.GetByCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("字典类型不存在",
				zap.String("code", code),
				zap.String("operation", "get_dict_type_by_code"))
		} else {
			logger.Error("查询字典类型失败",
				zap.String("code", code),
				zap.Error(err),
				zap.String("operation", "get_dict_type_by_code"))
			return nil, apperrors.NewDictTypeNotFoundError()
		}
		return nil, apperrors.NewDictTypeGetFailedError()
	}

	return dictType, nil
}

// Update 更新字典类型
func (s *DictTypeService) Update(id uint, req *model.UpdateDictTypeRequest) (*model.DictType, error) {
	logger.Info("开始更新字典类型",
		zap.Uint("id", id),
		zap.String("operation", "update_dict_type"))

	// 查询字典类型
	dictType, err := s.dictTypeRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("字典类型不存在",
				zap.Uint("id", id),
				zap.String("operation", "update_dict_type"))
		} else {
			logger.Error("查询字典类型失败",
				zap.Uint("id", id),
				zap.Error(err),
				zap.String("operation", "update_dict_type"))
			return nil, apperrors.NewDictTypeNotFoundError()
		}
		return nil, apperrors.NewDictTypeGetFailedError()
	}

	// 更新字段
	if req.Name != "" {
		dictType.Name = req.Name
	}
	if req.Description != "" {
		dictType.Description = req.Description
	}
	if req.Status != "" {
		dictType.Status = req.Status
	}
	if req.SortOrder >= 0 {
		dictType.SortOrder = req.SortOrder
	}

	err = s.dictTypeRepo.Update(dictType)
	if err != nil {
		logger.Error("更新字典类型失败",
			zap.Uint("id", id),
			zap.Error(err),
			zap.String("operation", "update_dict_type"))
		return nil, apperrors.NewDictTypeUpdateFailedError()
	}

	logger.Info("字典类型更新成功",
		zap.Uint("id", id),
		zap.String("code", dictType.Code),
		zap.String("operation", "update_dict_type"))

	return dictType, nil
}

// Delete 删除字典类型
func (s *DictTypeService) Delete(id uint) error {
	logger.Info("开始删除字典类型",
		zap.Uint("id", id),
		zap.String("operation", "delete_dict_type"))

	// 查询字典类型
	dictType, err := s.dictTypeRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("字典类型不存在",
				zap.Uint("id", id),
				zap.String("operation", "delete_dict_type"))
		} else {
			logger.Error("查询字典类型失败",
				zap.Uint("id", id),
				zap.Error(err),
				zap.String("operation", "delete_dict_type"))
			return apperrors.NewDictTypeNotFoundError()
		}
		return apperrors.NewDictTypeGetFailedError()
	}

	// 检查是否为系统内置
	if dictType.IsSystem {
		logger.Warn("系统内置字典类型不可删除",
			zap.Uint("id", id),
			zap.String("code", dictType.Code),
			zap.String("operation", "delete_dict_type"))
		return apperrors.NewPermissionDeniedErrorWithCode("系统内置字典类型不可删除")
	}

	// 检查是否有关联的字典项
	count, err := s.dictItemRepo.CountByTypeCode(dictType.Code)
	if err != nil {
		logger.Error("统计字典项数量失败",
			zap.String("code", dictType.Code),
			zap.Error(err),
			zap.String("operation", "delete_dict_type"))
		return apperrors.NewDictTypeCountFailedError()
	}
	if count > 0 {
		logger.Warn("字典类型下存在字典项，不可删除",
			zap.Uint("id", id),
			zap.String("code", dictType.Code),
			zap.Int64("item_count", count),
			zap.String("operation", "delete_dict_type"))
		return apperrors.NewDictTypeInUseError(count)
	}

	err = s.dictTypeRepo.Delete(id)
	if err != nil {
		logger.Error("删除字典类型失败",
			zap.Uint("id", id),
			zap.Error(err),
			zap.String("operation", "delete_dict_type"))
		return apperrors.NewDictTypeDeleteFailedError()
	}

	logger.Info("字典类型删除成功",
		zap.Uint("id", id),
		zap.String("code", dictType.Code),
		zap.String("operation", "delete_dict_type"))

	return nil
}

// List 获取字典类型列表
func (s *DictTypeService) List(page, pageSize int, status, keyword string) ([]model.DictType, int64, error) {
	logger.Debug("查询字典类型列表",
		zap.Int("page", page),
		zap.Int("page_size", pageSize),
		zap.String("status", status),
		zap.String("keyword", keyword),
		zap.String("operation", "list_dict_types"))

	offset := (page - 1) * pageSize
	dictTypes, total, err := s.dictTypeRepo.List(offset, pageSize, status, keyword)
	if err != nil {
		logger.Error("查询字典类型列表失败",
			zap.Int("page", page),
			zap.Int("page_size", pageSize),
			zap.Error(err),
			zap.String("operation", "list_dict_types"))
		return nil, 0, apperrors.NewDictTypeListFailedError()
	}

	logger.Debug("字典类型列表查询成功",
		zap.Int("page", page),
		zap.Int("page_size", pageSize),
		zap.Int64("total", total),
		zap.Int("returned_count", len(dictTypes)),
		zap.String("operation", "list_dict_types"))

	return dictTypes, total, nil
}

// GetAll 获取所有字典类型
func (s *DictTypeService) GetAll() ([]model.DictType, error) {
	logger.Debug("查询所有字典类型", zap.String("operation", "get_all_dict_types"))

	dictTypes, err := s.dictTypeRepo.GetAll()
	if err != nil {
		logger.Error("查询所有字典类型失败",
			zap.Error(err),
			zap.String("operation", "get_all_dict_types"))
		return nil, apperrors.NewDictTypeListFailedError()
	}

	return dictTypes, nil
}

// ==================== DictItemService ====================

// DictItemService 字典项业务服务
type DictItemService struct {
	dictTypeRepo DictTypeRepositoryInterface
	dictItemRepo DictItemRepositoryInterface
}

// NewDictItemService 创建 DictItemService 实例
func NewDictItemService(
	dictTypeRepo DictTypeRepositoryInterface,
	dictItemRepo DictItemRepositoryInterface,
) *DictItemService {
	return &DictItemService{
		dictTypeRepo: dictTypeRepo,
		dictItemRepo: dictItemRepo,
	}
}

// Create 创建字典项
func (s *DictItemService) Create(req *model.CreateDictItemRequest) (*model.DictItem, error) {
	logger.Info("开始创建字典项",
		zap.String("dict_type_code", req.DictTypeCode),
		zap.String("value", req.Value),
		zap.String("operation", "create_dict_item"))

	// 检查字典类型是否存在
	_, err := s.dictTypeRepo.GetByCode(req.DictTypeCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("字典类型不存在",
				zap.String("dict_type_code", req.DictTypeCode),
				zap.String("operation", "create_dict_item"))
			return nil, apperrors.NewDictTypeNotFoundError()
		}
		logger.Error("查询字典类型失败",
			zap.String("dict_type_code", req.DictTypeCode),
			zap.Error(err),
			zap.String("operation", "create_dict_item"))
		return nil, apperrors.NewDictItemTypeGetFailedError()
	}

	// 检查值是否已存在
	exists, err := s.dictItemRepo.CheckValueExists(req.DictTypeCode, req.Value)
	if err != nil {
		logger.Error("检查字典项值是否存在失败",
			zap.String("dict_type_code", req.DictTypeCode),
			zap.String("value", req.Value),
			zap.Error(err),
			zap.String("operation", "create_dict_item"))
		return nil, apperrors.NewDictItemValueCheckFailedError()
	}
	if exists {
		logger.Warn("字典项值已存在",
			zap.String("dict_type_code", req.DictTypeCode),
			zap.String("value", req.Value),
			zap.String("operation", "create_dict_item"))
		return nil, apperrors.NewDictItemExistsError()
	}

	// 转换 Extra
	extraJSON, err := repository.ConvertMapToJSON(req.Extra)
	if err != nil {
		logger.Error("转换Extra失败",
			zap.Error(err),
			zap.String("operation", "create_dict_item"))
		return nil, apperrors.NewValidationError("Extra 格式错误")
	}

	// 如果设置为默认值，清除其他默认值
	if req.IsDefault {
		if err := s.dictItemRepo.ClearDefaultByType(req.DictTypeCode); err != nil {
			logger.Error("清除默认值失败",
				zap.String("dict_type_code", req.DictTypeCode),
				zap.Error(err),
				zap.String("operation", "create_dict_item"))
			return nil, apperrors.NewDictItemDefaultClearFailedError()
		}
	}

	// 创建字典项
	dictItem := &model.DictItem{
		DictTypeCode: req.DictTypeCode,
		Label:        req.Label,
		Value:        req.Value,
		Extra:        extraJSON,
		Description:  req.Description,
		Status:       req.Status,
		SortOrder:    req.SortOrder,
		IsDefault:    req.IsDefault,
		IsSystem:     false, // 用户创建的都不是系统内置
	}

	if dictItem.Status == "" {
		dictItem.Status = "active"
	}

	err = s.dictItemRepo.Create(dictItem)
	if err != nil {
		logger.Error("创建字典项失败",
			zap.String("dict_type_code", req.DictTypeCode),
			zap.String("value", req.Value),
			zap.Error(err),
			zap.String("operation", "create_dict_item"))
		return nil, apperrors.NewDictItemCreateFailedError()
	}

	logger.Info("字典项创建成功",
		zap.Uint("id", dictItem.ID),
		zap.String("dict_type_code", dictItem.DictTypeCode),
		zap.String("value", dictItem.Value),
		zap.String("operation", "create_dict_item"))

	return dictItem, nil
}

// GetByID 根据ID获取字典项
func (s *DictItemService) GetByID(id uint) (*model.DictItem, error) {
	logger.Debug("查询字典项",
		zap.Uint("id", id),
		zap.String("operation", "get_dict_item"))

	dictItem, err := s.dictItemRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("字典项不存在",
				zap.Uint("id", id),
				zap.String("operation", "get_dict_item"))
			return nil, apperrors.NewDictItemNotFoundError()
		}
		logger.Error("查询字典项失败",
			zap.Uint("id", id),
			zap.Error(err),
			zap.String("operation", "get_dict_item"))
		return nil, apperrors.NewDictItemGetFailedError()
	}

	return dictItem, nil
}

// Update 更新字典项
func (s *DictItemService) Update(id uint, req *model.UpdateDictItemRequest) (*model.DictItem, error) {
	logger.Info("开始更新字典项",
		zap.Uint("id", id),
		zap.String("operation", "update_dict_item"))

	// 查询字典项
	dictItem, err := s.dictItemRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("字典项不存在",
				zap.Uint("id", id),
				zap.String("operation", "update_dict_item"))
		} else {
			logger.Error("查询字典项失败",
				zap.Uint("id", id),
				zap.Error(err),
				zap.String("operation", "update_dict_item"))
			return nil, apperrors.NewDictItemNotFoundError()
		}
		return nil, apperrors.NewDictItemGetFailedError()
	}

	// 更新字段
	if req.Label != "" {
		dictItem.Label = req.Label
	}
	if req.Extra != nil {
		extraJSON, err := repository.ConvertMapToJSON(req.Extra)
		if err != nil {
			logger.Error("转换Extra失败",
				zap.Error(err),
				zap.String("operation", "update_dict_item"))
			return nil, apperrors.NewValidationError("Extra 格式错误")
		}
		dictItem.Extra = extraJSON
	}
	if req.Description != "" {
		dictItem.Description = req.Description
	}
	if req.Status != "" {
		dictItem.Status = req.Status
	}
	if req.SortOrder >= 0 {
		dictItem.SortOrder = req.SortOrder
	}

	// 如果设置为默认值，清除其他默认值
	if req.IsDefault && !dictItem.IsDefault {
		if err := s.dictItemRepo.ClearDefaultByType(dictItem.DictTypeCode); err != nil {
			logger.Error("清除默认值失败",
				zap.String("dict_type_code", dictItem.DictTypeCode),
				zap.Error(err),
				zap.String("operation", "update_dict_item"))
			return nil, apperrors.NewDictItemDefaultClearFailedError()
		}
		dictItem.IsDefault = true
	} else if !req.IsDefault && dictItem.IsDefault {
		dictItem.IsDefault = false
	}

	err = s.dictItemRepo.Update(dictItem)
	if err != nil {
		logger.Error("更新字典项失败",
			zap.Uint("id", id),
			zap.Error(err),
			zap.String("operation", "update_dict_item"))
		return nil, apperrors.NewDictItemUpdateFailedError()
	}

	logger.Info("字典项更新成功",
		zap.Uint("id", id),
		zap.String("dict_type_code", dictItem.DictTypeCode),
		zap.String("operation", "update_dict_item"))

	return dictItem, nil
}

// Delete 删除字典项
func (s *DictItemService) Delete(id uint) error {
	logger.Info("开始删除字典项",
		zap.Uint("id", id),
		zap.String("operation", "delete_dict_item"))

	// 查询字典项
	dictItem, err := s.dictItemRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("字典项不存在",
				zap.Uint("id", id),
				zap.String("operation", "delete_dict_item"))
		} else {
			logger.Error("查询字典项失败",
				zap.Uint("id", id),
				zap.Error(err),
				zap.String("operation", "delete_dict_item"))
			return apperrors.NewDictItemNotFoundError()
		}
		return apperrors.NewDictItemGetFailedError()
	}

	// 检查是否为系统内置
	if dictItem.IsSystem {
		logger.Warn("系统内置字典项不可删除",
			zap.Uint("id", id),
			zap.String("value", dictItem.Value),
			zap.String("operation", "delete_dict_item"))
		return apperrors.NewPermissionDeniedErrorWithCode("系统内置字典项不可删除")
	}

	err = s.dictItemRepo.Delete(id)
	if err != nil {
		logger.Error("删除字典项失败",
			zap.Uint("id", id),
			zap.Error(err),
			zap.String("operation", "delete_dict_item"))
		return apperrors.NewDictItemDeleteFailedError()
	}

	logger.Info("字典项删除成功",
		zap.Uint("id", id),
		zap.String("dict_type_code", dictItem.DictTypeCode),
		zap.String("operation", "delete_dict_item"))

	return nil
}

// List 获取字典项列表
func (s *DictItemService) List(page, pageSize int, typeCode, status string) ([]model.DictItem, int64, error) {
	logger.Debug("查询字典项列表",
		zap.Int("page", page),
		zap.Int("page_size", pageSize),
		zap.String("dict_type_code", typeCode),
		zap.String("status", status),
		zap.String("operation", "list_dict_items"))

	offset := (page - 1) * pageSize
	dictItems, total, err := s.dictItemRepo.List(offset, pageSize, typeCode, status)
	if err != nil {
		logger.Error("查询字典项列表失败",
			zap.Int("page", page),
			zap.Int("page_size", pageSize),
			zap.Error(err),
			zap.String("operation", "list_dict_items"))
		return nil, 0, apperrors.NewDictItemListFailedError()
	}

	logger.Debug("字典项列表查询成功",
		zap.Int("page", page),
		zap.Int("page_size", pageSize),
		zap.Int64("total", total),
		zap.Int("returned_count", len(dictItems)),
		zap.String("operation", "list_dict_items"))

	return dictItems, total, nil
}

// GetByTypeCode 根据类型代码获取字典项
func (s *DictItemService) GetByTypeCode(typeCode string, activeOnly bool) ([]model.DictItem, error) {
	logger.Debug("根据类型代码查询字典项",
		zap.String("dict_type_code", typeCode),
		zap.Bool("active_only", activeOnly),
		zap.String("operation", "get_dict_items_by_type"))

	// 检查字典类型是否存在
	_, err := s.dictTypeRepo.GetByCode(typeCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Warn("字典类型不存在",
				zap.String("dict_type_code", typeCode),
				zap.String("operation", "get_dict_items_by_type"))
			return nil, apperrors.NewDictTypeNotFoundError()
		}
		logger.Error("查询字典类型失败",
			zap.String("dict_type_code", typeCode),
			zap.Error(err),
			zap.String("operation", "get_dict_items_by_type"))
		return nil, apperrors.NewDictItemTypeGetFailedError()
	}

	dictItems, err := s.dictItemRepo.GetByTypeCode(typeCode, activeOnly)
	if err != nil {
		logger.Error("根据类型获取字典项失败",
			zap.String("dict_type_code", typeCode),
			zap.Error(err),
			zap.String("operation", "get_dict_items_by_type"))
		return nil, apperrors.NewDictItemsByTypeGetFailedError()
	}

	return dictItems, nil
}
