package setup

import (
	"gin_notes/database"
	"gin_notes/models"
)

func LoadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Note{})
}
