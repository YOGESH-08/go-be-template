package utils

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type cfg struct {
	Env              string   `env:"ENV" envDefault:"development"`
	Port             string   `env:"PORT" envDefault:"8080"`
	JWTSecret        string   `env:"JWT_SECRET,notEmpty"`
	PostgresHost     string   `env:"POSTGRES_HOST,notEmpty"`
	PostgresPort     string   `env:"POSTGRES_PORT,notEmpty"`
	PostgresUser     string   `env:"POSTGRES_USER,notEmpty"`
	PostgresPassword string   `env:"POSTGRES_PASSWORD,notEmpty"`
	PostgresDB       string   `env:"POSTGRES_DB,notEmpty"`
	PostgresSSLMode  string   `env:"POSTGRES_SSLMODE" envDefault:"disable"`
	RedisHost        string   `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort        string   `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword    string   `env:"REDIS_PASSWORD"`
	FrontendURL      string   `env:"FRONTEND_URL" envDefault:"http://localhost:3000"`
	AdminURL         string   `env:"ADMIN_URL" envDefault:"http://localhost:3067"`
	CORSOrigins      []string // Computed field
}

var Config cfg

func LoadConfig() {
	// Load .env file if it exists; ignore error if it doesn't
	_ = godotenv.Load()

	if err := env.Parse(&Config); err != nil {
		log.Fatalf("Error parsing configuration: %v", err)
	}

	// Compute allowed CORS origins
	if Config.FrontendURL != "" {
		Config.CORSOrigins = append(Config.CORSOrigins, Config.FrontendURL)
	}
	if Config.AdminURL != "" {
		Config.CORSOrigins = append(Config.CORSOrigins, Config.AdminURL)
	}
}
