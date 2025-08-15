package models

import (
	"time"

	"gorm.io/gorm"
)

// 基礎模型
type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// 用戶模型
type User struct {
	BaseModel
	Username string `json:"username" gorm:"uniqueIndex;not null;size:50"`
	Email    string `json:"email" gorm:"uniqueIndex;not null;size:100"`
	Password string `json:"-" gorm:"not null"`
	Role     string `json:"role" gorm:"default:user;size:20"`
	Avatar   string `json:"avatar" gorm:"size:255"`
	Status   string `json:"status" gorm:"default:active;size:20"`
	Posts    []Post `json:"posts,omitempty" gorm:"foreignKey:AuthorID"`
}

// 文章模型
type Post struct {
	BaseModel
	Title     string `json:"title" gorm:"not null;size:200"`
	Content   string `json:"content" gorm:"type:text"`
	Summary   string `json:"summary" gorm:"size:500"`
	Status    string `json:"status" gorm:"default:draft;size:20"`
	AuthorID  uint   `json:"author_id" gorm:"not null"`
	Author    User   `json:"author,omitempty" gorm:"foreignKey:AuthorID"`
	Tags      []Tag  `json:"tags,omitempty" gorm:"many2many:post_tags;"`
	ViewCount int    `json:"view_count" gorm:"default:0"`
}

// 標籤模型
type Tag struct {
	BaseModel
	Name  string `json:"name" gorm:"uniqueIndex;not null;size:50"`
	Color string `json:"color" gorm:"size:7"`
	Posts []Post `json:"posts,omitempty" gorm:"many2many:post_tags;"`
}

// 分類模型
type Category struct {
	BaseModel
	Name        string     `json:"name" gorm:"uniqueIndex;not null;size:50"`
	Description string     `json:"description" gorm:"size:255"`
	ParentID    *uint      `json:"parent_id" gorm:"index"`
	Parent      *Category  `json:"parent,omitempty" gorm:"foreignKey:ParentID"`
	Children    []Category `json:"children,omitempty" gorm:"foreignKey:ParentID"`
}

// 系統設置模型
type Setting struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Key   string `json:"key" gorm:"uniqueIndex;not null;size:100"`
	Value string `json:"value" gorm:"type:text"`
	Type  string `json:"type" gorm:"size:20"` // string, number, boolean, json
	Group string `json:"group" gorm:"size:50"`
}

// 日誌模型
type Log struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Level     string    `json:"level" gorm:"size:10"`
	Message   string    `json:"message" gorm:"type:text"`
	UserID    *uint     `json:"user_id" gorm:"index"`
	User      *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	IP        string    `json:"ip" gorm:"size:45"`
	UserAgent string    `json:"user_agent" gorm:"size:255"`
	Path      string    `json:"path" gorm:"size:255"`
	Method    string    `json:"method" gorm:"size:10"`
	CreatedAt time.Time `json:"created_at"`
}
