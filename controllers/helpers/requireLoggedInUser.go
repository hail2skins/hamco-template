package helpers

import (
	"hamco-template/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireLoggedInUser(c *gin.Context) *models.User {
	currentUser := GetUserFromRequest(c)
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"note/index.html",
			gin.H{
				"alert": "You must be logged in to perform this action",
			},
		)
		c.Abort() // Abort the request
		return nil
	}
	return currentUser
}
