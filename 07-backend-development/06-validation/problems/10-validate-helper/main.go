/*
============================================================
06 — VALIDATION
Problem 10 — Reusable Validate Helper
============================================================

Create a reusable validation helper used by multiple handlers.

Requirements:

    1. Package-level validator instance (single validator.New())

    2. func validateRequest(w http.ResponseWriter, req any) bool

           runs validate.Struct
           writes 400 JSON if invalid
           returns false on failure, true on success

    3. Handlers:

           POST /register   → RegisterRequest
           POST /products   → CreateProductRequest

    4. Both handlers use validateRequest helper

    5. Return 201 on success for each

    6. Run on :8080

============================================================
Goal
============================================================

Practice:

    - DRY validation across handlers
    - shared validator instance

============================================================
*/

package main

func main() {}
