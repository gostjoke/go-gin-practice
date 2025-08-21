package main

import (
	"fmt"
	"log"

	"backend/config"
	"backend/internal/database"
	"backend/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 載入配置
	config.LoadConfig()

	// 設置 Gin 模式
	gin.SetMode(config.AppConfig.GinMode)

	// 初始化數據庫
	database.InitDB()

	// 創建並初始化路由器
	r := router.NewRouter()
	engine := r.Initialize()

	// 啟動服務器
	addr := fmt.Sprintf(":%d", config.AppConfig.Port)
	log.Printf("服務器啟動在端口 %d", config.AppConfig.Port)

	if err := engine.Run(addr); err != nil {
		log.Fatal("服務器啟動失敗:", err)
	}
}
