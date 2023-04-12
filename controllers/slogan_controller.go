package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hail2skins/hamcois-new/controllers/helpers"
	"github.com/hail2skins/hamcois-new/models"
)

func SloganIndex(c *gin.Context) {
	currentUser := helpers.RequireLoggedInUser(c)

	if currentUser != nil {
		slogan := helpers.GetRandomSloganOrDefault()
		slogans, err := models.SloganAll()

		if err != nil {
			// Handle the error, e.g., log it and return an error response
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching notes"})
			return
		}

		c.HTML(
			http.StatusOK,
			"slogan/index.html",
			gin.H{
				// Pass the slice of slogans
				"slogans":   slogans,
				"title":     "Silly Slogans",
				"slogan":    slogan,
				"logged_in": c.MustGet("logged_in").(bool),
			},
		)
	}
}

func SloganNew(c *gin.Context) {
	slogan := helpers.GetRandomSloganOrDefault()

	c.HTML(
		http.StatusOK,
		"slogan/new.html",
		gin.H{
			"title":     "New Slogan",
			"slogan":    slogan,
			"logged_in": c.MustGet("logged_in").(bool),
		},
	)
}

func SloganCreate(c *gin.Context) {
	currentUser := helpers.RequireLoggedInUser(c)
	if currentUser != nil {
		// Can do this through a bind.  See go-gin-bootcamp repo for actual code
		slogan := c.PostForm("slogan")
		// call the model to create the note
		models.SloganCreate(currentUser, slogan)
		// redirect to the notes index
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
