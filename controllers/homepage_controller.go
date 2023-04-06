package controllers

import (
	"hamco-template/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"title":     "Hamco Internet Solutions",
		"logged_in": helpers.IsUserLoggedIn(c),
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
