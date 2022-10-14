package middleware

import (
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

func SetEntClient(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("database", client)
		c.Next()
	}
}

func SetRedisClient(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redis", client)
		c.Next()
	}
}

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		rc, _ := c.Keys["redis"].(*redis.Client)
		// validate token
		accessData, err := jwt.ExtractTokenMetadata(c)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		// validate in redis token
		userID, err := jwt.GetUserIDFromRedis(rc, accessData.AccessUUID)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.Set("x-request-id", userID)
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
