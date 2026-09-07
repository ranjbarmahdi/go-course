/*
============================================================
48 — DATABASE/SQL
Problem 9 — openDB Helper
============================================================

Create a reusable database connection helper.

Requirements:

    1. Create function:

           func openDB(dsn string) (*sql.DB, error)

    2. openDB must:

           sql.Open("pgx", dsn)
           db.Ping()
           return db on success

    3. If Ping fails, close db and return error

    4. Use fmt.Errorf with %w for wrapping errors:

           fmt.Errorf("ping db: %w", err)

    5. In main(), call openDB with the standard DSN

    6. Print:

           Connected to PostgreSQL

    7. defer db.Close()

DSN:

    postgres://postgres:secret@localhost:5432/postgres?sslmode=disable

============================================================
Goal
============================================================

Practice:

    - reusable connection helper
    - error wrapping
    - clean startup code

============================================================
*/

package main

func main() {}
