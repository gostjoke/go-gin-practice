package middleware

import (
	"fmt"
	"time"

	"backend/internal/database"
	"backend/internal/models"

	"github.com/gin-gonic/gin"
)

// 日誌中間件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		// 處理請求
		c.Next()

		// 獲取響應狀態碼
		statusCode := c.Writer.Status()

		// 獲取處理時間
		latency := time.Since(start)

		// 獲取用戶ID（如果已登錄）
		var userID *uint
		if id, exists := c.Get("user_id"); exists {
			if uid, ok := id.(uint); ok {
				userID = &uid
			}
		}

		// 記錄日誌到數據庫
		logLevel := "info"
		if statusCode >= 400 {
			logLevel = "error"
		}

		message := fmt.Sprintf("%s %s %d %v", method, path, statusCode, latency)

		logEntry := models.Log{
			Level:     logLevel,
			Message:   message,
			UserID:    userID,
			IP:        clientIP,
			UserAgent: userAgent,
			Path:      path,
			Method:    method,
			CreatedAt: time.Now(),
		}

		// 異步保存日誌，避免影響響應性能
		go func() {
			database.DB.Create(&logEntry)
		}()
	}
}

// CORS 中間件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
