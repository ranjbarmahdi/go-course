/*
Problem:

Create a function:

func ParseProduct(data []byte) (map[string]json.Number, error)


The function converts JSON numbers into json.Number
instead of the default float64.


Requirements:

- Create ParseProduct().

- Create a map:

    map[string]json.Number

- Use json.Decoder instead of json.Unmarshal().

- Create a decoder from the JSON data.

- Call:

    decoder.UseNumber()

- Decode the JSON into the map.

- Return the map.
- Return the error.

- In main(), create this JSON:

{
    "price": 19.99,
    "quantity": 10
}

- Call ParseProduct().

- Print the complete map.

- Print the type of "price".

- Print the type of "quantity".

- Convert the price into a float64 using:

    Float64()

- Print the price.


Important:

When JSON is decoded into:

    map[string]any

numbers are normally decoded as float64.

But:

    decoder.UseNumber()

causes JSON numbers to be kept as json.Number.

json.Number is a string type that represents
a JSON number.

You can later convert it using:

    Int64()

or:

    Float64()


Expected output:

map[price:19.99 quantity:10]

price type: json.Number

quantity type: json.Number

Price: 19.99
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ParseProduct(data []byte) (map[string]json.Number, error) {
	productMap := map[string]json.Number{}
	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	err := decoder.Decode(&productMap)
	return productMap, err
}

func main() {
	jsonString :=
		`
	{
		"price": 19.99,
		"quantity": 10
	}
	`

	if res, err := ParseProduct([]byte(jsonString)); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
		fmt.Printf("price type: %T\n", res["price"])
		fmt.Printf("quantity type: %T\n", res["quantity"])
	}
}
