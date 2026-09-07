/*
============================================================
06 — VALIDATION
Problem 1 — Basic validate.Struct
============================================================

Use github.com/go-playground/validator/v10.

Requirements:

    1. Create validator instance with validator.New()

    2. Create RegisterRequest struct:

           Email    string `validate:"required,email"`
           Password string `validate:"required,min=8"`

    3. In main(), validate two requests:

           bad  → email "bad", password "123"
           good → email "me@example.com", password "secret123"

    4. Print "invalid" or "valid" for each

    5. Pass pointer to validate.Struct(&req)

============================================================
Goal
============================================================

Practice:

    - validator setup
    - struct tags
    - validate.Struct

============================================================
*/

package main

func main() {}
