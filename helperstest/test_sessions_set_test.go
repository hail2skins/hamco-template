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

func TestSessionSet(t *testing.T) {
	router := testhelpers.SetupTestRouter()

	router.GET("/test", func(c *gin.Context) {
		userID := uint64(42)
		helpers.SessionSet(c, userID)

		session := sessions.Default(c)
		savedID := session.Get("id")

		assert.Equal(t, userID, savedID.(uint64))
		c.Status(200)
	})

	// Send a request to the test route to trigger the session set
	req, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
