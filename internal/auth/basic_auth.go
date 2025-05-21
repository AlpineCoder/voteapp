package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BasicAuthMiddleware is a middleware that checks for basic authentication.
func BasicAuthMiddleware(username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, ok := c.Request.BasicAuth()
		if !ok || user != username || pass != password {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

func ApiKeyAuthMiddleware(expectedKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")
		if key != expectedKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing API key"})
			return
		}
		c.Next()
	}
}

func AlternateApiKeyAuthMiddleware(c *gin.Context) {
	expectedKey := "key"
	key := c.GetHeader("X-API-Key")
	if key != expectedKey {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing API key"})
		return
	}
	c.Next()
}
