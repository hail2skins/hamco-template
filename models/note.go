package models

import (
	"errors"
	"log"

	"github.com/hail2skins/hamcois-new/database"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title   string `gorm:"size:255;notnull" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserID  uint   `gorm:"notnull" json:"-"`
}

func NotesAll() ([]Note, error) {
	var notes []Note
	result := database.Database.Where("deleted_at is NULL").Order("updated_at desc").Find(&notes)
	if result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}

func NotesCreate(user *User, title string, content string) (*Note, error) {
	entry := Note{Title: title, Content: content, UserID: user.ID}
	result := database.Database.Create(&entry)
	if result.Error != nil {
		log.Printf("Error creating note: %v", result.Error)
		return nil, result.Error
	}

	log.Printf("Note created: ID: %d, Title: %s, Content: %s, UserID: %d\n", entry.ID, entry.Title, entry.Content, entry.UserID)
	return &entry, nil
}

func NotesFind(id uint64) (*Note, error) {
	var note Note
	result := database.Database.Where("id = ?", id).First(&note)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Note not found")
		}
		log.Printf("Error finding note: %v", result.Error)
		return nil, errors.New("Error finding note")
	}
	return &note, nil
}

func NotesFindByUser(user *User, id uint64) (*Note, error) {
	var note Note
	result := database.Database.Where("id = ? AND user_id = ?", id, user.ID).First(&note)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Note not found")
		}
		log.Printf("Error finding note: %v", result.Error)
		return nil, errors.New("Error finding note")
	}

	return &note, nil
}

// NotesLastFive returns the last five notes and is designed for the home index page.
func NotesLastFive() (*[]Note, error) {
	var notes []Note
	result := database.Database.Where("deleted_at is NULL").Order("updated_at desc").Limit(5).Find(&notes)
	if result.Error != nil {
		return nil, result.Error
	}

	return &notes, nil
}

func (note *Note) Update(title string, content string) error {
	note.Title = title
	note.Content = content
	result := database.Database.Save(&note)
	if result.Error != nil {
		log.Printf("Error updating note: %v", result.Error)
		return result.Error
	}
	return nil
}

func NotesMarkDelete(user *User, id uint64) error {
	// Update notes set deleted_at == Current Time> Where id = id and user_id = user_id
	result := database.Database.Where("id = ? AND user_id = ?", id, user.ID).Delete(&Note{})
	if result.Error != nil {
		log.Printf("Error deleting note: %v", result.Error)
		return result.Error
	}
	return nil
}
