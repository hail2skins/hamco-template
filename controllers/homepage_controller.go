package controllers

import (
	"net/http"

	"github.com/hail2skins/hamco-new/controllers/helpers"
	"github.com/hail2skins/hamco-new/models"

	"github.com/gin-gonic/gin"
)

// NoteView is a struct to hold the note and the published date which is formatted in the Index function
// This struct exists in the notes_controller.go file as NoteView.  Perhaps it should be moved to a common file.
type HomeNoteView struct {
	models.Note
	Published string
}

func Index(c *gin.Context) {
	notes := models.NotesLastFive()
	noteViews := helpers.NotesToNoteViews(notes)
	c.HTML(
		http.StatusOK,
		"home/index.html",
		gin.H{
			"title":     "Hamco Internet Solutions",
			"logged_in": c.MustGet("logged_in").(bool),
			"notes":     noteViews, // Pass the slice of NoteView structs to the template rather than the notes directly
		})
	//	fmt.Println(c.GetUint64("user_id")) // Unnecssary code check but leaving for later possible use
	//	fmt.Println(c.Get("user_id")) // Also unnecessary code check but leaving for later possible use
	// fmt.Println(c.GetUint("user_id")) // Also unnecessary code check but leaving for later possible use

}

func About(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/about.html",
		gin.H{
			"title":     "About",
			"logged_in": c.MustGet("logged_in").(bool), // This is the correct way to get the value from the gin.Context
		},
	)
}

func Contact(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/contact.html",
		gin.H{
			"title": "Contact",
		},
	)
}
