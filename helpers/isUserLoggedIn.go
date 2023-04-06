package helpers

import "github.com/gin-gonic/gin"

func IsUserLoggedIn(c *gin.Context) bool {
	_, exists := c.Get("user_id")
	return exists
}

// can also do it this way apparently because GetUint64 returns 0 if not found
// func IsUserLoggedIn(c *gin.Context) bool {
//	return (c.GetUint("user_id") > 0))
//}
