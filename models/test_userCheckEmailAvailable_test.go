package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestCheckEmailAvailable(t *testing.T) {
	LoadEnv()
	database.Connect()

	// Create a test user
	testUser := &User{Username: "testuser@example.com", Password: "testpassword"}
	database.Database.Create(testUser)
	defer database.Database.Unscoped().Delete(testUser)

	// Test email availability
	isAvailable := CheckEmailAvailable("testuser@example.com")
	assert.False(t, isAvailable, "Email should not be available")

	isAvailable = CheckEmailAvailable("nonexistentemail@example.com")
	assert.True(t, isAvailable, "Email should be available")
}
