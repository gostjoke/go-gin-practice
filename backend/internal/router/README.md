# Router 路由結構說明

## 概述

後端路由已經重構為模組化的結構，所有路由邏輯都分離到 `internal/router` 包中。

## 文件結構

```
internal/router/
├── router.go      # 路由器主結構和初始化
├── health.go      # 健康檢查路由
├── auth.go        # 認證相關路由
├── protected.go   # 需要認證的路由
└── admin.go       # 管理員路由
```

## 路由器結構

### Router 結構體
```go
type Router struct {
    engine      *gin.Engine
    authHandler *handlers.AuthHandler
    userHandler *handlers.UserHandler
    postHandler *handlers.PostHandler
}
```

### 主要方法
- `NewRouter()` - 創建新的路由實例
- `SetupMiddleware()` - 設置中間件
- `SetupRoutes()` - 設置所有路由
- `Initialize()` - 初始化路由器並返回 Gin Engine

## 路由分組

### 1. 健康檢查路由 (health.go)
- `GET /api/health` - 系統健康檢查

### 2. 認證路由 (auth.go)
不需要登錄的路由：
- `POST /api/auth/login` - 用戶登錄
- `POST /api/auth/register` - 用戶註冊

### 3. 受保護路由 (protected.go)
需要認證的路由：

#### 用戶相關
- `GET /api/user/profile` - 獲取用戶資料
- `PUT /api/user/profile` - 更新用戶資料
- `POST /api/user/change-password` - 修改密碼

#### 文章相關
- `GET /api/posts` - 獲取文章列表
- `GET /api/posts/my` - 獲取我的文章
- `GET /api/posts/search` - 搜索文章
- `GET /api/posts/:id` - 獲取單篇文章
- `POST /api/posts` - 創建文章
- `PUT /api/posts/:id` - 更新文章
- `DELETE /api/posts/:id` - 刪除文章

### 4. 管理員路由 (admin.go)
需要管理員權限的路由：

#### 用戶管理
- `GET /api/admin/users` - 獲取所有用戶
- `GET /api/admin/users/:id` - 獲取指定用戶
- `POST /api/admin/users` - 創建用戶
- `PUT /api/admin/users/:id` - 更新用戶
- `DELETE /api/admin/users/:id` - 刪除用戶

#### 文章管理
- `GET /api/admin/posts` - 獲取所有文章
- `GET /api/admin/posts/:id` - 獲取指定文章
- `PUT /api/admin/posts/:id` - 更新文章
- `DELETE /api/admin/posts/:id` - 刪除文章

## 使用方式

在 `main.go` 中：

```go
// 創建並初始化路由器
r := router.NewRouter()
engine := r.Initialize()

// 啟動服務器
engine.Run(":8080")
```

## 優勢

1. **模組化**: 路由按功能分離到不同文件
2. **可維護性**: 每個文件職責單一，易於維護
3. **可擴展性**: 新增路由時只需在對應文件中添加
4. **清晰結構**: 路由組織清晰，易於理解
5. **復用性**: Router 結構可以在測試中復用

## 中間件

路由器自動設置以下中間件：
- Logger 中間件 - 請求日誌記錄
- CORS 中間件 - 跨域請求處理
- Recovery 中間件 - 錯誤恢復
- Auth 中間件 - 認證檢查（受保護路由）
- Admin 中間件 - 管理員權限檢查（管理員路由）
