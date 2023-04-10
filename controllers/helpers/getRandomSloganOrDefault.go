package helpers

import (
	"log"

	"github.com/hail2skins/hamcois-new/models"
)

func GetRandomSloganOrDefault() string {
	slogan, err := models.RandomSlogan()
	if err != nil {
		log.Printf("Error getting slogan: %v", err)
		return "This is the default subhead."
	}
	return slogan
}
