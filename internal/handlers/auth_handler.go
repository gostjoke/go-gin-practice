package handlers

import (
	"net/http"

	"backend/internal/services"
	"backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userService: services.NewUserService(),
	}
}

// 登錄請求結構
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// 註冊請求結構
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// 登錄響應結構
type LoginResponse struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// 用戶登錄
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "請求參數錯誤: "+err.Error())
		return
	}

	user, token, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// 不返回密碼
	user.Password = ""

	utils.SuccessResponse(c, LoginResponse{
		User:  user,
		Token: token,
	})
}

// 用戶註冊
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "請求參數錯誤: "+err.Error())
		return
	}

	user, err := h.userService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 生成 token
	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "生成令牌失敗")
		return
	}

	// 不返回密碼
	user.Password = ""

	utils.SuccessResponse(c, LoginResponse{
		User:  user,
		Token: token,
	})
}

// 獲取當前用戶信息
func (h *AuthHandler) Profile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未登錄")
		return
	}

	user, err := h.userService.GetUserByID(userID.(uint))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "用戶不存在")
		return
	}

	// 不返回密碼
	user.Password = ""

	utils.SuccessResponse(c, user)
}

// 更新個人資料
type UpdateProfileRequest struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未登錄")
		return
	}

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "請求參數錯誤: "+err.Error())
		return
	}

	updates := make(map[string]interface{})
	if req.Username != "" {
		updates["username"] = req.Username
	}
	if req.Avatar != "" {
		updates["avatar"] = req.Avatar
	}

	user, err := h.userService.UpdateUser(userID.(uint), updates)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 不返回密碼
	user.Password = ""

	utils.SuccessResponse(c, user)
}

// 修改密碼
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

func (h *AuthHandler) ChangePassword(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未登錄")
		return
	}

	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "請求參數錯誤: "+err.Error())
		return
	}

	// 獲取當前用戶
	user, err := h.userService.GetUserByID(userID.(uint))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "用戶不存在")
		return
	}

	// 驗證舊密碼
	if !utils.CheckPasswordHash(req.OldPassword, user.Password) {
		utils.ErrorResponse(c, http.StatusBadRequest, "舊密碼錯誤")
		return
	}

	// 更新密碼
	updates := map[string]interface{}{
		"password": req.NewPassword,
	}

	_, err = h.userService.UpdateUser(userID.(uint), updates)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "密碼修改失敗")
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "密碼修改成功"})
}
