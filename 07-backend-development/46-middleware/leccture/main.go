package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

/*
============================================================
46 — MIDDLEWARE
============================================================

Topics
------------------------------------------------------------
 1. What Is Middleware
 2. Middleware Pattern
 3. Handler Chain
 4. Logging Middleware
 5. Global vs Route-Specific Middleware
 6. Middleware Chaining
 7. Recovery Middleware
 8. Authentication Middleware
 9. Context in Middleware
10. CORS Middleware
11. Complete Middleware Server
12. Middleware Mental Model
============================================================
*/

// ============================================================
// 1. WHAT IS MIDDLEWARE
// ============================================================

/*
Middleware is a function that wraps another http.Handler.

It sits between the server and your route handler.

    Client
       ↓
    Middleware
       ↓
    Handler
       ↓
    Response


Middleware runs shared logic without duplicating it
inside every handler.


Common uses:

    Logging
    Recovery (panic handling)
    Authentication
    CORS
    Request ID
    Rate limiting
    Timeout


You already saw the basic idea in 43 — HTTP Server.

This lecture goes deeper with real production patterns.
*/

// ============================================================
// 2. MIDDLEWARE PATTERN
// ============================================================

/*
Every middleware follows the same signature:

    func middleware(next http.Handler) http.Handler


Inside, convert to HandlerFunc:

    return http.HandlerFunc(func(w, r) {

        // BEFORE handler

        next.ServeHTTP(w, r)

        // AFTER handler
    })


Three rules:

    1. Accept  next http.Handler
    2. Return  http.Handler
    3. Call    next.ServeHTTP(w, r) to continue the chain


If you forget next.ServeHTTP, the request stops
and your route handler never runs.
*/

func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		fmt.Println("before handler")

		next.ServeHTTP(w, r)

		fmt.Println("after handler")
	})
}

// ============================================================
// 3. HANDLER CHAIN
// ============================================================

/*
Think of middleware as a chain of handlers.

    Client
       ↓
    Middleware A
       ↓
    Middleware B
       ↓
    Middleware C
       ↓
    Route Handler


Request flows DOWN the chain.
Response flows BACK UP the chain.


Each middleware can:

    - Run code BEFORE the next handler
    - Run code AFTER the next handler
    - Stop the chain early (auth failure, validation, etc.)
*/

// ============================================================
// 4. LOGGING MIDDLEWARE
// ============================================================

/*
Logging is the most common middleware.

It records:

    - HTTP method
    - URL path
    - Duration


Typical output:

    → GET /users
    ← GET /users (1.2ms)
*/

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		start := time.Now()

		fmt.Println("→", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		fmt.Printf(
			"← %s %s (%v)\n",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

// ============================================================
// 5. GLOBAL VS ROUTE-SPECIFIC MIDDLEWARE
// ============================================================

/*
There are two levels of middleware.


GLOBAL MIDDLEWARE
-----------------

Applies to ALL routes.

Wrap the router (ServeMux):

    handler := loggingMiddleware(mux)


Every registered route passes through it.


ROUTE-SPECIFIC MIDDLEWARE
-------------------------

Applies to ONE route only.

Wrap a single handler:

    protected := authMiddleware(
        http.HandlerFunc(usersHandler),
    )

    mux.Handle("GET /users", protected)


Example flows:

    GET /users

        logging → recovery → auth → usersHandler


    GET /products

        logging → recovery → productsHandler

        (no auth)
*/

// ============================================================
// 6. MIDDLEWARE CHAINING
// ============================================================

/*
Multiple middleware are composed by nesting:

    handler := loggingMiddleware(
        recoveryMiddleware(
            mux,
        ),
    )


Order matters.

The first wrapped middleware is the OUTERMOST layer.


Request direction (in):

    logging → recovery → mux → handler


Response direction (out):

    handler → mux → recovery → logging


Think of it like layers of an onion.


Alternative — chain helper (useful with many middleware):

    func chain(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
        return func(final http.Handler) http.Handler {
            for i := len(middlewares) - 1; i >= 0; i-- {
                final = middlewares[i](final)
            }
            return final
        }
    }


Usage:

    handler := chain(
        corsMiddleware,
        loggingMiddleware,
        recoveryMiddleware,
    )(mux)
*/

func chain(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(final http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}

// ============================================================
// 7. RECOVERY MIDDLEWARE
// ============================================================

/*
If a handler panics, the entire server can crash.

Recovery middleware catches panics and returns 500 instead.


Uses:

    defer func() {
        if err := recover(); err != nil {
            // handle panic
        }
    }()


Without recovery:

    panic in handler → server crash


With recovery:

    panic in handler → 500 Internal Server Error
*/

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("PANIC:", err)

				http.Error(
					w,
					"Internal Server Error",
					http.StatusInternalServerError,
				)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// ============================================================
// 8. AUTHENTICATION MIDDLEWARE
// ============================================================

/*
Auth middleware checks if the client is allowed to access
a route.


Common check:

    Authorization header


If missing or invalid:

    HTTP 401 Unauthorized
    DO NOT call next.ServeHTTP


If valid:

    continue to the next handler


Important:

    return early on failure.
    Never call next after sending an error response.
*/

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(
				w,
				"Unauthorized",
				http.StatusUnauthorized,
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ============================================================
// 9. CONTEXT IN MIDDLEWARE
// ============================================================

/*
Middleware can attach data to a request using context.

Typical use cases:

    User ID after authentication
    Request ID for tracing
    Roles / permissions


Pattern:

    ctx := context.WithValue(r.Context(), key, value)
    r = r.WithContext(ctx)


In the handler:

    userID := r.Context().Value(key)


Never use string keys directly — use a custom type
to avoid collisions.
*/

type contextKey string

const userIDKey contextKey = "userID"

func authWithContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(
				w,
				"Unauthorized",
				http.StatusUnauthorized,
			)
			return
		}

		// In a real app, parse JWT and extract user ID.
		// Here we use a simplified example.
		ctx := context.WithValue(r.Context(), userIDKey, "user-123")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(userIDKey)

	fmt.Fprintf(w, "Profile for user: %v\n", userID)
}

// ============================================================
// 10. CORS MIDDLEWARE
// ============================================================

/*
CORS = Cross-Origin Resource Sharing

Browsers block requests from a different origin unless
the server sends CORS headers.


Example:

    Frontend:  http://localhost:3000
    Backend:   http://localhost:8080

These are different origins. The browser requires CORS.


Required headers:

    Access-Control-Allow-Origin
    Access-Control-Allow-Methods
    Access-Control-Allow-Headers


Browsers send an OPTIONS preflight request before
the actual request. Handle it separately.
*/

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, PATCH, DELETE, OPTIONS",
		)
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization",
		)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ============================================================
// 11. ROUTE HANDLERS
// ============================================================

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users endpoint")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Products endpoint")
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("something went wrong")
}

// ============================================================
// 12. COMPLETE MIDDLEWARE SERVER
// ============================================================

func createRouter() *http.ServeMux {
	mux := http.NewServeMux()

	/*
		Protected route — auth required
	*/
	mux.Handle(
		"GET /users",
		authWithContextMiddleware(
			http.HandlerFunc(usersHandler),
		),
	)

	/*
		Protected route — auth + context
	*/
	mux.Handle(
		"GET /profile",
		authWithContextMiddleware(
			http.HandlerFunc(profileHandler),
		),
	)

	/*
		Public route — no auth
	*/
	mux.HandleFunc(
		"GET /products",
		productsHandler,
	)

	/*
		Test recovery — triggers a panic
	*/
	mux.HandleFunc(
		"GET /panic",
		panicHandler,
	)

	return mux
}

func main() {
	mux := createRouter()

	/*
		Global middleware chain:

		    cors → logging → recovery → router
	*/
	handler := chain(
		corsMiddleware,
		loggingMiddleware,
		recoveryMiddleware,
	)(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println()
	fmt.Println("Try these requests:")
	fmt.Println()
	fmt.Println("  curl http://localhost:8080/products")
	fmt.Println("  curl http://localhost:8080/users")
	fmt.Println(`  curl -H "Authorization: Bearer token" http://localhost:8080/users`)
	fmt.Println(`  curl -H "Authorization: Bearer token" http://localhost:8080/profile`)
	fmt.Println("  curl http://localhost:8080/panic")
	fmt.Println()

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

/*
============================================================
MIDDLEWARE MENTAL MODEL
============================================================

Question 1:

    Where does shared HTTP logic go?

Answer:

    Middleware


Question 2:

    How do I chain middleware?

Answer:

    Nest wrappers — outermost runs first on the way in


Question 3:

    Global or per-route?

Answer:

    Global     → wrap the router (mux)
    Per-route  → wrap individual handlers


Question 4:

    How do I block a request?

Answer:

    Return early — do NOT call next.ServeHTTP


Question 5:

    How do I pass data to handlers?

Answer:

    context.WithValue in middleware
    r.Context().Value in handler


============================================================
REQUEST FLOW — GET /users (with Authorization header)
============================================================

    Client
       ↓
    corsMiddleware
       ↓
    loggingMiddleware        → GET /users
       ↓
    recoveryMiddleware
       ↓
    ServeMux
       ↓
    authWithContextMiddleware
       ↓
    usersHandler
       ↓
    Response
       ↓
    loggingMiddleware        ← GET /users (duration)


============================================================
END OF 46 — MIDDLEWARE
============================================================
*/
