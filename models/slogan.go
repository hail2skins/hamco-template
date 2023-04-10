package models

import (
	"errors"
	"log"

	"github.com/hail2skins/hamcois-new/database"
	"gorm.io/gorm"
)

type Slogan struct {
	gorm.Model
	Slogan string `gorm:"size:255;notnull;unique" json:"slogan"`
	UserID uint   `gorm:"notnull" json:"-"`
}

func SloganCreate(user *User, slogan string) (*Slogan, error) {
	entry := Slogan{Slogan: slogan, UserID: user.ID}
	result := database.Database.Create(&entry)
	if result.Error != nil {
		log.Printf("Error creating slogan: %v", result.Error)
		return nil, result.Error
	}

	log.Printf("Slogan created: ID: %d, Slogan: %s, UserID: %d\n", entry.ID, entry.Slogan, entry.UserID)
	return &entry, nil
}

func SloganFind(id uint64) (*Slogan, error) {
	var slogan Slogan
	result := database.Database.Where("id = ?", id).First(&slogan)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Slogan not found")
		}
		log.Printf("Error finding slogan: %v", result.Error)
		return nil, errors.New("Error finding slogan")
	}
	return &slogan, nil
}
