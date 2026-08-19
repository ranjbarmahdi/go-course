package main

import "fmt"

func main() {

	// ============================================================
	// 1. Create a Map
	// ============================================================
	// A map stores key-value pairs.
	//
	// Syntax:
	//
	// map[KeyType]ValueType
	//
	// In this example:
	// key   = string
	// value = int

	users := map[string]int{
		"Mahdi": 27,
		"Ali":   30,
		"Sara":  25,
	}

	fmt.Println(users)

	// ============================================================
	// 2. Read a Value
	// ============================================================
	// Access a value using its key.

	fmt.Println(users["Mahdi"])

	// ============================================================
	// 3. Add a Value
	// ============================================================
	// Assigning a value to a new key adds it to the map.

	users["Reza"] = 35

	fmt.Println(users)

	// ============================================================
	// 4. Update a Value
	// ============================================================
	// If the key already exists, assigning a new value
	// updates the existing value.

	users["Ali"] = 31

	fmt.Println(users)

	// ============================================================
	// 5. Delete a Value
	// ============================================================
	// delete() removes a key-value pair.
	//
	// Syntax:
	//
	// delete(map, key)

	delete(users, "Mahdi")

	fmt.Println(users)

	// ============================================================
	// 6. Check if a Key Exists
	// ============================================================
	// Map lookup can return two values:
	//
	// value, exists := map[key]
	//
	// exists is true if the key exists.
	// exists is false if the key does not exist.

	age, exists := users["Ali"]

	fmt.Println("Age:", age)
	fmt.Println("Exists:", exists)

	// ============================================================
	// 7. Missing Keys
	// ============================================================
	// If a key does not exist, the map returns the zero value
	// of the value type.
	//
	// For int, the zero value is 0.

	age = users["Unknown"]

	fmt.Println("Unknown age:", age)

	// ============================================================
	// 8. range
	// ============================================================
	// Use range to iterate over a map.
	//
	// Map iteration order is NOT guaranteed.

	for name, age := range users {
		fmt.Println(name, age)
	}

	// ============================================================
	// 9. Map Length
	// ============================================================
	// len() returns the number of key-value pairs.

	fmt.Println("Number of users:", len(users))

	// ============================================================
	// 10. Create an Empty Map
	// ============================================================
	// make() creates an initialized map that can be written to.

	products := make(map[string]int)

	products["Laptop"] = 1200
	products["Phone"] = 800

	fmt.Println(products)

	// ============================================================
	// 11. Nil Map
	// ============================================================
	// A nil map can be read from,
	// but you cannot add values to it.

	var scores map[string]int

	fmt.Println(scores)
	fmt.Println(scores["Mahdi"])

	// The following would panic:
	//
	// scores["Mahdi"] = 100
}
