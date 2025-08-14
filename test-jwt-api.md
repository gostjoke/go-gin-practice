# JWT API 測試指南

## 1. 用戶登錄獲取 Token

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d "{
    \"email\": \"admin@example.com\",
    \"password\": \"admin123\"
  }"
```

**預期回應：**
```json
{
  "success": true,
  "data": {
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "role": "admin",
      "status": "active",
      "created_at": "2025-01-01T00:00:00Z",
      "updated_at": "2025-01-01T00:00:00Z"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

## 2. 測試需要認證的 API

### 獲取用戶資料 (需要 JWT Token)

```bash
curl -X GET http://localhost:8080/api/user/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

### 沒有 Token 的情況 (應該返回 401)

```bash
curl -X GET http://localhost:8080/api/user/profile
```

**預期回應：**
```json
{
  "success": false,
  "message": "請提供認證令牌"
}
```

### 無效 Token 的情況 (應該返回 401)

```bash
curl -X GET http://localhost:8080/api/user/profile \
  -H "Authorization: Bearer invalid_token"
```

## 3. 測試管理員路由

### 獲取用戶列表 (需要管理員權限)

```bash
curl -X GET http://localhost:8080/api/admin/users \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

### 非管理員用戶訪問 (應該返回 403)

首先註冊一個普通用戶：
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d "{
    \"username\": \"testuser\",
    \"email\": \"test@example.com\",
    \"password\": \"password123\"
  }"
```

然後用普通用戶的 token 訪問管理員路由：
```bash
curl -X GET http://localhost:8080/api/admin/users \
  -H "Authorization: Bearer NORMAL_USER_TOKEN"
```

**預期回應：**
```json
{
  "success": false,
  "message": "需要管理員權限"
}
```

## 4. 測試文章相關 API

### 創建文章 (需要認證)

```bash
curl -X POST http://localhost:8080/api/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE" \
  -d "{
    \"title\": \"測試文章\",
    \"content\": \"這是一篇測試文章的內容...\",
    \"summary\": \"測試文章摘要\",
    \"status\": \"published\"
  }"
```

### 獲取我的文章 (需要認證)

```bash
curl -X GET http://localhost:8080/api/posts/my \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

## 5. JWT Token 格式說明

JWT Token 包含三個部分，用點號分隔：
- Header: 包含加密算法信息
- Payload: 包含用戶信息 (user_id, username, role)
- Signature: 用於驗證 token 的簽名

Token 會在 24 小時後過期。

## 6. 測試 Token 過期

Token 過期後訪問受保護的路由會返回：
```json
{
  "success": false,
  "message": "令牌已過期"
}
```

## 7. 健康檢查 (無需認證)

```bash
curl -X GET http://localhost:8080/api/health
```

**預期回應：**
```json
{
  "status": "ok",
  "message": "服務運行正常"
}
```

## 8. 公開路由 (無需認證)

- `GET /api/health` - 健康檢查
- `POST /api/auth/login` - 用戶登錄
- `POST /api/auth/register` - 用戶註冊

## 9. 受保護路由 (需要 JWT 認證)

- `/api/user/*` - 用戶相關操作
- `/api/posts/*` - 文章相關操作

## 10. 管理員路由 (需要管理員權限)

- `/api/admin/*` - 所有管理功能

使用這些測試案例可以驗證 JWT 認證和授權功能是否正常工作。
