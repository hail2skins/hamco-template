package model

import (
	"web/database"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content string `gorm:"type:text" json:"content"`
	UserID  uint
}

func (post *Post) Save() (*Post, error) {
	err := database.Database.Create(&post).Error
	if err != nil {
		return &Post{}, err
	}
	return post, nil
}
