/*
============================================================
06 — VALIDATION
Problem 3 — Format Validation Errors
============================================================

Create formatValidationError(err error) map[string]string

Requirements:

    1. Use errors.As to detect validator.ValidationErrors

    2. Map field name → human-readable message

    3. Handle at least these tags in messages:

           required, email, min

    4. Demo in main():

           invalid RegisterRequest → print JSON-like map output

    5. Example output shape:

           map[Email:Email must be a valid email Password:Password must be at least 8 characters]

============================================================
Goal
============================================================

Practice:

    - reading ValidationErrors
    - user-friendly error messages

============================================================
*/

package main

func main() {}
