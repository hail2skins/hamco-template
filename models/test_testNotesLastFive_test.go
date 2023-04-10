package models

import (
	"fmt"
	"testing"

	"github.com/hail2skins/hamcois-new/database"
	"github.com/stretchr/testify/assert"
)

func TestNotesLastFive(t *testing.T) {
	LoadEnv()
	database.Connect()

	// Create a test user
	testUser := &User{Username: "testuser@example.com", Password: "testpassword"}
	database.Database.Create(testUser)

	// Create 6 test notes
	for i := 1; i <= 6; i++ {
		title := fmt.Sprintf("Test Note %d", i)
		content := fmt.Sprintf("Content for Test Note %d", i)
		_, err := NotesCreate(testUser, title, content)
		assert.NoError(t, err, "Note creation should not return an error")
	}

	// Call NotesLastFive function
	notes, err := NotesLastFive()
	assert.NoError(t, err, "NotesLastFive should not return an error")

	// Check if we got the last 5 notes
	assert.Len(t, *notes, 5, "NotesLastFive should return 5 notes")

	// Check if the notes are sorted in the correct order
	for i := 0; i < 4; i++ {
		assert.True(t, (*notes)[i].UpdatedAt.After((*notes)[i+1].UpdatedAt), "Notes should be sorted in descending order by UpdatedAt")
	}

	// Cleanup
	database.Database.Exec("DELETE FROM notes WHERE user_id = ?", testUser.ID)
	database.Database.Unscoped().Delete(testUser)
}
