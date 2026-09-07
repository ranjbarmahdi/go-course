/*
============================================================
06 — VALIDATION
Problem 2 — Common Validation Tags
============================================================

Create CreateProductRequest with these tags:

    Name     string  `validate:"required,min=2,max=100"`
    Price    float64 `validate:"required,gte=0"`
    Category string  `validate:"required,oneof=electronics clothing food"`

Requirements:

    1. Validate and print errors for:

           name too short, negative price, invalid category

    2. Validate and print "valid" for:

           Name: Phone, Price: 999, Category: electronics

    3. No HTTP server — CLI only

============================================================
Goal
============================================================

Practice:

    - min, max, gte, oneof tags

============================================================
*/

package main

func main() {}
