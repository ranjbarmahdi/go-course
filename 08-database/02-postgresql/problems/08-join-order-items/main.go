/*
============================================================
49 — POSTGRESQL
Problem 8 — JOIN Order Items and Products
============================================================

Get order details with product names.

Requirements:

    1. order_items and products must exist with data

    2. Query:

           SELECT p.name, oi.quantity, oi.unit_price
           FROM order_items oi
           INNER JOIN products p ON p.id = oi.product_id
           WHERE oi.order_id = $1

    3. Create function:

           func getOrderItems(ctx context.Context, db *sql.DB, orderID string) error

    4. Print each line:

           <product_name>  qty=<quantity>  price=<unit_price>

============================================================
Goal
============================================================

Practice:

    - JOIN across three-table schema
    - order detail queries

============================================================
*/

package main

func main() {}
