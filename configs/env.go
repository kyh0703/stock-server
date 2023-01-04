package configs

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
	DBType string `env:"DB_TYPE" envDefault:"mysql"`
	DBName string `env:"DB_NAME" envDefault:"stock"`
	DBUrl  string `env:"DB_URL" envDefault:"root:1234@tcp(localhost:3306)/stock?parseTime=true"`

	// Cache
	CacheUrl string `env:"CACHE_URL" envDefault:""`
}

func init() {
	if err := godotenv.Load("./configs/env/.development.env"); err != nil {
		log.Println("dose not exist .env file")
	} else {
		log.Println("load .env file complete")
	}
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
	log.Println("DB_TYPE             = ", Env.DBType)
	log.Println("DB_NAME             = ", Env.DBName)
	log.Println("DB_URL              = ", Env.DBUrl)
	log.Println("────────────────────────────────────")
	log.Println("[CACHE]")
	log.Println("CACHE_URL           = ", Env.CacheUrl)
	log.Println("────────────────────────────────────")
}
