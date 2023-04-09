package models

import (
	"errors"
	"log"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/hail2skins/hamcois-new/helpers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;notnull;unique" json:"username"`
	Password string `gorm:"size:255;notnull" json:"-"`
	Notes    []Note `gorm:"foreignKey:UserID"`
}

func CheckEmailAvailable(email string) bool {
	var user User
	result := database.Database.Where("username = ?", email).Find(&user)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Printf("Error checking email availability: %v", result.Error)
		return false
	}
	return result.RowsAffected == 0
}

func UserCreate(email string, password string) (*User, error) {
	hshPasswd, err := helpers.HashPassword(password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, errors.New("Error hashing password")
	}

	entry := User{Username: email, Password: hshPasswd}
	result := database.Database.Create(&entry)

	if result.Error != nil {
		log.Printf("Error creating user: %v", result.Error)
		return nil, errors.New("Error creating user")
	}

	return &entry, nil
}

func UserFind(id uint64) (*User, error) {
	var user User
	result := database.Database.Preload("Notes").Where("id = ?", id).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		log.Printf("Error finding user: %v", result.Error)
		return nil, errors.New("Error finding user")
	}
	return &user, nil
}

func UserFindByEmailAndPassword(email string, password string) (*User, error) {
	var user User
	result := database.Database.Where("username = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		log.Printf("Error finding user by email: %v", result.Error)
		return nil, errors.New("Error finding user by email")
	}

	match := helpers.CheckPasswordHash(password, user.Password)
	if match {
		return &user, nil
	} else {
		return nil, errors.New("Incorrect password")
	}
}
