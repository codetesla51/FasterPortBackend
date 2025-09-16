package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() *pgxpool.Pool {
	// Reuse existing pool if available
	if DB != nil {
		return DB
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set in environment")
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}

	// Ping to verify connection
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Cannot ping DB: %v", err)
	}

	DB = pool
	return DB
}