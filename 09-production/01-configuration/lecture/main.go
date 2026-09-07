package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

/*
============================================================
01 — CONFIGURATION
============================================================

Topics
------------------------------------------------------------
 1. Why Configuration
 2. Hardcoded vs Environment
 3. os.Getenv and getEnv
 4. Config Struct
 5. LoadConfig
 6. Validate and Fail Fast
 7. .env for Local Development
 8. Production Config (Docker/K8s/CI)
 9. .env.example
10. Wire Config to Server
11. Complete Demo
12. Configuration Mental Model
============================================================
*/

// ============================================================
// 1. WHY CONFIGURATION
// ============================================================

/*
Your previous lectures hardcoded values:

    dsn := "postgres://postgres:secret@localhost:5432/..."
    jwtSecret := "super-secret-key"
    http.ListenAndServe(":8080", mux)


Problems:

    Different values per environment (dev/staging/prod)
    Secrets must not live in source code
    Same binary must run everywhere with different config


Solution:

    Read configuration from environment variables at startup.


12-factor app rule:

    Store config in the environment.
*/

// ============================================================
// 2. HARDCODED VS ENVIRONMENT
// ============================================================

/*
Bad (hardcoded):

    port := "8080"
    dbURL := "postgres://..."


Good (environment):

    port := os.Getenv("PORT")
    dbURL := os.Getenv("DATABASE_URL")


Who sets env vars?

    Local dev     → .env file (godotenv) or shell
    Docker        → docker run -e KEY=value
    Kubernetes    → env / secrets in manifest
    CI/CD         → pipeline environment variables


Go reads the same way everywhere:

    os.Getenv("DATABASE_URL")
*/

// ============================================================
// 3. CONFIG STRUCT
// ============================================================

/*
Centralize all settings in one struct.

Load once at startup.
Pass cfg to the rest of the app.


Do NOT call os.Getenv scattered across handlers.
*/

type Config struct {
	Env         string
	Port        string
	DatabaseURL string
	JWTSecret   string
}

// ============================================================
// 4. getEnv HELPER
// ============================================================

/*
getEnv returns env value or fallback default.

Used for optional settings with safe defaults:

    PORT default 8080
    ENV  default development


Required secrets (DATABASE_URL, JWT_SECRET) should NOT
use silent defaults — fail if missing.
*/

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value != "" {
		return value
	}

	return fallback
}

// ============================================================
// 5. LOAD .env (LOCAL DEV ONLY)
// ============================================================

/*
Go does NOT read .env files automatically.

godotenv.Load() reads .env into os.Environ().

In production there is usually NO .env file.
Docker/Kubernetes inject env vars directly.
If .env is missing, that is normal — ignore the error.


Try common paths so `go run` works from repo root or lecture folder.
*/

func loadDotEnv() {
	paths := []string{
		".env",
		"../../.env",
		"../../../.env",
	}

	for _, path := range paths {
		err := godotenv.Load(path)
		if err == nil {
			log.Println("Loaded env file:", path)
			return
		}
	}

	log.Println("No .env file — using system environment")
}

// ============================================================
// 6. LOAD CONFIG
// ============================================================

func LoadConfig() (Config, error) {
	loadDotEnv()

	cfg := Config{
		Env:         getEnv("ENV", "development"),
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}

	err := cfg.Validate()
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

// ============================================================
// 7. VALIDATE — FAIL FAST
// ============================================================

/*
Validate required config at startup.

If DATABASE_URL is missing → log.Fatal
Better to crash immediately than run broken.


Never log secrets (JWT_SECRET, passwords).
*/

func (c Config) Validate() error {
	if c.DatabaseURL == "" {
		return errors.New("DATABASE_URL is required")
	}

	if c.JWTSecret == "" {
		return errors.New("JWT_SECRET is required")
	}

	if len(c.JWTSecret) < 32 {
		return errors.New("JWT_SECRET must be at least 32 characters")
	}

	if c.Port == "" {
		return errors.New("PORT is required")
	}

	return nil
}

// ============================================================
// 8. os.LookupEnv
// ============================================================

/*
LookupEnv tells you if variable EXISTS.

    value, ok := os.LookupEnv("REDIS_URL")

    ok == false  → not configured
    ok == true && value == ""  → configured but empty


Use for optional services (Redis, RabbitMQ).
*/

func optionalRedisURL() (string, bool) {
	return os.LookupEnv("REDIS_URL")
}

// ============================================================
// 9. WIRE CONFIG TO SERVER
// ============================================================

func createRouter(cfg Config, db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "env=%s port=%s db_connected=%t\n", cfg.Env, cfg.Port, db != nil)
	})

	return mux
}

func openDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return db, nil
}

// ============================================================
// 10. COMPLETE DEMO
// ============================================================

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting server")
	log.Println("  ENV: ", cfg.Env)
	log.Println("  PORT:", cfg.Port)
	log.Println("  DB:  configured (URL hidden)")

	db, err := openDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("  DB:  connected")

	if redisURL, ok := optionalRedisURL(); ok {
		log.Println("  REDIS_URL: configured (", len(redisURL), "chars)")
	} else {
		log.Println("  REDIS_URL: not set (optional)")
	}

	mux := createRouter(cfg, db)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	log.Println("Server running on http://localhost:" + cfg.Port)
	log.Println("Try: curl http://localhost:" + cfg.Port + "/health")

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

/*
============================================================
PRODUCTION — NO .env FILE
============================================================

Local:

    .env in repo root
    godotenv.Load() for convenience


Production:

    No .env in Docker image
    Kubernetes Secret → env var
    Same LoadConfig() code


Example Docker:

    docker run \
      -e PORT=8080 \
      -e DATABASE_URL=postgres://... \
      -e JWT_SECRET=... \
      myapp


LoadConfig works the same — only the source changes.


============================================================
.env.example
============================================================

Commit .env.example to git (placeholders only).
Add .env to .gitignore (real secrets).

See repo root: .env.example


============================================================
CONFIGURATION MENTAL MODEL
============================================================

Question 1:

    Where do secrets live?

Answer:

    Environment variables — never in git


Question 2:

    Does Go read .env automatically?

Answer:

    No — use godotenv locally only


Question 3:

    What if DATABASE_URL missing?

Answer:

    Fail fast at startup


Question 4:

    Same code dev and prod?

Answer:

    Yes — LoadConfig everywhere


Question 5:

    Defaults for PORT?

Answer:

    Yes — getEnv("PORT", "8080")


============================================================
RUN
============================================================

From repo root:

    go run ./09-production/01-configuration/lecture/

Requires .env or system env with:

    DATABASE_URL
    JWT_SECRET (min 32 chars)
    PORT (optional)


============================================================
END OF 01 — CONFIGURATION
============================================================
*/
