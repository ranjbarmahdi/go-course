/*
============================================================
48 — DATABASE/SQL
Problem 14 — Complete User Repository
============================================================

Build a complete UserRepository with all CRUD operations.

Requirements:

    1. Connect with openDB helper (ping included)

    2. Configure connection pool

    3. Create users table if not exists

    4. User struct:

           ID, Email, PasswordHash, CreatedAt

    5. UserRepository methods:

           Create(user User) error
           FindByEmail(email string) (User, error)
           List() ([]User, error)
           UpdateEmail(id, newEmail string) error
           DeleteByID(id string) error
           Count() (int, error)

    6. In main(), demonstrate full flow:

           a. Create user-99 / repo@example.com
           b. FindByEmail → print found
           c. List → print all users
           d. Count → print total
           e. UpdateEmail → repo.updated@example.com
           f. DeleteByID user-99
           g. FindByEmail → User not found

    7. Use pgx driver

    8. DSN:

           postgres://postgres:secret@localhost:5432/postgres?sslmode=disable

    9. Handle sql.ErrNoRows where appropriate

============================================================
Goal
============================================================

Practice:

    - complete database/sql CRUD
    - repository pattern
    - connecting all lecture concepts

============================================================
*/

package main

func main() {}
