package setup

import (
	"github.com/hail2skins/hamco-new/database"
	"github.com/hail2skins/hamco-new/models"
)

func LoadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&models.User{})
	database.Database.AutoMigrate(&models.Note{})
}
