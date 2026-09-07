/*
============================================================
49 — POSTGRESQL
Problem 4 — Order Items Table
============================================================

Create order_items table linking orders and products.

Requirements:

    1. orders and products tables must exist

    2. Create table:

           CREATE TABLE IF NOT EXISTS order_items (
               id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
               order_id    UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
               product_id  UUID NOT NULL REFERENCES products(id),
               quantity    INTEGER NOT NULL CHECK (quantity > 0),
               unit_price  NUMERIC(10, 2) NOT NULL CHECK (unit_price > 0)
           )

    3. Seed one user, one product, one order if needed

    4. Insert one order_item

    5. Print:

           Order item created

    6. ON DELETE CASCADE:

       deleting order should delete its order_items
       (verify manually in psql optional)

============================================================
Goal
============================================================

Practice:

    - many-to-many via junction table
    - ON DELETE CASCADE

============================================================
*/

package main

func main() {}
