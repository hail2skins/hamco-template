package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/hail2skins/hamcois-new/controllers/helpers"
	"github.com/hail2skins/hamcois-new/models"
	"github.com/russross/blackfriday/v2"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	notes := models.NotesAll()
	noteViews := helpers.NotesToNoteViews(notes)
	c.HTML(
		http.StatusOK,
		"note/index.html",
		gin.H{
			// Pass the slice of NoteView structs to the template rather than the notes directly
			"notes":     noteViews,
			"logged_in": c.MustGet("logged_in").(bool),
		},
	)
}

func NotesNew(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"note/new.html",
		gin.H{
			"title":     "New Note",
			"logged_in": c.MustGet("logged_in").(bool),
		},
	)
}

func NotesCreate(c *gin.Context) {
	currentUser := helpers.RequireLoggedInUser(c)
	if currentUser != nil {
		// Can do this through a bind.  See go-gin-bootcamp repo for actual code
		title := c.PostForm("title")
		content := c.PostForm("content")
		// call the model to create the note
		models.NotesCreate(currentUser, title, content)
		// redirect to the notes index
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func NotesShow(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		fmt.Printf("Error parsing note id: %v\n", err)
	}
	note := models.NotesFind(id)
	published := note.UpdatedAt.Format("Jan 2, 2006")

	// Render the Markdown content with Blackfriday
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{})
	htmlContent := blackfriday.Run([]byte(note.Content),
		blackfriday.WithRenderer(renderer),
		blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.HardLineBreak|blackfriday.AutoHeadingIDs|blackfriday.Autolink),
	)

	c.HTML(
		http.StatusOK,
		"note/show.html",
		gin.H{
			"note":      note,
			"content":   template.HTML(htmlContent),
			"published": published,
			"logged_in": c.MustGet("logged_in").(bool),
		},
	)
}

func NotesEditPage(c *gin.Context) {
	currentUser := helpers.RequireLoggedInUser(c)
	if currentUser != nil {
		idStr := c.Param("id")
		//fmt.Printf("ID string: %s\n", idStr) // Debugging statement
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing note id: %v\n", err)
		}
		//fmt.Printf("Parsed ID: %d\n", id) // Debugging statement
		note := models.NotesFindByUser(currentUser, id)
		c.HTML(
			http.StatusOK,
			"note/edit.html",
			gin.H{
				"note":      note,
				"logged_in": c.MustGet("logged_in").(bool),
			},
		)
	}
}

func NotesUpdate(c *gin.Context) {
	currentUser := helpers.RequireLoggedInUser(c)
	if currentUser != nil {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing note id: %v\n", err)
		}
		note := models.NotesFindByUser(currentUser, id)
		title := c.PostForm("title")
		content := c.PostForm("content")
		note.Update(title, content)
		c.Redirect(http.StatusMovedPermanently, "/notes/"+idStr)
	}
}

func NotesDelete(c *gin.Context) {
	currentUser := helpers.RequireLoggedInUser(c)
	if currentUser != nil {
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing note id: %v\n", err)
		}
		models.NotesMarkDelete(currentUser, id)
		c.Redirect(http.StatusMovedPermanently, "/notes")
	}
}
