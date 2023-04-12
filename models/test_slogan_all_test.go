package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestSloganAll(t *testing.T) {
	LoadEnv()
	database.Connect()

	// Create a test user
	email := "testuser@example.com"
	password := "testpassword"
	user, _ := UserCreate(email, password)

	// Create test Slogans
	slogan1 := Slogan{Slogan: "Test Slogan 1", UserID: user.ID}
	slogan2 := Slogan{Slogan: "Test Slogan 2", UserID: user.ID}
	database.Database.Create(&slogan1)
	database.Database.Create(&slogan2)

	// Test SlogansAll function
	Slogans, err := SloganAll()
	assert.NoError(t, err, "SloganAll should not return an error")
	assert.NotNil(t, Slogans, "Slogans should not be nil")

	// Check if the returned Slogans contain the test Slogans
	slogan1Found := false
	slogan2Found := false
	for _, Slogan := range Slogans {
		if Slogan.ID == slogan1.ID {
			slogan1Found = true
		}
		if Slogan.ID == slogan2.ID {
			slogan2Found = true
		}
	}
	assert.True(t, slogan1Found, "Test Slogan 1 should be in the returned Slogans")
	assert.True(t, slogan2Found, "Test Slogan 2 should be in the returned Slogans")

	// Cleanup
	database.Database.Unscoped().Delete(&slogan1)
	database.Database.Unscoped().Delete(&slogan2)
	database.Database.Unscoped().Delete(user)
}
