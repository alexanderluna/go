package main

import "fmt"

type User struct {
	firstName, lastName string
}

func (u *User) fullName() string {
	return fmt.Sprintf("Full name: %s %s", u.firstName, u.lastName)
}

type Person interface {
	fullName() string
}

func Greet(p Person) string {
	return fmt.Sprintf("Greeting: %s", p.fullName())
}

func main() {
	fmt.Printf("Hello")
	user := &User{"Alexander", "Luna"}
	fmt.Printf(Greet(user))
}
