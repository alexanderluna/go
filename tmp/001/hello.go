package main

import (
	"fmt"
	"os"
)

type User struct {
	firstName, lastName string
}

func (u *User) fullName() string {
	return fmt.Sprintf("%s %s", u.firstName, u.lastName)
}

type Person interface {
	fullName() string
}

func Greet(p Person) string {
	return fmt.Sprintf("His full name is: %s", p.fullName())
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	} else {
		fmt.Println("No argument here")
	}

	cities := []string{}
	otherCities := []string{"Bremen", "Berlin"}
	cities = append(cities, otherCities...)

	for _, v := range cities {
		fmt.Printf("Cities: %s\n", v)
	}

	user := &User{"Alexander", "Luna"}
	fmt.Println(Greet(user))
}
