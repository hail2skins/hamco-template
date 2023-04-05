package main

import (
	"log"

	"hamco-template/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	//setup.LoadEnv()
	//setup.LoadDatabase()
	serveApplication()
}

func serveApplication() {
	r := gin.Default()
	//r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.Static("/scss", "./static/scss")
	r.Static("/vendor", "./static/vendor")
	r.Static("/js", "./static/js")
	r.StaticFile("/favicon.ico", "./img/favicon.ico")

	r.LoadHTMLGlob("templates/**/**")

	controllers.Router(r)

	log.Println("Server started")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//r.Run(":8000")
	//fmt.Println("Server running on port 8000")
}
