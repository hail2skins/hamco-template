package models

import (
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
