package main

import (
	"fmt"
	"os"
)

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
}
