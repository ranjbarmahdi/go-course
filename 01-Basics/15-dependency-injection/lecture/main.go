package main

import "fmt"

// ============================================================
// 1. Interface
// ============================================================
// The service depends on the interface instead of a concrete
// repository implementation.
//
// This makes the service easier to test and replace.

type UserRepository interface {
	FindUser(id int) string
}

// ============================================================
// 2. Concrete Implementation
// ============================================================
// PostgresUserRepository implements UserRepository because it
// provides the FindUser method.

type PostgresUserRepository struct{}

func (p PostgresUserRepository) FindUser(id int) string {
	return "Mahdi"
}

// ============================================================
// 3. Service
// ============================================================
// The service depends on the UserRepository interface.
//
// It does not need to know whether the repository uses:
// - PostgreSQL
// - MySQL
// - MongoDB
// - In-memory data
// - A mock for testing

type UserService struct {
	repository UserRepository
}

// ============================================================
// 4. Constructor Injection
// ============================================================
// The repository is provided from outside the service.
//
// This is called Dependency Injection.
//
// The dependency is injected through the constructor.

func NewUserService(repository UserRepository) *UserService {

	return &UserService{
		repository: repository,
	}
}

// ============================================================
// 5. Using the Dependency
// ============================================================

func (s UserService) GetUser(id int) {

	user := s.repository.FindUser(id)

	fmt.Println(user)
}

func main() {

	// ============================================================
	// 6. Create the Dependency
	// ============================================================

	repository := PostgresUserRepository{}

	// ============================================================
	// 7. Inject the Dependency
	// ============================================================

	service := NewUserService(repository)

	// ============================================================
	// 8. Use the Service
	// ============================================================

	service.GetUser(1)
}
