package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestSloganUpdate(t *testing.T) {
	LoadEnv()
	database.Connect()

	// Create a test user
	testUser := &User{Username: "testuser@example.com", Password: "testpassword"}
	database.Database.Create(testUser)

	// Create a test note
	slogantext := "Test Slogan"
	slogan, err := SloganCreate(testUser, slogantext)
	assert.NoError(t, err, "Slogan creation should not return an error")

	// Update the note
	newSlogan := "Updated Test Slogan"
	err = slogan.Update(newSlogan)
	assert.NoError(t, err, "Slogan update should not return an error")

	// Check if the note was updated correctly
	var updatedSlogan Slogan
	result := database.Database.Where("id = ?", slogan.ID).First(&updatedSlogan)
	assert.NoError(t, result.Error, "Error finding updated slogan")
	assert.Equal(t, newSlogan, updatedSlogan.Slogan, "Slogan should be updated")

	// Cleanup
	database.Database.Unscoped().Delete(slogan)
	database.Database.Unscoped().Delete(testUser)
}
