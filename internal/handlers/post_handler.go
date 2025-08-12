package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"backend/internal/services"
	"backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService *services.PostService
}

func NewPostHandler() *PostHandler {
	return &PostHandler{
		postService: services.NewPostService(),
	}
}

// 創建文章請求結構
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"required"`
	Summary string `json:"summary"`
	Status  string `json:"status"`
	TagIDs  []uint `json:"tag_ids"`
}

// 創建文章
func (h *PostHandler) CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "請求參數錯誤: "+err.Error())
		return
	}

	// 獲取當前用戶ID
	authorID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未登錄")
		return
	}

	// 設置默認值
	if req.Status == "" {
		req.Status = "draft"
	}
	if req.Summary == "" && len(req.Content) > 200 {
		req.Summary = req.Content[:200] + "..."
	}

	post, err := h.postService.CreatePost(req.Title, req.Content, req.Summary, req.Status, authorID.(uint), req.TagIDs)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "創建文章失敗: "+err.Error())
		return
	}

	utils.SuccessResponse(c, post)
}

// 獲取文章列表
func (h *PostHandler) GetPosts(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)
	status := c.Query("status")
	authorIDStr := c.Query("author_id")

	var authorID uint
	if authorIDStr != "" {
		if id, err := strconv.ParseUint(authorIDStr, 10, 32); err == nil {
			authorID = uint(id)
		}
	}

	posts, total, err := h.postService.GetPosts(page, limit, status, authorID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "獲取文章列表失敗")
		return
	}

	utils.PaginatedSuccessResponse(c, posts, page, limit, total)
}

// 根據ID獲取文章
func (h *PostHandler) GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "無效的文章ID")
		return
	}

	post, err := h.postService.GetPostByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "文章不存在")
		return
	}

	// 增加瀏覽量
	go h.postService.IncrementViewCount(uint(id))

	utils.SuccessResponse(c, post)
}

// 更新文章請求結構
type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Summary string `json:"summary"`
	Status  string `json:"status"`
	TagIDs  []uint `json:"tag_ids"`
}

// 更新文章
func (h *PostHandler) UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "無效的文章ID")
		return
	}

	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "請求參數錯誤: "+err.Error())
		return
	}

	// 檢查文章是否存在
	existingPost, err := h.postService.GetPostByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "文章不存在")
		return
	}

	// 檢查權限：只有作者和管理員可以編輯
	currentUserID, _ := c.Get("user_id")
	currentRole, _ := c.Get("role")
	if currentRole != "admin" && existingPost.AuthorID != currentUserID.(uint) {
		utils.ErrorResponse(c, http.StatusForbidden, "沒有權限編輯此文章")
		return
	}

	updates := make(map[string]interface{})

	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}
	if req.Summary != "" {
		updates["summary"] = req.Summary
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}

	post, err := h.postService.UpdatePost(uint(id), updates, req.TagIDs)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "更新文章失敗: "+err.Error())
		return
	}

	utils.SuccessResponse(c, post)
}

// 刪除文章
func (h *PostHandler) DeletePost(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "無效的文章ID")
		return
	}

	// 檢查文章是否存在
	existingPost, err := h.postService.GetPostByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "文章不存在")
		return
	}

	// 檢查權限：只有作者和管理員可以刪除
	currentUserID, _ := c.Get("user_id")
	currentRole, _ := c.Get("role")
	if currentRole != "admin" && existingPost.AuthorID != currentUserID.(uint) {
		utils.ErrorResponse(c, http.StatusForbidden, "沒有權限刪除此文章")
		return
	}

	err = h.postService.DeletePost(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "刪除文章失敗")
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "文章刪除成功"})
}

// 獲取我的文章
func (h *PostHandler) GetMyPosts(c *gin.Context) {
	page, limit := utils.GetPaginationParams(c)
	status := c.Query("status")

	// 獲取當前用戶ID
	authorID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未登錄")
		return
	}

	posts, total, err := h.postService.GetPosts(page, limit, status, authorID.(uint))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "獲取文章列表失敗")
		return
	}

	utils.PaginatedSuccessResponse(c, posts, page, limit, total)
}

// 搜索文章
func (h *PostHandler) SearchPosts(c *gin.Context) {
	keyword := strings.TrimSpace(c.Query("keyword"))
	if keyword == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "搜索關鍵字不能為空")
		return
	}

	page, limit := utils.GetPaginationParams(c)

	// 這裡可以實現更複雜的搜索邏輯
	// 暫時使用簡單的標題和內容搜索
	posts, _, err := h.postService.GetPosts(page, limit, "published", 0)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "搜索失敗")
		return
	}

	// 過濾包含關鍵字的文章
	var filteredPosts []interface{}
	for _, post := range posts {
		if strings.Contains(strings.ToLower(post.Title), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(post.Content), strings.ToLower(keyword)) {
			filteredPosts = append(filteredPosts, post)
		}
	}

	utils.PaginatedSuccessResponse(c, filteredPosts, page, limit, int64(len(filteredPosts)))
}
