package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/index.html",
		gin.H{
			"title": "Hamco Internet Solutions",
		},
	)
}

func About(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/about.html",
		gin.H{
			"title": "About",
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

func Post(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/post.html",
		gin.H{
			"title": "Post",
		},
	)
}
