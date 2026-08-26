/*
Problem:

Create a function:

func JSONToMap(data []byte) (map[string]any, error)


The function converts a JSON object into a
map[string]any using json.Unmarshal().


Requirements:

- Create the JSONToMap() function.
- Create a map[string]any variable.
- Use json.Unmarshal().
- Remember that json.Unmarshal() needs a pointer
  to the destination map.
- Return the map.
- Return the error if unmarshaling fails.

- In main(), create this JSON:

{
    "name": "Mahdi",
    "age": 27,
    "active": true
}

- Call JSONToMap().
- Print the complete map.
- Print:
    name
    age
    active


Expected output:

map[active:true age:27 name:Mahdi]
Mahdi
27
true
*/

package main

import (
	"encoding/json"
	"fmt"
)

func JSONToMap(data []byte) (map[string]any, error) {
	var dataMap map[string]any
	err := json.Unmarshal(data, &dataMap)
	return dataMap, err
}

func main() {
	jsonString := `
		{
			"name": "Mahdi",
			"age": 27,
			"active": true
		}
	`

	if res, err := JSONToMap([]byte(jsonString)); err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println(res)
		fmt.Println(res["name"])
		fmt.Println(res["age"])
		fmt.Println(res["active"])
	}

}
