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

func NotesCreate(user *User, title string, content string) *Note {
	entry := Note{Title: title, Content: content, UserID: user.ID}
	database.Database.Create(&entry)
	return &entry
}

func NotesFind(id uint64) *Note {
	var note Note
	database.Database.Where("id = ?", id).First(&note)
	return &note
}
