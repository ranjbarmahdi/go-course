package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

/*
============================================================
49 — POSTGRESQL
============================================================

Topics
------------------------------------------------------------
 1. Schema Design
 2. PostgreSQL Data Types
 3. Constraints
 4. Relationships and Foreign Keys
 5. Migrations
 6. Indexes
 7. JOIN Queries
 8. Pagination in SQL
 9. Transactions
10. Complete PostgreSQL Demo
11. PostgreSQL Mental Model
============================================================
*/

// ============================================================
// 1. SCHEMA DESIGN
// ============================================================

/*
Schema = structure of your database.

Tables, columns, types, constraints, relationships.


In 48-database-sql you created a simple users table from Go.

In production you design the schema carefully FIRST,
then apply it with migrations.


Example e-commerce schema:

    users
    products
    orders
    order_items


Design rules:

    Use meaningful table names (plural: users, orders)
    Every table should have a primary key
    Use proper PostgreSQL types (UUID, TIMESTAMPTZ, NUMERIC)
    Add created_at on most tables
    Use foreign keys for relationships
*/

const dsn = "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable"

// ============================================================
// 2. POSTGRESQL DATA TYPES
// ============================================================

/*
Common types in production:

    UUID           → primary keys (gen_random_uuid())
    VARCHAR(n)     → short text with limit (email, name)
    TEXT           → long text (password hash, description)
    BOOLEAN        → true/false flags
    INTEGER        → counts
    BIGINT         → large numbers
    NUMERIC(p,s)   → money (exact decimal, no float errors)
    TIMESTAMPTZ    → timestamps with timezone
    JSONB          → flexible JSON metadata


Money example:

    NUMERIC(10, 2)   → up to 99999999.99


Never use FLOAT for money — rounding errors.
*/

// ============================================================
// 3. CONSTRAINTS
// ============================================================

/*
Constraints enforce data rules at database level.


PRIMARY KEY
    Uniquely identifies each row

NOT NULL
    Column must have a value

UNIQUE
    No duplicate values (email)

DEFAULT
    Value if not provided (NOW(), 'pending')

CHECK
    Custom rule (price > 0)

REFERENCES
    Foreign key to another table
*/

// ============================================================
// 4. RELATIONSHIPS AND FOREIGN KEYS
// ============================================================

/*
One user has many orders.

    users (1) ──────< orders (many)


orders.user_id REFERENCES users(id)


One order has many order items.

    orders (1) ──────< order_items (many)


order_items.order_id   REFERENCES orders(id)
order_items.product_id REFERENCES products(id)


Foreign keys protect data integrity:

    Cannot create order for non-existent user
    Cannot delete user if orders exist (unless CASCADE)
*/

// ============================================================
// 5. MIGRATIONS
// ============================================================

/*
Do NOT run CREATE TABLE from application code in production.

Use versioned migration files:


    migrations/
        001_create_users.up.sql
        001_create_users.down.sql
        002_create_products.up.sql
        002_create_products.down.sql
        ...


    up   → apply change
    down → rollback change


Popular tools:

    golang-migrate
    goose


This lecture runs migrations from Go for learning.
In production use a migration tool in CI/CD.
*/

const migration001Up = `
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS users (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email         VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
`

const migration002Up = `
CREATE TABLE IF NOT EXISTS products (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(255) NOT NULL,
    price       NUMERIC(10, 2) NOT NULL CHECK (price > 0),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
`

const migration003Up = `
CREATE TABLE IF NOT EXISTS orders (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID NOT NULL REFERENCES users(id),
    total       NUMERIC(10, 2) NOT NULL DEFAULT 0,
    status      VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
`

const migration004Up = `
CREATE TABLE IF NOT EXISTS order_items (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id    UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id  UUID NOT NULL REFERENCES products(id),
    quantity    INTEGER NOT NULL CHECK (quantity > 0),
    unit_price  NUMERIC(10, 2) NOT NULL CHECK (unit_price > 0)
);
`

func runMigrations(db *sql.DB) error {
	migrations := []struct {
		name string
		sql  string
	}{
		{"001_create_users", migration001Up},
		{"002_create_products", migration002Up},
		{"003_create_orders", migration003Up},
		{"004_create_order_items", migration004Up},
	}

	for _, m := range migrations {
		_, err := db.Exec(m.sql)
		if err != nil {
			return fmt.Errorf("migration %s: %w", m.name, err)
		}

		fmt.Println("Migration applied:", m.name)
	}

	return nil
}

// ============================================================
// 6. INDEXES
// ============================================================

/*
Indexes speed up reads.

Without index → full table scan (slow on large tables)
With index    → fast lookup


Automatically created:

    PRIMARY KEY
    UNIQUE


Add manually for common queries:

    WHERE user_id = ?
    ORDER BY created_at
    WHERE status = 'pending'


Trade-off:

    Faster reads
    Slower writes (index must update)
*/

const migration005Indexes = `
CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);
CREATE INDEX IF NOT EXISTS idx_orders_created_at ON orders(created_at);
CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items(order_id);
`

func createIndexes(db *sql.DB) error {
	_, err := db.Exec(migration005Indexes)
	if err != nil {
		return fmt.Errorf("create indexes: %w", err)
	}

	fmt.Println("Indexes created")
	return nil
}

// ============================================================
// 7. JOIN QUERIES
// ============================================================

/*
JOIN combines rows from multiple tables.


INNER JOIN
    Only matching rows from both tables


Example — orders with user email:

    SELECT o.id, o.total, o.status, u.email
    FROM orders o
    INNER JOIN users u ON u.id = o.user_id
    WHERE o.user_id = $1


Example — order with items and product names:

    SELECT oi.quantity, oi.unit_price, p.name
    FROM order_items oi
    INNER JOIN products p ON p.id = oi.product_id
    WHERE oi.order_id = $1
*/

type OrderWithEmail struct {
	OrderID string
	Total   string
	Status  string
	Email   string
}

func listOrdersByUser(ctx context.Context, db *sql.DB, userID string) ([]OrderWithEmail, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT o.id, o.total, o.status, u.email
		FROM orders o
		INNER JOIN users u ON u.id = o.user_id
		WHERE o.user_id = $1
		ORDER BY o.created_at DESC
	`, userID)

	if err != nil {
		return nil, fmt.Errorf("list orders: %w", err)
	}

	defer rows.Close()

	orders := make([]OrderWithEmail, 0)

	for rows.Next() {
		var order OrderWithEmail

		err = rows.Scan(&order.OrderID, &order.Total, &order.Status, &order.Email)
		if err != nil {
			return nil, fmt.Errorf("scan order: %w", err)
		}

		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return orders, nil
}

type OrderItemDetail struct {
	ProductName string
	Quantity    int
	UnitPrice   string
}

func getOrderItems(ctx context.Context, db *sql.DB, orderID string) ([]OrderItemDetail, error) {
	rows, err := db.QueryContext(ctx, `
		SELECT p.name, oi.quantity, oi.unit_price
		FROM order_items oi
		INNER JOIN products p ON p.id = oi.product_id
		WHERE oi.order_id = $1
	`, orderID)

	if err != nil {
		return nil, fmt.Errorf("get order items: %w", err)
	}

	defer rows.Close()

	items := make([]OrderItemDetail, 0)

	for rows.Next() {
		var item OrderItemDetail

		err = rows.Scan(&item.ProductName, &item.Quantity, &item.UnitPrice)
		if err != nil {
			return nil, fmt.Errorf("scan item: %w", err)
		}

		items = append(items, item)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return items, nil
}

// ============================================================
// 8. PAGINATION IN SQL
// ============================================================

/*
Pagination in SQL:

    LIMIT   → how many rows
    OFFSET  → how many rows to skip


Formula (same as REST API):

    offset = (page - 1) * limit


Example — page 2, 10 per page:

    SELECT id, total, status
    FROM orders
    WHERE user_id = $1
    ORDER BY created_at DESC
    LIMIT 10 OFFSET 10
*/

func listOrdersPaginated(
	ctx context.Context,
	db *sql.DB,
	userID string,
	page, limit int,
) ([]OrderWithEmail, error) {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	rows, err := db.QueryContext(ctx, `
		SELECT o.id, o.total, o.status, u.email
		FROM orders o
		INNER JOIN users u ON u.id = o.user_id
		WHERE o.user_id = $1
		ORDER BY o.created_at DESC
		LIMIT $2 OFFSET $3
	`, userID, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("paginated orders: %w", err)
	}

	defer rows.Close()

	orders := make([]OrderWithEmail, 0)

	for rows.Next() {
		var order OrderWithEmail

		err = rows.Scan(&order.OrderID, &order.Total, &order.Status, &order.Email)
		if err != nil {
			return nil, fmt.Errorf("scan order: %w", err)
		}

		orders = append(orders, order)
	}

	return orders, rows.Err()
}

// ============================================================
// 9. TRANSACTIONS
// ============================================================

/*
Transaction = all queries succeed OR all fail.

    BEGIN
        INSERT order
        INSERT order_items
        UPDATE product stock
    COMMIT

If any step fails → ROLLBACK (nothing saved)


Use transactions when:

    Creating order + order items
    Transferring money
    Any multi-step write that must stay consistent


In Go:

    tx, err := db.BeginTx(ctx, nil)
    defer tx.Rollback()   // no-op after Commit

    tx.ExecContext(...)
    tx.ExecContext(...)

    tx.Commit()
*/

type CreateOrderInput struct {
	UserID    string
	ProductID string
	Quantity  int
	UnitPrice string
}

func createOrderWithItem(ctx context.Context, db *sql.DB, input CreateOrderInput) (string, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("begin tx: %w", err)
	}

	defer tx.Rollback()

	var orderID string

	total := input.UnitPrice // simplified: 1 item for demo

	err = tx.QueryRowContext(ctx, `
		INSERT INTO orders (user_id, total, status)
		VALUES ($1, $2, 'pending')
		RETURNING id
	`, input.UserID, total).Scan(&orderID)

	if err != nil {
		return "", fmt.Errorf("insert order: %w", err)
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO order_items (order_id, product_id, quantity, unit_price)
		VALUES ($1, $2, $3, $4)
	`, orderID, input.ProductID, input.Quantity, input.UnitPrice)

	if err != nil {
		return "", fmt.Errorf("insert order item: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return "", fmt.Errorf("commit tx: %w", err)
	}

	return orderID, nil
}

// ============================================================
// 10. SEED DEMO DATA
// ============================================================

func seedDemoData(ctx context.Context, db *sql.DB) (userID, productID string, err error) {
	err = db.QueryRowContext(ctx, `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		ON CONFLICT (email) DO UPDATE SET email = EXCLUDED.email
		RETURNING id
	`, "shop@example.com", "$2a$10$demo").Scan(&userID)

	if err != nil {
		return "", "", fmt.Errorf("seed user: %w", err)
	}

	err = db.QueryRowContext(ctx, `
		INSERT INTO products (name, price)
		VALUES ($1, $2)
		RETURNING id
	`, "Go Book", "29.99").Scan(&productID)

	if err != nil {
		err = db.QueryRowContext(ctx, `
			SELECT id FROM products WHERE name = $1 LIMIT 1
		`, "Go Book").Scan(&productID)
	}

	if err != nil {
		return "", "", fmt.Errorf("seed product: %w", err)
	}

	return userID, productID, nil
}

// ============================================================
// 11. HELPERS
// ============================================================

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return db, nil
}

// ============================================================
// 12. COMPLETE DEMO
// ============================================================

func main() {
	db, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to PostgreSQL")
	fmt.Println()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = runMigrations(db)
	if err != nil {
		log.Fatal(err)
	}

	err = createIndexes(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	userID, productID, err := seedDemoData(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Demo user ID:", userID)
	fmt.Println("Demo product ID:", productID)

	orderID, err := createOrderWithItem(ctx, db, CreateOrderInput{
		UserID:    userID,
		ProductID: productID,
		Quantity:  1,
		UnitPrice: "29.99",
	})

	if err != nil {
		log.Fatal("create order:", err)
	}

	fmt.Println("Order created:", orderID)
	fmt.Println()

	orders, err := listOrdersByUser(ctx, db, userID)
	if err != nil {
		log.Fatal("list orders:", err)
	}

	fmt.Println("Orders for user:")

	for _, o := range orders {
		fmt.Printf("  %s  total=%s  status=%s  email=%s\n", o.OrderID, o.Total, o.Status, o.Email)
	}

	fmt.Println()

	items, err := getOrderItems(ctx, db, orderID)
	if err != nil {
		log.Fatal("order items:", err)
	}

	fmt.Println("Order items:")

	for _, item := range items {
		fmt.Printf("  %s  qty=%d  price=%s\n", item.ProductName, item.Quantity, item.UnitPrice)
	}

	fmt.Println()

	pageOrders, err := listOrdersPaginated(ctx, db, userID, 1, 10)
	if err != nil {
		log.Fatal("paginated:", err)
	}

	fmt.Printf("Paginated orders (page 1): %d result(s)\n", len(pageOrders))
}

/*
============================================================
POSTGRESQL MENTAL MODEL
============================================================

Question 1:

    Where to define tables?

Answer:

    Schema design + migration files


Question 2:

    When to use transactions?

Answer:

    Multiple writes that must all succeed or all fail


Question 3:

    When to add indexes?

Answer:

    Columns in WHERE, JOIN, ORDER BY used frequently


Question 4:

    How to query related data?

Answer:

    JOIN tables on foreign keys


Question 5:

    How to paginate in SQL?

Answer:

    ORDER BY ... LIMIT $n OFFSET $m


============================================================
MIGRATION FILES (PRODUCTION)
============================================================

Create folder:

    08-databases/49-postgresql/migrations/


Example 001_create_users.up.sql:

    CREATE TABLE users (...);


Example 001_create_users.down.sql:

    DROP TABLE IF EXISTS users;


Run with golang-migrate:

    migrate -path ./migrations -database "$DSN" up


============================================================
EXPLAIN (DEBUG SLOW QUERIES)
============================================================

In psql:

    EXPLAIN ANALYZE
    SELECT * FROM orders WHERE user_id = '...';


Shows whether PostgreSQL uses an index or full table scan.


============================================================
DOCKER
============================================================

    docker exec -it postgres psql -U postgres

    \dt              -- list tables
    \d users         -- describe table
    \di              -- list indexes


Run lecture:

    go run main.go


Note:

    If you have an old TEXT-based users table from lecture 48,
    drop it first or use a fresh database:

    DROP TABLE IF EXISTS order_items, orders, products, users CASCADE;


============================================================
END OF 49 — POSTGRESQL
============================================================
*/
