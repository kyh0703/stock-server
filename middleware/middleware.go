package middleware

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/routes/auth"
	"github.com/kyh0703/stock-server/types"
)

func SetUserContext() fiber.Handler {
	// user context is not null
	return func(c *fiber.Ctx) error {
		c.SetUserContext(c.Context())
		return c.Next()
	}
}

func SetJson() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !c.Is("json") {
			return fiber.ErrBadRequest
		}
		return c.Next()
	}
}

func TokenAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			tokenString string
			authService auth.AuthService
		)

		// Get token string 'Header : Authorization: Bearer ${accessToken}'
		bearerToken := c.Get("Authorization")
		if bearerToken == "" {
			return c.App().ErrorHandler(c, types.ErrUnauthorized)
		}
		strArr := strings.Split(bearerToken, " ")
		if len(strArr) == 2 {
			tokenString = strArr[1]
		}
		if tokenString == "" {
			return c.App().ErrorHandler(c, types.ErrUnauthorized)
		}

		// Get access token
		uuid, err := authService.GetUUIDByAccessToken(tokenString)
		if err != nil {
			return c.App().ErrorHandler(c, types.ErrUnauthorized)
		}

		// Validate in redis token
		userID, err := authService.FindUserIDByUUID(uuid)
		if err != nil {
			return c.App().ErrorHandler(c, types.ErrUnauthorized)
		}

		// Set token to context
		c.SetUserContext(context.WithValue(c.UserContext(), "access_token", tokenString))
		c.SetUserContext(context.WithValue(c.UserContext(), "user_id", userID))
		return c.Next()
	}
}
