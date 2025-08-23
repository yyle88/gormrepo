package models

import (
	"gorm.io/gorm"
)

// Post 帖子模型 - 演示分页查询功能
type Post struct {
	gorm.Model
	Title     string `gorm:"size:200;not null;index"`
	Content   string `gorm:"type:text"`
	AuthorID  uint   `gorm:"index"`
	Category  string `gorm:"size:50;index"`
	Tags      string `gorm:"size:200"`
	ViewCount int    `gorm:"default:0;index"`
	LikeCount int    `gorm:"default:0;index"`
	Status    string `gorm:"size:20;default:'published';index"`
}

func (*Post) TableName() string {
	return "posts"
}
