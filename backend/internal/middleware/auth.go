package middleware

import (
	"net/http"
	"strings"

	"backend/internal/database"
	"backend/internal/models"
	"backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

// JWT 驗證中間件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 獲取 Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "缺少認證令牌")
			c.Abort()
			return
		}

		// 檢查 Bearer 前綴
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			utils.ErrorResponse(c, http.StatusUnauthorized, "令牌格式錯誤")
			c.Abort()
			return
		}

		// 解析令牌
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "令牌無效")
			c.Abort()
			return
		}

		// 檢查用戶是否存在
		var user models.User
		if err := database.DB.First(&user, claims.UserID).Error; err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "用戶不存在")
			c.Abort()
			return
		}

		// 檢查用戶狀態
		if user.Status != "active" {
			utils.ErrorResponse(c, http.StatusForbidden, "用戶已被禁用")
			c.Abort()
			return
		}

		// 設置用戶信息到上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("user", user)

		c.Next()
	}
}

// 管理員權限中間件
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			utils.ErrorResponse(c, http.StatusForbidden, "需要管理員權限")
			c.Abort()
			return
		}
		c.Next()
	}
}
