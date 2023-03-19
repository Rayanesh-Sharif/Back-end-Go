package api

import (
	"github.com/gin-gonic/gin"
	"strings"
)

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
