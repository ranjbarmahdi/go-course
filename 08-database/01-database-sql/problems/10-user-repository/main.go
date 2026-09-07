/*
============================================================
48 — DATABASE/SQL
Problem 10 — User Struct and Repository
============================================================

Create a User struct and basic repository functions.

Requirements:

    1. Define struct:

           type User struct {
               ID           string
               Email        string
               PasswordHash string
               CreatedAt    time.Time
           }

    2. Create UserRepository:

           type UserRepository struct {
               db *sql.DB
           }

           func NewUserRepository(db *sql.DB) *UserRepository

    3. Implement:

           func (r *UserRepository) Create(user User) error
           func (r *UserRepository) FindByEmail(email string) (User, error)

    4. Create must INSERT into users table

    5. FindByEmail must return sql.ErrNoRows when not found

    6. In main():

           create user-2 / mahdi2@example.com
           find by email and print id + email

    7. Ensure users table exists before insert

============================================================
Goal
============================================================

Practice:

    - struct mapping
    - simple repository pattern
    - separating DB logic from main

============================================================
*/

package main

func main() {}
