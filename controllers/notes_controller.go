package controllers

import (
	"hamco-template/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	notes := models.NotesAll()
	c.HTML(
		http.StatusOK,
		"note/index.html",
		gin.H{
			"notes": notes,
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
