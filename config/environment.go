package config

import "github.com/caarlos0/env"

var Config Environment

type Environment struct {
	Port             string `env:"HTTP_PORT" envDefault:"8000"`
	DatabaseType     string `env:"DATABASE_TYPE" envDefault:"sqllite3"`
	DatabaseHost     string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort     string `env:"DATABASE_PORT" envDefault:"3060"`
	DatabaseUser     string `env:"DATABASE_USER" envDefault:"root"`
	DatabasePassword string `env:"DATABASE_PASSWORD" envDefault:"1234"`
}

func init() {
	env.Parse(&Config)
}
