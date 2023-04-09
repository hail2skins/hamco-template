package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestSloganCreate(t *testing.T) {
	LoadEnv()
	database.Connect()

	// Create a test user
	testUser := &User{Username: "testuser@example.com", Password: "testpassword"}
	database.Database.Create(testUser)

	// Create a test slogan
	sloganText := "Test slogan for Go test"
	slogan := SloganCreate(testUser, sloganText)

	// Test the slogan
	assert.NotNil(t, slogan, "Slogan should not be nil")
	assert.Equal(t, sloganText, slogan.Slogan, "Slogan text should match")
	assert.Equal(t, testUser.ID, slogan.UserID, "User ID should match")

	// Test the user
	assert.NotNil(t, testUser.ID, "User ID should not be nil")
	assert.Equal(t, "testuser@example.com", testUser.Username, "User email should match")

	// Delete the slogan
	database.Database.Unscoped().Delete(slogan)

	// Delete the user
	database.Database.Unscoped().Delete(testUser)
}
