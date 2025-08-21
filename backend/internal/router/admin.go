package router

import (
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// setupAdminRoutes 設置管理員路由
func (r *Router) setupAdminRoutes(api *gin.RouterGroup) {
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminMiddleware())
	{
		// 用戶管理路由
		r.setupAdminUserRoutes(admin)

		// 文章管理路由
		r.setupAdminPostRoutes(admin)
	}
}

// setupAdminUserRoutes 設置管理員用戶管理路由
func (r *Router) setupAdminUserRoutes(admin *gin.RouterGroup) {
	adminUsers := admin.Group("/users")
	{
		adminUsers.GET("", r.userHandler.GetUsers)
		adminUsers.GET("/:id", r.userHandler.GetUser)
		adminUsers.POST("", r.userHandler.CreateUser)
		adminUsers.PUT("/:id", r.userHandler.UpdateUser)
		adminUsers.DELETE("/:id", r.userHandler.DeleteUser)
	}
}

// setupAdminPostRoutes 設置管理員文章管理路由
func (r *Router) setupAdminPostRoutes(admin *gin.RouterGroup) {
	adminPosts := admin.Group("/posts")
	{
		adminPosts.GET("", r.postHandler.GetPosts)
		adminPosts.GET("/:id", r.postHandler.GetPost)
		adminPosts.PUT("/:id", r.postHandler.UpdatePost)
		adminPosts.DELETE("/:id", r.postHandler.DeletePost)
	}
}
