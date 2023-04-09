package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/hail2skins/hamcois-new/controllers"
	"github.com/hail2skins/hamcois-new/middlewares"
	"github.com/hail2skins/hamcois-new/setup"
)

func main() {
	setup.LoadEnv()
	setup.LoadDatabase()
	serveApplication()
}

func serveApplication() {
	r := gin.Default()
	//r := gin.New() // I think default turns on logger and recovery automatically and new doesn't.
	//r.Use(gin.Logger())  // These are brought in by default
	//r.Use(gin.Recovery()) // These are brought in by default

	// Sessions init
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))
	r.Use(sessions.Sessions("slogans", store))

	r.Use(middlewares.AuthenticateUser())

	// Set up a map of functions for templates. This is where we can add our custom functions
	//r.SetFuncMap(template.FuncMap{
	//	"truncateWords": helpers.TruncateWords,
	//})
	r.LoadHTMLGlob("templates/**/**")

	r.GET("/", controllers.Index)
	r.GET("/contact", controllers.Contact)
	r.GET("/about", controllers.About)
	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)

	notes := r.Group("/notes")
	{
		notes.GET("/", controllers.NotesIndex)
		notes.GET("/new", controllers.NotesNew)
		notes.POST("/", controllers.NotesCreate)
		notes.GET("/:id", controllers.NotesShow)
		notes.GET("/edit/:id", controllers.NotesEditPage)
		notes.POST("/:id", controllers.NotesUpdate)
		notes.DELETE("/:id", controllers.NotesDelete)
	}

	slogans := r.Group("/slogans")
	{
		slogans.GET("/new", controllers.SloganNew)
		slogans.POST("/", controllers.SloganCreate)
	}

	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.Static("/scss", "./static/scss")
	r.Static("/vendor", "./static/vendor")
	r.Static("/js", "./static/js")
	r.StaticFile("/favicon.ico", "./img/favicon.ico")

	log.Println("Server started")
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
