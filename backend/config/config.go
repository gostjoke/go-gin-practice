package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       int
	GinMode    string
	DBType     string
	DBPath     string
	JWTSecret  string
	JWTExpires string
	AdminEmail string
	AdminPass  string
}

var AppConfig *Config

func LoadConfig() {
	// 載入 .env 檔案
	if err := godotenv.Load(); err != nil {
		log.Println("未找到 .env 檔案，使用環境變數")
	}

	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		port = 8080
	}

	AppConfig = &Config{
		Port:       port,
		GinMode:    getEnv("GIN_MODE", "debug"),
		DBType:     getEnv("DB_TYPE", "sqlite"),
		DBPath:     getEnv("DB_PATH", "./data/app.db"),
		JWTSecret:  getEnv("JWT_SECRET", "default-secret-key"),
		JWTExpires: getEnv("JWT_EXPIRES_IN", "24h"),
		AdminEmail: getEnv("ADMIN_EMAIL", "admin@example.com"),
		AdminPass:  getEnv("ADMIN_PASSWORD", "admin123"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
