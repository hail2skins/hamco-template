package models

import (
	"hamco-template/database"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title   string `gorm:"size:255;notnull" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserID  uint   `gorm:"notnull" json:"-"`
}

func NotesAll() *[]Note {
	var notes []Note
	database.Database.Where("deleted_at is NULL").Order("updated_at desc").Find(&notes)
	return &notes
}
