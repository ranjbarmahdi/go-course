/*
============================================================
49 — POSTGRESQL
Problem 10 — Transaction Create Order
============================================================

Create an order and order_item in a single transaction.

Requirements:

    1. Create function:

           func createOrderWithItem(
               ctx context.Context,
               db *sql.DB,
               userID, productID string,
               quantity int,
               unitPrice string,
           ) (orderID string, err error)

    2. Use db.BeginTx(ctx, nil)

    3. Steps inside transaction:

           a. INSERT INTO orders ... RETURNING id
           b. INSERT INTO order_items ...

    4. On any error → Rollback

    5. On success → Commit

    6. defer tx.Rollback() (safe after Commit)

    7. Print:

           Order created in transaction: <order_id>

============================================================
Goal
============================================================

Practice:

    - BeginTx / Commit / Rollback
    - atomic multi-table writes

============================================================
*/

package main

func main() {}
