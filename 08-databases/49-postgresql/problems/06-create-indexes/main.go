/*
============================================================
49 — POSTGRESQL
Problem 6 — Create Indexes
============================================================

Add indexes for common query patterns.

Requirements:

    1. orders and order_items tables must exist

    2. Create indexes:

           CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);
           CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);
           CREATE INDEX IF NOT EXISTS idx_orders_created_at ON orders(created_at);
           CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items(order_id);

    3. Print:

           Indexes created

    4. Verify in psql (optional):

           \di

============================================================
Goal
============================================================

Practice:

    - when and why to add indexes
    - CREATE INDEX IF NOT EXISTS

============================================================
*/

package main

func main() {}
