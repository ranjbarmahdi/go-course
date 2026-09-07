/*
============================================================
51 — HANDLING TRANSACTIONS IN GO
Problem 6 — CreateUserWithProfile Use Case
============================================================

Create use case that runs two repos in one transaction.

Requirements:

    1. CreateUserWithProfileUseCase with:

           txManager TransactionManager
           userRepo  UserRepository
           profileRepo ProfileRepository

    2. Execute(ctx, email, password, name) error

    3. Inside RunInTx:

           check email not exists
           userRepo.Create(ctx, user)
           profileRepo.Create(ctx, profile)

    4. Use case must NOT import database/sql for BeginTx

    5. Use case MAY use errors.Is(err, sql.ErrNoRows) OR map in repo

    6. Wire in main and create user + profile

    7. Print success

============================================================
Goal
============================================================

Practice:

    - RunInTx in use case
    - injected repos pattern you wanted

============================================================
Layer: APPLICATION
============================================================
*/

package main

func main() {}
