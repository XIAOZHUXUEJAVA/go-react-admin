package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/service"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
)

type PermissionHandler struct {
	permissionService *service.PermissionService
}

func NewPermissionHandler(permissionService *service.PermissionService) *PermissionHandler {
	return &PermissionHandler{permissionService: permissionService}
}

// CreatePermission godoc
// @Summary 创建权限
// @Description 创建新权限
// @Tags permissions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param permission body model.CreatePermissionRequest true "权限信息"
// @Success 201 {object} utils.APIResponse{data=model.PermissionResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Router /permissions [post]
func (h *PermissionHandler) CreatePermission(c *gin.Context) {
	var req model.CreatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	permission, err := h.permissionService.Create(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Created(c, permission)
}

// GetPermission godoc
// @Summary 获取权限详情
// @Description 根据ID获取权限详情
// @Tags permissions
// @Produce json
// @Security BearerAuth
// @Param id path int true "权限ID"
// @Success 200 {object} utils.APIResponse{data=model.PermissionResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /permissions/{id} [get]
func (h *PermissionHandler) GetPermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的权限ID")
		return
	}

	permission, err := h.permissionService.GetByID(uint(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, permission)
}

// UpdatePermission godoc
// @Summary 更新权限
// @Description 更新权限信息
// @Tags permissions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "权限ID"
// @Param permission body model.UpdatePermissionRequest true "权限信息"
// @Success 200 {object} utils.APIResponse{data=model.PermissionResponse}
// @Failure 400 {object} utils.APIResponse
// @Router /permissions/{id} [put]
func (h *PermissionHandler) UpdatePermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的权限ID")
		return
	}

	var req model.UpdatePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	permission, err := h.permissionService.Update(uint(id), &req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, permission)
}

// DeletePermission godoc
// @Summary 删除权限
// @Description 删除权限
// @Tags permissions
// @Produce json
// @Security BearerAuth
// @Param id path int true "权限ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /permissions/{id} [delete]
func (h *PermissionHandler) DeletePermission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的权限ID")
		return
	}

	if err := h.permissionService.Delete(uint(id)); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, gin.H{"message": "权限删除成功"})
}

// ListPermissions godoc
// @Summary 获取权限列表
// @Description 分页获取权限列表
// @Tags permissions
// @Produce json
// @Security BearerAuth
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} utils.PaginatedResponse{data=[]model.PermissionResponse}
// @Failure 400 {object} utils.APIResponse
// @Router /permissions [get]
func (h *PermissionHandler) ListPermissions(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	permissions, total, err := h.permissionService.List(page, pageSize)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))
	pagination := utils.PaginationMeta{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}

	utils.PaginatedSuccess(c, permissions, pagination)
}

// GetAllPermissions godoc
// @Summary 获取所有权限
// @Description 获取所有权限（不分页）
// @Tags permissions
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse{data=[]model.PermissionResponse}
// @Failure 500 {object} utils.APIResponse
// @Router /permissions/all [get]
func (h *PermissionHandler) GetAllPermissions(c *gin.Context) {
	permissions, err := h.permissionService.GetAll()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, permissions)
}

// GetPermissionsByResource godoc
// @Summary 根据资源获取权限
// @Description 根据资源类型获取权限列表
// @Tags permissions
// @Produce json
// @Security BearerAuth
// @Param resource path string true "资源类型"
// @Success 200 {object} utils.APIResponse{data=[]model.PermissionResponse}
// @Failure 500 {object} utils.APIResponse
// @Router /permissions/resource/{resource} [get]
func (h *PermissionHandler) GetPermissionsByResource(c *gin.Context) {
	resource := c.Param("resource")
	if resource == "" {
		utils.BadRequest(c, "资源类型不能为空")
		return
	}

	permissions, err := h.permissionService.GetByResource(resource)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, permissions)
}

// GetPermissionsByType godoc
// @Summary 根据类型获取权限
// @Description 根据权限类型获取权限列表
// @Tags permissions
// @Produce json
// @Security BearerAuth
// @Param type path string true "权限类型" Enums(api, menu, button)
// @Success 200 {object} utils.APIResponse{data=[]model.PermissionResponse}
// @Failure 500 {object} utils.APIResponse
// @Router /permissions/type/{type} [get]
func (h *PermissionHandler) GetPermissionsByType(c *gin.Context) {
	permType := c.Param("type")
	if permType == "" {
		utils.BadRequest(c, "权限类型不能为空")
		return
	}

	permissions, err := h.permissionService.GetByType(permType)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, permissions)
}

// GetPermissionTree godoc
// @Summary 获取权限树
// @Description 获取按资源分组的权限树
// @Tags permissions
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse{data=[]model.PermissionTree}
// @Failure 500 {object} utils.APIResponse
// @Router /permissions/tree [get]
func (h *PermissionHandler) GetPermissionTree(c *gin.Context) {
	tree, err := h.permissionService.GetPermissionTree()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, tree)
}
