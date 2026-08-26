package main

import "fmt"

// ============================================================
// Go Generics
// ============================================================
//
// Generics allow us to write functions and data structures
// that work with multiple types while keeping type safety.
//
// Instead of writing:
//
// func PrintInt(value int)
// func PrintString(value string)
// func PrintFloat(value float64)
//
// We can write:
//
// func Print[T any](value T)
//
// and use it with different types.
//
// Generics are especially useful for:
// - Reusable functions
// - Reusable data structures
// - Collections
// - Utility functions
// - Backend helper types
// ============================================================

func main() {

	// ============================================================
	// 1. Why Generics ⭐⭐⭐
	// ============================================================
	//
	// Without generics, we may need separate functions
	// for different types.
	//
	// Generics allow one function to work with multiple types.

	fmt.Println("1. Why Generics")

	fmt.Println("Generics allow reusable, type-safe code.")

	// ============================================================
	// 2. Generic Function ⭐⭐⭐
	// ============================================================
	//
	// Syntax:
	//
	// func Print[T any](value T)
	//
	// T is a type parameter.
	//
	// `any` means T can be any type.

	fmt.Println("\n2. Generic Function")

	Print(100)
	Print("Mahdi")
	Print(3.14)
	Print(true)

	// ============================================================
	// 3. Type Parameters
	// ============================================================
	//
	// T is a type parameter.
	//
	// The compiler determines what T is based on
	// the value passed to the function.

	fmt.Println("\n3. Type Parameters")

	Print("Hello")

	// Here:
	//
	// T = string
	//
	// So the compiler effectively treats it like:
	//
	// Print[string]("Hello")

	// ============================================================
	// 4. Generic Function With Multiple Type Parameters
	// ============================================================
	//
	// A generic function can have multiple type parameters.
	//
	// Example:
	//
	// func PrintPair[T any, U any](a T, b U)

	fmt.Println("\n4. Multiple Type Parameters")

	PrintPair(100, "Mahdi")
	PrintPair("Age", 27)

	// ============================================================
	// 5. Generic Return Values
	// ============================================================
	//
	// A generic function can return the same generic type.

	fmt.Println("\n5. Generic Return Values")

	number := Identity(100)
	name := Identity("Mahdi")

	fmt.Println(number)
	fmt.Println(name)

	// Identity(100):
	//
	// T = int
	//
	// Return type = int
	//
	// Identity("Mahdi"):
	//
	// T = string
	//
	// Return type = string

	// ============================================================
	// 6. Type Inference ⭐⭐⭐
	// ============================================================
	//
	// Go can usually infer the type parameter automatically.
	//
	// We do not need to explicitly write:
	//
	// Identity[int](100)
	//
	// We can simply write:
	//
	// Identity(100)

	fmt.Println("\n6. Type Inference")

	value := Identity(100)

	fmt.Println(value)

	// Explicit type parameter is also possible:

	value2 := Identity[int](100)

	fmt.Println(value2)

	// ============================================================
	// 7. Generic Slice Function
	// ============================================================
	//
	// Generic functions can work with slices of different types.

	fmt.Println("\n7. Generic Slice Function")

	numbers := []int{10, 20, 30}
	names := []string{"Ali", "Mahdi", "Sara"}

	fmt.Println(First(numbers))
	fmt.Println(First(names))

	// ============================================================
	// 8. comparable Constraint ⭐⭐⭐
	// ============================================================
	//
	// Some operations are not available for every type.
	//
	// For example:
	//
	// a == b
	//
	// cannot be used with every possible Go type.
	//
	// `comparable` restricts T to types that can be compared
	// using == and !=.

	fmt.Println("\n8. comparable Constraint")

	fmt.Println(Contains([]int{1, 2, 3, 4}, 3))
	fmt.Println(Contains([]string{"Go", "Java", "Python"}, "Go"))

	// ============================================================
	// 9. Generic Search Function
	// ============================================================
	//
	// A very common generic utility:
	//
	// Search a slice for a value.
	//
	// Because we use ==, T must satisfy comparable.

	fmt.Println("\n9. Generic Search")

	fmt.Println(Contains([]int{10, 20, 30}, 20))
	fmt.Println(Contains([]string{"Ali", "Mahdi"}, "Reza"))

	// ============================================================
	// 10. Generic Constraints ⭐⭐⭐
	// ============================================================
	//
	// A constraint tells Go what types are allowed for T.
	//
	// Example:
	//
	// T comparable
	//
	// means T must be comparable.
	//
	// Another example:
	//
	// T Number
	//
	// means T must satisfy our custom Number constraint.

	fmt.Println("\n10. Generic Constraints")

	fmt.Println(Add(10, 20))
	fmt.Println(Add(1.5, 2.5))

	// ============================================================
	// 11. Custom Constraints ⭐⭐⭐
	// ============================================================
	//
	// We can create our own constraint.
	//
	// Example:
	//
	// type Number interface {
	//     int | int64 | float64
	// }
	//
	// This means T can only be:
	//
	// int
	// int64
	// float64

	fmt.Println("\n11. Custom Constraints")

	fmt.Println(Add(10, 20))
	fmt.Println(Add(int64(10), int64(20)))
	fmt.Println(Add(1.5, 2.5))

	// ============================================================
	// 12. Type Sets
	// ============================================================
	//
	// A constraint such as:
	//
	// type Number interface {
	//     int | int64 | float64
	// }
	//
	// defines a type set.
	//
	// T must belong to that set.
	//
	// This allows operations that are valid
	// for all types in the constraint.

	fmt.Println("\n12. Type Sets")

	fmt.Println("Number constraint allows int, int64 and float64.")

	// ============================================================
	// 13. Approximation ~ ⭐⭐⭐
	// ============================================================
	//
	// The `~` means:
	//
	// The underlying type must match.
	//
	// Example:
	//
	// type UserID int
	//
	// UserID is NOT exactly int.
	//
	// But its underlying type is int.
	//
	// Therefore:
	//
	// ~int
	//
	// accepts UserID.

	fmt.Println("\n13. Approximation ~")

	type UserID int

	userID := UserID(100)

	fmt.Println(Sum([]UserID{
		userID,
		200,
	}))

	// ============================================================
	// 14. Generic Structs ⭐⭐⭐
	// ============================================================
	//
	// Structs can also use type parameters.
	//
	// Example:
	//
	// type Box[T any] struct {
	//     Value T
	// }

	fmt.Println("\n14. Generic Structs")

	intBox := Box[int]{
		Value: 100,
	}

	stringBox := Box[string]{
		Value: "Hello",
	}

	fmt.Println(intBox)
	fmt.Println(stringBox)

	// ============================================================
	// 15. Generic Data Structures ⭐⭐⭐
	// ============================================================
	//
	// We can build reusable data structures using generics.
	//
	// Example:
	//
	// Stack[int]
	// Stack[string]
	//
	// The same Stack implementation can work with
	// different types.

	fmt.Println("\n15. Generic Data Structures")

	intStack := Stack[int]{}

	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Println(intStack)

	value3, ok := intStack.Pop()

	fmt.Println(value3)
	fmt.Println(ok)

	// ============================================================
	// 16. Generic Methods
	// ============================================================
	//
	// A generic type can have methods.
	//
	// Example:
	//
	// type Box[T any] struct {
	//     Value T
	// }
	//
	// func (b Box[T]) Get() T {
	//     return b.Value
	// }

	fmt.Println("\n16. Generic Methods")

	box := Box2[int]{
		Value: 100,
	}

	fmt.Println(box.Get())

	// ============================================================
	// 17. Generic Type Aliases vs Defined Types
	// ============================================================
	//
	// A defined type:
	//
	// type UserID int
	//
	// creates a new distinct type.
	//
	// It has underlying type int.
	//
	// This is important when using constraints
	// with the `~` operator.

	fmt.Println("\n17. Defined Types")

	type ProductID int

	var productID ProductID = 10

	fmt.Println(productID)

	// ============================================================
	// 18. Generics vs any ⭐⭐⭐
	// ============================================================
	//
	// `any` means:
	//
	// "I accept any type."
	//
	// But when using any, we lose compile-time knowledge
	// about the actual type.
	//
	// Generics preserve the relationship between
	// input and output types.

	fmt.Println("\n18. Generics vs any")

	fmt.Println("any accepts any value.")
	fmt.Println("Generics preserve type information.")

	// Example:
	//
	// func IdentityAny(value any) any
	//
	// The caller gets an any value.
	//
	// With:
	//
	// func Identity[T any](value T) T
	//
	// the returned value has the same concrete type as T.

	// ============================================================
	// 19. Generics vs Interfaces ⭐⭐⭐
	// ============================================================
	//
	// Interfaces are useful when we want to work with
	// different types through shared behavior.
	//
	// Generics are useful when we want reusable code
	// while preserving type information.
	//
	// Example:
	//
	// Interface:
	//
	// type Repository interface {
	//     FindByID(id int) User
	// }
	//
	// Generic:
	//
	// type Result[T any] struct {
	//     Data  T
	//     Error error
	// }
	//
	// Both are useful, but they solve different problems.

	fmt.Println("\n19. Generics vs Interfaces")

	fmt.Println("Interfaces describe behavior.")
	fmt.Println("Generics provide type-parameterized code.")

	// ============================================================
	// 20. Generics in Backend Development ⭐⭐⭐
	// ============================================================
	//
	// Generics can be useful in backend systems for:
	//
	// - Pagination
	// - API results
	// - Reusable utilities
	// - Generic collections
	// - Generic service helpers
	// - Generic response structures
	//
	// Example:
	//
	// Page[User]
	// Page[Product]
	//
	// Same structure, different item type.

	fmt.Println("\n20. Generics in Backend Development")

	usersPage := Page[User]{
		Items: []User{
			{Name: "Mahdi", Age: 27},
			{Name: "Ali", Age: 30},
		},
		Total: 2,
	}

	fmt.Println(usersPage)

	// ============================================================
	// 21. Generic Result Type ⭐⭐⭐
	// ============================================================
	//
	// A generic Result type can represent:
	//
	// Data + Error
	//
	// Example:
	//
	// Result[User]
	// Result[Product]
	//
	// This can be useful in application/service layers.

	fmt.Println("\n21. Generic Result Type")

	result := Result[User]{
		Data: User{
			Name: "Mahdi",
			Age:  27,
		},
	}

	fmt.Println(result.Data)

	// ============================================================
	// 22. Generic Functions + Constraints
	// ============================================================
	//
	// Generics become especially powerful when combined
	// with constraints.
	//
	// Example:
	//
	// Sum[T NumberApproximation]
	//
	// allows us to sum:
	//
	// int
	// int64
	// float64
	//
	// and also custom types whose underlying type
	// is one of these types.

	fmt.Println("\n22. Generic Functions + Constraints")

	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Sum([]float64{1.5, 2.5, 3.0}))

	// ============================================================
	// 23. More Realistic Backend Example ⭐⭐⭐
	// ============================================================
	//
	// Imagine an API that returns paginated data.
	//
	// Instead of creating:
	//
	// UserPage
	// ProductPage
	// OrderPage
	//
	// we can create one generic type:
	//
	// Page[T]
	//
	// Then:
	//
	// Page[User]
	// Page[Product]
	// Page[Order]

	fmt.Println("\n23. Realistic Backend Example")

	productPage := Page[Product]{
		Items: []Product{
			{Name: "Laptop", Price: 1200},
			{Name: "Phone", Price: 800},
		},
		Total: 2,
	}

	fmt.Println(productPage)

	// ============================================================
	// 24. Important Mental Model ⭐⭐⭐
	// ============================================================
	//
	// Remember these concepts:
	//
	// Generic function:
	//
	// func Print[T any](value T)
	//
	// T:
	//     Type parameter
	//
	// any:
	//     Any type
	//
	// Constraint:
	//     Restricts allowed types
	//
	// comparable:
	//     Allows == and !=
	//
	// Custom constraint:
	//     Defines allowed types
	//
	// ~:
	//     Matches types by underlying type
	//
	// Generic struct:
	//
	// type Box[T any] struct {
	//     Value T
	// }
	//
	// Type inference:
	//     Go determines T automatically.
	//
	// Main idea:
	//
	// Generics
	//     ↓
	// Reusable + type-safe code
	//
	// ============================================================

	fmt.Println("\n24. Important Mental Model")

	fmt.Println(`
Generics:
    Reusable and type-safe code.

Type parameter:
    T

any:
    Allows any type.

Constraint:
    Restricts which types T can be.

comparable:
    Allows types that support == and !=.

Custom constraint:
    Defines your own allowed type set.

~:
    Matches types by underlying type.

Generic struct:
    Reusable struct for different types.

Type inference:
    Go automatically determines T.

Backend:
    Useful for reusable utilities,
    pagination, results, collections,
    and shared application components.
`)
}

// ============================================================
// Generic Functions
// ============================================================

func Print[T any](value T) {
	fmt.Println(value)
}

func PrintPair[T any, U any](a T, b U) {
	fmt.Println(a)
	fmt.Println(b)
}

func Identity[T any](value T) T {
	return value
}

func First[T any](values []T) T {
	return values[0]
}

// ============================================================
// comparable Constraint
// ============================================================

func Contains[T comparable](values []T, target T) bool {

	for _, value := range values {

		if value == target {
			return true
		}
	}

	return false
}

// ============================================================
// Custom Constraint
// ============================================================

type Number interface {
	int | int64 | float64
}

func Add[T Number](a T, b T) T {
	return a + b
}

// ============================================================
// Approximation Constraint
// ============================================================
//
// `~` allows types whose underlying type matches.
//
// Example:
//
// type UserID int
//
// UserID has underlying type int.
//
// Therefore UserID satisfies:
//
// ~int
//

type NumberApproximation interface {
	~int | ~int64 | ~float64
}

func Sum[T NumberApproximation](values []T) T {

	var total T

	for _, value := range values {
		total += value
	}

	return total
}

// ============================================================
// Generic Struct
// ============================================================

type Box[T any] struct {
	Value T
}

// ============================================================
// Generic Data Structure
// ============================================================

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() (T, bool) {

	if len(s.values) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.values) - 1

	value := s.values[lastIndex]

	s.values = s.values[:lastIndex]

	return value, true
}

// ============================================================
// Generic Method
// ============================================================

type Box2[T any] struct {
	Value T
}

func (b Box2[T]) Get() T {
	return b.Value
}

// ============================================================
// Generic Pagination
// ============================================================

type Page[T any] struct {
	Items []T
	Total int
}

// ============================================================
// Generic Result
// ============================================================

type Result[T any] struct {
	Data  T
	Error error
}

// ============================================================
// Backend Example Types
// ============================================================

type User struct {
	Name string
	Age  int
}

type Product struct {
	Name  string
	Price float64
}
