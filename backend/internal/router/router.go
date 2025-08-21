package router

import (
	"backend/internal/handlers"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Router 路由結構體
type Router struct {
	engine      *gin.Engine
	authHandler *handlers.AuthHandler
	userHandler *handlers.UserHandler
	postHandler *handlers.PostHandler
}

// NewRouter 創建新的路由實例
func NewRouter() *Router {
	return &Router{
		engine:      gin.New(),
		authHandler: handlers.NewAuthHandler(),
		userHandler: handlers.NewUserHandler(),
		postHandler: handlers.NewPostHandler(),
	}
}

// GetEngine 獲取 Gin Engine
func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}

// SetupMiddleware 設置中間件
func (r *Router) SetupMiddleware() {
	r.engine.Use(middleware.LoggerMiddleware())
	r.engine.Use(middleware.CORSMiddleware())
	r.engine.Use(gin.Recovery())
}

// SetupRoutes 設置所有路由
func (r *Router) SetupRoutes() {
	// API 根路由
	api := r.engine.Group("/api")

	// 設置各個路由組
	r.setupHealthRoutes(api)
	r.setupAuthRoutes(api)
	r.setupProtectedRoutes(api)
	r.setupAdminRoutes(api)
}

// Initialize 初始化路由器
func (r *Router) Initialize() *gin.Engine {
	r.SetupMiddleware()
	r.SetupRoutes()
	return r.engine
}
