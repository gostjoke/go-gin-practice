# Gin 後台管理系統

這是一個基於 Gin 框架構建的完整後台管理系統，提供用戶管理、文章管理、認證授權等功能。

## 功能特性

- 🔐 **用戶認證與授權**：JWT Token 認證，角色權限控制
- 👥 **用戶管理**：用戶 CRUD 操作，密碼加密存儲
- 📝 **文章管理**：文章發布、編輯、刪除，標籤系統
- 🔍 **搜索功能**：文章內容搜索
- 📊 **日誌記錄**：完整的操作日誌記錄
- 🌐 **CORS 支持**：跨域請求支持
- 📖 **API 文檔**：RESTful API 設計

## 技術棧

- **後端框架**：Gin (Go)
- **數據庫**：SQLite（可擴展為 MySQL/PostgreSQL）
- **ORM**：GORM
- **認證**：JWT
- **配置管理**：環境變量 + .env 文件
- **密碼加密**：bcrypt

## 項目結構

```
├── cmd/                    # 應用程序入口
├── config/                 # 配置文件
├── internal/              # 內部業務邏輯
│   ├── database/          # 數據庫連接和初始化
│   ├── handlers/          # HTTP 處理器
│   ├── middleware/        # 中間件
│   ├── models/           # 數據模型
│   └── services/         # 業務邏輯服務
├── pkg/                  # 公共工具包
│   └── utils/           # 工具函數
├── data/                # 數據庫文件存儲
├── .env                 # 環境配置文件
├── .env.example        # 環境配置示例
├── go.mod              # Go 模塊配置
└── main.go             # 主程序入口
```

## 安裝和使用

### 1. 克隆項目

```bash
git clone <repository-url>
cd go-gin-practice
```

### 2. 安裝依賴

```bash
go mod tidy
```

### 3. 配置環境變量

複製 `.env.example` 為 `.env` 並修改配置：

```bash
cp .env.example .env
```

### 4. 啟動服務器

```bash
go run main.go
```

服務器將在 `http://localhost:8080` 啟動。

### 5. 測試 API

服務器啟動後，您可以使用以下默認管理員賬號：

- **郵箱**：admin@example.com
- **密碼**：admin123

## API 文檔

### 健康檢查

```
GET /api/health
```

### 認證 API

#### 用戶登錄
```
POST /api/auth/login
Content-Type: application/json

{
  "email": "admin@example.com",
  "password": "admin123"
}
```

#### 用戶註冊
```
POST /api/auth/register
Content-Type: application/json

{
  "username": "testuser",
  "email": "test@example.com",
  "password": "password123"
}
```

### 用戶 API （需要認證）

#### 獲取個人資料
```
GET /api/user/profile
Authorization: Bearer <token>
```

#### 更新個人資料
```
PUT /api/user/profile
Authorization: Bearer <token>
Content-Type: application/json

{
  "username": "new_username",
  "avatar": "avatar_url"
}
```

#### 修改密碼
```
POST /api/user/change-password
Authorization: Bearer <token>
Content-Type: application/json

{
  "old_password": "old_password",
  "new_password": "new_password"
}
```

### 文章 API （需要認證）

#### 獲取文章列表
```
GET /api/posts?page=1&limit=10&status=published&author_id=1
Authorization: Bearer <token>
```

#### 創建文章
```
POST /api/posts
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "文章標題",
  "content": "文章內容",
  "summary": "文章摘要",
  "status": "published",
  "tag_ids": [1, 2, 3]
}
```

#### 獲取單篇文章
```
GET /api/posts/:id
Authorization: Bearer <token>
```

#### 更新文章
```
PUT /api/posts/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "更新的標題",
  "content": "更新的內容",
  "status": "published"
}
```

#### 刪除文章
```
DELETE /api/posts/:id
Authorization: Bearer <token>
```

#### 獲取我的文章
```
GET /api/posts/my?page=1&limit=10&status=draft
Authorization: Bearer <token>
```

#### 搜索文章
```
GET /api/posts/search?keyword=關鍵字&page=1&limit=10
Authorization: Bearer <token>
```

### 管理員 API （需要管理員權限）

#### 用戶管理

```
GET /api/admin/users              # 獲取所有用戶
GET /api/admin/users/:id          # 獲取單個用戶
POST /api/admin/users             # 創建用戶
PUT /api/admin/users/:id          # 更新用戶
DELETE /api/admin/users/:id       # 刪除用戶
```

#### 文章管理

```
GET /api/admin/posts              # 獲取所有文章
GET /api/admin/posts/:id          # 獲取單篇文章
PUT /api/admin/posts/:id          # 更新文章
DELETE /api/admin/posts/:id       # 刪除文章
```

## 數據模型

### 用戶模型 (User)
- ID, 用戶名, 郵箱, 密碼, 角色, 頭像, 狀態
- 角色：admin（管理員）、user（普通用戶）
- 狀態：active（活躍）、inactive（不活躍）

### 文章模型 (Post)
- ID, 標題, 內容, 摘要, 狀態, 作者ID, 標籤, 瀏覽量
- 狀態：draft（草稿）、published（已發布）、archived（歸檔）

### 標籤模型 (Tag)
- ID, 名稱, 顏色

### 分類模型 (Category)
- ID, 名稱, 描述, 父分類ID（支持嵌套分類）

### 系統設置模型 (Setting)
- ID, 鍵名, 值, 類型, 分組

### 日誌模型 (Log)
- ID, 級別, 消息, 用戶ID, IP地址, User Agent, 路徑, 方法

## 開發說明

### 環境配置

項目使用 `.env` 文件進行配置管理，主要配置項包括：

- `PORT`：服務器端口
- `GIN_MODE`：Gin 運行模式（debug/release）
- `DB_TYPE`：數據庫類型
- `DB_PATH`：數據庫文件路徑
- `JWT_SECRET`：JWT 簽名密鑰
- `JWT_EXPIRES_IN`：JWT 過期時間
- `ADMIN_EMAIL`：默認管理員郵箱
- `ADMIN_PASSWORD`：默認管理員密碼

### 數據庫遷移

項目使用 GORM 自動遷移功能，首次啟動時會自動創建數據表並初始化默認數據。

### 中間件

- **認證中間件**：驗證 JWT Token
- **權限中間件**：檢查用戶角色權限
- **日誌中間件**：記錄請求日誌
- **CORS 中間件**：處理跨域請求

### 密碼安全

用戶密碼使用 bcrypt 算法進行加密存儲，確保數據安全。

## 擴展功能

您可以基於此項目擴展以下功能：

1. **文件上傳**：用戶頭像、文章圖片上傳
2. **郵件功能**：註冊驗證、密碼重置
3. **緩存機制**：Redis 緩存熱點數據
4. **搜索引擎**：Elasticsearch 全文搜索
5. **消息隊列**：RabbitMQ/Kafka 異步處理
6. **監控告警**：Prometheus + Grafana 監控
7. **API 限流**：Redis + Lua 實現限流
8. **數據備份**：定時備份數據庫

## 貢獻

歡迎提交 Issue 和 Pull Request 來改善這個項目。

## 許可證

MIT License