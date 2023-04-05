package setup

import (
	"hamco-template/database"
	//"hamco-template/models"
)

func LoadDatabase() {
	database.Connect()
	//database.Database.AutoMigrate(&models.User{})
	//database.Database.AutoMigrate(&models.Note{})
}
