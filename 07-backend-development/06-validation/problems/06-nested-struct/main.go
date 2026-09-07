/*
============================================================
06 — VALIDATION
Problem 6 — Nested Struct Validation
============================================================

Create:

    POST /users

Request body:

    {
        "name": "Mahdi",
        "email": "mahdi@example.com",
        "address": {
            "city": "Tehran",
            "country": "IR"
        }
    }

Requirements:

    1. CreateUserRequest with nested Address struct

    2. Address tags:

           city    → required
           country → required, len=2

    3. Validate nested fields automatically

    4. Return 400 with field errors if address.city or country invalid

    5. Return 201 on success

    6. Run on :8080

Test:

    curl -X POST http://localhost:8080/users \
      -H "Content-Type: application/json" \
      -d '{"name":"Mahdi","email":"mahdi@example.com","address":{"city":"Tehran","country":"IR"}}'

============================================================
Goal
============================================================

Practice:

    - nested struct validation

============================================================
*/

package main

func main() {}
