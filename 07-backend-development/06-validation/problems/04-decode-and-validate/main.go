/*
============================================================
06 — VALIDATION
Problem 4 — Decode JSON and Validate
============================================================

Create a function:

    func parseAndValidateRegister(body io.Reader) (RegisterRequest, error)

Requirements:

    1. RegisterRequest with validate tags (required, email, min=8)

    2. Decode JSON from body

    3. Run validate.Struct

    4. Return wrapped errors:

           invalid JSON
           validation failed

    5. In main(), test with strings.NewReader for:

           valid JSON
           invalid JSON
           valid JSON but bad email

    6. Print result or error for each test

    7. No HTTP server yet

============================================================
Goal
============================================================

Practice:

    - combining json.Decode + validator
    - reusable parse helper

============================================================
*/

package main

func main() {}
