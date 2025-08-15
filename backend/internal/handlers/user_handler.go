package handlers

import (
	"net/http"
	"strconv"

	"backend/internal/services"
	"backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: services.NewUserService(),
	}
}

// 獲取用戶列表
func (h *UserHandler) GetUsers(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)

	users, total, err := h.userService.GetUsers(page, limit)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "獲取用戶列表失敗")
		return
	}

	// 清除密碼字段
	for i := range users {
		users[i].Password = ""
	}

	utils.PaginatedSuccessResponse(c, users, page, limit, total)
}

// 根據 ID 獲取用戶
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "無效的用戶ID")
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "用戶不存在")
		return
	}

	// 清除密碼字段
	user.Password = ""

	utils.SuccessResponse(c, user)
}

// 創建用戶
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "請求參數錯誤: "+err.Error())
		return
	}

	// 設置默認值
	if req.Role == "" {
		req.Role = "user"
	}
	if req.Status == "" {
		req.Status = "active"
	}

	user, err := h.userService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 如果指定了角色和狀態，更新它們
	if req.Role != "user" || req.Status != "active" {
		updates := map[string]interface{}{
			"role":   req.Role,
			"status": req.Status,
		}
		user, err = h.userService.UpdateUser(user.ID, updates)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "更新用戶信息失敗")
			return
		}
	}

	// 清除密碼字段
	user.Password = ""

	utils.SuccessResponse(c, user)
}

// 更新用戶
type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "無效的用戶ID")
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "請求參數錯誤: "+err.Error())
		return
	}

	updates := make(map[string]interface{})

	if req.Username != "" {
		updates["username"] = req.Username
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Role != "" {
		updates["role"] = req.Role
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}
	if req.Password != "" {
		updates["password"] = req.Password
	}

	user, err := h.userService.UpdateUser(uint(id), updates)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 清除密碼字段
	user.Password = ""

	utils.SuccessResponse(c, user)
}

// 刪除用戶
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "無效的用戶ID")
		return
	}

	// 檢查是否為當前用戶
	currentUserID, exists := c.Get("user_id")
	if exists && currentUserID.(uint) == uint(id) {
		utils.ErrorResponse(c, http.StatusBadRequest, "不能刪除自己")
		return
	}

	err = h.userService.DeleteUser(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "刪除用戶失敗")
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "用戶刪除成功"})
}
