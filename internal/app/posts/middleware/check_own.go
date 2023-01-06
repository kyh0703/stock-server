package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyh0703/stock-server/internal/types"
)

func CheckOwnPost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, ok := c.Context().UserValue(types.ContextKeyUserID).(int)
		if !ok {
			return c.App().ErrorHandler(c, types.ErrUnauthorized)
		}

		return c.Next()
	}
}
