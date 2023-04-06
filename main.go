package main

import (
	"log"

	"hamco-template/controllers"
	"hamco-template/setup"

	"github.com/gin-gonic/gin"
)

func main() {
	setup.LoadEnv()
	setup.LoadDatabase()
	serveApplication()
}

func serveApplication() {
	r := gin.Default()
	//r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("templates/**/**")

	r.GET("/", controllers.Index)
	r.GET("/contact", controllers.Contact)
	r.GET("/about", controllers.About)
	r.GET("/post", controllers.Post)

	notes := r.Group("/notes")
	{
		notes.GET("/", controllers.NotesIndex)
		notes.GET("/new", controllers.NotesNew)
	}
	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.Static("/scss", "./static/scss")
	r.Static("/vendor", "./static/vendor")
	r.Static("/js", "./static/js")
	r.StaticFile("/favicon.ico", "./img/favicon.ico")

	log.Println("Server started")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//r.Run(":8000")
	//fmt.Println("Server running on port 8000")
}
