/*
============================================================
49 — POSTGRESQL
Problem 3 — Foreign Key Orders
============================================================

Create orders table with foreign key to users.

Requirements:

    1. users table must exist (Problem 1)

    2. Create orders table:

           CREATE TABLE IF NOT EXISTS orders (
               id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
               user_id     UUID NOT NULL REFERENCES users(id),
               total       NUMERIC(10, 2) NOT NULL DEFAULT 0,
               status      VARCHAR(50) NOT NULL DEFAULT 'pending',
               created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
           )

    3. Insert a user if not exists:

           email: shop@example.com

    4. Insert an order for that user

    5. Print:

           Order created: <order_id>

    6. Insert with invalid user_id must fail (foreign key)

============================================================
Goal
============================================================

Practice:

    - foreign keys
    - REFERENCES
    - relational integrity

============================================================
*/

package main

func main() {}
