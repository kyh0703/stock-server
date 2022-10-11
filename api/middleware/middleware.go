package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kyh0703/stock-server/api/auth"
	"github.com/kyh0703/stock-server/ent"
)

func SetJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func SetAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access-token")
		if err != nil {
			c.Next()
			return
		}
		if err := auth.ValidateTokenFromCookie(accessToken); err != nil {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		// if err := auth.ValidateToken(c); err != nil {
		// 	c.AbortWithError(http.StatusUnauthorized, errors.New(("Unauthorized")))
		// 	return
		// }
		c.Next()
	}
}

func SetDatabase(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", client)
		c.Next()
	}
}

func SetContext(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("context", ctx)
		c.Next()
	}
}

func CheckLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("access-token")
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Next()
	}
}
