/*
============================================================
Problem 8 — Unknown JSON Fields
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

 1. Decode the request using json.Decoder.

 2. Do not enable strict field checking.

 3. The request must still be accepted.

 4. Return 201 Created.

============================================================
Goal
============================================================

Practice:

  - json.Decoder
  - Unknown JSON fields
  - Default decoder behavior

============================================================
*/
package main

func main() {}
