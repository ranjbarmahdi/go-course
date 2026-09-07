package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserRequest struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

type UserResponse struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest
	var response UserResponse

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)

	fmt.Println(err)
	if err != nil {
		http.Error(
			w,
			"Invalid Input",
			http.StatusBadRequest,
		)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
	data, err := json.Marshal(response)

	fmt.Println("Name:", request.Name)
	fmt.Println("Age", request.Age)

	w.Write(data)
	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	w.WriteHeader(http.StatusCreated)
}

func main() {

	http.HandleFunc("POST /users", createUserHandler)

	http.ListenAndServe(":8080", nil)
}
