package helpers

import (
	"github.com/hail2skins/hamcois-new/models"

	"github.com/gin-gonic/gin"
)

func GetUserFromRequest(c *gin.Context) *models.User {
	userID := c.GetUint("user_id")

	var currentUser *models.User
	if userID > 0 {
		currentUser = models.UserFind(uint64(userID))
	} else {
		currentUser = nil
	}
	return currentUser
}
func SetPayload(c *gin.Context, h gin.H) gin.H {
	email := c.GetString("email")
	if len(email) > 0 {
		h["email"] = email
	}
	return h
}
