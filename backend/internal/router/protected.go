package router

import (
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// setupProtectedRoutes 設置需要認證的路由
func (r *Router) setupProtectedRoutes(api *gin.RouterGroup) {
	// 需要認證的路由
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// 用戶相關路由
		r.setupUserRoutes(protected)

		// 文章相關路由
		r.setupPostRoutes(protected)
	}
}

// setupUserRoutes 設置用戶相關路由
func (r *Router) setupUserRoutes(protected *gin.RouterGroup) {
	user := protected.Group("/user")
	{
		user.GET("/profile", r.authHandler.Profile)
		user.PUT("/profile", r.authHandler.UpdateProfile)
		user.POST("/change-password", r.authHandler.ChangePassword)
	}
}

// setupPostRoutes 設置文章相關路由
func (r *Router) setupPostRoutes(protected *gin.RouterGroup) {
	posts := protected.Group("/posts")
	{
		posts.GET("", r.postHandler.GetPosts)
		posts.GET("/my", r.postHandler.GetMyPosts)
		posts.GET("/search", r.postHandler.SearchPosts)
		posts.GET("/:id", r.postHandler.GetPost)
		posts.POST("", r.postHandler.CreatePost)
		posts.PUT("/:id", r.postHandler.UpdatePost)
		posts.DELETE("/:id", r.postHandler.DeletePost)
	}
}
