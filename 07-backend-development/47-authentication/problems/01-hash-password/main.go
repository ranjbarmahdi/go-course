/*
============================================================
47 — AUTHENTICATION
Problem 1 — Hash Password
============================================================

Create two functions:

    func hashPassword(password string) (string, error)
    func verifyPassword(password, hashedPassword string) error

Requirements:

    1. Use golang.org/x/crypto/bcrypt

    2. hashPassword must use bcrypt.DefaultCost

    3. verifyPassword must return nil if password matches

    4. verifyPassword must return an error if password is wrong

    5. In main(), demonstrate:

           hash "secret123"
           verify correct password → success
           verify wrong password   → error

    6. Print results to console

Allowed packages:

    fmt
    golang.org/x/crypto/bcrypt

============================================================
Goal
============================================================

Practice:

    - bcrypt.GenerateFromPassword
    - bcrypt.CompareHashAndPassword
    - never store plain passwords

============================================================
*/

package main

func main() {}
