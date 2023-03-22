package api

import (
	"RayaneshBackend/pkg/session"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// userIDAuthedContext is the key in context which we use to store the user ID in it
const userIDAuthedContext = "userID"

// getAccessTokenFromHeaders will get the access token in the Authorization header.
// If header is empty or does not have the Bearer prefix, returns an empty string.
func getAccessTokenFromHeaders(c *gin.Context) string {
	const headerName = "Authorization"
	const prefix = "Bearer "
	header := c.Request.Header.Get(headerName)
	if !strings.HasPrefix(header, prefix) {
		return ""
	}
	return header[len(prefix):]
}

func (api *API) AuthorizeUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the header
		header := getAccessTokenFromHeaders(c)
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{"empty auth"})
			return
		}
		// Authorize
		userID, err := api.Session.Get(header)
		if err == session.ErrNotFound {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{"bad auth"})
			return
		} else if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{errInternalError})
			return
		}
		// Set in map
		c.Set(userIDAuthedContext, userID)
	}
}

// CORS adds the CORS headers to request
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
