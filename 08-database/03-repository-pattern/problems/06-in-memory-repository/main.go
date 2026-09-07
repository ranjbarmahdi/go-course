/*
============================================================
50 — REPOSITORY PATTERN
Problem 6 — In-Memory Mock Repository
============================================================

Create an in-memory UserRepository for testing.

Requirements:

    1. type InMemoryUserRepository struct {
           users map[string]User   // key: email
       }

    2. Implement UserRepository interface:

           Create
           FindByEmail
           FindByID

    3. FindByEmail returns sql.ErrNoRows when not found

    4. Create returns error if email already exists

    5. In main(), use InMemoryUserRepository with RegisterUserUseCase

    6. Do NOT use PostgreSQL in this problem

    7. Print success after register + find

============================================================
Goal
============================================================

Practice:

    - mock repository
    - why interfaces enable testing without DB

============================================================
Layer: INFRASTRUCTURE (test double)
============================================================
*/

package main

func main() {}
