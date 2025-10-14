package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/model"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/service"
	"github.com/XIAOZHUXUEJAVA/go-manage-starter/manage-backend/internal/utils"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with username, email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.CreateUserRequest true "User registration data"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}

	user, err := h.userService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "bad request",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "created successfully",
		"data":    user,
	})
}

// Login godoc
// @Summary User login
// @Description Login with username, password and captcha verification
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body model.LoginRequest true "Login credentials with captcha"
// @Success 200 {object} utils.APIResponse{data=model.LoginResponse} "登录成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误"
// @Failure 401 {object} utils.APIResponse "认证失败或验证码错误"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request format")
		return
	}

	// Extract client information
	deviceInfo := c.GetHeader("X-Device-Info")
	if deviceInfo == "" {
		deviceInfo = "Unknown Device"
	}
	
	ipAddress := c.ClientIP()
	userAgent := c.GetHeader("User-Agent")

	response, err := h.userService.LoginWithContext(c.Request.Context(), &req, deviceInfo, ipAddress, userAgent)
	if err != nil {
		utils.Unauthorized(c, err.Error())
		return
	}

	utils.Success(c, response)
}

// RefreshToken godoc
// @Summary Refresh access token
// @Description Refresh access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param refresh body model.RefreshTokenRequest true "Refresh token request"
// @Success 200 {object} utils.APIResponse{data=model.RefreshTokenResponse} "刷新成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误"
// @Failure 401 {object} utils.APIResponse "刷新token无效"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /auth/refresh [post]
func (h *UserHandler) RefreshToken(c *gin.Context) {
	var req model.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request format")
		return
	}

	response, err := h.userService.RefreshToken(c.Request.Context(), &req)
	if err != nil {
		utils.Unauthorized(c, err.Error())
		return
	}

	utils.Success(c, response)
}

// Logout godoc
// @Summary User logout
// @Description Logout user and invalidate tokens
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param logout body model.LogoutRequest false "Logout request (refresh token optional)"
// @Success 200 {object} utils.APIResponse "登出成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误"
// @Failure 401 {object} utils.APIResponse "未授权"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /auth/logout [post]
func (h *UserHandler) Logout(c *gin.Context) {
	userID := c.GetUint("user_id")
	accessToken := c.GetString("access_token")

	var req model.LogoutRequest
	// Logout request body is optional
	c.ShouldBindJSON(&req)

	err := h.userService.Logout(c.Request.Context(), userID, accessToken, &req)
	if err != nil {
		utils.InternalServerError(c, "Failed to logout")
		return
	}

	utils.Success(c, gin.H{"message": "Logged out successfully"})
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get current user profile
// @Tags users
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.User
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	user, err := h.userService.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "not found",
			"error":   "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    user,
	})
}

// GetUserPermissions godoc
// @Summary Get user permissions
// @Description Get current user's permissions
// @Tags users
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.APIResponse{data=model.UserPermissionsResponse} "获取成功"
// @Failure 401 {object} utils.APIResponse "未授权"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /users/permissions [get]
func (h *UserHandler) GetUserPermissions(c *gin.Context) {
	userID := c.GetUint("user_id")
	
	permissions, err := h.userService.GetUserPermissions(userID)
	if err != nil {
		utils.InternalServerError(c, "failed to get user permissions")
		return
	}

	utils.Success(c, permissions)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Description Update current user profile
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body model.UpdateUserRequest true "User update data"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/profile [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Update(userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// ListUsers godoc
// @Summary List users
// @Description Get list of users with pagination (需要认证)
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(10)
// @Success 200 {object} utils.PaginatedResponse{data=[]model.UserResponse} "获取成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误"
// @Failure 401 {object} utils.APIResponse "未授权"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	// 解析查询参数
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
	if pageSize > 50 { // 限制最大每页数量
		pageSize = 50
	}

	// 调用服务层的 List 方法
	users, total, err := h.userService.List(page, pageSize)
	if err != nil {
		utils.InternalServerError(c, "failed to get user list")
		return
	}

	// 计算分页信息
	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))
	pagination := utils.PaginationMeta{
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: totalPages,
	}

	utils.PaginatedSuccess(c, users, pagination)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user (admin only)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body model.CreateUserRequest true "User creation data"
// @Success 201 {object} utils.APIResponse{data=model.User} "创建成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误"
// @Failure 401 {object} utils.APIResponse "未授权"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "invalid request data")
		return
	}

	user, err := h.userService.Register(&req)
	if err != nil {
		if err.Error() == "username already exists" || err.Error() == "email already exists" {
			utils.BadRequest(c, err.Error())
			return
		}
		utils.InternalServerError(c, "failed to create user")
		return
	}

	utils.Created(c, user)
}

// UpdateUser godoc
// @Summary Update user
// @Description Update user information
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Param user body model.UpdateUserRequest true "User update data"
// @Success 200 {object} utils.APIResponse{data=model.User} "更新成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误"
// @Failure 401 {object} utils.APIResponse "未授权"
// @Failure 404 {object} utils.APIResponse "用户不存在"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "invalid user id")
		return
	}

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "invalid request data")
		return
	}

	user, err := h.userService.Update(uint(id), &req)
	if err != nil {
		if err.Error() == "user not found" {
			utils.NotFound(c, "user not found")
			return
		}
		utils.InternalServerError(c, "failed to update user")
		return
	}

	utils.Success(c, user)
}

// CheckUsernameAvailable godoc
// @Summary Check username availability
// @Description Check if username is available for registration
// @Tags users
// @Produce json
// @Param username path string true "Username to check"
// @Success 200 {object} utils.APIResponse{data=model.SimpleAvailabilityResponse} "检查成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /users/check-username/{username} [get]
func (h *UserHandler) CheckUsernameAvailable(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		utils.BadRequest(c, "username is required")
		return
	}

	available, err := h.userService.CheckUsernameAvailable(username)
	if err != nil {
		utils.InternalServerError(c, "failed to check username availability")
		return
	}

	response := model.SimpleAvailabilityResponse{
		Available: available,
	}

	utils.Success(c, response)
}

// CheckEmailAvailable godoc
// @Summary Check email availability
// @Description Check if email is available for registration
// @Tags users
// @Produce json
// @Param email path string true "Email to check"
// @Success 200 {object} utils.APIResponse{data=model.SimpleAvailabilityResponse} "检查成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /users/check-email/{email} [get]
func (h *UserHandler) CheckEmailAvailable(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		utils.BadRequest(c, "email is required")
		return
	}

	available, err := h.userService.CheckEmailAvailable(email)
	if err != nil {
		utils.InternalServerError(c, "failed to check email availability")
		return
	}

	response := model.SimpleAvailabilityResponse{
		Available: available,
	}

	utils.Success(c, response)
}

// CheckUserDataAvailability godoc
// @Summary Check user data availability
// @Description Batch check username and email availability
// @Tags users
// @Accept json
// @Produce json
// @Param data body model.CheckAvailabilityRequest true "Data to check"
// @Success 200 {object} utils.APIResponse{data=model.CheckAvailabilityResponse} "检查成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /users/check-availability [post]
func (h *UserHandler) CheckUserDataAvailability(c *gin.Context) {
	var req model.CheckAvailabilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "invalid request data")
		return
	}

	// 至少需要检查一个字段
	if req.Username == "" && req.Email == "" {
		utils.BadRequest(c, "username or email is required")
		return
	}

	response, err := h.userService.CheckUserDataAvailability(&req)
	if err != nil {
		utils.InternalServerError(c, "failed to check data availability")
		return
	}

	utils.Success(c, response)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user by ID
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} utils.APIResponse "删除成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误"
// @Failure 401 {object} utils.APIResponse "未授权"
// @Failure 404 {object} utils.APIResponse "用户不存在"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "invalid user id")
		return
	}

	err = h.userService.Delete(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			utils.NotFound(c, "user not found")
			return
		}
		utils.InternalServerError(c, "failed to delete user")
		return
	}

	utils.Success(c, gin.H{"message": "user deleted successfully"})
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get user information by ID
// @Tags users
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} utils.APIResponse{data=model.User} "获取成功"
// @Failure 400 {object} utils.APIResponse "请求参数错误"
// @Failure 401 {object} utils.APIResponse "未授权"
// @Failure 404 {object} utils.APIResponse "用户不存在"
// @Failure 500 {object} utils.APIResponse "服务器内部错误"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequest(c, "invalid user id")
		return
	}

	user, err := h.userService.GetByID(uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			utils.NotFound(c, "user not found")
			return
		}
		utils.InternalServerError(c, "failed to get user")
		return
	}

	utils.Success(c, user)
}