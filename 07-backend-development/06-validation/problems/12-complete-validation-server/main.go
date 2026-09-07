/*
============================================================
06 — VALIDATION
Problem 12 — Complete Validation Server
============================================================

Build a complete API with validation on all write endpoints.

Routes:

    GET  /health
    POST /register
    POST /products
    POST /users

Requirements:

    1. RegisterRequest — required, email, min=8 password

    2. CreateProductRequest — name, price, category, tags (dive)

    3. CreateUserRequest — name, email, nested address

    4. Shared validator instance + formatValidationError helper

    5. All POST handlers:

           invalid JSON     → 400
           validation error → 400 with fields
           success          → 201

    6. GET /health → 200 ok

    7. Run on :8080

Test:

    curl http://localhost:8080/health

    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"me@example.com","password":"secret123"}'

    curl -X POST http://localhost:8080/products \
      -H "Content-Type: application/json" \
      -d '{"name":"Phone","price":999,"category":"electronics","tags":["new","sale"]}'

    curl -X POST http://localhost:8080/users \
      -H "Content-Type: application/json" \
      -d '{"name":"Mahdi","email":"mahdi@example.com","address":{"city":"Tehran","country":"IR"}}'

============================================================
Goal
============================================================

Practice:

    - complete validation architecture
    - multiple DTOs and endpoints together

============================================================
*/

package main

func main() {}
