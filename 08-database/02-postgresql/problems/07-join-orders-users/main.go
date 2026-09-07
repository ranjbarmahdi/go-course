/*
============================================================
49 — POSTGRESQL
Problem 7 — JOIN Orders and Users
============================================================

Query orders with user email using INNER JOIN.

Requirements:

    1. Schema must exist with at least one order

    2. Query:

           SELECT o.id, o.total, o.status, u.email
           FROM orders o
           INNER JOIN users u ON u.id = o.user_id
           WHERE o.user_id = $1
           ORDER BY o.created_at DESC

    3. Create function:

           func listOrdersByUser(ctx context.Context, db *sql.DB, userID string) error

    4. Print each order:

           <order_id>  total=<total>  status=<status>  email=<email>

    5. Use QueryContext and rows.Next

============================================================
Goal
============================================================

Practice:

    - INNER JOIN
    - querying related tables

============================================================
*/

package main

func main() {}
