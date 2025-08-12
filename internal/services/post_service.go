package services

import (
	"backend/internal/database"
	"backend/internal/models"
	"backend/pkg/utils"

	"gorm.io/gorm"
)

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}

// 創建文章
func (s *PostService) CreatePost(title, content, summary, status string, authorID uint, tagIDs []uint) (*models.Post, error) {
	post := models.Post{
		Title:    title,
		Content:  content,
		Summary:  summary,
		Status:   status,
		AuthorID: authorID,
	}

	// 開始事務
	tx := database.DB.Begin()

	// 創建文章
	if err := tx.Create(&post).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 關聯標籤
	if len(tagIDs) > 0 {
		var tags []models.Tag
		if err := tx.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		if err := tx.Model(&post).Association("Tags").Append(tags); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	tx.Commit()

	// 重新查詢包含關聯數據的文章
	return s.GetPostByID(post.ID)
}

// 獲取文章列表
func (s *PostService) GetPosts(page, limit int, status string, authorID uint) ([]models.Post, int64, error) {
	var posts []models.Post
	var total int64

	offset := utils.GetOffset(page, limit)
	query := database.DB.Model(&models.Post{}).Preload("Author").Preload("Tags")

	// 條件過濾
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if authorID > 0 {
		query = query.Where("author_id = ?", authorID)
	}

	// 獲取總數
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 獲取文章列表
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// 根據 ID 獲取文章
func (s *PostService) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := database.DB.Preload("Author").Preload("Tags").First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// 更新文章
func (s *PostService) UpdatePost(id uint, updates map[string]interface{}, tagIDs []uint) (*models.Post, error) {
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		return nil, err
	}

	// 開始事務
	tx := database.DB.Begin()

	// 更新文章基本信息
	if err := tx.Model(&post).Updates(updates).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 更新標籤關聯
	if tagIDs != nil {
		// 清除現有標籤關聯
		if err := tx.Model(&post).Association("Tags").Clear(); err != nil {
			tx.Rollback()
			return nil, err
		}

		// 添加新標籤關聯
		if len(tagIDs) > 0 {
			var tags []models.Tag
			if err := tx.Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
				tx.Rollback()
				return nil, err
			}
			if err := tx.Model(&post).Association("Tags").Append(tags); err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	tx.Commit()

	// 重新查詢包含關聯數據的文章
	return s.GetPostByID(post.ID)
}

// 刪除文章
func (s *PostService) DeletePost(id uint) error {
	return database.DB.Delete(&models.Post{}, id).Error
}

// 增加瀏覽量
func (s *PostService) IncrementViewCount(id uint) error {
	return database.DB.Model(&models.Post{}).Where("id = ?", id).Update("view_count", gorm.Expr("view_count + ?", 1)).Error
}
