/*
============================================================
Problem 16 — Complete User JSON API
============================================================

Create:

	POST /users
	GET  /users

Requirements:

 1. POST /users accepts JSON:

    {
    "name": "Mahdi",
    "age": 27
    }

 2. Decode using json.Decoder.

 3. Return 400 for invalid JSON.

 4. Return 201 for successful creation.

 5. GET /users returns a JSON array.

 6. Use separate request and response DTOs.

 7. Set Content-Type:
    application/json

 8. Use json.Encoder for responses.

 9. Store users in memory.

 10. Do NOT use a database.

 11. Do NOT use third-party packages.

============================================================
Goal
============================================================

Practice:

  - JSON request/response
  - DTOs
  - json.Decoder
  - json.Encoder
  - JSON arrays
  - HTTP status codes
  - Content-Type
  - net/http
  - Complete JSON API flow

============================================================
*/

package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

var UserDb []User = []User{}

type CreateUserDto struct {
	Name string `json:"name"`
	Age  *int   `json:"age"`
}

type UserResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func createUsersHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody CreateUserDto
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(
			w,
			"Invalid Input",
			400,
		)
		return
	}

	if requestBody.Name == "" {
		http.Error(
			w,
			"Name Is Required",
			400,
		)
		return
	}

	if requestBody.Age == nil {
		http.Error(
			w,
			"Age is required",
			400,
		)
		return
	} else if *(requestBody.Age) <= 0 {
		http.Error(
			w,
			"Age must be positive",
			400,
		)
		return
	}

	user := User{
		Age:  *requestBody.Age,
		Name: requestBody.Name,
	}

	UserDb = append(UserDb, user)

	w.WriteHeader(http.StatusCreated)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := make([]UserResponse, 0, len(UserDb))

	for _, user := range UserDb {
		response = append(response, UserResponse{
			Name: user.Name,
			Age:  user.Age,
		})
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode users", http.StatusInternalServerError)
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", getUsersHandler)
	mux.HandleFunc("POST /users", createUsersHandler)

	http.ListenAndServe(":8080", mux)
}
