/*
============================================================
48 — DATABASE/SQL
Problem 2 — Connection Pool
============================================================

Connect to PostgreSQL and configure the connection pool.

Requirements:

    1. Connect using pgx (same DSN as Problem 1)

    2. Configure pool:

           SetMaxOpenConns(25)
           SetMaxIdleConns(5)
           SetConnMaxLifetime(5 * time.Minute)

    3. Create function:

           func configurePool(db *sql.DB)

    4. After configuring, print:

           Pool configured

    5. Ping to verify connection still works

    6. defer db.Close()

============================================================
Goal
============================================================

Practice:

    - connection pool settings
    - *sql.DB as a pool

============================================================
*/

package main

func main() {}
