package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestUserFind(t *testing.T) {
	LoadEnv()
	database.Connect()

	// Create a test user
	email := "testuserfind@example.com"
	password := "testpassword"
	testUser, err := UserCreate(email, password)
	assert.NoError(t, err, "User creation should not return an error")

	// Test UserFind
	foundUser, err := UserFind(uint64(testUser.ID))
	assert.NoError(t, err, "UserFind should not return an error")
	assert.NotNil(t, foundUser, "Found user should not be nil")
	assert.Equal(t, testUser.ID, foundUser.ID, "User ID should match")

	// Test UserFind with non-existent ID
	_, err = UserFind(0)
	assert.Error(t, err, "UserFind should return an error for non-existent user")

	// Cleanup
	database.Database.Unscoped().Delete(testUser)
}
