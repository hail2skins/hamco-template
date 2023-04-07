package models

import (
	"github.com/hail2skins/hamco-new/database"

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

func NotesFindByUser(user *User, id uint64) *Note {
	var note Note
	database.Database.Where("id = ? AND user_id = ?", id, user.ID).First(&note)
	return &note
}

// NotesLastFive returns the last five notes and is designed for the home index page.
func NotesLastFive() *[]Note {
	var notes []Note
	database.Database.Where("deleted_at is NULL").Order("updated_at desc").Limit(5).Find(&notes)
	return &notes
}

func (note *Note) Update(title string, content string) {
	note.Title = title
	note.Content = content
	database.Database.Save(&note)
}

func NotesMarkDelete(user *User, id uint64) {
	// Update notes set deleted_at == Current Time> Where id = id and user_id = user_id
	database.Database.Where("id = ? AND user_id = ?", id, user.ID).Delete(&Note{})
}
