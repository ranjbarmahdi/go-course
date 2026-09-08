package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

/*
============================================================
47 — AUTHENTICATION
============================================================

Topics
------------------------------------------------------------
 1. Authentication vs Authorization
 2. Why Hash Passwords
 3. bcrypt Password Hashing
 4. User Model
 5. Register
 6. Login
 7. JWT Structure
 8. Generate JWT
 9. Validate JWT
 10. JWT Auth Middleware
 11. Context From JWT
 12. Access Token vs Refresh Token
 13. Refresh Token Flow
 14. Complete Auth Server
 15. Authentication Mental Model
============================================================
*/

// ============================================================
// 1. AUTHENTICATION VS AUTHORIZATION
// ============================================================

/*
Authentication (AuthN)
----------------------

    "Who are you?"


Examples:

    Login with email + password
    Verify JWT token
    Check API key


Authorization (AuthZ)
-------------------

    "What are you allowed to do?"


Examples:

    Admin can delete users
    User can only edit own profile
    Guest cannot access /admin


In middleware (46), you checked:

    Authorization header exists?


That was a simplified gate.

This lecture builds real authentication:

    Register
    Login
    Password hashing
    JWT access token
    JWT refresh token
    Protected routes
*/

// ============================================================
// 2. WHY HASH PASSWORDS
// ============================================================

/*
NEVER store plain passwords in a database.


Bad:

    password = "mypassword123"


If the database leaks, every password is exposed.


Good:

    password_hash = "$2a$10$..."


Hashing is one-way:

    plain password  →  hash
    hash            ↗  plain password   (NOT possible)


On login:

    hash the input password
    compare with stored hash


Use bcrypt for password hashing in Go:

    golang.org/x/crypto/bcrypt
*/

// ============================================================
// 3. BCRYPT PASSWORD HASHING
// ============================================================

/*
Hash a password:

    hash, err := bcrypt.GenerateFromPassword(
        []byte(password),
        bcrypt.DefaultCost,
    )


Verify a password:

    err := bcrypt.CompareHashAndPassword(
        []byte(storedHash),
        []byte(inputPassword),
    )


If err == nil  → password matches
If err != nil  → wrong password


DefaultCost controls slowness.

Higher cost = safer but slower.
*/

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}

// ============================================================
// 4. USER MODEL
// ============================================================

/*
A user stores identity data.

For authentication we need at least:

    ID
    Email
    PasswordHash


Never expose PasswordHash in API responses.


Register request DTO:

    email
    password


Login request DTO:

    email
    password


Auth response DTO:

    access_token
    refresh_token
    token_type
*/

type User struct {
	ID           string
	Email        string
	PasswordHash string
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type ProfileResponse struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

// ============================================================
// 5. IN-MEMORY USER STORE
// ============================================================

/*
For learning, we use an in-memory store.

In production, this would be PostgreSQL via repository pattern.
*/

type UserStore struct {
	mu    sync.RWMutex
	users map[string]User
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[string]User),
	}
}

func (store *UserStore) Create(user User) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, exists := store.users[user.Email]; exists {
		return errors.New("email already exists")
	}

	store.users[user.Email] = user
	return nil
}

func (store *UserStore) FindByEmail(email string) (User, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	user, ok := store.users[email]
	return user, ok
}

// ============================================================
// 6. REGISTER
// ============================================================

/*
Register flow:

    POST /register
    {
        "email": "mahdi@example.com",
        "password": "secret123"
    }


Steps:

    1. Decode JSON body
    2. Validate email and password
    3. Hash password with bcrypt
    4. Save user
    5. Return 201 Created
*/

func validateRegisterRequest(req RegisterRequest) error {
	if req.Email == "" {
		return errors.New("email is required")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

func registerHandler(store *UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{
				Error: "invalid JSON",
			})
			return
		}

		err = validateRegisterRequest(req)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		passwordHash, err := hashPassword(req.Password)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, ErrorResponse{
				Error: "failed to hash password",
			})
			return
		}

		user := User{
			ID:           fmt.Sprintf("user-%d", time.Now().UnixNano()),
			Email:        req.Email,
			PasswordHash: passwordHash,
		}

		err = store.Create(user)
		if err != nil {
			writeJSON(w, http.StatusConflict, ErrorResponse{
				Error: err.Error(),
			})
			return
		}

		writeJSON(w, http.StatusCreated, map[string]string{
			"message": "user registered",
			"email":   user.Email,
		})
	}
}

// ============================================================
// 7. LOGIN
// ============================================================

/*
Login flow:

    POST /login
    {
        "email": "mahdi@example.com",
        "password": "secret123"
    }


Steps:

    1. Decode JSON body
    2. Find user by email
    3. Verify password with bcrypt
    4. Generate access token
    5. Generate refresh token
    6. Return tokens
*/

func loginHandler(
	store *UserStore,
	tokenService *TokenService,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{
				Error: "invalid JSON",
			})
			return
		}

		user, ok := store.FindByEmail(req.Email)
		if !ok {
			writeJSON(w, http.StatusUnauthorized, ErrorResponse{
				Error: "invalid email or password",
			})
			return
		}

		err = verifyPassword(req.Password, user.PasswordHash)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, ErrorResponse{
				Error: "invalid email or password",
			})
			return
		}

		accessToken, err := tokenService.GenerateAccessToken(user)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, ErrorResponse{
				Error: "failed to generate access token",
			})
			return
		}

		refreshToken, err := tokenService.GenerateRefreshToken(user)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, ErrorResponse{
				Error: "failed to generate refresh token",
			})
			return
		}

		writeJSON(w, http.StatusOK, AuthResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			TokenType:    "Bearer",
		})
	}
}

// ============================================================
// 8. JWT STRUCTURE
// ============================================================

/*
JWT = JSON Web Token

Format:

    header.payload.signature


Example:

    eyJhbGciOiJIUzI1NiIs...
    .eyJ1c2VyX2lkIjoiMTIzIi...
    .SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c


Payload (claims) example:

    {
        "user_id": "user-123",
        "email": "mahdi@example.com",
        "exp": 1710000000
    }


Common claims:

    user_id   → who the token belongs to
    email     → user email
    exp       → expiration timestamp
    iat       → issued at


Client sends token in header:

    Authorization: Bearer <token>
*/

// ============================================================
// 9. TOKEN SERVICE
// ============================================================

/*
Centralize token logic in a TokenService.

Two token types:

    Access token   → short-lived (15 minutes)
    Refresh token  → long-lived (7 days)


Both are JWTs, but refresh token has a different purpose.
*/

type TokenService struct {
	secret          []byte
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewTokenService(secret string) *TokenService {
	return &TokenService{
		secret:          []byte(secret),
		accessTokenTTL:  15 * time.Minute,
		refreshTokenTTL: 7 * 24 * time.Hour,
	}
}

type TokenClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

func (service *TokenService) GenerateAccessToken(user User) (string, error) {
	claims := TokenClaims{
		UserID: user.ID,
		Email:  user.Email,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(service.accessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(service.secret)
}

func (service *TokenService) GenerateRefreshToken(user User) (string, error) {
	claims := TokenClaims{
		UserID: user.ID,
		Email:  user.Email,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(service.refreshTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(service.secret)
}

func (service *TokenService) ParseToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&TokenClaims{},
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}

			return service.secret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// ============================================================
// 10. JWT AUTH MIDDLEWARE
// ============================================================

/*
Auth middleware for JWT:

    1. Read Authorization header
    2. Extract Bearer token
    3. Parse and validate JWT
    4. Attach user info to context
    5. Call next handler


If token is missing or invalid:

    HTTP 401 Unauthorized
*/

type contextKey string

const userIDKey contextKey = "userID"
const userEmailKey contextKey = "userEmail"

func jwtAuthMiddleware(tokenService *TokenService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				writeJSON(w, http.StatusUnauthorized, ErrorResponse{
					Error: "missing authorization header",
				})
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				writeJSON(w, http.StatusUnauthorized, ErrorResponse{
					Error: "invalid authorization header format",
				})
				return
			}

			claims, err := tokenService.ParseToken(parts[1])
			if err != nil {
				writeJSON(w, http.StatusUnauthorized, ErrorResponse{
					Error: "invalid or expired token",
				})
				return
			}

			if claims.Type != "access" {
				writeJSON(w, http.StatusUnauthorized, ErrorResponse{
					Error: "access token required",
				})
				return
			}

			ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
			ctx = context.WithValue(ctx, userEmailKey, claims.Email)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// ============================================================
// 11. PROTECTED ROUTE — PROFILE
// ============================================================

/*
Protected route reads user info from context.

    GET /profile
    Authorization: Bearer <access_token>


Response:

    {
        "user_id": "user-123",
        "email": "mahdi@example.com"
    }
*/

func profileHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := r.Context().Value(userIDKey).(string)
	email, _ := r.Context().Value(userEmailKey).(string)

	writeJSON(w, http.StatusOK, ProfileResponse{
		UserID: userID,
		Email:  email,
	})
}

// ============================================================
// 12. ACCESS TOKEN VS REFRESH TOKEN
// ============================================================

/*
Access Token
------------

    Short-lived (15 minutes)
    Sent with every API request
    Used to access protected routes


Refresh Token
-------------

    Long-lived (7 days)
    Used ONLY to get a new access token
    NOT used on normal API routes


Why two tokens?

    If access token is stolen, damage is limited (expires fast)
    User stays logged in via refresh token
    Refresh token can be revoked in production
*/

// ============================================================
// 13. REFRESH TOKEN FLOW
// ============================================================

/*
When access token expires:

    POST /refresh
    {
        "refresh_token": "<refresh_token>"
    }


Steps:

    1. Parse refresh token
    2. Verify type == "refresh"
    3. Find user
    4. Generate new access token
    5. Optionally rotate refresh token
*/

func refreshHandler(
	store *UserStore,
	tokenService *TokenService,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RefreshRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{
				Error: "invalid JSON",
			})
			return
		}

		if req.RefreshToken == "" {
			writeJSON(w, http.StatusBadRequest, ErrorResponse{
				Error: "refresh_token is required",
			})
			return
		}

		claims, err := tokenService.ParseToken(req.RefreshToken)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, ErrorResponse{
				Error: "invalid or expired refresh token",
			})
			return
		}

		if claims.Type != "refresh" {
			writeJSON(w, http.StatusUnauthorized, ErrorResponse{
				Error: "refresh token required",
			})
			return
		}

		user, ok := store.FindByEmail(claims.Email)
		if !ok {
			writeJSON(w, http.StatusUnauthorized, ErrorResponse{
				Error: "user not found",
			})
			return
		}

		accessToken, err := tokenService.GenerateAccessToken(user)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, ErrorResponse{
				Error: "failed to generate access token",
			})
			return
		}

		writeJSON(w, http.StatusOK, map[string]string{
			"access_token": accessToken,
			"token_type":   "Bearer",
		})
	}
}

// ============================================================
// 14. JSON HELPER
// ============================================================

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

// ============================================================
// 15. COMPLETE AUTH SERVER
// ============================================================

func createRouter(
	store *UserStore,
	tokenService *TokenService,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", registerHandler(store))
	mux.HandleFunc("POST /login", loginHandler(store, tokenService))
	mux.HandleFunc("POST /refresh", refreshHandler(store, tokenService))

	mux.Handle(
		"GET /profile",
		jwtAuthMiddleware(tokenService)(
			http.HandlerFunc(profileHandler),
		),
	)

	return mux
}

func main() {
	store := NewUserStore()
	tokenService := NewTokenService("super-secret-key-change-in-production")

	mux := createRouter(store, tokenService)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Auth server running on http://localhost:8080")
	fmt.Println()
	fmt.Println("Try these requests:")
	fmt.Println()
	fmt.Println(`  curl -X POST http://localhost:8080/register \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"email":"mahdi@example.com","password":"secret123"}'`)
	fmt.Println()
	fmt.Println(`  curl -X POST http://localhost:8080/login \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"email":"mahdi@example.com","password":"secret123"}'`)
	fmt.Println()
	fmt.Println(`  curl http://localhost:8080/profile \`)
	fmt.Println(`    -H "Authorization: Bearer <access_token>"`)
	fmt.Println()
	fmt.Println(`  curl -X POST http://localhost:8080/refresh \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"refresh_token":"<refresh_token>"}'`)
	fmt.Println()

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

/*
============================================================
AUTHENTICATION MENTAL MODEL
============================================================

Question 1:

    Where do passwords go?

Answer:

    Never store plain text — hash with bcrypt


Question 2:

    How does login work?

Answer:

    Find user → verify password → return tokens


Question 3:

    What is JWT used for?

Answer:

    Stateless proof of identity after login


Question 4:

    Access vs refresh token?

Answer:

    Access  → short-lived, used on API routes
    Refresh → long-lived, used to get new access token


Question 5:

    How do protected routes work?

Answer:

    JWT middleware validates token → puts user in context


============================================================
FULL FLOW
============================================================

    REGISTER
        POST /register
            ↓
        hash password
            ↓
        save user


    LOGIN
        POST /login
            ↓
        verify password
            ↓
        return access + refresh tokens


    PROTECTED ROUTE
        GET /profile
            ↓
        jwtAuthMiddleware
            ↓
        profileHandler


    REFRESH
        POST /refresh
            ↓
        validate refresh token
            ↓
        return new access token


============================================================
END OF 47 — AUTHENTICATION
============================================================
*/
