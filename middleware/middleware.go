package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/lib/jwt"
)

func SetJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func SetDatabase(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", client)
		c.Next()
	}
}

func SetRedis(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redis", client)
		c.Next()
	}
}

func SetContext(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("context", ctx)
		c.Next()
	}
}

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := jwt.ValidateToken(c); err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
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
