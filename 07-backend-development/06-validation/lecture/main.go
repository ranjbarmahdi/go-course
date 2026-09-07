package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"unicode"

	"github.com/go-playground/validator/v10"
)

/*
============================================================
06 — VALIDATION
============================================================

Topics
------------------------------------------------------------
 1. Why Validation
 2. Three Layers of Validation
 3. Manual vs validator/v10
 4. Install and Setup
 5. Struct Tags
 6. validate.Struct
 7. Common Tags
 8. Validation Errors
 9. HTTP Handler Pattern
10. Nested Structs
11. Slices with dive
12. omitempty for Updates
13. Custom Validators
14. Complete Demo
15. Validation Mental Model
============================================================
*/

// ============================================================
// 1. WHY VALIDATION
// ============================================================

/*
Clients send bad data all the time:

    missing fields
    wrong types
    invalid email format
    negative prices


Without validation:

    bad data reaches your database
    cryptic errors later
    security issues


Validate at the boundary — right after JSON decode, before business logic.
*/

// ============================================================
// 2. THREE LAYERS OF VALIDATION
// ============================================================

/*
Layer 1 — Config validation (startup)

    DATABASE_URL required
    JWT_SECRET min 32 chars

    Tool: manual Validate() in LoadConfig
    Topic: 09-production/01-configuration


Layer 2 — Input validation (HTTP DTOs)

    email format, password min length, price >= 0

    Tool: validator/v10 struct tags
    Topic: this lecture


Layer 3 — Business rules (use case / domain)

    email already exists
    insufficient stock
    user cannot delete own account

    Tool: plain Go in use case layer
    NOT struct tags


Flow:

    JSON decode → validator → use case → domain rules → response
*/

// ============================================================
// 3. MANUAL VS validator/v10
// ============================================================

/*
In 05-authentication you wrote:

    func validateRegisterRequest(req RegisterRequest) error {
        if req.Email == "" {
            return errors.New("email is required")
        }
        if len(req.Password) < 8 {
            return errors.New("password must be at least 8 characters")
        }
        ...
    }


Fine for 2 fields. Painful for 20 fields across many endpoints.


validator/v10 replaces if-chains with struct tags:

    type RegisterRequest struct {
        Email    string `json:"email" validate:"required,email"`
        Password string `json:"password" validate:"required,min=8"`
    }

    err := validate.Struct(&req)
*/

// ============================================================
// 4. SETUP
// ============================================================

/*
Install:

    go get github.com/go-playground/validator/v10


Create ONE validator instance and reuse it.
It caches struct metadata via reflection.
*/

var validate = validator.New()

// ============================================================
// 5. REQUEST DTOs
// ============================================================

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type CreateProductRequest struct {
	Name     string   `json:"name" validate:"required,min=2,max=100"`
	Price    float64  `json:"price" validate:"required,gte=0"`
	Category string   `json:"category" validate:"required,oneof=electronics clothing food"`
	Tags     []string `json:"tags" validate:"dive,min=1,max=20"`
}

type Address struct {
	City    string `json:"city" validate:"required"`
	Country string `json:"country" validate:"required,len=2"`
}

type CreateUserRequest struct {
	Name    string  `json:"name" validate:"required,min=2"`
	Email   string  `json:"email" validate:"required,email"`
	Address Address `json:"address" validate:"required"`
}

type UpdateProfileRequest struct {
	Age int    `json:"age" validate:"omitempty,gte=0,lte=120"`
	Bio string `json:"bio" validate:"omitempty,max=500"`
}

type ErrorResponse struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

// ============================================================
// 6. validate.Struct
// ============================================================

/*
Always pass a POINTER:

    validate.Struct(&req)   ✓
    validate.Struct(req)    ✗


Returns nil if valid.
Returns validator.ValidationErrors if tags fail.
*/

func validateStruct(value any) error {
	return validate.Struct(value)
}

// ============================================================
// 7. COMMON TAGS
// ============================================================

/*
Tag          Meaning
────────────────────────────────────────────
required     field must be non-zero
email        valid email format
min=8        min length (string) or value (number)
max=100      max length / value
gte=0        greater than or equal
lte=120      less than or equal
oneof=a b c  must be one of listed values
omitempty    skip validation if empty (updates)
dive         validate each slice element
len=2        exact length
url          valid URL
uuid         valid UUID


Combine with commas:

    validate:"required,email"
    validate:"required,min=8,max=72"
*/

// ============================================================
// 8. FORMAT VALIDATION ERRORS
// ============================================================

/*
Never return raw validator errors to clients in production.
Map them to field → message for 400 Bad Request.
*/

func formatValidationError(err error) map[string]string {
	fields := make(map[string]string)

	var validationErrors validator.ValidationErrors
	if !errors.As(err, &validationErrors) {
		fields["error"] = err.Error()
		return fields
	}

	for _, e := range validationErrors {
		fields[e.Field()] = formatFieldError(e)
	}

	return fields
}

func formatFieldError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "email":
		return e.Field() + " must be a valid email"
	case "min":
		return e.Field() + " must be at least " + e.Param() + " characters"
	case "max":
		return e.Field() + " must be at most " + e.Param() + " characters"
	case "gte":
		return e.Field() + " must be >= " + e.Param()
	case "oneof":
		return e.Field() + " must be one of: " + e.Param()
	default:
		return e.Field() + " failed " + e.Tag() + " validation"
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// ============================================================
// 9. HTTP HANDLER PATTERN
// ============================================================

/*
Every validated endpoint follows the same steps:

    1. Decode JSON        → 400 invalid JSON
    2. validate.Struct    → 400 validation failed
    3. Business logic     → 409 conflict, 201 created, etc.


Order matters — decode first, then validate, then rules.
*/

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid JSON"})
		return
	}

	err = validate.Struct(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{
			Error:  "validation failed",
			Fields: formatValidationError(err),
		})
		return
	}

	// Business rule example (NOT a validator tag):
	// if emailAlreadyExists(req.Email) → 409 Conflict

	writeJSON(w, http.StatusCreated, map[string]string{
		"message": "user registered",
		"email":   req.Email,
	})
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid JSON"})
		return
	}

	err = validate.Struct(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{
			Error:  "validation failed",
			Fields: formatValidationError(err),
		})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"message":  "product created",
		"name":     req.Name,
		"price":    req.Price,
		"category": req.Category,
	})
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid JSON"})
		return
	}

	err = validate.Struct(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{
			Error:  "validation failed",
			Fields: formatValidationError(err),
		})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{
		"message": "user created",
		"name":    req.Name,
		"city":    req.Address.City,
	})
}

// ============================================================
// 10. NESTED STRUCTS
// ============================================================

/*
Validator walks nested structs automatically.

    type CreateUserRequest struct {
        Address Address `json:"address" validate:"required"`
    }

    type Address struct {
        City string `json:"city" validate:"required"`
    }


If Address.City is empty → validation fails on nested field.
*/

// ============================================================
// 11. SLICES WITH dive
// ============================================================

/*
Tags on slice:

    Tags []string `json:"tags" validate:"dive,min=1,max=20"`


dive validates EACH element.
Empty tag on slice itself means slice can be nil/empty.
Add `required` before dive if slice must exist:

    validate:"required,dive,min=1"
*/

// ============================================================
// 12. omitempty FOR UPDATES
// ============================================================

/*
PATCH /profile — client may send only changed fields.

    type UpdateProfileRequest struct {
        Age int    `json:"age" validate:"omitempty,gte=0,lte=120"`
        Bio string `json:"bio" validate:"omitempty,max=500"`
    }


omitempty = skip validation when field is zero value.
If Age is sent as 0, gte=0 still runs.
If Age is omitted (0 default), validation is skipped.
*/

// ============================================================
// 13. CUSTOM VALIDATORS
// ============================================================

/*
Built-in tags cover most cases.
Register custom validators for domain-specific rules.

Example: password must contain at least one digit.
*/

func passwordHasDigit(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	for _, ch := range password {
		if unicode.IsDigit(ch) {
			return true
		}
	}

	return false
}

func init() {
	_ = validate.RegisterValidation("has_digit", passwordHasDigit)
}

type SecureRegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,has_digit"`
}

// ============================================================
// 14. COMPLETE DEMO
// ============================================================

func createRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})

	mux.HandleFunc("POST /register", registerHandler)
	mux.HandleFunc("POST /products", createProductHandler)
	mux.HandleFunc("POST /users", createUserHandler)

	return mux
}

func demoCLIValidation() {
	fmt.Println("=== CLI validation demo ===")

	bad := RegisterRequest{Email: "not-an-email", Password: "123"}
	err := validate.Struct(&bad)
	if err != nil {
		fmt.Println("invalid register:", formatValidationError(err))
	}

	good := RegisterRequest{Email: "me@example.com", Password: "secret123"}
	err = validate.Struct(&good)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("valid register request")

	secureBad := SecureRegisterRequest{Email: "me@example.com", Password: "secretonly"}
	err = validate.Struct(&secureBad)
	if err != nil {
		fmt.Println("invalid secure register:", formatValidationError(err))
	}

	fmt.Println()
}

func main() {
	demoCLIValidation()

	mux := createRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server running on http://localhost:8080")
	log.Println("Try:")
	log.Println(`  curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d "{\"email\":\"bad\",\"password\":\"123\"}"`)
	log.Println(`  curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d "{\"email\":\"me@example.com\",\"password\":\"secret123\"}"`)
	log.Println(`  curl -X POST http://localhost:8080/products -H "Content-Type: application/json" -d "{\"name\":\"Phone\",\"price\":999,\"category\":\"electronics\",\"tags\":[\"new\"]}"`)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

/*
============================================================
VALIDATION MENTAL MODEL
============================================================

Question 1:

    Where does validator/v10 run?

Answer:

    On HTTP request DTOs, after JSON decode


Question 2:

    Replace config Validate()?

Answer:

    No — keep manual validation for startup config


Question 3:

    Replace "email already exists"?

Answer:

    No — business rule in use case (409 Conflict)


Question 4:

    What HTTP status for tag validation failure?

Answer:

    400 Bad Request


Question 5:

    One validator or new per request?

Answer:

    One shared validator.New() instance


Question 6:

    validate.Struct(&req) or validate.Struct(req)?

Answer:

    Pointer — always &req


============================================================
RUN
============================================================

    go run ./07-backend-development/06-validation/lecture/

Test:

    curl http://localhost:8080/health

    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"bad","password":"123"}'

    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"me@example.com","password":"secret123"}'


============================================================
END OF 06 — VALIDATION
============================================================
*/
