package main

import "fmt"

// ==========================
// Domain Layer
// ==========================

// Entity
type User struct {
	ID   int
	Name string
}

// Repository abstraction
// Application depends on this
type UserRepository interface {
	GetUser(id int) User
}

// ==========================
// Infrastructure Layer
// ==========================

// Lower level database abstraction
type PostgresRepository interface {
	Query(query string) string
}

// Concrete PostgreSQL implementation
type PostgresDB struct {
}

func (p PostgresDB) Query(query string) string {
	return "Postgres result"
}

// User repository implementation
// It uses PostgreSQL internally
type UserRepositoryImpl struct {
	db PostgresRepository
}

func NewUserRepository(
	db PostgresRepository,
) *UserRepositoryImpl {

	return &UserRepositoryImpl{
		db: db,
	}
}

func (r UserRepositoryImpl) GetUser(id int) User {

	result := r.db.Query("SELECT * FROM users")

	fmt.Println(result)

	return User{
		ID:   id,
		Name: "Mahdi",
	}
}

// ==========================
// Application Layer
// ==========================

type UserService struct {
	userRepo UserRepository
}

func NewUserService(
	userRepo UserRepository,
) *UserService {

	return &UserService{
		userRepo: userRepo,
	}
}

func (s UserService) GetUser(id int) User {
	return s.userRepo.GetUser(id)
}

// ==========================
// Composition Root
// main creates dependencies
// ==========================

func main() {

	// Create database
	postgres := PostgresDB{}

	// Inject postgres into repository
	userRepository := NewUserRepository(
		postgres,
	)

	// Inject repository into service
	userService := NewUserService(
		userRepository,
	)

	// Use service
	user := userService.GetUser(1)

	fmt.Println(user)

}
