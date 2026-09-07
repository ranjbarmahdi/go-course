/*
============================================================
01 — CONFIGURATION
Problem 9 — Production Style LoadConfig
============================================================

LoadConfig that works WITHOUT .env (production style).

Requirements:

    1. LoadConfig must NOT require godotenv

    2. Read only from os.Getenv (simulates Docker/K8s)

    3. In main(), set env with os.Setenv before LoadConfig:

           DATABASE_URL
           JWT_SECRET (32+ chars)
           PORT

    4. Validate all required fields

    5. Print:

           Config loaded for production-style env

    6. Comment in code explaining:

           In production, Docker/K8s sets env — no .env file

============================================================
Goal
============================================================

Practice:

    - same LoadConfig for dev and prod
    - production mental model

============================================================
*/

package main

func main() {}
