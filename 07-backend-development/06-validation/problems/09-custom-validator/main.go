/*
============================================================
06 — VALIDATION
Problem 9 — Custom Validator
============================================================

Register a custom validator: has_digit

Requirements:

    1. Password must contain at least one digit (0-9)

    2. Register with:

           validate.RegisterValidation("has_digit", passwordHasDigit)

    3. SecureRegisterRequest:

           Email    validate:"required,email"
           Password validate:"required,min=8,has_digit"

    4. "secret123" → valid

    5. "secretonly" → invalid

    6. POST /register using SecureRegisterRequest

    7. Return 400 with field error if password has no digit

    8. Run on :8080

Test:

    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"me@example.com","password":"secretonly"}'

    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"me@example.com","password":"secret123"}'

============================================================
Goal
============================================================

Practice:

    - custom validation rules

============================================================
*/

package main

func main() {}
