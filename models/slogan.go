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

func SloganAll() ([]Slogan, error) {
	var slogans []Slogan
	result := database.Database.Where("deleted_at is NULL").Order("updated_at desc").Find(&slogans)
	if result.Error != nil {
		return nil, result.Error
	}
	return slogans, nil
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

// RandomSlogan returns a random slogan from the database designed for use in views through controllers
func RandomSlogan() (string, error) {
	var slogan Slogan
	if err := database.Database.Order("random()").First(&slogan).Error; err != nil {
		return "", err
	}
	return slogan.Slogan, nil
}
