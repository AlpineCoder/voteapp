package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	cookieName       = "voter_id"
	cookiePath       = "/"
	cookieAge        = 5 * 365 * 24 * time.Hour // ~5 years
	cookieDomain     = ""                       // set to your domain if you need it, e.g. "party.example.com"
	useSecureCookies = true                     // set true in production (HTTPS), false only for local http testing
)

// EnsureVoterID sets a voter_id cookie if missing and stores it on the context.
func EnsureVoterID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// SameSite=Lax is a nice default for simple apps
		ctx.SetSameSite(http.SameSiteLaxMode)

		id, err := ctx.Cookie(cookieName)
		if err != nil || id == "" {
			id = uuid.NewString()
			// Gin's SetCookie wraps http.SetCookie
			ctx.SetCookie(
				cookieName,
				id,
				int(cookieAge.Seconds()), // Max-Age in seconds
				cookiePath,
				cookieDomain,
				useSecureCookies, // Secure
				true,             // HttpOnly
			)
		}
		// Make it easy for handlers to access
		ctx.Set(cookieName, id)

		ctx.Next()
	}
}
