package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/service"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
	apperrors "github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/pkg/errors"
)

type MenuHandler struct {
	menuService *service.MenuService
	roleRepo    service.RoleRepositoryInterface
}

func NewMenuHandler(menuService *service.MenuService, roleRepo service.RoleRepositoryInterface) *MenuHandler {
	return &MenuHandler{
		menuService: menuService,
		roleRepo:    roleRepo,
	}
}

// CreateMenu godoc
// @Summary 创建菜单
// @Description 创建新菜单
// @Tags menus
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param menu body model.CreateMenuRequest true "菜单信息"
// @Success 201 {object} utils.APIResponse{data=model.MenuResponse}
// @Failure 400 {object} utils.APIResponse
// @Router /menus [post]
func (h *MenuHandler) CreateMenu(c *gin.Context) {
	var req model.CreateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	menu, err := h.menuService.Create(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Created(c, menu)
}

// GetMenu godoc
// @Summary 获取菜单详情
// @Description 根据ID获取菜单详情
// @Tags menus
// @Produce json
// @Security BearerAuth
// @Param id path int true "菜单ID"
// @Success 200 {object} utils.APIResponse{data=model.MenuResponse}
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /menus/{id} [get]
func (h *MenuHandler) GetMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的菜单ID")
		return
	}

	menu, err := h.menuService.GetByID(uint(id))
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, menu)
}

// UpdateMenu godoc
// @Summary 更新菜单
// @Description 更新菜单信息
// @Tags menus
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "菜单ID"
// @Param menu body model.UpdateMenuRequest true "菜单信息"
// @Success 200 {object} utils.APIResponse{data=model.MenuResponse}
// @Failure 400 {object} utils.APIResponse
// @Router /menus/{id} [put]
func (h *MenuHandler) UpdateMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的菜单ID")
		return
	}

	var req model.UpdateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	menu, err := h.menuService.Update(uint(id), &req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, menu)
}

// DeleteMenu godoc
// @Summary 删除菜单
// @Description 删除菜单
// @Tags menus
// @Produce json
// @Security BearerAuth
// @Param id path int true "菜单ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /menus/{id} [delete]
func (h *MenuHandler) DeleteMenu(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "无效的菜单ID")
		return
	}

	if err := h.menuService.Delete(uint(id)); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, gin.H{"message": "菜单删除成功"})
}

// GetMenuTree godoc
// @Summary 获取菜单树
// @Description 获取完整的菜单树结构
// @Tags menus
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse{data=[]model.MenuResponse}
// @Failure 500 {object} utils.APIResponse
// @Router /menus/tree [get]
func (h *MenuHandler) GetMenuTree(c *gin.Context) {
	tree, err := h.menuService.GetMenuTree()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, tree)
}

// GetVisibleMenuTree godoc
// @Summary 获取可见菜单树
// @Description 获取所有可见的菜单树
// @Tags menus
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse{data=[]model.MenuResponse}
// @Failure 500 {object} utils.APIResponse
// @Router /menus/tree/visible [get]
func (h *MenuHandler) GetVisibleMenuTree(c *gin.Context) {
	tree, err := h.menuService.GetVisibleMenuTree()
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, tree)
}

// GetUserMenuTree godoc
// @Summary 获取用户菜单树
// @Description 根据用户权限获取菜单树
// @Tags menus
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse{data=[]model.MenuResponse}
// @Failure 401 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /menus/user [get]
func (h *MenuHandler) GetUserMenuTree(c *gin.Context) {
	userID := c.GetUint("user_id")
	if userID == 0 {
		utils.HandleError(c, apperrors.NewUnauthorizedErrorWithCode(""))
		return
	}

	tree, err := h.menuService.GetUserMenuTree(userID, h.roleRepo)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, tree)
}

// UpdateMenuOrder godoc
// @Summary 批量更新菜单顺序
// @Description 批量更新菜单的排序和父级关系
// @Tags menus
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param menus body []model.MenuOrderUpdate true "菜单顺序列表"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /menus/order [put]
func (h *MenuHandler) UpdateMenuOrder(c *gin.Context) {
	var req struct {
		Menus []model.MenuOrderUpdate `json:"menus" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	if err := h.menuService.UpdateMenuOrder(req.Menus); err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.Success(c, gin.H{"message": "顺序更新成功"})
}
