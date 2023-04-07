package middlewares

import (
	"github.com/hail2skins/hamco-new/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if user is authenticated
		// If not, redirect to login page
		// If yes, continue
		session := sessions.Default(c)
		sessionID := session.Get("id")
		var user *models.User
		userPresent := true

		if sessionID == nil {
			userPresent = false
		} else {
			user = models.UserFind(sessionID.(uint64))
			userPresent = (user.ID > 0)
		}

		if userPresent {
			c.Set("user_id", user.ID)
			c.Set("email", user.Username)
		}

		// Set the logged_in value in the gin.Context for use in templates and controllers
		c.Set("logged_in", userPresent)

		c.Next()
	}
}
