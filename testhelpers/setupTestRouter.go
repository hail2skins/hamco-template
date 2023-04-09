package testhelpers

import (
	"io"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/hail2skins/hamcois-new/middlewares"
)

func SetupTestRouter(logOutput io.Writer) *gin.Engine {
	gin.SetMode(gin.TestMode) // Set Gin to Test Mode

	router := gin.New() // Create a new Gin instance without the default logger

	// Add the logger and recovery middleware
	router.Use(gin.LoggerWithWriter(logOutput), gin.Recovery())

	// Sessions init
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("slogans", store))

	router.Use(middlewares.AuthenticateUser())

	router.LoadHTMLGlob("../templates/**/**")

	return router
}
