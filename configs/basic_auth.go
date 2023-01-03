package config

import "github.com/gofiber/fiber/v2/middleware/basicauth"

func BasicAuth() basicauth.Config {
	return basicauth.Config{
		Users: map[string]string{
			"admin": "1234",
		},
	}
}
