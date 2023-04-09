package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	LoadEnv()
	database.Connect()

	// Test user creation
	email := "testusercreate@example.com"
	password := "testpassword"
	user, err := UserCreate(email, password)
	assert.NoError(t, err, "User creation should not return an error")
	assert.NotNil(t, user, "User should not be nil")

	// Check if the user was created correctly
	var createdUser User
	result := database.Database.Where("username = ?", email).Find(&createdUser)
	assert.NoError(t, result.Error, "Error finding created user")
	assert.Equal(t, email, createdUser.Username, "User email should match")

	// Cleanup
	database.Database.Unscoped().Delete(user)
}
