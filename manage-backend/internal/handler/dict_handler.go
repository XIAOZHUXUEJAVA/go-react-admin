package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/repository"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/service"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
)

// DictTypeHandler 字典类型处理器
type DictTypeHandler struct {
	dictTypeService *service.DictTypeService
}

// NewDictTypeHandler 创建 DictTypeHandler 实例
func NewDictTypeHandler(dictTypeService *service.DictTypeService) *DictTypeHandler {
	return &DictTypeHandler{dictTypeService: dictTypeService}
}

// CreateDictType godoc
// @Summary 创建字典类型
// @Description 创建新的字典类型
// @Tags dict-types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param dict_type body model.CreateDictTypeRequest true "字典类型信息"
// @Success 201 {object} utils.APIResponse{data=model.DictTypeResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /dict-types [post]
func (h *DictTypeHandler) CreateDictType(c *gin.Context) {
	var req model.CreateDictTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	dictType, err := h.dictTypeService.Create(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Created(c, dictType)
}

// GetDictType godoc
// @Summary 获取字典类型详情
// @Description 根据ID获取字典类型详情
// @Tags dict-types
// @Produce json
// @Security BearerAuth
// @Param id path int true "字典类型ID"
// @Success 200 {object} utils.APIResponse{data=model.DictTypeResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /dict-types/{id} [get]
func (h *DictTypeHandler) GetDictType(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的字典类型ID")
		return
	}

	dictType, err := h.dictTypeService.GetByID(uint(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, dictType)
}

// UpdateDictType godoc
// @Summary 更新字典类型
// @Description 更新字典类型信息
// @Tags dict-types
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "字典类型ID"
// @Param dict_type body model.UpdateDictTypeRequest true "字典类型信息"
// @Success 200 {object} utils.APIResponse{data=model.DictTypeResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /dict-types/{id} [put]
func (h *DictTypeHandler) UpdateDictType(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的字典类型ID")
		return
	}

	var req model.UpdateDictTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	dictType, err := h.dictTypeService.Update(uint(id), &req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, dictType)
}

// DeleteDictType godoc
// @Summary 删除字典类型
// @Description 删除字典类型
// @Tags dict-types
// @Produce json
// @Security BearerAuth
// @Param id path int true "字典类型ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /dict-types/{id} [delete]
func (h *DictTypeHandler) DeleteDictType(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的字典类型ID")
		return
	}

	err = h.dictTypeService.Delete(uint(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, nil)
}

// ListDictTypes godoc
// @Summary 获取字典类型列表
// @Description 获取字典类型列表（分页）
// @Tags dict-types
// @Produce json
// @Security BearerAuth
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param status query string false "状态筛选" Enums(active, inactive)
// @Param keyword query string false "关键字搜索"
// @Success 200 {object} utils.APIResponse{data=[]model.DictTypeResponse}
// @Failure 400 {object} utils.APIResponse
// @Router /dict-types [get]
func (h *DictTypeHandler) ListDictTypes(c *gin.Context) {
	var req model.DictTypeListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	dictTypes, total, err := h.dictTypeService.List(req.Page, req.PageSize, req.Status, req.Keyword)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.SuccessWithPagination(c, dictTypes, req.Page, req.PageSize, int(total))
}

// GetAllDictTypes godoc
// @Summary 获取所有字典类型
// @Description 获取所有启用的字典类型（不分页）
// @Tags dict-types
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse{data=[]model.DictTypeResponse}
// @Failure 500 {object} utils.APIResponse
// @Router /dict-types/all [get]
func (h *DictTypeHandler) GetAllDictTypes(c *gin.Context) {
	dictTypes, err := h.dictTypeService.GetAll()
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	utils.Success(c, dictTypes)
}

// ==================== DictItemHandler ====================

// DictItemHandler 字典项处理器
type DictItemHandler struct {
	dictItemService *service.DictItemService
}

// NewDictItemHandler 创建 DictItemHandler 实例
func NewDictItemHandler(dictItemService *service.DictItemService) *DictItemHandler {
	return &DictItemHandler{dictItemService: dictItemService}
}

// CreateDictItem godoc
// @Summary 创建字典项
// @Description 创建新的字典项
// @Tags dict-items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param dict_item body model.CreateDictItemRequest true "字典项信息"
// @Success 201 {object} utils.APIResponse{data=model.DictItemResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /dict-items [post]
func (h *DictItemHandler) CreateDictItem(c *gin.Context) {
	var req model.CreateDictItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	dictItem, err := h.dictItemService.Create(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	// 转换为响应格式
	response := repository.ConvertDictItemToResponse(dictItem)
	utils.Created(c, response)
}

// GetDictItem godoc
// @Summary 获取字典项详情
// @Description 根据ID获取字典项详情
// @Tags dict-items
// @Produce json
// @Security BearerAuth
// @Param id path int true "字典项ID"
// @Success 200 {object} utils.APIResponse{data=model.DictItemResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /dict-items/{id} [get]
func (h *DictItemHandler) GetDictItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的字典项ID")
		return
	}

	dictItem, err := h.dictItemService.GetByID(uint(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	// 转换为响应格式
	response := repository.ConvertDictItemToResponse(dictItem)
	utils.Success(c, response)
}

// UpdateDictItem godoc
// @Summary 更新字典项
// @Description 更新字典项信息
// @Tags dict-items
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "字典项ID"
// @Param dict_item body model.UpdateDictItemRequest true "字典项信息"
// @Success 200 {object} utils.APIResponse{data=model.DictItemResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /dict-items/{id} [put]
func (h *DictItemHandler) UpdateDictItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的字典项ID")
		return
	}

	var req model.UpdateDictItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	dictItem, err := h.dictItemService.Update(uint(id), &req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	// 转换为响应格式
	response := repository.ConvertDictItemToResponse(dictItem)
	utils.Success(c, response)
}

// DeleteDictItem godoc
// @Summary 删除字典项
// @Description 删除字典项
// @Tags dict-items
// @Produce json
// @Security BearerAuth
// @Param id path int true "字典项ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /dict-items/{id} [delete]
func (h *DictItemHandler) DeleteDictItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的字典项ID")
		return
	}

	err = h.dictItemService.Delete(uint(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, nil)
}

// ListDictItems godoc
// @Summary 获取字典项列表
// @Description 获取字典项列表（分页）
// @Tags dict-items
// @Produce json
// @Security BearerAuth
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param dict_type_code query string false "字典类型代码"
// @Param status query string false "状态筛选" Enums(active, inactive)
// @Success 200 {object} utils.APIResponse{data=[]model.DictItemResponse}
// @Failure 400 {object} utils.APIResponse
// @Router /dict-items [get]
func (h *DictItemHandler) ListDictItems(c *gin.Context) {
	var req model.DictItemListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	dictItems, total, err := h.dictItemService.List(req.Page, req.PageSize, req.DictTypeCode, req.Status)
	if err != nil {
		utils.InternalServerError(c, err.Error())
		return
	}

	// 转换为响应格式
	responses := make([]*model.DictItemResponse, len(dictItems))
	for i, item := range dictItems {
		responses[i] = repository.ConvertDictItemToResponse(&item)
	}

	utils.SuccessWithPagination(c, responses, req.Page, req.PageSize, int(total))
}

// GetDictItemsByType godoc
// @Summary 根据类型获取字典项
// @Description 根据字典类型代码获取所有字典项（不分页）
// @Tags dict-items
// @Produce json
// @Security BearerAuth
// @Param code path string true "字典类型代码"
// @Param active_only query bool false "仅返回启用的" default(true)
// @Success 200 {object} utils.APIResponse{data=[]model.DictItemResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /dict-types/{code}/items [get]
func (h *DictItemHandler) GetDictItemsByType(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		utils.BadRequest(c, "字典类型代码不能为空")
		return
	}

	// 获取 active_only 参数，默认为 true
	activeOnly := true
	if activeOnlyStr := c.Query("active_only"); activeOnlyStr == "false" {
		activeOnly = false
	}

	dictItems, err := h.dictItemService.GetByTypeCode(code, activeOnly)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	// 转换为响应格式
	responses := make([]*model.DictItemResponse, len(dictItems))
	for i, item := range dictItems {
		responses[i] = repository.ConvertDictItemToResponse(&item)
	}

	utils.Success(c, responses)
}
