package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/routes/auth"
)

func SetUserContext() fiber.Handler {
	// user context is not null
	return func(c *fiber.Ctx) error {
		c.SetUserContext(c.Context())
		return c.Next()
	}
}

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
		var (
			tokenString string
			authService auth.AuthService
		)

		// Get token string 'Header : Authorization: Bearer ${accessToken}'
		bearerToken := c.Get("Authorization")
		if bearerToken == "" {
			return fiber.NewError(fiber.StatusBadRequest, "token is invalid")
		}
		strArr := strings.Split(bearerToken, " ")
		if len(strArr) == 2 {
			tokenString = strArr[1]
		}
		if tokenString == "" {
			return fiber.NewError(fiber.StatusBadRequest, "token is invalid")
		}

		// Get access token
		uuid, err := authService.GetUUIDByAccessToken(tokenString)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "token is invalid")
		}

		// Validate in redis token
		userID, err := authService.FindUserIDByUUID(uuid)
		if err != nil {
			return c.Status(http.StatusUnauthorized).SendString(err.Error())
		}

		// Set token to context
		c.SetUserContext(context.WithValue(c.UserContext(), "access_token", tokenString))
		c.SetUserContext(context.WithValue(c.UserContext(), "user_id", userID))
		return c.Next()
	}
}
