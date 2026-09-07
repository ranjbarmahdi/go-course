/*
============================================================
Problem 6 — Nested JSON
============================================================

Create an endpoint:

	GET /users/1

Create:

	User
	Address

The User must contain:

	id
	name
	address

The Address must contain:

	city
	country

Return nested JSON.

Expected response:

	{
		"id": 1,
		"name": "Mahdi",
		"address": {
			"city": "Baku",
			"country": "Azerbaijan"
		}
	}

============================================================
Goal
============================================================

Practice:

	- Nested structs
	- Nested JSON
	- JSON tags
	- json.Encoder
============================================================
*/

package main

func main() {}
