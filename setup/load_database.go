package setup

import (
	"github.com/hail2skins/hamcois-new/database"
	"github.com/hail2skins/hamcois-new/models"
)

func LoadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Note{})
}
