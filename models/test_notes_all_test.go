package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestNotesAll(t *testing.T) {
	LoadEnv()
	database.Connect()

	// Create a test user
	email := "testuser@example.com"
	password := "testpassword"
	user, _ := UserCreate(email, password)

	// Create test notes
	note1 := Note{Title: "Test Note 1", Content: "Test Content 1", UserID: user.ID}
	note2 := Note{Title: "Test Note 2", Content: "Test Content 2", UserID: user.ID}
	database.Database.Create(&note1)
	database.Database.Create(&note2)

	// Test NotesAll function
	notes, err := NotesAll()
	assert.NoError(t, err, "NotesAll should not return an error")
	assert.NotNil(t, notes, "Notes should not be nil")

	// Check if the returned notes contain the test notes
	note1Found := false
	note2Found := false
	for _, note := range notes {
		if note.ID == note1.ID {
			note1Found = true
		}
		if note.ID == note2.ID {
			note2Found = true
		}
	}
	assert.True(t, note1Found, "Test Note 1 should be in the returned notes")
	assert.True(t, note2Found, "Test Note 2 should be in the returned notes")

	// Cleanup
	database.Database.Unscoped().Delete(&note1)
	database.Database.Unscoped().Delete(&note2)
	database.Database.Unscoped().Delete(user)
}
