package models

import (
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
	database.Database.Where("username = ?", email).First(&user)
	return (user.ID == 0)
}

func UserCreate(email string, password string) *User {
	hshPasswd, _ := helpers.HashPassword(password)
	entry := User{Username: email, Password: hshPasswd}
	database.Database.Create(&entry)
	return &entry
}

func UserFind(id uint64) *User {
	var user User
	database.Database.Preload("Notes").Where("id = ?", id).First(&user)
	return &user
}

func UserFindByEmailAndPassword(email string, password string) *User {
	var user User
	database.Database.Where("username = ?", email).First(&user)
	if user.ID == 0 {
		return nil
	}
	match := helpers.CheckPasswordHash(password, user.Password)
	if match {
		return &user
	} else {
		return nil
	}
}
