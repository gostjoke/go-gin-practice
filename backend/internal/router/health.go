package router

import (
	"github.com/gin-gonic/gin"
)

// setupHealthRoutes 設置健康檢查路由
func (r *Router) setupHealthRoutes(api *gin.RouterGroup) {
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "服務運行正常",
		})
	})
}
