/*
============================================================
43 — HTTP SERVER
============================================================

Package:
    net/http

Topics:
    1. HTTP Server
    2. Handler
    3. HandlerFunc
    4. Router / ServeMux
    5. Request
    6. HTTP Methods
    7. URL
    8. Query Parameters
    9. Request Headers
   10. Request Body
   11. ResponseWriter
   12. Response Headers
   13. Status Codes
   14. http.Error
   15. Request Context
   16. Middleware
   17. Handler vs HandlerFunc vs ServeMux
   18. Backend Request Flow

Run:

    go run main.go

The program starts a local HTTP server and then makes
requests to it automatically.

============================================================
*/

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

/*
============================================================
1. HANDLER
============================================================

A Handler processes an HTTP request.

The fundamental interface is:

    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }

A function can also be used as a handler through
http.HandlerFunc / http.HandleFunc.
*/

// ----------------------------------------------------------
// Handler using a normal function
// ----------------------------------------------------------

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home handler called")

	fmt.Fprintln(w, "Welcome to Home")
}

// ----------------------------------------------------------
// Users handler
// ----------------------------------------------------------

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Users handler called")

	fmt.Fprintln(w, "Users endpoint")
}

// ----------------------------------------------------------
// Create user handler
// ----------------------------------------------------------

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	fmt.Fprintln(w, "User created")
}

/*
============================================================
2. REQUEST
============================================================

Every handler receives:

    r *http.Request

The Request contains information about the incoming request.

Important fields:

    r.Method
    r.URL
    r.Header
    r.Body
    r.Context()
    r.Host
*/

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----- Request Information -----")

	fmt.Println("Method:", r.Method)
	fmt.Println("URL:", r.URL)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Host:", r.Host)

	fmt.Fprintln(w, "Request received")
}

/*
============================================================
3. HTTP METHODS
============================================================

Common methods:

    GET
    POST
    PUT
    PATCH
    DELETE

Go provides constants:

    http.MethodGet
    http.MethodPost
    http.MethodPut
    http.MethodPatch
    http.MethodDelete
*/

func methodHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		fmt.Fprintln(w, "GET request")

	case http.MethodPost:
		fmt.Fprintln(w, "POST request")

	case http.MethodPut:
		fmt.Fprintln(w, "PUT request")

	case http.MethodPatch:
		fmt.Fprintln(w, "PATCH request")

	case http.MethodDelete:
		fmt.Fprintln(w, "DELETE request")

	default:
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

/*
============================================================
4. URL
============================================================

Example:

    /users?page=2&limit=10

r.URL.Path:

    /users

r.URL.RawQuery:

    page=2&limit=10
*/

func urlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("RawQuery:", r.URL.RawQuery)

	fmt.Fprintln(w, "URL information printed")
}

/*
============================================================
5. QUERY PARAMETERS
============================================================

Example:

    /users?page=2&limit=10

Use:

    r.URL.Query()

Then:

    query.Get("page")
    query.Get("limit")
*/

func queryHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	page := query.Get("page")
	limit := query.Get("limit")

	fmt.Println("Page:", page)
	fmt.Println("Limit:", limit)

	fmt.Fprintln(w, "Query parameters:")
	fmt.Fprintln(w, "page =", page)
	fmt.Fprintln(w, "limit =", limit)
}

/*
============================================================
6. REQUEST HEADERS
============================================================

Incoming HTTP headers are available through:

    r.Header

For one header:

    r.Header.Get("Authorization")
*/

func headerHandler(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Authorization")
	contentType := r.Header.Get("Content-Type")

	fmt.Println("Authorization:", authorization)
	fmt.Println("Content-Type:", contentType)

	fmt.Fprintln(w, "Headers received")
}

/*
============================================================
7. REQUEST BODY
============================================================

The request body is:

    r.Body

It implements io.ReadCloser.

For simple demonstration:

    io.ReadAll(r.Body)

Later, for JSON APIs, we will normally use:

    json.NewDecoder(r.Body).Decode(...)
*/

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(
			w,
			"Could not read body",
			http.StatusBadRequest,
		)
		return
	}

	fmt.Println("Request body:", string(body))

	fmt.Fprintln(w, "Body received:", string(body))
}

/*
============================================================
8. RESPONSE WRITER
============================================================

The first argument of a handler is:

    w http.ResponseWriter

It is used to create the HTTP response.

For example:

    fmt.Fprintln(w, "Hello")
*/

func responseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from ResponseWriter")
}

/*
============================================================
9. RESPONSE HEADERS
============================================================

Set response headers with:

    w.Header().Set(...)
*/

func responseHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/plain",
	)

	fmt.Fprintln(w, "Response header was set")
}

/*
============================================================
10. STATUS CODES
============================================================

Set a response status using:

    w.WriteHeader(...)

Example:

    http.StatusCreated

Important:

    Headers/status should be set before writing
    the response body.
*/

func createdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/plain",
	)

	w.WriteHeader(http.StatusCreated)

	fmt.Fprintln(w, "Resource created")
}

/*
============================================================
11. http.Error
============================================================

Go provides:

    http.Error(...)

Example:

    http.Error(
        w,
        "User not found",
        http.StatusNotFound,
    )
*/

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(
		w,
		"User not found",
		http.StatusNotFound,
	)
}

/*
============================================================
12. REQUEST CONTEXT
============================================================

Every HTTP request has a Context:

    r.Context()

The context represents the lifetime of the request.

It can be passed to:

    Handler
       ↓
    Service
       ↓
    Repository
       ↓
    Database

This connects directly to:

    42 — Context
*/

func contextHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Println("Request context:", ctx)

	fmt.Fprintln(w, "Request context received")
}

/*
============================================================
13. SERVICE + CONTEXT
============================================================
*/

func service(ctx context.Context) error {
	fmt.Println("Service started")

	select {

	case <-time.After(500 * time.Millisecond):
		fmt.Println("Service completed")
		return nil

	case <-ctx.Done():
		fmt.Println("Service cancelled")
		return ctx.Err()
	}
}

/*
============================================================
14. CONTEXT IN HTTP HANDLER
============================================================

We can derive a timeout from the HTTP request context.

    r.Context()
          ↓
    WithTimeout()
          ↓
       Service
*/

func contextTimeoutHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		r.Context(),
		2*time.Second,
	)
	defer cancel()

	err := service(ctx)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusGatewayTimeout,
		)
		return
	}

	fmt.Fprintln(w, "Service completed")
}

/*
============================================================
15. MIDDLEWARE
============================================================

Middleware wraps another Handler.

Basic structure:

    func middleware(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w, r) {

            // before

            next.ServeHTTP(w, r)

            // after
        })
    }

Middleware can be used for:

    Logging
    Recovery
    Authentication
    CORS
    Request IDs
*/

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		fmt.Println("----- Middleware BEFORE -----")
		fmt.Println(r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		fmt.Println("----- Middleware AFTER -----")
	})
}

/*
============================================================
16. HANDLER TYPE
============================================================

A Handler is an interface:

    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }

Example custom Handler:
*/

type HelloHandler struct{}

func (HelloHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintln(w, "Hello from custom Handler")
}

/*
============================================================
17. ROUTER / SERVEMUX
============================================================

ServeMux is Go's standard router.

Create one:

    mux := http.NewServeMux()

Register handlers:

    mux.HandleFunc(...)
    mux.Handle(...)

Then give the router to the server.
*/

/*
============================================================
18. HTTP SERVER
============================================================

A server can be started using:

    http.ListenAndServe(...)

For simple applications:

    http.ListenAndServe(":8080", mux)

For more control, Go provides:

    http.Server
*/

func createServer() *http.Server {
	mux := http.NewServeMux()

	/*
		Normal routes
	*/
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/create", createUserHandler)

	/*
		Request information
	*/
	mux.HandleFunc("/request", requestHandler)
	mux.HandleFunc("/method", methodHandler)
	mux.HandleFunc("/url", urlHandler)
	mux.HandleFunc("/query", queryHandler)
	mux.HandleFunc("/headers", headerHandler)
	mux.HandleFunc("/body", bodyHandler)

	/*
		Response
	*/
	mux.HandleFunc("/response", responseHandler)
	mux.HandleFunc("/response-header", responseHeaderHandler)
	mux.HandleFunc("/created", createdHandler)
	mux.HandleFunc("/error", errorHandler)

	/*
		Context
	*/
	mux.HandleFunc("/context", contextHandler)
	mux.HandleFunc("/context-timeout", contextTimeoutHandler)

	/*
		Custom Handler
	*/
	mux.Handle("/custom-handler", HelloHandler{})

	/*
		Middleware

		Instead of:

			server → mux

		we have:

			server
			   ↓
			logging middleware
			   ↓
			mux
			   ↓
			handler
	*/
	handler := loggingMiddleware(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	return server
}

/*
============================================================
19. HTTP CLIENT
============================================================

We use the HTTP client below only so this lecture can
demonstrate the server automatically.

Normally the client could be:

    Browser
    Frontend
    Mobile App
    Another Backend
*/

func makeRequest(
	method string,
	url string,
	body io.Reader,
) {
	req, err := http.NewRequest(
		method,
		url,
		body,
	)

	if err != nil {
		fmt.Println("Request creation error:", err)
		return
	}

	/*
		Set request headers.
	*/
	req.Header.Set(
		"Authorization",
		"Bearer demo-token",
	)

	req.Header.Set(
		"Content-Type",
		"text/plain",
	)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Request error:", err)
		return
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Response read error:", err)
		return
	}

	fmt.Println("HTTP Status:", resp.Status)
	fmt.Println("Response:", string(responseBody))
}

/*
============================================================
20. MAIN
============================================================
*/

func main() {
	fmt.Println("==================================================")
	fmt.Println("43 — HTTP SERVER")
	fmt.Println("==================================================")

	server := createServer()

	/*
		Start server in a goroutine.

		Why?

		http.Server blocks while it is running.

		We want main() to continue so we can
		make test requests automatically.
	*/
	go func() {
		fmt.Println("HTTP server running on :8080")

		err := server.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error:", err)
		}
	}()

	/*
		Give the server a moment to start.

		This is ONLY for this executable lecture/demo.

		In production, do not use Sleep for synchronization.
	*/
	time.Sleep(200 * time.Millisecond)

	/*
		========================================================
			1. GET /
		========================================================
	*/

	fmt.Println("\n1. GET /")

	makeRequest(
		http.MethodGet,
		"http://localhost:8080/",
		nil,
	)

	/*
		========================================================
			2. GET /users
		========================================================
	*/

	fmt.Println("\n2. GET /users")

	makeRequest(
		http.MethodGet,
		"http://localhost:8080/users",
		nil,
	)

	/*
		========================================================
			3. Query Parameters
		========================================================
	*/

	fmt.Println("\n3. GET /query?page=2&limit=10")

	makeRequest(
		http.MethodGet,
		"http://localhost:8080/query?page=2&limit=10",
		nil,
	)

	/*
		========================================================
			4. Request Headers
		========================================================
	*/

	fmt.Println("\n4. GET /headers")

	makeRequest(
		http.MethodGet,
		"http://localhost:8080/headers",
		nil,
	)

	/*
		========================================================
			5. POST + Request Body
		========================================================
	*/

	fmt.Println("\n5. POST /body")

	makeRequest(
		http.MethodPost,
		"http://localhost:8080/body",
		io.Reader(
			&stringReader{
				value: `{"name":"Mahdi"}`,
			},
		),
	)

	/*
		========================================================
			6. Created Response
		========================================================
	*/

	fmt.Println("\n6. POST /created")

	makeRequest(
		http.MethodPost,
		"http://localhost:8080/created",
		nil,
	)

	/*
		========================================================
			7. Error Response
		========================================================
	*/

	fmt.Println("\n7. GET /error")

	makeRequest(
		http.MethodGet,
		"http://localhost:8080/error",
		nil,
	)

	/*
		========================================================
			8. Custom Handler
		========================================================
	*/

	fmt.Println("\n8. GET /custom-handler")

	makeRequest(
		http.MethodGet,
		"http://localhost:8080/custom-handler",
		nil,
	)

	/*
		========================================================
			9. Context
		========================================================
	*/

	fmt.Println("\n9. GET /context")

	makeRequest(
		http.MethodGet,
		"http://localhost:8080/context",
		nil,
	)

	/*
		========================================================
		10. Context + Timeout
		========================================================
	*/

	fmt.Println("\n10. GET /context-timeout")

	makeRequest(
		http.MethodGet,
		"http://localhost:8080/context-timeout",
		nil,
	)

	/*
		========================================================
		SHUTDOWN
		========================================================

		Graceful shutdown will be studied later in more detail.

		Here we simply stop the server so the lecture
		program can exit cleanly.
	*/

	fmt.Println("\nStopping server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Second,
	)
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		fmt.Println("Shutdown error:", err)
	}

	fmt.Println("Server stopped")

	fmt.Println("\n==================================================")
	fmt.Println("43 — HTTP SERVER COMPLETE")
	fmt.Println("==================================================")
}

/*
============================================================
21. SIMPLE STRING READER
============================================================

This is only used to demonstrate sending a request body
without adding another package.

In real applications we can use:

	strings.NewReader(...)
*/
type stringReader struct {
	value string
}

func (s *stringReader) Read(p []byte) (int, error) {
	if len(s.value) == 0 {
		return 0, io.EOF
	}

	n := copy(p, s.value)
	s.value = s.value[n:]

	return n, nil
}
