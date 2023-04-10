package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestRandomSlogan(t *testing.T) {
	LoadEnv()
	// Connect to the database
	database.Connect()

	// Create a test user
	user, err := UserCreate("test@example.com", "testpassword")
	assert.NoError(t, err, "User creation should not result in an error")

	// Create a few test slogans
	slogans := []string{"Test slogan 1", "Test slogan 2", "Test slogan 3"}
	for _, s := range slogans {
		_, err := SloganCreate(user, s)
		assert.NoError(t, err, "Slogan creation should not result in an error")
	}

	// Test the RandomSlogan function
	randomSlogan, err := RandomSlogan()
	assert.NoError(t, err, "RandomSlogan should not result in an error")
	assert.Contains(t, slogans, randomSlogan, "RandomSlogan should return one of the test slogans")

	// Cleanup
	for _, s := range slogans {
		var slogan Slogan
		database.Database.Where("slogan = ?", s).First(&slogan)
		database.Database.Unscoped().Delete(&slogan)
	}
	database.Database.Unscoped().Delete(user)
}
