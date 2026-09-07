package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// 24-encoding-json
// │
// ├── 01 json.Marshal()
// ├── 02 Struct JSON tags
// ├── 03 omitempty
// ├── 04 json.Unmarshal()
// ├── 05 JSON array → Go slice
// ├── 06 json.MarshalIndent()
// ├── 07 json.NewEncoder()
// ├── 08 json.NewDecoder()
// ├── 09 Marshal vs Unmarshal
// ├── 10 Encoder vs Decoder
// ├── 11 JSON in Backend Development
// ├── 12 JSON → map[string]any
// ├── 13 Unknown JSON Fields
// ├── 14 DisallowUnknownFields()
// ├── 15 JSON null
// ├── 16 Custom JSON Marshaling
// ├── 17 Custom JSON Unmarshaling
// ├── 18 json.RawMessage
// ├── 19 json.Number
// ├── 20 json.Valid()
// ├── 21 JSON "-" Tag
// ├── 22 JSON Field Matching
// └── 23 JSON Syntax Errors

// ============================================================
// 16. Custom JSON Marshaling ⭐⭐⭐
// ============================================================
// A type can customize how it is converted into JSON.
//
// To do this, implement:
//
// MarshalJSON() ([]byte, error)
//
// The method must return valid JSON.

type Product struct {
	Name  string
	Price float64
}

func (p Product) MarshalJSON() ([]byte, error) {

	type ProductJSON struct {
		ProductName string  `json:"product_name"`
		Price       float64 `json:"price"`
	}

	return json.Marshal(ProductJSON{
		ProductName: p.Name,
		Price:       p.Price,
	})
}

// ============================================================
// 17. Custom JSON Unmarshaling ⭐⭐⭐
// ============================================================

func (p *Product) UnmarshalJSON(data []byte) error {

	type ProductJSON struct {
		ProductName string  `json:"product_name"`
		Price       float64 `json:"price"`
	}

	var value ProductJSON

	err := json.Unmarshal(data, &value)

	if err != nil {
		return err
	}

	p.Name = value.ProductName
	p.Price = value.Price

	return nil
}

func main() {
	// ============================================================
	// 1. json.Marshal() ⭐⭐⭐
	// ============================================================
	// `json.Marshal()` converts a Go value into JSON.
	//
	// Go value → JSON bytes
	//
	// Function:
	//
	// json.Marshal(value)
	//
	// Returns:
	//
	// []byte
	// error
	//
	// The returned data is []byte because JSON is encoded as bytes.

	fmt.Println("1. json.Marshal()")

	user := struct {
		Name string
		Age  int
	}{
		Name: "Mahdi",
		Age:  27,
	}

	data, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}

	fmt.Println(data)
	fmt.Println(string(data))

	// ============================================================
	// 2. Struct JSON Tags ⭐⭐⭐
	// ============================================================
	// Struct tags control how struct fields are represented
	// in JSON.
	//
	// Example:
	//
	// `json:"first_name"`
	//
	// Go field:
	//
	// FirstName
	//
	// JSON field:
	//
	// first_name

	fmt.Println("\n2. Struct JSON tags")

	user2 := struct {
		FirstName string `json:"first_name"`
		Age       int    `json:"age"`
	}{
		FirstName: "Mahdi",
		Age:       25,
	}

	data, err = json.Marshal(user2)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// ============================================================
	// 3. omitempty ⭐⭐⭐
	// ============================================================
	// `omitempty` tells the JSON encoder to omit a field
	// when it has its zero value.
	//
	// Example:
	//
	// int    → 0
	// string → ""
	// bool   → false
	// pointer → nil
	//
	// Example:
	//
	// `json:"age,omitempty"`

	fmt.Println("\n3. omitempty")

	user3 := struct {
		Name string `json:"name"`
		Age  int    `json:"age,omitempty"`
	}{
		Name: "Mahdi",
		Age:  0,
	}

	data, err = json.Marshal(user3)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// Age is omitted because its value is 0.

	// ============================================================
	// 4. json.Unmarshal() ⭐⭐⭐
	// ============================================================
	// `json.Unmarshal()` converts JSON bytes into a Go value.
	//
	// JSON bytes → Go value
	//
	// Function:
	//
	// json.Unmarshal(data, &value)
	//
	// A pointer is passed because Unmarshal needs to modify
	// the destination value.

	fmt.Println("\n4. json.Unmarshal()")

	jsonData := `
	{
		"first_name": "Ali",
		"age": 30
	}
	`

	var user4 struct {
		FirstName string `json:"first_name"`
		Age       int    `json:"age"`
	}

	err = json.Unmarshal(
		[]byte(jsonData),
		&user4,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(user4)
	fmt.Println(user4.FirstName)
	fmt.Println(user4.Age)

	// ============================================================
	// 5. JSON Array → Go Slice ⭐⭐⭐
	// ============================================================
	// A JSON array can be decoded into a Go slice.
	//
	// JSON array:
	//
	// [
	//     {...},
	//     {...}
	// ]
	//
	// Go:
	//
	// []User

	fmt.Println("\n5. JSON array → Go slice")

	jsonData = `
	[
		{
			"first_name": "Ali",
			"age": 20
		},
		{
			"first_name": "Sara",
			"age": 22
		}
	]
	`

	type User struct {
		FirstName string `json:"first_name"`
		Age       int    `json:"age"`
	}

	users := []User{}

	err = json.Unmarshal(
		[]byte(jsonData),
		&users,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(users)

	// ============================================================
	// 6. json.MarshalIndent()
	// ============================================================
	// `json.MarshalIndent()` converts a Go value into
	// formatted JSON.
	//
	// It is useful when JSON needs to be human-readable.
	//
	// json.MarshalIndent(value, prefix, indent)

	fmt.Println("\n6. json.MarshalIndent()")

	user5 := struct {
		FirstName string `json:"first_name"`
		Age       int    `json:"age"`
	}{
		FirstName: "Mahdi",
		Age:       25,
	}

	data, err = json.MarshalIndent(
		user5,
		"",
		"    ",
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// ============================================================
	// 7. json.NewEncoder() ⭐⭐⭐
	// ============================================================
	// `json.NewEncoder()` creates an Encoder that writes
	// JSON directly to an io.Writer.
	//
	// Go value → io.Writer
	//
	// Example:
	//
	// json.NewEncoder(os.Stdout)
	//
	// This is commonly used when sending JSON through
	// HTTP responses.

	fmt.Println("\n7. json.NewEncoder()")

	encoder := json.NewEncoder(os.Stdout)

	err = encoder.Encode(user5)

	if err != nil {
		panic(err)
	}

	// ============================================================
	// 8. json.NewDecoder() ⭐⭐⭐
	// ============================================================
	// `json.NewDecoder()` creates a Decoder that reads
	// JSON from an io.Reader.
	//
	// io.Reader → Go value
	//
	// This is commonly used when reading JSON from
	// HTTP request bodies.

	fmt.Println("\n8. json.NewDecoder()")

	jsonData = `
	{
		"first_name": "Reza",
		"age": 35
	}
	`

	reader := strings.NewReader(jsonData)

	decoder := json.NewDecoder(reader)

	var user6 User

	err = decoder.Decode(&user6)

	if err != nil {
		panic(err)
	}

	fmt.Println(user6)
	fmt.Println(user6.FirstName)
	fmt.Println(user6.Age)

	// ============================================================
	// 9. Marshal vs Unmarshal ⭐⭐⭐
	// ============================================================
	//
	// Marshal:
	//
	// Go value → JSON
	//
	// Unmarshal:
	//
	// JSON → Go value
	//
	// Think:
	//
	// Go
	//  ↓
	// Marshal
	//  ↓
	// JSON
	//
	// JSON
	//  ↓
	// Unmarshal
	//  ↓
	// Go

	fmt.Println("\n9. Marshal vs Unmarshal")

	fmt.Println("Marshal:")
	fmt.Println("Go value → JSON")

	fmt.Println("Unmarshal:")
	fmt.Println("JSON → Go value")

	// ============================================================
	// 10. Encoder vs Decoder ⭐⭐⭐
	// ============================================================
	//
	// Encoder:
	//
	// Go value → io.Writer
	//
	// Decoder:
	//
	// io.Reader → Go value
	//
	// Think:
	//
	// Go value
	//     ↓
	// Encoder
	//     ↓
	// Writer
	//
	// Reader
	//     ↓
	// Decoder
	//     ↓
	// Go value

	fmt.Println("\n10. Encoder vs Decoder")

	fmt.Println("Encoder:")
	fmt.Println("Go value → io.Writer")

	fmt.Println("Decoder:")
	fmt.Println("io.Reader → Go value")

	// ============================================================
	// 11. JSON in Backend Development ⭐⭐⭐
	// ============================================================
	// JSON is commonly used in REST APIs.
	//
	// HTTP Request:
	//
	// Client
	//     ↓
	// JSON
	//     ↓
	// HTTP Request Body
	//     ↓
	// json.Decoder
	//     ↓
	// Go struct
	//
	// HTTP Response:
	//
	// Go struct
	//     ↓
	// json.Encoder
	//     ↓
	// JSON
	//     ↓
	// HTTP Response
	//
	// Example:
	//
	// json.NewDecoder(r.Body).Decode(&user)
	//
	// json.NewEncoder(w).Encode(user)

	fmt.Println("\n11. JSON in Backend Development")

	fmt.Println(`
	HTTP Request:

	Client
	    ↓
	   JSON
	    ↓
	HTTP Request Body
	    ↓
	json.Decoder
	    ↓
	Go struct


	HTTP Response:

	Go struct
	    ↓
	json.Encoder
	    ↓
	   JSON
	    ↓
	HTTP Response
	`)

	// ============================================================
	// 12. JSON → map[string]any ⭐⭐
	// ============================================================
	// JSON objects can be decoded into a map.
	//
	// JSON:
	//
	// {
	//     "name": "Mahdi",
	//     "age": 27
	// }
	//
	// Go:
	//
	// map[string]any
	//
	// This is useful when the JSON structure is dynamic
	// or not known in advance.

	fmt.Println("\n12. JSON → map[string]any")

	jsonData = `
	{
		"name": "Mahdi",
		"age": 27,
		"active": true
	}
	`

	var dataMap map[string]any

	err = json.Unmarshal(
		[]byte(jsonData),
		&dataMap,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(dataMap)
	fmt.Println(dataMap["name"])
	fmt.Println(dataMap["age"])
	fmt.Println(dataMap["active"])

	// ============================================================
	// 13. Unknown JSON Fields ⭐⭐
	// ============================================================
	// By default, json.Unmarshal ignores JSON fields that
	// do not exist in the destination struct.
	//
	// JSON:
	//
	// {
	//     "name": "Mahdi",
	//     "age": 27,
	//     "city": "Baku"
	// }
	//
	// If the Go struct does not contain `city`,
	// the field is ignored.

	fmt.Println("\n13. Unknown JSON fields")

	jsonData = `
	{
		"name": "Mahdi",
		"age": 27,
		"city": "Baku"
	}
	`

	var user7 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	err = json.Unmarshal(
		[]byte(jsonData),
		&user7,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(user7)

	// ============================================================
	// 14. DisallowUnknownFields() ⭐⭐⭐
	// ============================================================
	// `DisallowUnknownFields()` makes the JSON decoder return
	// an error when the JSON contains a field that does not
	// exist in the destination struct.
	//
	// This can be useful for strict API request validation.

	fmt.Println("\n14. DisallowUnknownFields()")

	jsonData = `
	{
		"name": "Mahdi",
		"age": 27,
		"city": "Baku"
	}
	`

	reader = strings.NewReader(jsonData)

	decoder = json.NewDecoder(reader)

	decoder.DisallowUnknownFields()

	var user8 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	err = decoder.Decode(&user8)

	fmt.Println("User:", user8)
	fmt.Println("Error:", err)

	// ============================================================
	// 15. JSON null ⭐⭐⭐
	// ============================================================
	// JSON has a special value:
	//
	// null
	//
	// When decoded into pointer/map/slice/interface fields,
	// null can result in a nil value.

	fmt.Println("\n15. JSON null")

	jsonData = `
	{
		"name": null,
		"age": null
	}
	`

	var user9 struct {
		Name *string `json:"name"`
		Age  *int    `json:"age"`
	}

	err = json.Unmarshal(
		[]byte(jsonData),
		&user9,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(user9)
	fmt.Println(user9.Name)
	fmt.Println(user9.Age)

	// ============================================================
	// 16. Custom JSON Marshaling ⭐⭐⭐
	// ============================================================
	// A type can customize how it is converted into JSON.
	//
	// Implement:
	//
	// MarshalJSON() ([]byte, error)
	//
	// This allows complete control over the JSON output.

	fmt.Println("\n16. Custom JSON Marshaling")

	product := Product{
		Name:  "Laptop",
		Price: 1200,
	}

	data, err = json.Marshal(product)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// ============================================================
	// 17. Custom JSON Unmarshaling ⭐⭐⭐
	// ============================================================
	// A type can customize how JSON is converted into it.
	//
	// Implement:
	//
	// UnmarshalJSON([]byte) error
	//
	// The method receives the JSON bytes and can control
	// how the fields are assigned.

	fmt.Println("\n17. Custom JSON Unmarshaling")

	jsonData = `
	{
		"product_name": "Phone",
		"price": 800
	}
	`

	var product2 Product

	err = json.Unmarshal(
		[]byte(jsonData),
		&product2,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(product2)
	fmt.Println(product2.Name)
	fmt.Println(product2.Price)

	// ============================================================
	// 18. json.RawMessage ⭐⭐⭐
	// ============================================================
	// `json.RawMessage` stores raw JSON data.
	//
	// It is useful when part of a JSON structure is dynamic
	// or needs to be decoded later.
	//
	// Example:
	//
	// {
	//     "type": "user",
	//     "data": {...}
	// }

	fmt.Println("\n18. json.RawMessage")

	jsonData = `
	{
		"type": "user",
		"data": {
			"name": "Mahdi",
			"age": 27
		}
	}
	`

	var message struct {
		Type string          `json:"type"`
		Data json.RawMessage `json:"data"`
	}

	err = json.Unmarshal(
		[]byte(jsonData),
		&message,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Type:", message.Type)
	fmt.Println("Raw data:", string(message.Data))

	// ============================================================
	// 19. json.Number ⭐⭐⭐
	// ============================================================
	// When JSON is decoded into map[string]any,
	// numbers are normally decoded as float64.
	//
	// `Decoder.UseNumber()` makes JSON numbers become
	// `json.Number` instead.
	//
	// This is useful when exact numeric representation matters.

	fmt.Println("\n19. json.Number")

	jsonData = `
	{
		"price": 19.99,
		"quantity": 10
	}
	`

	reader = strings.NewReader(jsonData)

	decoder = json.NewDecoder(reader)

	decoder.UseNumber()

	var productData map[string]any

	err = decoder.Decode(&productData)

	if err != nil {
		panic(err)
	}

	fmt.Println(productData)

	fmt.Printf("price type: %T\n", productData["price"])
	fmt.Printf("quantity type: %T\n", productData["quantity"])

	price := productData["price"].(json.Number)

	fmt.Println("Price:", price)

	// ============================================================
	// 20. json.Valid() ⭐⭐
	// ============================================================
	// `json.Valid()` checks whether a byte slice contains
	// valid JSON.
	//
	// Returns:
	//
	// true  → valid JSON
	// false → invalid JSON

	fmt.Println("\n20. json.Valid()")

	validJSON := []byte(`{"name":"Mahdi","age":27}`)
	invalidJSON := []byte(`{"name":"Mahdi","age":}`)

	fmt.Println("Valid:", json.Valid(validJSON))
	fmt.Println("Invalid:", json.Valid(invalidJSON))

	// ============================================================
	// 21. JSON "-" Tag ⭐⭐
	// ============================================================
	// `json:"-"` completely ignores a struct field
	// during JSON marshaling and unmarshaling.
	//
	// This can be useful for fields that should never
	// appear in API responses.

	fmt.Println("\n21. JSON '-' tag")

	user11 := struct {
		Name     string `json:"name"`
		Password string `json:"-"`
	}{
		Name:     "Mahdi",
		Password: "secret",
	}

	data, err = json.Marshal(user11)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// ============================================================
	// 22. JSON Field Matching ⭐⭐
	// ============================================================
	// JSON fields can match Go struct fields without requiring
	// the exact capitalization.
	//
	// JSON:
	//
	// {
	//     "Name": "Mahdi"
	// }
	//
	// Go:
	//
	// Name string

	fmt.Println("\n22. JSON Field Matching")

	jsonData = `
	{
		"NAME": "Mahdi"
	}
	`

	var user12 struct {
		Name string
	}

	err = json.Unmarshal(
		[]byte(jsonData),
		&user12,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(user12)

	// However, in backend code, use explicit JSON tags rather than relying on this behavior:

	// ============================================================
	// 23. JSON Syntax Errors ⭐⭐⭐
	// ============================================================
	// json.Unmarshal() returns an error when the JSON
	// is syntactically invalid.
	//
	// Always check the returned error.

	fmt.Println("\n23. JSON Syntax Errors")

	invalidData := []byte(`
	{
		"name": "Mahdi",
		"age":
	}
	`)

	var user13 struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	err = json.Unmarshal(
		invalidData,
		&user13,
	)

	fmt.Println("User:", user13)
	fmt.Println("Error:", err)
}
