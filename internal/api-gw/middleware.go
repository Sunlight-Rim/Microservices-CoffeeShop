package gw

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/// MIDDLEWARE

const jwtKey = "verySecretKey" // TODO: move to config

var h = hmac.New(sha256.New, []byte(jwtKey))

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := strings.Split(strings.Split(c.Request.Header["Authorization"][0], " ")[1], ".")
		// verify signature
		h.Write([]byte(token[0] + "." + token[1]))
		defer h.Reset()
		if base64.RawURLEncoding.EncodeToString(h.Sum(nil)) != token[2] {
			c.String(http.StatusUnauthorized, "token is invalid")
			c.Abort()
			return
		}
		c.Next()
	}
}
