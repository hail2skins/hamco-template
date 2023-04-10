package helpers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hail2skins/hamcois-new/helpers"
	"github.com/hail2skins/hamcois-new/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestSessionGet(t *testing.T) {
	router := testhelpers.SetupTestRouter()

	userID := uint64(42)

	router.GET("/test", func(c *gin.Context) {
		// Set the session value
		session := sessions.Default(c)
		session.Set("id", interface{}(userID))
		session.Save()

		// Retrieve the session value using SessionGet
		retrievedID := helpers.SessionGet(c)

		assert.Equal(t, userID, retrievedID)
		c.Status(200)
	})

	// Send a request to the test route to trigger the session get
	req, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
