package config

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

// https://docs.gofiber.io/api/middleware/logger
func Logger() logger.Config {
	return logger.Config{
		Next:         nil,
		Format:       "${white} ${time} ${ip}:${port} > ${yellow}${path} │ ${white}${latency} │ ${cyan}${status} │ ${blue}${method} │ ${magenta}${route} ${red}${error}${reset}\n",
		TimeFormat:   "2006/01/02 15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       os.Stderr,
	}
}
