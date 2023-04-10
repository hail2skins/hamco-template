package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestNotesMarkDelete(t *testing.T) {
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

	// Mark the note as deleted
	err = NotesMarkDelete(testUser, uint64(note.ID))
	assert.NoError(t, err, "Note deletion should not return an error")

	// Check if the note was marked as deleted
	var deletedNote Note
	result := database.Database.Unscoped().Where("id = ?", note.ID).First(&deletedNote)
	assert.NoError(t, result.Error, "Error finding deleted note")
	assert.NotNil(t, deletedNote.DeletedAt, "Note should be marked as deleted")

	// Cleanup
	database.Database.Unscoped().Delete(&deletedNote)
	database.Database.Unscoped().Delete(testUser)
}
