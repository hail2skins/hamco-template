package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestNotesCreate(t *testing.T) {
	LoadEnv()
	database.Connect()

	// Create a test user
	testUser := &User{Username: "testuser@example.com", Password: "testpassword"}
	database.Database.Create(testUser)

	// Test note creation
	title := "Test note title"
	content := "This is a test note content."
	note, err := NotesCreate(testUser, title, content)
	assert.NoError(t, err, "Note creation should not return an error")
	assert.NotNil(t, note, "Note should not be nil")

	// Check if the note was created correctly
	var createdNote Note
	result := database.Database.Where("title = ? AND content = ?", title, content).Find(&createdNote)
	assert.NoError(t, result.Error, "Error finding created note")
	assert.Equal(t, title, createdNote.Title, "Note title should match")
	assert.Equal(t, content, createdNote.Content, "Note content should match")
	assert.Equal(t, testUser.ID, createdNote.UserID, "User ID should match")

	// Cleanup
	database.Database.Unscoped().Delete(note)
	database.Database.Unscoped().Delete(testUser)
}
