package middleware

import (
	"github.com/gofiber/fiber/v2"
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
		// rc, _ := c.Keys["redis"].(*redis.Client)
		// // validate token
		// accessData, err := jwt.ExtractTokenMetadata(c)
		// if err != nil {
		// 	c.AbortWithError(http.StatusUnauthorized, err)
		// 	return
		// }
		// // validate in redis token
		// userID, err := jwt.GetUserIDFromRedis(rc, accessData.AccessUUID)
		// if err != nil {
		// 	c.AbortWithError(http.StatusUnauthorized, err)
		// 	return
		// }
		// c.Set("x-request-id", userID)
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
