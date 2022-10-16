package middleware

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/lib/jwt"
)

func SetJSON() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !c.Is("json") {
			return fiber.ErrBadRequest
		}
		return c.Next()
	}
}

func TokenAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// validate token
		accessData, err := jwt.ExtractTokenMetadata(c)
		if err != nil {
			return c.Status(http.StatusUnauthorized).SendString(err.Error())
		}
		// validate in redis token
		userID, err := jwt.GetUserIDFromRedis(accessData.AccessUUID)
		if err != nil {
			return c.Status(http.StatusUnauthorized).SendString(err.Error())
		}
		ctx := context.WithValue(c.Context(), "user_id", userID)
		c.SetUserContext(ctx)
		return c.Next()
	}
}

func CheckLoggedIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// _, err := c.Cookie("access-token")
		// if err != nil {
		// 	c.AbortWithError(http.StatusUnauthorized, err)
		// 	return
		// }
		return c.Next()
	}
}
