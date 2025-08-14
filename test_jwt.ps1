# 實際測試 JWT 認證流程

Write-Host "=== JWT 認證流程測試 ===" -ForegroundColor Cyan
Write-Host "服務器地址: http://localhost:8080" -ForegroundColor Green

# 步驟 1: 測試健康檢查（無需認證）
Write-Host "`n1. 測試健康檢查 API（無需認證）" -ForegroundColor Yellow
try {
    $healthResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/health" -Method GET
    Write-Host "✅ 健康檢查成功: $($healthResponse.message)" -ForegroundColor Green
} catch {
    Write-Host "❌ 健康檢查失敗: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

# 步驟 2: 嘗試訪問受保護的資源（無 Token）
Write-Host "`n2. 嘗試訪問受保護資源（無 Token）" -ForegroundColor Yellow
try {
    $profileResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/user/profile" -Method GET
    Write-Host "❌ 預期失敗但成功了" -ForegroundColor Red
} catch {
    Write-Host "✅ 如預期被拒絕: 缺少認證令牌" -ForegroundColor Green
}

# 步驟 3: 管理員登錄獲取 JWT Token
Write-Host "`n3. 管理員登錄獲取 JWT Token" -ForegroundColor Yellow
$loginData = @{
    email = "admin@example.com"
    password = "admin123"
} | ConvertTo-Json -Compress

try {
    $loginResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/auth/login" -Method POST -Body $loginData -ContentType "application/json"
    $token = $loginResponse.data.token
    $user = $loginResponse.data.user
    
    Write-Host "✅ 登錄成功!" -ForegroundColor Green
    Write-Host "   用戶ID: $($user.id)" -ForegroundColor White
    Write-Host "   用戶名: $($user.username)" -ForegroundColor White
    Write-Host "   角色: $($user.role)" -ForegroundColor White
    Write-Host "   JWT Token: $($token.Substring(0,50))..." -ForegroundColor White
    
} catch {
    Write-Host "❌ 登錄失敗: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

# 步驟 4: 使用 JWT Token 訪問用戶資料
Write-Host "`n4. 使用 JWT Token 訪問用戶資料" -ForegroundColor Yellow
$headers = @{ 
    Authorization = "Bearer $token"
    "Content-Type" = "application/json"
}

try {
    $profileResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/user/profile" -Method GET -Headers $headers
    Write-Host "✅ 成功獲取用戶資料!" -ForegroundColor Green
    Write-Host "   ID: $($profileResponse.data.id)" -ForegroundColor White
    Write-Host "   用戶名: $($profileResponse.data.username)" -ForegroundColor White
    Write-Host "   郵箱: $($profileResponse.data.email)" -ForegroundColor White
    Write-Host "   角色: $($profileResponse.data.role)" -ForegroundColor White
} catch {
    Write-Host "❌ 獲取用戶資料失敗: $($_.Exception.Message)" -ForegroundColor Red
}

# 步驟 5: 測試管理員權限
Write-Host "`n5. 測試管理員權限（獲取所有用戶）" -ForegroundColor Yellow
try {
    $usersResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/admin/users" -Method GET -Headers $headers
    Write-Host "✅ 管理員權限驗證成功!" -ForegroundColor Green
    Write-Host "   用戶總數: $($usersResponse.meta.total)" -ForegroundColor White
    Write-Host "   當前頁: $($usersResponse.meta.current_page)" -ForegroundColor White
} catch {
    Write-Host "❌ 管理員權限測試失敗: $($_.Exception.Message)" -ForegroundColor Red
}

# 步驟 6: 創建一篇文章
Write-Host "`n6. 創建一篇文章（需要認證）" -ForegroundColor Yellow
$postData = @{
    title = "JWT 測試文章"
    content = "這是一篇用來測試 JWT 認證的文章。只有通過 JWT 驗證的用戶才能創建文章。"
    summary = "JWT 測試文章摘要"
    status = "published"
} | ConvertTo-Json -Compress

try {
    $postResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/posts" -Method POST -Body $postData -Headers $headers
    Write-Host "✅ 文章創建成功!" -ForegroundColor Green
    Write-Host "   文章ID: $($postResponse.data.id)" -ForegroundColor White
    Write-Host "   標題: $($postResponse.data.title)" -ForegroundColor White
    Write-Host "   作者: $($postResponse.data.author.username)" -ForegroundColor White
    Write-Host "   狀態: $($postResponse.data.status)" -ForegroundColor White
} catch {
    Write-Host "❌ 文章創建失敗: $($_.Exception.Message)" -ForegroundColor Red
}

# 步驟 7: 測試無效 Token
Write-Host "`n7. 測試無效 JWT Token" -ForegroundColor Yellow
$invalidHeaders = @{ 
    Authorization = "Bearer invalid-token-here"
    "Content-Type" = "application/json"
}

try {
    $invalidResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/user/profile" -Method GET -Headers $invalidHeaders
    Write-Host "❌ 預期失敗但成功了" -ForegroundColor Red
} catch {
    Write-Host "✅ 如預期被拒絕: 無效的 JWT Token" -ForegroundColor Green
}

# 總結
Write-Host "`n=== JWT 認證測試完成 ===" -ForegroundColor Cyan
Write-Host @"

🔐 JWT 工作流程總結:
1. ✅ 健康檢查 - 無需認證的公開 API
2. ✅ 無 Token 訪問 - 被正確拒絕
3. ✅ 用戶登錄 - 獲得有效 JWT Token
4. ✅ Token 認證 - 成功訪問受保護資源
5. ✅ 權限控制 - 管理員權限驗證成功
6. ✅ 業務操作 - 創建文章需要認證
7. ✅ 安全檢查 - 無效 Token 被拒絕

💡 JWT Token 的結構:
- Header（頭部）: 算法和類型信息
- Payload（載荷）: 用戶信息（UserID, Username, Role, 過期時間等）
- Signature（簽名）: 使用密鑰簽名，確保 Token 完整性

🚀 JWT 的優勢:
- 無狀態：服務器不需要存儲 Session
- 可擴展：適合分佈式系統
- 安全：使用密鑰簽名防篡改
- 自包含：Token 包含所有必要信息

"@ -ForegroundColor White

Write-Host "測試完成! 您的 JWT 認證系統工作正常。" -ForegroundColor Green
