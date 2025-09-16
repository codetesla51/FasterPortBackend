package config

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "github.com/jackc/pgx/v5"     
    "github.com/jackc/pgx/v5/pgxpool"
)
var DB *pgxpool.Pool

func ConnectDB() {
    if err := godotenv.Load(); err != nil {
        log.Println("Warning: .env file not found, using environment variables")
    }

    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL not set in environment")
    }

    fmt.Println("Trying to connect to DB...")

    cfg, err := pgxpool.ParseConfig(dbURL)
    if err != nil {
        log.Fatalf("Failed to parse DB config: %v", err)
    }

    cfg.MaxConns = 10
    cfg.MinConns = 2
    cfg.MaxConnLifetime = time.Hour
cfg.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe

    pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
    if err != nil {
        log.Fatalf("Unable to connect to DB: %v", err)
    }

    if err := pool.Ping(context.Background()); err != nil {
        log.Fatalf("Cannot ping DB: %v", err)
    }

    DB = pool
    fmt.Println("Connected to Postgres!")
}