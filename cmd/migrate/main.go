package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/CodeChefVIT/go-backend-template/internal/utils"
	_ "github.com/jackc/pgx/v5/stdlib" // SQL driver wrapper for pgx/v5
	"github.com/pressly/goose/v3"
)

const migrationsDir = "database/schema"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/migrate/main.go [up|down|status]")
	}

	command := os.Args[1]

	// Load environment configuration
	utils.LoadConfig()

	// Parse database DSN from config
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		utils.Config.PostgresUser,
		utils.Config.PostgresPassword,
		utils.Config.PostgresHost,
		utils.Config.PostgresPort,
		utils.Config.PostgresDB,
	)

	// Open connection using pgx stdlib compatibility layer
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Set dialect to postgres
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	log.Printf("Executing migration command: %s", command)

	// Execute migrations
	switch command {
	case "up":
		if err := goose.Up(db, migrationsDir); err != nil {
			log.Fatalf("Goose migration UP failed: %v", err)
		}
	case "down":
		if err := goose.Down(db, migrationsDir); err != nil {
			log.Fatalf("Goose migration DOWN failed: %v", err)
		}
	case "status":
		if err := goose.Status(db, migrationsDir); err != nil {
			log.Fatalf("Goose migration STATUS failed: %v", err)
		}
	default:
		log.Fatalf("Unknown migration command: %s. Use 'up', 'down', or 'status'.", command)
	}

	log.Println("Migration command executed successfully")
}
