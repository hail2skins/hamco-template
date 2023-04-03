package main

import (
	"log"
	"web/controller"
	"web/database"
	"web/model"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	loadEnv()
	loadDatabase()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.Static("/scss", "./static/scss")
	r.Static("/vendor", "./static/vendor")
	r.Static("/js", "./static/js")
	r.StaticFile("/favicon.ico", "./img/favicon.ico")

	r.LoadHTMLGlob("templates/**/*")
	controller.Router(r)

	log.Println("Server started")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Post{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
