package helpers

import (
	"hamco-template/models"

	"github.com/gin-gonic/gin"
)

func GetUserFromRequest(c *gin.Context) *models.User {
	userID, _ := c.Get("user_id")

	var currentUser *models.User
	if userID != nil {
		currentUser = models.UserFind(userID.(uint64))
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
