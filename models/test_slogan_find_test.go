package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestSloganFind(t *testing.T) {
	// Load env and connect to database
	LoadEnv()
	database.Connect()

	// Create a test user
	testUser := &User{Username: "testuser@example.com", Password: "testpassword"}
	database.Database.Create(testUser)

	// Create a test slogan
	sloganContent := "This is a test slogan content."
	slogan, _ := SloganCreate(testUser, sloganContent)

	// Test slogan finding
	foundSlogan, err := SloganFind(uint64(slogan.ID))
	assert.NoError(t, err, "Finding slogan should not return an error")
	assert.NotNil(t, foundSlogan, "Found slogan should not be nil")
	assert.Equal(t, slogan.ID, foundSlogan.ID, "Slogan ID should match")
	assert.Equal(t, sloganContent, foundSlogan.Slogan, "Slogan content should match")
	assert.Equal(t, testUser.ID, foundSlogan.UserID, "User ID should match")

	// Test slogan not found
	nonExistentID := uint64(999999999)
	_, err = SloganFind(nonExistentID)
	assert.Error(t, err, "Finding non-existent slogan should return an error")
	assert.Equal(t, "Slogan not found", err.Error(), "Error message should match")

	// Cleanup
	database.Database.Unscoped().Delete(slogan)
	database.Database.Unscoped().Delete(testUser)
}
