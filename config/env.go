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
	Mode             string        `env:"GIN_MODE" envDefault:"debug"`
	Port             string        `env:"APP_PORT" envDefault:"8000"`
	AccessSecretKey  string        `env:"ACCESS_SECRET_KEY" envDefault:"secret"`
	RefreshSecretKey string        `env:"REFRESH_SECRET_KEY" envDefault:"refresh"`
	ReadTimeout      time.Duration `env:"READ_TIME_OUT" envDefault:"3s"`

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
	log.Println("Environment")
	log.Println("────────────────────────────────────")
	log.Println("[APP]")
	log.Println("GIN_MODE            = ", Env.Mode)
	log.Println("APP_PORT            = ", Env.Port)
	log.Println("────────────────────────────────────")
	log.Println("[DATABASE]")
	log.Println("DB_NAME             = ", Env.DBType)
	log.Println("DB_HOST             = ", Env.DBHost)
	log.Println("DB_PORT             = ", Env.DBPort)
	log.Println("DB_USER             = ", Env.DBUser)
	log.Println("DB_PASSWORD         = ", Env.DBPassword)
	log.Println("DB_OPTIONS          = ", Env.DBOptions)
	log.Println("────────────────────────────────────")
}
