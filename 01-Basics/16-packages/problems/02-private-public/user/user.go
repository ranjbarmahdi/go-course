package user

type User struct {
	Name string
	age  int
}

func NewUser(name string, age int) User {
	return User{
		Name: name,
		age:  age,
	}
}

func (u User) GetAge() int {
	return u.age
}
