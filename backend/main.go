package main

import (
	"fmt"
	"log"

	"backend/config"
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// 載入配置
	config.LoadConfig()

	// 設置 Gin 模式
	gin.SetMode(config.AppConfig.GinMode)

	// 初始化數據庫
	database.InitDB()

	// 創建 Gin 實例
	r := gin.New()

	// 添加中間件
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Recovery())

	// 設置路由
	setupRoutes(r)

	// 啟動服務器
	addr := fmt.Sprintf(":%d", config.AppConfig.Port)
	log.Printf("服務器啟動在端口 %d", config.AppConfig.Port)

	if err := r.Run(addr); err != nil {
		log.Fatal("服務器啟動失敗:", err)
	}
}

func setupRoutes(r *gin.Engine) {
	// 創建處理器實例
	authHandler := handlers.NewAuthHandler()
	userHandler := handlers.NewUserHandler()
	postHandler := handlers.NewPostHandler()

	// API 根路由
	api := r.Group("/api")

	// 健康檢查
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "服務運行正常",
		})
	})

	// 認證路由（不需要登錄）
	auth := api.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/register", authHandler.Register)
	}

	// 需要認證的路由
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// 用戶相關
		user := protected.Group("/user")
		{
			user.GET("/profile", authHandler.Profile)
			user.PUT("/profile", authHandler.UpdateProfile)
			user.POST("/change-password", authHandler.ChangePassword)
		}

		// 文章相關
		posts := protected.Group("/posts")
		{
			posts.GET("", postHandler.GetPosts)
			posts.GET("/my", postHandler.GetMyPosts)
			posts.GET("/search", postHandler.SearchPosts)
			posts.GET("/:id", postHandler.GetPost)
			posts.POST("", postHandler.CreatePost)
			posts.PUT("/:id", postHandler.UpdatePost)
			posts.DELETE("/:id", postHandler.DeletePost)
		}
	}

	// 管理員路由
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminMiddleware())
	{
		// 用戶管理
		adminUsers := admin.Group("/users")
		{
			adminUsers.GET("", userHandler.GetUsers)
			adminUsers.GET("/:id", userHandler.GetUser)
			adminUsers.POST("", userHandler.CreateUser)
			adminUsers.PUT("/:id", userHandler.UpdateUser)
			adminUsers.DELETE("/:id", userHandler.DeleteUser)
		}

		// 文章管理
		adminPosts := admin.Group("/posts")
		{
			adminPosts.GET("", postHandler.GetPosts)
			adminPosts.GET("/:id", postHandler.GetPost)
			adminPosts.PUT("/:id", postHandler.UpdatePost)
			adminPosts.DELETE("/:id", postHandler.DeletePost)
		}
	}
}
