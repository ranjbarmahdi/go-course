/*
============================================================
50 — REPOSITORY PATTERN
Problem 2 — UserRepository Port
============================================================

Create the application port (interface).

Requirements:

    1. Define UserRepository interface with:

           Create(ctx context.Context, user User) error
           FindByEmail(ctx context.Context, email string) (User, error)
           FindByID(ctx context.Context, id string) (User, error)

    2. Every method must accept context.Context as first param

    3. Interface lives in application layer

    4. Do NOT implement PostgreSQL yet

    5. Add comment above interface:

           // LAYER: APPLICATION — PORT

============================================================
Goal
============================================================

Practice:

    - port interfaces
    - repository contract

============================================================
Layer: APPLICATION
============================================================
*/

package main

func main() {}
