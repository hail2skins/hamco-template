package testhelpers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("test", store))
	return router
}
