# Gin 後台管理系統 - 使用示例

## 項目已成功搭建完成！

您的 Gin 後台管理系統現在已經包含以下功能：

### 🎯 核心功能
✅ **用戶認證系統** - JWT Token 認證
✅ **用戶管理** - 完整的 CRUD 操作
✅ **文章管理** - 發布、編輯、刪除文章
✅ **角色權限** - 管理員和普通用戶權限分離
✅ **數據庫** - SQLite 數據庫，自動遷移
✅ **中間件** - 認證、日誌、CORS 中間件
✅ **API 文檔** - RESTful API 設計

### 🚀 啟動服務器

```bash
cd "c:\Users\gostj\Desktop\Golang\go-gin-practice"
go run main.go
```

服務器將在 http://localhost:8080 啟動

### 🔑 默認管理員賬號

- **郵箱**: admin@example.com
- **密碼**: admin123

### 📋 API 測試示例

#### 1. 健康檢查
```bash
GET http://localhost:8080/api/health
```

#### 2. 管理員登錄
```bash
POST http://localhost:8080/api/auth/login
Content-Type: application/json

{
  "email": "admin@example.com",
  "password": "admin123"
}
```

#### 3. 獲取個人資料（需要 Token）
```bash
GET http://localhost:8080/api/user/profile
Authorization: Bearer YOUR_TOKEN
```

#### 4. 創建文章（需要 Token）
```bash
POST http://localhost:8080/api/posts
Authorization: Bearer YOUR_TOKEN
Content-Type: application/json

{
  "title": "我的第一篇文章",
  "content": "這是文章內容...",
  "summary": "文章摘要",
  "status": "published"
}
```

#### 5. 獲取文章列表（需要 Token）
```bash
GET http://localhost:8080/api/posts?page=1&limit=10
Authorization: Bearer YOUR_TOKEN
```

### 🗂️ 項目結構

```
go-gin-practice/
├── cmd/                    # 應用程序入口
├── config/                 # 配置管理
│   └── config.go          # 環境變數配置
├── internal/              # 內部業務邏輯
│   ├── database/          # 數據庫連接和初始化
│   │   └── database.go
│   ├── handlers/          # HTTP 處理器
│   │   ├── auth_handler.go    # 認證相關
│   │   ├── user_handler.go    # 用戶管理
│   │   └── post_handler.go    # 文章管理
│   ├── middleware/        # 中間件
│   │   ├── auth.go           # 認證中間件
│   │   └── logger.go         # 日誌中間件
│   ├── models/           # 數據模型
│   │   └── models.go         # 所有數據模型
│   └── services/         # 業務邏輯服務
│       ├── user_service.go   # 用戶服務
│       └── post_service.go   # 文章服務
├── pkg/                  # 公共工具包
│   └── utils/           # 工具函數
│       ├── auth.go          # JWT 工具
│       └── response.go      # 響應工具
├── data/                # 數據庫文件
│   └── app.db              # SQLite 數據庫
├── .env                 # 環境配置
├── .env.example        # 環境配置示例
├── go.mod              # Go 模塊
├── main.go             # 主程序入口
└── README.md           # 項目說明
```

### 📊 數據模型

- **User**: 用戶表（ID、用戶名、郵箱、密碼、角色、頭像、狀態）
- **Post**: 文章表（ID、標題、內容、摘要、狀態、作者ID、瀏覽量）
- **Tag**: 標籤表（ID、名稱、顏色）
- **Category**: 分類表（ID、名稱、描述、父分類ID）
- **Setting**: 系統設置表（ID、鍵名、值、類型、分組）
- **Log**: 操作日誌表（ID、級別、消息、用戶ID、IP、路徑、方法）

### 🔧 下一步擴展

1. **文件上傳功能** - 用戶頭像、文章圖片
2. **標籤管理** - 創建、編輯、刪除標籤
3. **分類管理** - 文章分類系統
4. **系統設置** - 網站配置管理
5. **操作日誌查看** - 管理員查看系統日誌
6. **郵件功能** - 註冊驗證、密碼重置
7. **緩存系統** - Redis 緩存
8. **搜索功能** - 全文搜索

### 🛠️ 使用工具推薦

- **API 測試**: Postman、Insomnia、Thunder Client (VS Code)
- **數據庫查看**: DB Browser for SQLite
- **代碼編輯**: VS Code with Go extension

### 🎉 恭喜！

您的 Gin 後台管理系統已經搭建完成，具備完整的用戶認證、文章管理等功能。
可以在此基礎上繼續開發更多功能！
