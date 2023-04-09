package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hail2skins/hamcois-new/controllers/helpers"
	"github.com/hail2skins/hamcois-new/models"
)

func SloganNew(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"slogan/new.html",
		gin.H{
			"title":     "New Slogan",
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
