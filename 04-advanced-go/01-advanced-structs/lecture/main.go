package main

import "fmt"

// ============================================================
// 1. Struct Composition ⭐⭐⭐
// ============================================================
//
// Go does not have classical class inheritance.
//
// Instead, Go prefers composition.
//
// Composition means:
// One struct contains another struct.
//
// Example:
//
// type User struct {
//     Name    string
//     Address Address
// }
//
// User "has an" Address.
//
// This is different from inheritance.
//
// Inheritance:
//     User extends Person
//
// Go:
//     User contains Person
//

type Address struct {
	City    string
	Country string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

// ============================================================
// 2. Struct Embedding ⭐⭐⭐
// ============================================================
//
// Go provides a special form of composition called embedding.
//
// Instead of:
//
// type User struct {
//     Name    string
//     Address Address
// }
//
// We can write:
//
// type User struct {
//     Name string
//     Address
// }
//
// Address is now embedded in User.
//
// An embedded struct allows its fields and methods
// to be promoted to the outer struct.
//

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Salary int
}

// ============================================================
// 4. Method Promotion ⭐⭐⭐
// ============================================================
//
// Methods of an embedded struct are also promoted.
//

type Person2 struct {
	Name string
}

func (p Person2) SayHello() {
	fmt.Println("Hello", p.Name)
}

type Employee2 struct {
	Person2
}

// ============================================================
// 5. Method Shadowing ⭐⭐⭐
// ============================================================
//
// The outer struct can define a method with the same name
// as a promoted method.
//
// The outer method takes priority when called directly.
//

type Person3 struct {
	Name string
}

func (p Person3) SayHello() {
	fmt.Println("Hello from Person")
}

type Employee3 struct {
	Person3
}

func (e Employee3) SayHello() {
	fmt.Println("Hello from Employee")
}

// ============================================================
// 6. Embedded Struct Initialization ⭐⭐⭐
// ============================================================

type Contact struct {
	Email string
	Phone string
}

type Customer struct {
	Name string
	Contact
}

// ============================================================
// 7. Embedding vs Named Field ⭐⭐⭐
// ============================================================

type Address2 struct {
	City string
}

type UserWithNamedAddress struct {
	Address Address2
}

type UserWithEmbeddedAddress struct {
	Address2
}

// ============================================================
// 8. Embedding Interfaces ⭐⭐⭐
// ============================================================
//
// Interfaces can also be embedded.
//
// Example:
//
// type ReadCloser interface {
//     Reader
//     Closer
// }
//
// ReadCloser now requires both:
//
// Read()
// Close()
//

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type ReadCloser interface {
	Reader
	Closer
}

type File struct{}

func (File) Read() {
	fmt.Println("Reading...")
}

func (File) Close() {
	fmt.Println("Closing...")
}

// ============================================================
// 9. Multiple Struct Embedding ⭐⭐⭐
// ============================================================

type Person4 struct {
	Name string
}

type Contact4 struct {
	Email string
}

type User4 struct {
	Person4
	Contact4
}

// ============================================================
// 10. Name Conflicts ⭐⭐⭐
// ============================================================

type Person5 struct {
	Name string
}

type Company struct {
	Name string
}

type Employee5 struct {
	Person5
	Company
}

// ============================================================
// 11. Composition in Backend Development ⭐⭐⭐
// ============================================================

type BaseResponse struct {
	RequestID string `json:"request_id"`
	Status    string `json:"status"`
}

type UserResponse struct {
	BaseResponse

	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ============================================================
// 12. Composition and Clean Architecture ⭐⭐⭐
// ============================================================

type UserRepository interface {
	FindByID(id int) string
}

type UserService struct {
	repo UserRepository
}

type UserRepositoryImpl struct{}

func (UserRepositoryImpl) FindByID(id int) string {
	return "Mahdi"
}

// ============================================================
// MAIN
// ============================================================

func main() {

	// ============================================================
	// 1. Struct Composition ⭐⭐⭐
	// ============================================================

	fmt.Println("1. Struct Composition")

	user := User{
		Name: "Mahdi",
		Age:  27,
		Address: Address{
			City:    "Baku",
			Country: "Azerbaijan",
		},
	}

	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.Address.City)
	fmt.Println(user.Address.Country)

	// ============================================================
	// 2. Struct Embedding ⭐⭐⭐
	// ============================================================

	fmt.Println("\n2. Struct Embedding")

	employee := Employee{
		Person: Person{
			Name: "Mahdi",
			Age:  27,
		},
		Salary: 5000,
	}

	fmt.Println(employee.Person.Name)
	fmt.Println(employee.Person.Age)
	fmt.Println(employee.Salary)

	// ============================================================
	// 3. Field Promotion ⭐⭐⭐
	// ============================================================
	//
	// When a struct is embedded,
	// its fields can be accessed directly
	// through the outer struct.
	//
	// employee.Person.Name
	//
	// can also be accessed as:
	//
	// employee.Name
	//

	fmt.Println("\n3. Field Promotion")

	fmt.Println(employee.Name)
	fmt.Println(employee.Age)
	fmt.Println(employee.Salary)

	// ============================================================
	// 4. Method Promotion ⭐⭐⭐
	// ============================================================

	fmt.Println("\n4. Method Promotion")

	employee2 := Employee2{
		Person2: Person2{
			Name: "Mahdi",
		},
	}

	employee2.SayHello()

	// The following also works:
	//
	// employee2.Person2.SayHello()

	// ============================================================
	// 5. Method Shadowing ⭐⭐⭐
	// ============================================================

	fmt.Println("\n5. Method Shadowing")

	employee3 := Employee3{
		Person3: Person3{
			Name: "Mahdi",
		},
	}

	employee3.SayHello()

	// The Person method still exists:
	//
	// employee3.Person3.SayHello()

	// ============================================================
	// 6. Embedded Struct Initialization ⭐⭐⭐
	// ============================================================

	fmt.Println("\n6. Embedded Struct Initialization")

	customer := Customer{
		Name: "Mahdi",

		Contact: Contact{
			Email: "mahdi@example.com",
			Phone: "123456789",
		},
	}

	fmt.Println(customer.Name)
	fmt.Println(customer.Email)
	fmt.Println(customer.Phone)

	// ============================================================
	// 7. Embedding vs Named Field ⭐⭐⭐
	// ============================================================

	fmt.Println("\n7. Embedding vs Named Field")

	user1 := UserWithNamedAddress{
		Address: Address2{
			City: "Baku",
		},
	}

	user2 := UserWithEmbeddedAddress{
		Address2: Address2{
			City: "Baku",
		},
	}

	// Named field:
	//
	// user1.Address.City

	fmt.Println(user1.Address.City)

	// Embedded field:
	//
	// user2.City

	fmt.Println(user2.City)

	// ============================================================
	// 8. Embedding Interfaces ⭐⭐⭐
	// ============================================================

	fmt.Println("\n8. Embedding Interfaces")

	var file ReadCloser = File{}

	file.Read()
	file.Close()

	// ============================================================
	// 9. Multiple Struct Embedding ⭐⭐⭐
	// ============================================================

	fmt.Println("\n9. Multiple Struct Embedding")

	user4 := User4{
		Person4: Person4{
			Name: "Mahdi",
		},
		Contact4: Contact4{
			Email: "mahdi@example.com",
		},
	}

	fmt.Println(user4.Name)
	fmt.Println(user4.Email)

	// ============================================================
	// 10. Name Conflicts ⭐⭐⭐
	// ============================================================
	//
	// If two embedded structs contain fields with the same name,
	// the field cannot be accessed directly.
	//
	// employee5.Name
	//
	// would be ambiguous.
	//
	// We must explicitly specify the embedded struct.
	//

	fmt.Println("\n10. Name Conflicts")

	employee5 := Employee5{
		Person5: Person5{
			Name: "Mahdi",
		},
		Company: Company{
			Name: "Vardast",
		},
	}

	fmt.Println(employee5.Person5.Name)
	fmt.Println(employee5.Company.Name)

	// ============================================================
	// 11. Composition in Backend Development ⭐⭐⭐
	// ============================================================

	fmt.Println("\n11. Composition in Backend Development")

	response := UserResponse{
		BaseResponse: BaseResponse{
			RequestID: "abc123",
			Status:    "success",
		},

		ID:   1,
		Name: "Mahdi",
	}

	fmt.Println(response.RequestID)
	fmt.Println(response.Status)
	fmt.Println(response.ID)
	fmt.Println(response.Name)

	// ============================================================
	// 12. Composition and Clean Architecture ⭐⭐⭐
	// ============================================================

	fmt.Println("\n12. Composition and Clean Architecture")

	service := UserService{
		repo: UserRepositoryImpl{},
	}

	fmt.Println(service.repo.FindByID(1))

	// ============================================================
	// 13. Composition vs Inheritance ⭐⭐⭐
	// ============================================================
	//
	// Traditional inheritance:
	//
	//     Person
	//       ↑
	//    Employee
	//       ↑
	//   Manager
	//
	// This creates a hierarchy.
	//
	// Go prefers composition:
	//
	//     Employee
	//       |
	//       +── Person
	//       +── Contact
	//       +── Permissions
	//
	// Each component has a specific responsibility.
	//
	// The main idea:
	//
	// "Build complex types by composing smaller types."
	//

	fmt.Println("\n13. Composition vs Inheritance")

	fmt.Println("Go prefers composition over inheritance.")

	// ============================================================
	// 14. Important Mental Model ⭐⭐⭐
	// ============================================================
	//
	// Struct composition:
	//
	// type User struct {
	//     Address Address
	// }
	//
	// Means:
	//
	// User HAS an Address.
	//
	// Struct embedding:
	//
	// type User struct {
	//     Address
	// }
	//
	// Means:
	//
	// User HAS an Address
	//
	// and its fields/methods can be promoted.
	//
	// Remember:
	//
	// Composition
	//     ↓
	// Struct contains another struct
	//
	// Embedding
	//     ↓
	// Special composition with promotion
	//
	// Field promotion
	//     ↓
	// Access embedded fields directly
	//
	// Method promotion
	//     ↓
	// Access embedded methods directly
	//
	// Interfaces
	//     ↓
	// Can also be composed using embedding
	//

	fmt.Println("\n14. Important Mental Model")

	fmt.Println(`
Composition:
    Struct contains another struct.

Embedding:
    Special form of composition.

Field promotion:
    Embedded fields can be accessed directly.

Method promotion:
    Embedded methods can be accessed directly.

Interface embedding:
    Multiple interfaces can be composed.

Go:
    Composition over inheritance.
`)
}
