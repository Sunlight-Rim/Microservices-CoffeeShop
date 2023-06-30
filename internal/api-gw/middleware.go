package gw

import (
	"encoding/base64"
	"hash"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/// MIDDLEWARE

func Auth(tokenHash hash.Hash) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader, ok := c.Request.Header["Authorization"]
		if !ok {
			c.String(http.StatusUnauthorized, "token was not specified")
			c.Abort()
			return
		}
		bearerToken := strings.Split(authHeader[0], " ")
		if len(bearerToken) != 2 {
			c.String(http.StatusUnauthorized, "token format is incorrect")
			c.Abort()
			return
		}
		if bearerToken[0] != "Bearer" {
			c.String(http.StatusUnauthorized, "'Bearer' keyword in token is needed")
			c.Abort()
			return
		}
		token := strings.Split(bearerToken[1], ".")
		if len(token) != 3 {
			c.String(http.StatusUnauthorized, "token format is incorrect")
			c.Abort()
			return
		}
		// verify signature
		tokenHash.Write([]byte(token[0] + "." + token[1]))
		defer tokenHash.Reset()
		if base64.RawURLEncoding.EncodeToString(tokenHash.Sum(nil)) != token[2] {
			c.String(http.StatusUnauthorized, "token is invalid")
			c.Abort()
			return
		}
		c.Next()
	}
}
