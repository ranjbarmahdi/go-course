/*
============================================================
49 — POSTGRESQL
Problem 9 — Pagination in SQL
============================================================

List orders with LIMIT and OFFSET pagination.

Requirements:

    1. Create function:

           func listOrdersPaginated(
               ctx context.Context,
               db *sql.DB,
               userID string,
               page, limit int,
           ) error

    2. Calculate offset:

           offset = (page - 1) * limit

    3. Query:

           SELECT o.id, o.total, o.status
           FROM orders o
           WHERE o.user_id = $1
           ORDER BY o.created_at DESC
           LIMIT $2 OFFSET $3

    4. Defaults:

           page  < 1  → page = 1
           limit < 1  → limit = 10

    5. Print results for page 1, limit 10

    6. Print count of returned rows

============================================================
Goal
============================================================

Practice:

    - LIMIT OFFSET pagination
    - same formula as REST API

============================================================
*/

package main

func main() {}
