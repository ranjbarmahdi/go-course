/*
============================================================
Problem 9 — Strict JSON Validation
============================================================

Create:

	POST /users

The request struct contains:

	name
	age

The client sends:

	{
		"name": "Mahdi",
		"age": 27,
		"unknown": "test"
	}

Requirements:

	1. Create a json.Decoder.

	2. Enable:

		DisallowUnknownFields()

	3. Decode the request.

	4. Return 400 Bad Request when
	   an unknown field exists.

============================================================
Goal
============================================================

Practice:

	- json.Decoder
	- DisallowUnknownFields()
	- Strict JSON APIs
============================================================
*/

package main

func main() {}
