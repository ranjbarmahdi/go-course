/*
============================================================
01 — CONFIGURATION
Problem 3 — Required Environment Variables
============================================================

Load required configuration and fail if missing.

Requirements:

    1. Config must include:

           DatabaseURL string
           JWTSecret   string

    2. LoadConfig() returns error if:

           DATABASE_URL is empty → "DATABASE_URL is required"
           JWT_SECRET is empty   → "JWT_SECRET is required"

    3. Do NOT use godotenv in this problem

    4. Test by setting env vars in main before LoadConfig:

           os.Setenv("DATABASE_URL", "postgres://...")
           os.Setenv("JWT_SECRET", "secret-with-at-least-32-characters-long")

    5. Print "Config loaded" on success

============================================================
Goal
============================================================

Practice:

    - required vs optional config
    - fail fast

============================================================
*/

package main

func main() {}
