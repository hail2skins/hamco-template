package models

import (
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestNotesFindByUser(t *testing.T) {
	LoadEnv()
	database.Connect()

	// Create a test user
	testUser := &User{Username: "testuser@example.com", Password: "testpassword"}
	database.Database.Create(testUser)

	// Create another test user
	anotherUser := &User{Username: "anotheruser@example.com", Password: "testpassword"}
	database.Database.Create(anotherUser)

	// Create a test note for the first user
	title := "Test note title"
	content := "This is a test note content."
	note, _ := NotesCreate(testUser, title, content)

	// Test note finding by user
	foundNote, err := NotesFindByUser(testUser, uint64(note.ID))
	assert.NoError(t, err, "Finding note by user should not return an error")
	assert.NotNil(t, foundNote, "Found note should not be nil")
	assert.Equal(t, note.ID, foundNote.ID, "Note ID should match")
	assert.Equal(t, title, foundNote.Title, "Note title should match")
	assert.Equal(t, content, foundNote.Content, "Note content should match")
	assert.Equal(t, testUser.ID, foundNote.UserID, "User ID should match")

	// Test note not found for another user
	_, err = NotesFindByUser(anotherUser, uint64(note.ID))
	assert.Error(t, err, "Finding note by another user should return an error")
	assert.Equal(t, "Note not found", err.Error(), "Error message should match")

	// Cleanup
	database.Database.Unscoped().Delete(note)
	database.Database.Unscoped().Delete(testUser)
	database.Database.Unscoped().Delete(anotherUser)
}
