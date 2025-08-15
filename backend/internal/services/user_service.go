package services

import (
	"errors"

	"backend/internal/database"
	"backend/internal/models"
	"backend/pkg/utils"

	"gorm.io/gorm"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// 用戶登錄
func (s *UserService) Login(email, password string) (*models.User, string, error) {
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("用戶不存在")
		}
		return nil, "", err
	}

	// 檢查用戶狀態
	if user.Status != "active" {
		return nil, "", errors.New("用戶已被禁用")
	}

	// 驗證密碼
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, "", errors.New("密碼錯誤")
	}

	// 生成 JWT token
	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}

// 用戶註冊
func (s *UserService) Register(username, email, password string) (*models.User, error) {
	// 檢查用戶名是否已存在
	var count int64
	database.DB.Model(&models.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return nil, errors.New("用戶名已存在")
	}

	// 檢查郵箱是否已存在
	database.DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return nil, errors.New("郵箱已存在")
	}

	// 加密密碼
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// 創建用戶
	user := models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Role:     "user",
		Status:   "active",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// 獲取用戶列表
func (s *UserService) GetUsers(page, limit int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	offset := utils.GetOffset(page, limit)

	// 獲取總數
	if err := database.DB.Model(&models.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 獲取用戶列表
	if err := database.DB.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// 根據 ID 獲取用戶
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 更新用戶
func (s *UserService) UpdateUser(id uint, updates map[string]interface{}) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	// 如果包含密碼，需要加密
	if password, ok := updates["password"]; ok {
		if passwordStr, ok := password.(string); ok && passwordStr != "" {
			hashedPassword, err := utils.HashPassword(passwordStr)
			if err != nil {
				return nil, err
			}
			updates["password"] = hashedPassword
		} else {
			delete(updates, "password")
		}
	}

	if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// 刪除用戶
func (s *UserService) DeleteUser(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}
