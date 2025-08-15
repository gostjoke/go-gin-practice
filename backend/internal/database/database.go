package database

import (
	"log"
	"os"
	"path/filepath"

	"backend/config"
	"backend/internal/models"
	"backend/pkg/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// 初始化數據庫
func InitDB() {
	var err error

	// 確保數據目錄存在
	dbPath := config.AppConfig.DBPath
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal("創建數據庫目錄失敗:", err)
	}

	// 連接數據庫
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("連接數據庫失敗:", err)
	}

	log.Println("數據庫連接成功")

	// 自動遷移
	autoMigrate()

	// 初始化數據
	initData()
}

// 自動遷移
func autoMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Tag{},
		&models.Category{},
		&models.Setting{},
		&models.Log{},
	)
	if err != nil {
		log.Fatal("數據庫遷移失敗:", err)
	}
	log.Println("數據庫遷移完成")
}

// 初始化數據
func initData() {
	// 檢查是否已有管理員用戶
	var count int64
	DB.Model(&models.User{}).Where("role = ?", "admin").Count(&count)

	if count == 0 {
		// 創建默認管理員
		hashedPassword, err := utils.HashPassword(config.AppConfig.AdminPass)
		if err != nil {
			log.Fatal("密碼加密失敗:", err)
		}

		admin := models.User{
			Username: "admin",
			Email:    config.AppConfig.AdminEmail,
			Password: hashedPassword,
			Role:     "admin",
			Status:   "active",
		}

		if err := DB.Create(&admin).Error; err != nil {
			log.Fatal("創建管理員失敗:", err)
		}
		log.Println("默認管理員創建成功")
	}

	// 初始化系統設置
	settings := []models.Setting{
		{Key: "site_name", Value: "Gin Admin", Type: "string", Group: "basic"},
		{Key: "site_description", Value: "基於 Gin 的後台管理系統", Type: "string", Group: "basic"},
		{Key: "posts_per_page", Value: "10", Type: "number", Group: "content"},
		{Key: "allow_registration", Value: "true", Type: "boolean", Group: "user"},
	}

	for _, setting := range settings {
		var count int64
		DB.Model(&models.Setting{}).Where("key = ?", setting.Key).Count(&count)
		if count == 0 {
			DB.Create(&setting)
		}
	}

	// 創建默認分類
	var categoryCount int64
	DB.Model(&models.Category{}).Count(&categoryCount)
	if categoryCount == 0 {
		categories := []models.Category{
			{Name: "技術", Description: "技術相關文章"},
			{Name: "生活", Description: "日常生活分享"},
			{Name: "隨筆", Description: "隨意記錄"},
		}
		DB.Create(&categories)
		log.Println("默認分類創建成功")
	}

	log.Println("數據初始化完成")
}
