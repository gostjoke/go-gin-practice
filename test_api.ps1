# API 測試腳本

# 1. 測試健康檢查
Write-Host "=== 測試健康檢查 API ===" -ForegroundColor Green
try {
    $healthResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/health" -Method GET
    Write-Host "健康檢查成功: $($healthResponse | ConvertTo-Json -Depth 2)" -ForegroundColor Green
} catch {
    Write-Host "健康檢查失敗: $($_.Exception.Message)" -ForegroundColor Red
}

# 2. 測試管理員登錄
Write-Host "`n=== 測試管理員登錄 ===" -ForegroundColor Green
$loginData = @{
    email = "admin@example.com"
    password = "admin123"
} | ConvertTo-Json

try {
    $loginResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/auth/login" -Method POST -Body $loginData -ContentType "application/json"
    $token = $loginResponse.data.token
    Write-Host "登錄成功，獲得 Token: $($token.Substring(0,20))..." -ForegroundColor Green
    
    # 3. 測試獲取用戶資料
    Write-Host "`n=== 測試獲取用戶資料 ===" -ForegroundColor Green
    $headers = @{ Authorization = "Bearer $token" }
    $profileResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/user/profile" -Method GET -Headers $headers
    Write-Host "用戶資料: $($profileResponse.data | ConvertTo-Json -Depth 2)" -ForegroundColor Green
    
    # 4. 測試創建文章
    Write-Host "`n=== 測試創建文章 ===" -ForegroundColor Green
    $postData = @{
        title = "我的第一篇文章"
        content = "這是文章的內容，用來測試後台系統的功能。"
        summary = "測試文章摘要"
        status = "published"
    } | ConvertTo-Json
    
    $postResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/posts" -Method POST -Body $postData -ContentType "application/json" -Headers $headers
    Write-Host "文章創建成功: $($postResponse.data | ConvertTo-Json -Depth 2)" -ForegroundColor Green
    
    # 5. 測試獲取文章列表
    Write-Host "`n=== 測試獲取文章列表 ===" -ForegroundColor Green
    $postsResponse = Invoke-RestMethod -Uri "http://localhost:8080/api/posts" -Method GET -Headers $headers
    Write-Host "文章列表: $($postsResponse | ConvertTo-Json -Depth 3)" -ForegroundColor Green
    
} catch {
    Write-Host "登錄失敗: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n=== 測試完成 ===" -ForegroundColor Green
