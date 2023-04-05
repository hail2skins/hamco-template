package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:255;notnull;unique" json:"username"`
	Password string `gorm:"size:255;notnull" json:"-"`
	Notes    []Note `gorm:"foreignKey:UserID"`
}
