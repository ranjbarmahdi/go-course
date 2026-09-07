/*
============================================================
01 — CONFIGURATION
Problem 8 — Config + Database
============================================================

Connect to PostgreSQL using DATABASE_URL from config.

Requirements:

    1. LoadConfig (DATABASE_URL, JWT_SECRET, PORT)

    2. openDB(cfg.DatabaseURL) with Ping

    3. GET /health returns:

           db=connected

       if ping succeeds

    4. Fail startup if DATABASE_URL invalid

    5. Use pgx driver

    6. Docker PostgreSQL must be running

    7. Use .env or os.Setenv

============================================================
Goal
============================================================

Practice:

    - wiring database URL from config

============================================================
*/

package main

func main() {}
