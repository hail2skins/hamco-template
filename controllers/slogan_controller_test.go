package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hail2skins/hamcois-new/testhelpers"
)

func TestSloganNew(t *testing.T) {
	logBuffer := &bytes.Buffer{}
	router := testhelpers.SetupTestRouter(logBuffer)

	slogans := router.Group("/slogans")
	{
		slogans.GET("/new", SloganNew)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/slogans/new", nil)

	router.ServeHTTP(w, req)

	responseBody := w.Body.String()

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	expectedText := "New Slogan"
	testhelpers.CustomAssertContains(t, responseBody, expectedText, "TestSloganNew - Response title does not contain '%s'. Status code: %d", expectedText, w.Code)

	t.Log("Gin log output:")
	t.Log(logBuffer.String())
}
