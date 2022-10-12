package config

import (
	"log"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

var Env Environment

type Environment struct {
	// App
	Mode              string        `env:"APP_MODE" envDefault:"debug"`
	Port              string        `env:"APP_PORT" envDefault:"8000"`
	APISecret         string        `env:"API_SECRET" envDefault:"secret"`
	APISecretLifeTime time.Duration `env:"API_SECRET_LIFETIME" envDefault:"1h"`
	// Database
	DBType     string `env:"DB_NAME" envDefault:"mysql"`
	DBHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort     string `env:"DB_PORT" envDefault:"3306"`
	DBUser     string `env:"DB_USER" envDefault:"root"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"1234"`
	DBOptions  string `env:"DB_OPTIONS" envDefault:"1234"`
}

func init() {
	godotenv.Load(".env")
	env.Parse(&Env)
	log.Fatal(Env)
}
