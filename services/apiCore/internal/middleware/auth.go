package middleware

import (
	"errors"
	"go-mail-sender/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrTokenExpired = errors.New("Token has expired")

func (mw *MiddlewareManager) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the JWT token from the request header
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		claims, err := jwt.ValidateJWTToken(token, mw.cfg)
		if err != nil {
			if err == ErrTokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			}
			c.Abort()
			return
		}

		c.Set("userClaims", claims)
		c.Next()
	}
}
