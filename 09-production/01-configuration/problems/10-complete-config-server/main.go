/*
============================================================
01 — CONFIGURATION
Problem 10 — Complete Configurable Server
============================================================

Build a complete configurable API server.

Requirements:

    1. Config fields:

           ENV, PORT, DATABASE_URL, JWT_SECRET

    2. loadDotEnv() — try .env paths, ignore if missing

    3. LoadConfig + Validate

    4. Connect PostgreSQL with DATABASE_URL

    5. Routes:

           GET /health → env, port, db_connected (no secrets)

    6. Log startup info (never log JWT_SECRET or full DATABASE_URL)

    7. Fail fast on missing/invalid config

    8. Run from repo root with .env

Test:

    go run ./09-production/01-configuration/problems/10-complete-config-server/
    curl http://localhost:8080/health

============================================================
Goal
============================================================

Practice:

    - complete configuration flow
    - local .env + production-ready LoadConfig

============================================================
*/

package main

func main() {}
