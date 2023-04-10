package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestNotesUpdate(t *testing.T) {
	LoadEnv()
	database.Connect()

	// Create a test user
	testUser := &User{Username: "testuser@example.com", Password: "testpassword"}
	database.Database.Create(testUser)

	// Create a test note
	title := "Test Note"
	content := "Content for Test Note"
	note, err := NotesCreate(testUser, title, content)
	assert.NoError(t, err, "Note creation should not return an error")

	// Update the note
	newTitle := "Updated Test Note"
	newContent := "Updated content for Test Note"
	err = note.Update(newTitle, newContent)
	assert.NoError(t, err, "Note update should not return an error")

	// Check if the note was updated correctly
	var updatedNote Note
	result := database.Database.Where("id = ?", note.ID).First(&updatedNote)
	assert.NoError(t, result.Error, "Error finding updated note")
	assert.Equal(t, newTitle, updatedNote.Title, "Note title should be updated")
	assert.Equal(t, newContent, updatedNote.Content, "Note content should be updated")

	// Cleanup
	database.Database.Unscoped().Delete(note)
	database.Database.Unscoped().Delete(testUser)
}
