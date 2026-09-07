/*
============================================================
06 — VALIDATION
Problem 7 — Slice Validation with dive
============================================================

Create CreateProductRequest:

    Tags []string `json:"tags" validate:"dive,min=1,max=20"`

Requirements:

    1. Validate product with empty tag in slice → fails

    2. Validate product with tag longer than 20 chars → fails

    3. Validate product with valid tags → passes

    4. CLI demo in main() — print errors or "valid"

    5. Optional: POST /products handler returning 400/201

============================================================
Goal
============================================================

Practice:

    - dive tag for slice elements

============================================================
*/

package main

func main() {}
