/*
============================================================
50 — REPOSITORY PATTERN
Problem 9 — FindByID
============================================================

Add FindByID to the repository and use case.

Requirements:

    1. PostgresUserRepository.FindByID(ctx, id) (User, error)

    2. GetUserByIDUseCase:

           Execute(ctx, id) (User, error)

    3. sql.ErrNoRows → ErrUserNotFound

    4. In main():

           register user
           find by id
           print user email

============================================================
Goal
============================================================

Practice:

    - extending repository interface
    - new use case with same pattern

============================================================
*/

package main

func main() {}
