package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FakeApiKeyAuthMiddleware is a middleware that checks for a fake API key in the request header.
func FakeApiKeyAuthMiddleware(c *gin.Context) {
	expectedKey := "key"
	key := c.GetHeader("X-API-Key")
	if key != expectedKey {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing API key"})
		return
	}
	c.Next()
}

// GetBasicAuthUsers returns a map of usernames and passwords for basic authentication.
func GetBasicAuthUsers() map[string]string {
	// This is a hardcoded example. In a real application, you would retrieve this from a database or configuration file.
	return map[string]string{
		"admin": "password",
	}
}
