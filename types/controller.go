package types

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Path() string
	Routes(fiber.Router)
}
