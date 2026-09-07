/*
============================================================
01 — CONFIGURATION
Problem 1 — getEnv Helper
============================================================

Create a getEnv helper function.

Requirements:

    1. func getEnv(key, fallback string) string

    2. If os.Getenv(key) is not empty → return it

    3. If empty → return fallback

    4. In main(), demonstrate:

           getEnv("PORT", "8080")        → 8080 (if PORT unset)
           getEnv("ENV", "development")  → development

    5. Set PORT=9090 in environment and show it returns 9090

    6. Use only standard library (os)

============================================================
Goal
============================================================

Practice:

    - environment variables with defaults

============================================================
*/

package main

func main() {}
