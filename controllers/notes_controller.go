package controllers

import (
	"fmt"
	"hamco-template/controllers/helpers"
	"hamco-template/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// NoteView is a struct to hold the note and the published date which is formatted in the NotesIndex function
type NoteView struct {
	models.Note
	Published string
}

func NotesIndex(c *gin.Context) {
	notes := models.NotesAll()
	// Create a slice of NoteView structs
	var noteViews []NoteView
	// Loop through the notes and create a NoteView struct for each note
	for _, note := range *notes {
		// Format the date
		published := note.UpdatedAt.Format("Jan 2, 2006")
		noteView := NoteView{
			Note:      note,
			Published: published,
		}
		noteViews = append(noteViews, noteView)
	}
	c.HTML(
		http.StatusOK,
		"note/index.html",
		gin.H{
			// Pass the slice of NoteView structs to the template rather than the notes directly
			"notes": noteViews,
		},
	)
}

func NotesNew(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"note/new.html",
		gin.H{
			"title": "New Note",
		},
	)
}

func NotesCreate(c *gin.Context) {
	// Declare the currentUser variable
	currentUser := helpers.GetUserFromRequest(c)
	// Check if the user is logged in
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"note/index.html",
			gin.H{
				"alert": "You must be logged in to create a note",
			},
		)
		return
	}
	// Can do this through a bind.  See go-gin-bootcamp repo for actual code
	title := c.PostForm("title")
	content := c.PostForm("content")
	// call the model to create the note
	models.NotesCreate(currentUser, title, content)
	// redirect to the notes index
	c.Redirect(http.StatusMovedPermanently, "/")
}

func NotesShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error parsing note id: %v\n", err)
	}
	note := models.NotesFind(id)
	published := note.UpdatedAt.Format("Jan 2, 2006")
	c.HTML(
		http.StatusOK,
		"note/show.html",
		gin.H{
			"note":      note,
			"published": published,
		},
	)
}
