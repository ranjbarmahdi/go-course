/*
Problem:

Create a function:

func ParseEvent(data []byte) (Event, error)


The function decodes JSON using json.Unmarshal()
while keeping one field as json.RawMessage.


Requirements:

- Create an Event struct with:

    Type string
    Data json.RawMessage

- Add JSON tags:

    Type → "type"
    Data → "data"

- Create ParseEvent().

- Use json.Unmarshal().

- Return the Event.
- Return the error.

- In main(), create this JSON:

{
    "type": "user",
    "data": {
        "name": "Mahdi",
        "age": 27
    }
}

- Call ParseEvent().

- Print the event Type.

- Print the raw Data.

Important:

json.RawMessage allows you to keep
JSON data without immediately decoding it
into a specific Go type.

Expected output:

Type: user

Raw data:
{
    "name": "Mahdi",
    "age": 27
}
*/

package main

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func ParseEvent(data []byte) (Event, error) {
	event := Event{}

	err := json.Unmarshal(data, &event)

	return event, err
}

func main() {
	jsonString := `
	{
		"type": "user",
		"data": {
			"name": "Mahdi",
			"age": 27
		}
	}
	`

	res, err := ParseEvent([]byte(jsonString))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Type:", res.Type)
	fmt.Println("Raw data:", string(res.Data))
}
