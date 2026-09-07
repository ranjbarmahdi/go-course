/*
============================================================
01 — CONFIGURATION
Problem 6 — os.LookupEnv
============================================================

Handle optional configuration with LookupEnv.

Requirements:

    1. Create function:

           func redisConfig() (url string, enabled bool)

    2. Use os.LookupEnv("REDIS_URL")

    3. If not set:

           enabled = false

    4. If set (even empty string):

           enabled = true, url = value

    5. In main(), demonstrate both cases with os.Setenv / Unsetenv

    6. Print:

           Redis enabled: true/false

============================================================
Goal
============================================================

Practice:

    - optional services
    - LookupEnv vs Getenv

============================================================
*/

package main

func main() {}
