package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hail2skins/hamcois-new/database"
	"github.com/hail2skins/hamcois-new/helpers"
	"github.com/hail2skins/hamcois-new/models"
	"github.com/hail2skins/hamcois-new/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticateUser(t *testing.T) {
	models.LoadEnv()
	database.Connect()
	router := testhelpers.SetupTestRouter()

	// Create a unique user
	user, err := models.UserCreate(fmt.Sprintf("test%d@example.com", time.Now().UnixNano()), "password")
	if err != nil {
		t.Fatalf("Error creating user: %v", err)
	}

	router.Use(AuthenticateUser())

	router.GET("/test", func(c *gin.Context) {
		loggedIn := c.GetBool("logged_in")
		userID, _ := c.Get("user_id")
		email, _ := c.Get("email")

		if loggedIn {
			assert.Equal(t, user.ID, userID.(uint64))
			assert.Equal(t, user.Username, email.(string))
		} else {
			assert.False(t, loggedIn)
		}
		c.Status(200)
	})

	// Set session route for testing
	router.POST("/set-session", func(c *gin.Context) {
		userID := c.PostForm("user_id")
		userIDUint, _ := strconv.ParseUint(userID, 10, 64)
		helpers.SessionSet(c, userIDUint)
		c.Status(200)
	})

	// Test the middleware when the user is not logged in
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Set the session for the user
	w = httptest.NewRecorder()
	form := url.Values{}
	form.Add("user_id", fmt.Sprintf("%d", user.ID))
	req, _ = http.NewRequest("POST", "/set-session", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Test the middleware when the user is logged in
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Cleanup
	database.Database.Unscoped().Delete(user)
}
