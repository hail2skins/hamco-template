package helpers

import (
	"html/template"
	"net/http"

	"github.com/hail2skins/hamcois-new/models"

	"github.com/gin-gonic/gin"
)

func RequireLoggedInUser(c *gin.Context) *models.User {
	currentUser := GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		notes, _ := models.NotesLastFive()
		noteViews := NotesToNoteViews(notes)

		// Render the content of each note using Goldmark
		for i := range noteViews {
			truncatedContent := TruncateWords(string(noteViews[i].Content), 50) // Limit to 50 words, for example
			noteViews[i].Content = template.HTML(truncatedContent)
		}

		c.HTML(
			http.StatusUnauthorized,
			"home/index.html",
			gin.H{
				"alert":     "You must be logged in to perform this action",
				"notes":     noteViews,
				"logged_in": false,
			},
		)

		c.Abort() // Abort the request
		return nil
	}
	return currentUser
}
