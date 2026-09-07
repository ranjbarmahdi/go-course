/*
============================================================
06 — VALIDATION
Problem 8 — omitempty for Updates
============================================================

Create UpdateProfileRequest:

    Age int    `json:"age" validate:"omitempty,gte=0,lte=120"`
    Bio string `json:"bio" validate:"omitempty,max=500"`

Requirements:

    1. Empty struct {} → valid (nothing to update)

    2. Age: 25, Bio omitted → valid

    3. Age: -1 → invalid

    4. Bio longer than 500 chars → invalid

    5. POST /profile/update handler (no auth needed):

           400 on validation error
           200 { "message": "profile updated" } on success

    6. Run on :8080

Test:

    curl -X POST http://localhost:8080/profile/update \
      -H "Content-Type: application/json" \
      -d '{"age":25}'

============================================================
Goal
============================================================

Practice:

    - omitempty for partial updates

============================================================
*/

package main

func main() {}
