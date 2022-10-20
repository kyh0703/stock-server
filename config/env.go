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
	Port             string        `env:"APP_PORT" envDefault:"8000"`
	AccessSecretKey  string        `env:"ACCESS_SECRET_KEY" envDefault:"secret"`
	RefreshSecretKey string        `env:"REFRESH_SECRET_KEY" envDefault:"refresh"`
	ReadTimeout      time.Duration `env:"READ_TIME_OUT" envDefault:"3s"`
	// Database
	DBType string `env:"DB_NAME" envDefault:"mysql"`
	DBUrl  string `env:"DB_URL" envDefault:"root:1234@tcp(localhost:3306)/stock?parseTime=true"`
}

func init() {
	godotenv.Load("./env/.development.env")
	env.Parse(&Env)
	PrintEnvironment()
}

func PrintEnvironment() {
	log.Println("────────────────────────────────────")
	log.Println("Environment")
	log.Println("────────────────────────────────────")
	log.Println("[APP]")
	log.Println("APP_PORT            = ", Env.Port)
	log.Println("────────────────────────────────────")
	log.Println("[DATABASE]")
	log.Println("DB_NAME             = ", Env.DBType)
	log.Println("DB_URL              = ", Env.DBUrl)
	log.Println("────────────────────────────────────")
}
