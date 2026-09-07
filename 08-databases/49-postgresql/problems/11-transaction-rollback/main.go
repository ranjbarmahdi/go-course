/*
============================================================
49 — POSTGRESQL
Problem 11 — Transaction Rollback
============================================================

Demonstrate rollback when a step fails.

Requirements:

    1. Start a transaction

    2. Insert a valid order

    3. Attempt insert order_item with INVALID product_id
       (UUID that does not exist in products)

    4. Foreign key violation must cause error

    5. Call tx.Rollback()

    6. Verify order was NOT saved:

           SELECT COUNT(*) FROM orders WHERE id = $1
           → must be 0

    7. Print:

           Transaction rolled back — order not saved

============================================================
Goal
============================================================

Practice:

    - rollback on failure
    - why transactions matter

============================================================
*/

package main

func main() {}
