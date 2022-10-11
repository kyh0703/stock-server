package config

import (
	"time"

	"github.com/caarlos0/env"
)

var Env Environment

type Environment struct {
	// App
	Mode         string        `env:"APP_MODE" envDefault:"debug"`
	Port         string        `env:"APP_PORT" envDefault:"8000"`
	APISecret    string        `env:"API_SECRET" envDefault:"98hbun98h"`
	APITokenLife time.Duration `env:"API_SECRET_LIFE" envDefault:"15m"`
	// Database
	DBType     string `env:"DB_NAME" envDefault:"mysql"`
	DBHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort     string `env:"DB_PORT" envDefault:"3306"`
	DBUser     string `env:"DB_USER" envDefault:"root"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"1234"`
	DBOptions  string `env:"DB_OPTION" envDefault:"1234"`
}

func init() {
	env.Parse(&Env)
}
