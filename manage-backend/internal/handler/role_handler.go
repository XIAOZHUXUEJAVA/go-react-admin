package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/service"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
)

type RoleHandler struct {
	roleService *service.RoleService
}

func NewRoleHandler(roleService *service.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService}
}

// CreateRole godoc
// @Summary 创建角色
// @Description 创建新角色
// @Tags roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param role body model.CreateRoleRequest true "角色信息"
// @Success 201 {object} utils.APIResponse{data=model.RoleResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 401 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /roles [post]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req model.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	role, err := h.roleService.Create(&req)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Created(c, role)
}

// GetRole godoc
// @Summary 获取角色详情
// @Description 根据ID获取角色详情
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Param id path int true "角色ID"
// @Success 200 {object} utils.APIResponse{data=model.RoleResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /roles/{id} [get]
func (h *RoleHandler) GetRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的角色ID")
		return
	}

	role, err := h.roleService.GetByID(uint(id))
	if err != nil {
		utils.NotFound(c, err.Error())
		return
	}

	utils.Success(c, role)
}

// UpdateRole godoc
// @Summary 更新角色
// @Description 更新角色信息
// @Tags roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "角色ID"
// @Param role body model.UpdateRoleRequest true "角色信息"
// @Success 200 {object} utils.APIResponse{data=model.RoleResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /roles/{id} [put]
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的角色ID")
		return
	}

	var req model.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	role, err := h.roleService.Update(uint(id), &req)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, role)
}

// DeleteRole godoc
// @Summary 删除角色
// @Description 删除角色
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Param id path int true "角色ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /roles/{id} [delete]
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的角色ID")
		return
	}

	if err := h.roleService.Delete(uint(id)); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "角色删除成功"})
}

// ListRoles godoc
// @Summary 获取角色列表
// @Description 分页获取角色列表
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} utils.PaginatedResponse{data=[]model.RoleResponse}
// @Failure 400 {object} utils.APIResponse
// @Router /roles [get]
func (h *RoleHandler) ListRoles(c *gin.Context) {
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

	roles, total, err := h.roleService.List(page, pageSize)
	if err != nil {
		utils.InternalServerError(c, "获取角色列表失败")
		return
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))
	pagination := utils.PaginationMeta{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}

	utils.PaginatedSuccess(c, roles, pagination)
}

// GetAllRoles godoc
// @Summary 获取所有角色
// @Description 获取所有角色（不分页）
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse{data=[]model.RoleResponse}
// @Failure 500 {object} utils.APIResponse
// @Router /roles/all [get]
func (h *RoleHandler) GetAllRoles(c *gin.Context) {
	roles, err := h.roleService.GetAll()
	if err != nil {
		utils.InternalServerError(c, "获取角色列表失败")
		return
	}

	utils.Success(c, roles)
}

// AssignPermissions godoc
// @Summary 为角色分配权限
// @Description 为角色分配权限列表
// @Tags roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "角色ID"
// @Param permissions body model.AssignRolePermissionsRequest true "权限ID列表"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /roles/{id}/permissions [put]
func (h *RoleHandler) AssignPermissions(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的角色ID")
		return
	}

	var req model.AssignRolePermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	if err := h.roleService.AssignPermissions(uint(id), &req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "权限分配成功"})
}

// GetRolePermissions godoc
// @Summary 获取角色权限
// @Description 获取角色的权限列表
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Param id path int true "角色ID"
// @Success 200 {object} utils.APIResponse{data=model.RoleWithPermissions}
// @Failure 400 {object} utils.APIResponse
// @Router /roles/{id}/permissions [get]
func (h *RoleHandler) GetRolePermissions(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的角色ID")
		return
	}

	roleWithPerms, err := h.roleService.GetRolePermissions(uint(id))
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, roleWithPerms)
}

// AssignRolesToUser godoc
// @Summary 为用户分配角色
// @Description 为用户分配角色列表
// @Tags roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "用户ID"
// @Param roles body model.AssignUserRolesRequest true "角色ID列表"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /users/{id}/roles [put]
func (h *RoleHandler) AssignRolesToUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的用户ID")
		return
	}

	var req model.AssignUserRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	// 获取当前操作用户ID
	assignedBy := c.GetUint("user_id")

	if err := h.roleService.AssignRolesToUser(uint(userID), &req, assignedBy); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.Success(c, gin.H{"message": "角色分配成功"})
}

// GetUserRoles godoc
// @Summary 获取用户角色
// @Description 获取用户的角色列表
// @Tags roles
// @Produce json
// @Security BearerAuth
// @Param id path int true "用户ID"
// @Success 200 {object} utils.APIResponse{data=[]model.RoleResponse}
// @Failure 400 {object} utils.APIResponse
// @Router /users/{id}/roles [get]
func (h *RoleHandler) GetUserRoles(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的用户ID")
		return
	}

	roles, err := h.roleService.GetUserRoles(uint(userID))
	if err != nil {
		utils.InternalServerError(c, "获取用户角色失败")
		return
	}

	utils.Success(c, roles)
}
