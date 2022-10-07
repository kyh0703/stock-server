package config

import "github.com/caarlos0/env"

var Env Environment

type Environment struct {
	// App
	Port      string `env:"APP_PORT" envDefault:"8000"`
	APISecret string `env:"API_SECRET" envDefault:"98hbun98h"`
	// Database
	DBType     string `env:"DB_NAME" envDefault:"mysql"`
	DBHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort     string `env:"DB_PORT" envDefault:"3306"`
	DBUser     string `env:"DB_USER" envDefault:"root"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"1234"`
}

func init() {
	env.Parse(&Env)
}
