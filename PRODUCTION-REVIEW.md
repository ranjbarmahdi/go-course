# Go Backend — Production Review

Fast cheat sheet from your lectures. Focus: what you use daily in real backends.

---

## Architecture (memorize this)

```
HTTP Request
     ↓
Middleware (logging, recovery, CORS, auth)
     ↓
Handler        → parse path/query/body, return status + JSON
     ↓
Service        → business logic, validation rules
     ↓
Repository     → database/sql (interface in prod)
     ↓
PostgreSQL
```

**Rules:**
- Handlers = HTTP only (no SQL in handlers)
- Services = business logic (no `ResponseWriter`)
- Repository = persistence (no HTTP types)
- Depend on **interfaces**, inject via **constructors**

---

## 1. Errors (17-error-handling)

```go
// Always return error as last value
func GetUser(id string) (User, error)

// Wrap with context
return fmt.Errorf("find user %s: %w", id, err)

// Sentinel errors
var ErrNotFound = errors.New("not found")

// Check sentinels
if errors.Is(err, ErrNotFound) { /* 404 */ }

// Typed errors
var ve *ValidationError
if errors.As(err, &ve) { /* 400 with details */ }
```

| HTTP | When |
|------|------|
| 400 | Bad input / validation |
| 401 | Missing or invalid auth |
| 404 | `sql.ErrNoRows` / not found |
| 409 | Duplicate (email exists) |
| 500 | Unexpected DB/system error |

---

## 2. Interfaces + DI (14, 15)

```go
// Small interface at consumer
type UserRepository interface {
    FindByEmail(ctx context.Context, email string) (User, error)
    Create(ctx context.Context, user User) error
}

// Service depends on interface
type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}
```

- Implicit interface satisfaction (no `implements` keyword)
- Mock repos in tests by implementing the same interface

---

## 3. JSON (24-encoding-json, 45-json-api)

```go
// Request
var req CreateUserRequest
decoder := json.NewDecoder(r.Body)
decoder.DisallowUnknownFields()  // strict API
if err := decoder.Decode(&req); err != nil {
    http.Error(w, "invalid JSON", 400)
    return
}

// Response
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(response)
```

**Tags you use constantly:**
```go
Name  string  `json:"name"`
Email string  `json:"email,omitempty"`
Pass  string  `json:"-"`              // never in JSON
Age   *int    `json:"age"`            // null vs omitted
```

**Order matters:** set headers → `WriteHeader` → write body

---

## 4. HTTP Server (43-http-server)

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /users", listUsers)
mux.HandleFunc("GET /users/{id}", getUser)
mux.HandleFunc("POST /users", createUser)

server := &http.Server{Addr: ":8080", Handler: mux}
server.ListenAndServe()
```

**Request:**
```go
r.Method
r.URL.Path
r.PathValue("id")           // Go 1.22+
r.URL.Query().Get("page")   // ?page=2
r.Header.Get("Authorization")
r.Body                      // read once
r.Context()                 // pass downstream
```

**Response:**
```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
fmt.Fprintln(w, "text")
http.Error(w, "msg", code)  // sets status + plain text
```

---

## 5. REST API (44-rest-api)

| Method | Path | Action |
|--------|------|--------|
| GET | `/users` | List (query: page, limit, sort) |
| POST | `/users` | Create |
| GET | `/users/{id}` | Read one |
| PUT | `/users/{id}` | Replace |
| PATCH | `/users/{id}` | Partial update |
| DELETE | `/users/{id}` | Delete |

**Path vs query:**
- Path → **which** resource: `/users/123`
- Query → **how** to fetch: `/users?page=2&limit=20&sort=-createdAt`

**Pagination:**
```go
offset = (page - 1) * limit
// defaults: page=1, limit=20, max limit=100
```

**DTOs:** separate request/response structs from domain models. Never expose `PasswordHash` in API.

---

## 6. Middleware (46-middleware)

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)  // MUST call or request stops
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

// Global: wrap mux
handler := chain(corsMiddleware, loggingMiddleware, recoveryMiddleware)(mux)

// Per-route: wrap handler
mux.Handle("GET /profile", authMiddleware(http.HandlerFunc(profileHandler)))
```

**Recovery (always use in prod):**
```go
defer func() {
    if err := recover(); err != nil {
        http.Error(w, "Internal Server Error", 500)
    }
}()
```

**Chain order:** first wrapped = outermost = runs first on way in.

---

## 7. Authentication (47-authentication)

```go
// Hash — never store plain password
hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(input))

// JWT access token (short: 15 min)
claims := jwt.MapClaims{
    "user_id": id,
    "type":    "access",
    "exp":     time.Now().Add(15 * time.Minute).Unix(),
}
token.SignedString([]byte(secret))

// Client sends
Authorization: Bearer <access_token>
```

| Token | TTL | Used on |
|-------|-----|---------|
| Access | ~15 min | All protected API routes |
| Refresh | ~7 days | `POST /refresh` only |

**Auth middleware flow:**
1. Read `Authorization: Bearer ...`
2. Parse + validate JWT
3. Check `type == "access"`
4. Attach user to `context.WithValue`
5. `next.ServeHTTP(w, r.WithContext(ctx))`

**On failure:** return early — do NOT call `next`.

---

## 8. Context (42-context)

```go
// Always first parameter
func (r *Repo) FindByEmail(ctx context.Context, email string) (User, error)

// From HTTP
ctx := r.Context()

// Timeout for DB/external calls
ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
defer cancel()

db.QueryRowContext(ctx, query, args...)
```

- Pass `ctx` handler → service → repository
- Cancelled request = stop DB work
- `WithValue` for request-scoped data (user ID, request ID) — not for DI

---

## 9. Database (48-database-sql)

```go
import (
    "database/sql"
    _ "github.com/jackc/pgx/v5/stdlib"
)

db, _ := sql.Open("pgx", "postgres://user:pass@localhost:5432/db?sslmode=disable")
db.Ping()
defer db.Close()

// Pool tuning (production)
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(5 * time.Minute)
```

| Method | Use |
|--------|-----|
| `Exec` | INSERT, UPDATE, DELETE, CREATE |
| `QueryRow` | SELECT one row |
| `Query` | SELECT many rows |

```go
// One row
err := db.QueryRow("SELECT id, email FROM users WHERE email = $1", email).Scan(&id, &email)
if err == sql.ErrNoRows { /* 404 */ }

// Many rows
rows, _ := db.Query("SELECT id, email FROM users")
defer rows.Close()
for rows.Next() { rows.Scan(&id, &email) }
rows.Err()
```

**Placeholders:** PostgreSQL uses `$1, $2` (not `?`).

---

## 10. Concurrency (37–41) — production patterns

```go
// Wait for goroutines
var wg sync.WaitGroup
wg.Add(1)
go func() { defer wg.Done(); work() }()
wg.Wait()

// Protect shared state
var mu sync.Mutex
mu.Lock()
counter++
mu.Unlock()

// Worker pool (cap concurrency)
jobs := make(chan Job, 100)
for i := 0; i < 5; i++ {
    go worker(jobs)
}

// Timeout
select {
case result := <-ch:
    use(result)
case <-time.After(3 * time.Second):
    return errors.New("timeout")
case <-ctx.Done():
    return ctx.Err()
}
```

**Rules:**
- `main` exit kills all goroutines — always sync
- Only sender closes channels
- Prefer `context` over raw `time.After` in HTTP handlers

---

## 11. OS + I/O (22-os, 23-io)

```go
// Config
password := os.Getenv("DB_PASSWORD")
value, ok := os.LookupEnv("PORT")  // distinguish missing vs empty

// Streams
io.Copy(dst, src)
data, _ := io.ReadAll(r.Body)
io.LimitReader(r.Body, 1<<20)  // max 1MB upload
```

---

## 12. Full request lifecycle (put it together)

```
POST /register { email, password }
  → decode JSON
  → validate DTO
  → bcrypt hash
  → repo.Create(ctx, user)
  → 201 JSON

POST /login { email, password }
  → repo.FindByEmail(ctx, email)
  → ErrNoRows → 401
  → bcrypt compare
  → generate access + refresh JWT
  → 200 JSON

GET /profile
  → jwtAuthMiddleware
  → read user from context
  → 200 JSON
```

---

## Quick decision table

| Need | Use |
|------|-----|
| Shared HTTP logic | Middleware |
| Who is the user? | JWT + auth middleware |
| Store password | bcrypt hash |
| API input/output shape | DTO + json tags |
| Not found in DB | `sql.ErrNoRows` → 404 |
| Cancel long operation | `context.Context` |
| Many DB connections | `*sql.DB` pool settings |
| Test without DB | Interface + mock repository |
| Strict JSON input | `DisallowUnknownFields()` |
| Panic in handler | Recovery middleware → 500 |

---

## Common mistakes (avoid these)

1. Storing plain-text passwords
2. Forgetting `next.ServeHTTP` in middleware
3. Setting headers after `WriteHeader` / `Write`
4. Treating `sql.ErrNoRows` as 500
5. Using `sql.Open` without `Ping`
6. Not closing `rows` after `Query`
7. Putting SQL inside HTTP handlers
8. Using refresh token on protected routes
9. Not passing `context` to DB calls
10. `main` returning while goroutines still run

---

## Packages you use most

```
net/http              HTTP server
encoding/json         API bodies
database/sql          SQL interface
github.com/jackc/pgx/v5/stdlib   PostgreSQL
golang.org/x/crypto/bcrypt       passwords
github.com/golang-jwt/jwt/v5       tokens
context               cancellation + values
errors, fmt           error handling
sync                  mutex, WaitGroup
time                  timeouts, JWT exp
os                    environment config
io                    streaming
```

---

## Your progress

| Done | Topic |
|------|-------|
| ✅ | HTTP Server, REST, JSON API |
| ✅ | Middleware, Authentication |
| ✅ | database/sql + pgx |
| ⬜ | PostgreSQL (schema, migrations, transactions) |
| ⬜ | Repository pattern |
| ⬜ | Production (config, slog, Redis, Docker, CI) |

---

*Generated from go-crash lectures. Re-read before interviews or starting a new feature.*
