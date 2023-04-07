package main

import (
	"log"
	"text/template"

	"hamco-template/controllers"
	"hamco-template/controllers/helpers"
	"hamco-template/middlewares"
	"hamco-template/setup"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
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

	// Sessions init
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))

	r.Use(middlewares.AuthenticateUser())

	// Set up a map of functions for templates. This is where we can add our custom functions
	r.SetFuncMap(template.FuncMap{
		"truncateWords": helpers.TruncateWords,
	})
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
