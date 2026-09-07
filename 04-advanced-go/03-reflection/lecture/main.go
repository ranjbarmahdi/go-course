package main

import (
	"fmt"
	"reflect"
)

/*
Reflection allows us to inspect and work with values
at runtime, even when we don't know their concrete
type at compile time.

Reflection allows us to ask:

- What type is this?
- What kind of value is this?
- What is its value?
- How many fields does this struct have?
- What is the name of this field?
- What is the type of this field?
- What are its struct tags?
- Can I change this value?
- Does this value have a method?

The `reflect` package provides reflection functionality.
*/

func main() {

	// ============================================================
	// 1. any ⭐⭐⭐
	// ============================================================
	// `any` is an alias for `interface{}`.
	//
	// It means the variable can hold a value of any type.
	//
	// Example:
	//
	//     var value any = "Mahdi"
	//
	// The variable has type `any`,
	// but the actual value stored inside it is a string.
	//
	// Think:
	//
	//     any
	//      │
	//      ▼
	//   "Mahdi"
	//    string

	fmt.Println("1. any")

	var value any = "Mahdi"

	fmt.Println(value)

	// We can later store another type in the same variable.

	value = 27

	fmt.Println(value)

	// Reflection becomes useful when we don't know
	// what concrete type is stored inside an `any` value.

	// ============================================================
	// 2. reflect.TypeOf() ⭐⭐⭐
	// ============================================================
	// `reflect.TypeOf()` returns the runtime type of a value.
	//
	// Example:
	//
	//     reflect.TypeOf(value)
	//
	// returns a `reflect.Type`.
	//
	// If value contains a string:
	//
	//     string
	//
	// If value contains an int:
	//
	//     int

	fmt.Println("\n2. reflect.TypeOf()")

	var value2 any = "Mahdi"

	fmt.Println(value2)
	fmt.Println(reflect.TypeOf(value2))

	value2 = 27

	fmt.Println(value2)
	fmt.Println(reflect.TypeOf(value2))

	// ============================================================
	// 3. reflect.ValueOf() ⭐⭐⭐
	// ============================================================
	// `reflect.ValueOf()` returns a `reflect.Value`
	// representing the actual runtime value.
	//
	// `reflect.TypeOf()`:
	//
	//     Gives information about the type.
	//
	// `reflect.ValueOf()`:
	//
	//     Gives a reflection representation of the value.
	//
	// Think:
	//
	//                  "Mahdi"
	//                     │
	//              ┌──────┴──────┐
	//              │             │
	//          TypeOf()       ValueOf()
	//              │             │
	//              ▼             ▼
	//           string        "Mahdi"

	fmt.Println("\n3. reflect.ValueOf()")

	var value3 any = "Mahdi"

	t := reflect.TypeOf(value3)
	v := reflect.ValueOf(value3)

	fmt.Println("Type:", t)
	fmt.Println("Value:", v)

	// ============================================================
	// 4. Type vs Kind ⭐⭐⭐
	// ============================================================
	// Reflection distinguishes between:
	//
	//     Type
	//     Kind
	//
	// Type is the specific Go type.
	//
	// Kind is the general category of the type.
	//
	// Example:
	//
	//     type UserID int
	//
	// The Type is:
	//
	//     main.UserID
	//
	// But the Kind is:
	//
	//     int
	//
	// Think:
	//
	//     Type → specific
	//
	//     Kind → general category

	fmt.Println("\n4. Type vs Kind")

	var value4 any = "Mahdi"

	t2 := reflect.TypeOf(value4)

	fmt.Println("Type:", t2)
	fmt.Println("Kind:", t2.Kind())

	type UserID int

	var id UserID = 10

	t3 := reflect.TypeOf(id)

	fmt.Println("Type:", t3)
	fmt.Println("Kind:", t3.Kind())

	// ============================================================
	// 5. Common Kinds
	// ============================================================
	// `reflect.Kind` represents the general category
	// of a Go value.
	//
	// Common kinds include:
	//
	//     Bool
	//     Int
	//     Int8
	//     Int16
	//     Int32
	//     Int64
	//
	//     Uint
	//     Uint8
	//     Uint16
	//     Uint32
	//     Uint64
	//
	//     Float32
	//     Float64
	//
	//     String
	//
	//     Array
	//     Slice
	//     Map
	//     Struct
	//
	//     Pointer
	//     Interface
	//     Func
	//     Chan
	//
	// Example:
	//
	//     []int
	//
	// has:
	//
	//     Type → []int
	//     Kind → slice
	//
	// Another example:
	//
	//     map[string]int
	//
	// has:
	//
	//     Type → map[string]int
	//     Kind → map

	fmt.Println("\n5. Common Kinds")

	values := []any{
		"Mahdi",
		27,
		true,
		3.14,
		[]int{1, 2, 3},
		map[string]int{
			"age": 27,
		},
	}

	for _, value := range values {

		t := reflect.TypeOf(value)

		fmt.Println("Type:", t)
		fmt.Println("Kind:", t.Kind())
		fmt.Println()
	}

	// ============================================================
	// 6. Why Kind() is Useful ⭐⭐⭐
	// ============================================================
	// `Kind()` is useful when we don't know
	// the concrete type of a value.
	//
	// We can inspect its kind and make a decision.
	//
	// Example:
	//
	//     string → string logic
	//     int    → int logic
	//     bool   → bool logic
	//     slice  → slice logic
	//     map    → map logic
	//     struct → struct logic
	//
	// This is useful when building generic utilities
	// that need to work with different runtime types.

	fmt.Println("\n6. Why Kind() is useful")

	Inspect("Mahdi")
	Inspect(27)
	Inspect(true)
	Inspect([]int{1, 2, 3})
	Inspect(map[string]int{"age": 27})

	// ============================================================
	// 7. Struct Reflection ⭐⭐⭐
	// ============================================================
	// Reflection becomes especially useful with structs.
	//
	// Suppose we have:
	//
	//     User
	//       ├── Name string
	//       └── Age  int
	//
	// `reflect.TypeOf()` allows us to inspect
	// the structure of the type.

	fmt.Println("\n7. Struct reflection")

	user := User{
		Name: "Mahdi",
		Age:  27,
	}

	t4 := reflect.TypeOf(user)

	fmt.Println("Type:", t4)
	fmt.Println("Kind:", t4.Kind())
	fmt.Println("Name:", t4.Name())

	// ============================================================
	// 8. Number of Struct Fields ⭐⭐⭐
	// ============================================================
	// `NumField()` returns the number of fields
	// in a struct.
	//
	// For:
	//
	//     type User struct {
	//         Name string
	//         Age  int
	//     }
	//
	// NumField() returns:
	//
	//     2

	fmt.Println("\n8. Number of struct fields")

	fmt.Println("NumField:", t4.NumField())

	// We can iterate through all fields.

	for i := 0; i < t4.NumField(); i++ {

		field := t4.Field(i)

		fmt.Println("Name:", field.Name)
		fmt.Println("Type:", field.Type)
		fmt.Println()
	}

	// ============================================================
	// 9. FieldByName()
	// ============================================================
	// `FieldByName()` allows us to find a struct field
	// using its name.
	//
	// This is usually more readable than:
	//
	//     Field(0)
	//
	// because we can explicitly say which field we want.

	fmt.Println("\n9. FieldByName()")

	nameField, ok := t4.FieldByName("Name")

	if !ok {
		fmt.Println("Field not found")
		return
	}

	fmt.Println("Name:", nameField.Name)
	fmt.Println("Type:", nameField.Type)

	// ============================================================
	// 10. Struct Tags ⭐⭐⭐
	// ============================================================
	// Struct tags are metadata attached to struct fields.
	//
	// Example:
	//
	//     Name string `json:"user_name"`
	//
	// Reflection can read these tags.
	//
	// This is important for understanding libraries such as:
	//
	// - encoding/json
	// - ORM libraries
	// - validation libraries
	// - dependency injection libraries
	//
	// Many libraries inspect struct tags at runtime.

	fmt.Println("\n10. Struct tags")

	user2 := UserWithTags{
		Name: "Mahdi",
		Age:  27,
	}

	t5 := reflect.TypeOf(user2)

	fmt.Println("Type:", t5)

	field := t5.Field(0)

	fmt.Println("Field name:", field.Name)
	fmt.Println("Field tag:", field.Tag)

	jsonTag := field.Tag.Get("json")

	fmt.Println("JSON tag:", jsonTag)

	// We can inspect the second field too.

	field2 := t5.Field(1)

	fmt.Println("Field name:", field2.Name)
	fmt.Println("JSON tag:", field2.Tag.Get("json"))

	// ============================================================
	// 11. Inspecting Actual Field Values ⭐⭐⭐
	// ============================================================
	// `reflect.Type` gives us metadata about the struct.
	//
	// `reflect.Value` gives us access to the actual values.
	//
	// Example:
	//
	//     User{
	//         Name: "Mahdi",
	//         Age:  27,
	//     }
	//
	// Type:
	//
	//     Name → string
	//     Age  → int
	//
	// Value:
	//
	//     Name → "Mahdi"
	//     Age  → 27

	fmt.Println("\n11. Inspecting actual field values")

	user3 := User{
		Name: "Mahdi",
		Age:  27,
	}

	v1 := reflect.ValueOf(user3)

	fmt.Println("Value:", v1)

	fmt.Println("Name:", v1.Field(0))
	fmt.Println("Age:", v1.Field(1))

	fmt.Println("Name type:", v1.Field(0).Type())
	fmt.Println("Age type:", v1.Field(1).Type())

	// We can also use FieldByName() with reflect.Value.

	fmt.Println("Name:", v1.FieldByName("Name"))
	fmt.Println("Age:", v1.FieldByName("Age"))

	// ============================================================
	// 12. Reading Values by Kind ⭐⭐⭐
	// ============================================================
	// `reflect.Value` provides methods to extract
	// the underlying value.
	//
	// Common methods:
	//
	//     Int()
	//     Uint()
	//     Float()
	//     String()
	//     Bool()
	//
	// You must use the correct method for the value's kind.
	//
	// Example:
	//
	//     int     → Int()
	//     string  → String()
	//     bool    → Bool()
	//     float64 → Float()

	fmt.Println("\n12. Reading values by kind")

	v2 := reflect.ValueOf(27)

	fmt.Println("Kind:", v2.Kind())
	fmt.Println("Value:", v2.Int())

	v3 := reflect.ValueOf("Mahdi")

	fmt.Println("Kind:", v3.Kind())
	fmt.Println("Value:", v3.String())

	v4 := reflect.ValueOf(true)

	fmt.Println("Kind:", v4.Kind())
	fmt.Println("Value:", v4.Bool())

	v5 := reflect.ValueOf(3.14)

	fmt.Println("Kind:", v5.Kind())
	fmt.Println("Value:", v5.Float())

	// ============================================================
	// 13. Reflection + Pointers ⭐⭐⭐
	// ============================================================
	// If we pass a pointer to `reflect.ValueOf()`,
	// the reflected value has kind `Pointer`.
	//
	// Example:
	//
	//     &user
	//
	// means:
	//
	//     pointer → User
	//
	// We can use `Elem()` to access the value
	// stored behind the pointer.
	//
	// Think:
	//
	//     reflect.ValueOf(&user)
	//             │
	//             ▼
	//           *User
	//             │
	//           Elem()
	//             │
	//             ▼
	//            User

	fmt.Println("\n13. Reflection + pointers")

	user4 := User{
		Name: "Mahdi",
		Age:  27,
	}

	v6 := reflect.ValueOf(&user4)

	fmt.Println("Kind:", v6.Kind())
	fmt.Println("Value:", v6)

	v6 = v6.Elem()

	fmt.Println("After Elem():", v6)
	fmt.Println("Kind:", v6.Kind())

	// ============================================================
	// 14. Can Reflection Modify Values? ⭐⭐⭐
	// ============================================================
	// Reflection can modify values.
	//
	// But the reflected value must be settable.
	//
	// We can check this with:
	//
	//     CanSet()
	//
	// A common pattern is:
	//
	//     reflect.ValueOf(&user).Elem()
	//
	// Why?
	//
	// `&user` gives reflection access to the address
	// of the original value.
	//
	// `Elem()` gets the actual value behind the pointer.
	//
	// The resulting value can be set.

	fmt.Println("\n14. Can reflection modify values?")

	user5 := User{
		Name: "Mahdi",
		Age:  27,
	}

	v7 := reflect.ValueOf(&user5).Elem()

	name := v7.FieldByName("Name")

	fmt.Println("Before:", user5.Name)
	fmt.Println("CanSet:", name.CanSet())

	name.SetString("Ali")

	fmt.Println("After:", user5.Name)

	// ============================================================
	// 15. Important Reflection Mental Model ⭐⭐⭐
	// ============================================================
	//
	// Reflection has two main sides:
	//
	//                    any
	//                     │
	//          ┌──────────┴──────────┐
	//          │                     │
	//      TypeOf()               ValueOf()
	//          │                     │
	//          ▼                     ▼
	//    reflect.Type          reflect.Value
	//          │                     │
	//          │                     ├── Field()
	//          │                     ├── FieldByName()
	//          │                     ├── Elem()
	//          │                     ├── Kind()
	//          │                     └── CanSet()
	//          │
	//          ├── Name()
	//          ├── Kind()
	//          ├── NumField()
	//          ├── Field()
	//          └── StructField.Tag
	//
	//
	// Think:
	//
	//     reflect.Type
	//         ↓
	//     "What is this?"
	//
	//     reflect.Value
	//         ↓
	//     "What value does it contain?"
	//
	// Struct example:
	//
	//     User
	//       │
	//       ├── Name string `json:"user_name"`
	//       └── Age  int    `json:"user_age"`
	//
	// Type information:
	//
	//     Name
	//     Type
	//     Kind
	//     Tags
	//     Number of fields
	//
	// Value information:
	//
	//     "Mahdi"
	//     27
	//
	// ============================================================
	// 16. When Reflection Is Useful ⭐⭐⭐
	// ============================================================
	//
	// Reflection is useful when code needs to work with
	// values whose types are not known at compile time.
	//
	// Common examples:
	//
	//     JSON serialization
	//     ORM libraries
	//     Validation libraries
	//     Dependency injection
	//     Generic utilities
	//     Framework internals
	//
	// For example, when a JSON library receives:
	//
	//     User{
	//         Name: "Mahdi",
	//         Age: 27,
	//     }
	//
	// it can inspect the struct and its tags:
	//
	//     Name → json:"user_name"
	//     Age  → json:"user_age"
	//
	// and use that information when converting
	// the value into JSON.
	//
	// ============================================================
	// Important Rule
	// ============================================================
	//
	// Reflection is powerful, but it should not be used
	// everywhere.
	//
	// Prefer normal Go code when the type is already known.
	//
	// Use reflection when runtime type inspection
	// is actually required.
}

type User struct {
	Name string
	Age  int
}

type UserWithTags struct {
	Name string `json:"user_name"`
	Age  int    `json:"user_age"`
}

func Inspect(value any) {

	kind := reflect.TypeOf(value).Kind()

	switch kind {

	case reflect.String:
		fmt.Println("This is a string")

	case reflect.Int:
		fmt.Println("This is an int")

	case reflect.Bool:
		fmt.Println("This is a bool")

	case reflect.Slice:
		fmt.Println("This is a slice")

	case reflect.Map:
		fmt.Println("This is a map")

	case reflect.Struct:
		fmt.Println("This is a struct")
	}
}
