/*
============================================================
01 — CONFIGURATION
Problem 4 — Validate Config
============================================================

Add Validate() method to Config.

Requirements:

    1. func (c Config) Validate() error

    2. Rules:

           DATABASE_URL required
           JWT_SECRET required
           JWT_SECRET min length 32

    3. Return clear error messages

    4. LoadConfig must call Validate()

    5. Test cases in main (using os.Setenv):

           valid config   → success
           short JWT      → error
           missing DB URL → error

============================================================
Goal
============================================================

Practice:

    - config validation
    - startup safety

============================================================
*/

package main

func main() {}
