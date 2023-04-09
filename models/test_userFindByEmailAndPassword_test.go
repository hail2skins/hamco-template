package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestUserFindByEmailAndPassword(t *testing.T) {
	LoadEnv()
	database.Connect()

	// Create a test user
	email := "testuserfind@example.com"
	password := "testpassword"
	testUser, err := UserCreate(email, password)
	assert.NoError(t, err, "User creation should not return an error")

	// Test UserFindByEmailAndPassword
	foundUser, err := UserFindByEmailAndPassword(email, password)
	assert.NoError(t, err, "UserFindByEmailAndPassword should not return an error")
	assert.NotNil(t, foundUser, "Found user should not be nil")
	assert.Equal(t, testUser.ID, foundUser.ID, "User ID should match")

	// Test UserFindByEmailAndPassword with incorrect password
	_, err = UserFindByEmailAndPassword(email, "wrongpassword")
	assert.Error(t, err, "UserFindByEmailAndPassword should return an error for incorrect password")

	// Test UserFindByEmailAndPassword with non-existent email
	_, err = UserFindByEmailAndPassword("nonexistent@example.com", password)
	assert.Error(t, err, "UserFindByEmailAndPassword should return an error for non-existent email")

	// Cleanup
	database.Database.Unscoped().Delete(testUser)
}
