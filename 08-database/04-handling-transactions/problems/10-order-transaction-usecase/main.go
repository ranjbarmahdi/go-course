/*
============================================================
51 — HANDLING TRANSACTIONS IN GO
Problem 10 — Order Transaction Use Case
============================================================

Create order + order item in one transaction.

Requirements:

    1. Tables (create if not exists):

           orders (id, user_id, total, status)
           order_items (id, order_id, product_id, quantity, unit_price)

    2. OrderRepository.Create(ctx, order)
       OrderItemRepository.Create(ctx, item)

    3. Both repos use getExecutor

    4. CreateOrderUseCase:

           return txManager.RunInTx(ctx, func(ctx) {
               orderRepo.Create(ctx, order)
               orderItemRepo.Create(ctx, item)
           })

    5. Use case injects txManager + both repos

    6. Print order ID on success

    7. DSN:

           postgres://postgres:secret@localhost:5432/postgres?sslmode=disable

============================================================
Goal
============================================================

Practice:

    - e-commerce style transaction
    - multiple repos in RunInTx

============================================================
*/

package main

func main() {}
