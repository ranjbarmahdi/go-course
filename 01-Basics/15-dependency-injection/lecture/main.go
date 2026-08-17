package main

import "fmt"

// Interface
type UserRepository interface {
	FindUser(id int) string
}

// Concrete implementation
type PostgresUserRepository struct{}

func (p PostgresUserRepository) FindUser(id int) string {
	return "Mahdi"
}

// Service
type UserService struct {
	repository UserRepository
}

// Constructor

func NewUserService(
	repository UserRepository,
) *UserService {

	return &UserService{
		repository: repository,
	}
}

func (s UserService) GetUser(id int) {
	user := s.repository.FindUser(id)

	fmt.Println(user)
}

func main() {

	repository := PostgresUserRepository{}

	service := NewUserService(repository)

	service.GetUser(1)

}
