package router

import (
	"github.com/gin-gonic/gin"
)

// setupAuthRoutes 設置認證路由（不需要登錄）
func (r *Router) setupAuthRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/login", r.authHandler.Login)
		auth.POST("/register", r.authHandler.Register)
	}
}
