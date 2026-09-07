/*
============================================================
Problem 1 — Create User JSON API
============================================================

Create a JSON API endpoint:

	POST /users

Requirements:

	1. Create a CreateUserRequest struct with:
		name
		age

	2. Use JSON tags.

	3. Decode the request body using:
		json.NewDecoder()

	4. Return 400 Bad Request if the JSON is invalid.

	5. Create a UserResponse struct.

	6. Return the response using:
		json.NewEncoder()

	7. Set Content-Type to:
		application/json

	8. Return:
		201 Created

	9. Run the server on:
		:8080

Important:

	Do NOT use json.Marshal().
	Do NOT use a database.
	Do NOT use third-party packages.

Expected request:

	{
		"name": "Mahdi",
		"age": 27
	}

Expected response:

	{
		"name": "Mahdi",
		"age": 27
	}

============================================================
Goal
============================================================

Practice:

	- JSON request/response
	- JSON tags
	- Request/Response DTOs
	- json.Decoder
	- json.Encoder
	- HTTP status codes
	- Content-Type
	- net/http

============================================================
*/

package main

func main() {}
