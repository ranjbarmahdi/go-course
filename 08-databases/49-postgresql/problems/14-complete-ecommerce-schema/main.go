/*
============================================================
49 — POSTGRESQL
Problem 14 — Complete E-Commerce Schema
============================================================

Build and demo the full e-commerce schema.

Requirements:

    1. Connect + ping PostgreSQL

    2. Run all migrations:

           users, products, orders, order_items, indexes

    3. Seed data:

           user:    shop@example.com
           product: Go Book — 29.99

    4. Create order with item in a transaction

    5. listOrdersByUser — JOIN with users, print orders

    6. getOrderItems — JOIN with products, print items

    7. listOrdersPaginated — page 1, limit 10

    8. Use context.WithTimeout for all queries

    9. DSN:

           postgres://postgres:secret@localhost:5432/postgres?sslmode=disable

============================================================
TEST OUTPUT (example)
============================================================

    Connected to PostgreSQL
    Migration applied: ...
    Order created: ...
    Orders for user: ...
    Order items: ...
    Paginated: N result(s)

============================================================
Goal
============================================================

Practice:

    - complete PostgreSQL schema
    - migrations, indexes, joins, transactions
    - production-style database demo

============================================================
*/

package main

func main() {}
