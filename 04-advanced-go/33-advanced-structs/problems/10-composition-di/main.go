/*
Problem:

Create a repository interface:

type UserRepository interface {
    FindByID(id int) string
}

Create a concrete repository:

type UserRepositoryImpl struct{}

Implement:

FindByID(id int) string

It should return:

"Mahdi"

Then create:

type UserService struct {
    repo UserRepository
}

Create a UserService using dependency injection.

Requirements:

- UserService must depend on the UserRepository interface,
  NOT UserRepositoryImpl.
- Inject UserRepositoryImpl into UserService.
- Call FindByID(1) through the service's repository.
- Print the returned username.

Expected output:

Mahdi

Important:

This demonstrates composition + interfaces + dependency injection.

The service depends on:

    UserRepository

not:

    UserRepositoryImpl

This makes the dependency replaceable and easier to test.
*/

package main

import "fmt"

type UserRepository interface {
	FindByID(id int) string
}

type UserRepositoryImpl struct{}

func (u UserRepositoryImpl) FindByID(id int) string {
	return "Mahdi"
}

type UserService struct {
	repo UserRepository
}

func main() {
	repo := UserRepositoryImpl{}
	service := UserService{
		repo: repo,
	}
	fmt.Println(service.repo.FindByID(10))
}
