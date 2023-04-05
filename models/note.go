package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title   string `gorm:"size:255;notnull" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserID  uint   `gorm:"notnull" json:"-"`
}
