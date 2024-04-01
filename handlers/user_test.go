package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// These test cases only checks the status codes
func TestEndpoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	r.POST("/", func(c *gin.Context) {
		c.Status(http.StatusCreated)
	})

	t.Run("GET /", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/hello", nil)
		if err != nil {
			t.Fatalf("Error creating GET request: %v", err)
		}
		respRecorder := httptest.NewRecorder()

		r.ServeHTTP(respRecorder, req)

		if status := respRecorder.Code; status != http.StatusOK {
			t.Errorf("GET /GetAllUsers returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("POST /", func(t *testing.T) {
		requestBody := []byte(`{"name": "rahul", "email": "rahul@gmail.com"}`)

		req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatalf("Error creating POST request: %v", err)
		}

		respRecorder := httptest.NewRecorder()
		r.ServeHTTP(respRecorder, req)

		if status := respRecorder.Code; status != http.StatusCreated {
			t.Errorf("POST /CreateUser returned wrong status code: got %v want %v",
				status, http.StatusCreated)
		}
	})
}
