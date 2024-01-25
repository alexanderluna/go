////////////////////////////////////////////////////////////////////////////////
// Go has several composite types which you can use to build on top of
// predeclared types:
// 	- Arrays
//	- Slices
//	- Maps
//	- Structs
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"slices"
)

func main() {
	////////////////////////////////////////////////////////////////////////////
	// While Go has Arrays usually you want to use Slices instead. You can
	// declare arrays either with a number of elements or with [...] when you
	// use an array literal.
	////////////////////////////////////////////////////////////////////////////

	var plantArray = [3]string{"Bisasam", "Bisaknosp", "Bisaflor"}
	var fireArray = [...]string{"Glumanda", "Glutexo", "Glurak"}
	var pokemons = [5]string{3: "Pikachu"}
	var stats [5][5]int

	fmt.Println(plantArray == fireArray) // false
	fmt.Println(len(pokemons))           // 5
	fmt.Println(pokemons)                // ["", "", "", "Pickachu", "", ""]
	fmt.Println(stats)                   // [[00000][00000][00000][00000][00000]]

	////////////////////////////////////////////////////////////////////////////
	// In Go, the size of an array is part of the type of the array. This makes
	// it impossible to dynamically set an array size or convert arrays at run
	// time. Hence, Slices are the go to choice instead. You declare them with
	// empty brackets. Most of the array behavior is the same.
	////////////////////////////////////////////////////////////////////////////

	var plantSlice = []string{"Bisasam", "Bisaknosp", "Bisaflor"}
	var waterSlice = []string{"Schiggy", "Schillok", "Turtok"}

	slices.Equal(plantSlice, waterSlice) // true if same length and values

	////////////////////////////////////////////////////////////////////////////
	// You can [append] to slices using the append function. Slices have a
	// capacity (number of consecutive memory spaces). As you append items the
	// capacity grows. You can also use the [make] function to set the type,
	// length and capacity of a slice. [clear] sets all values to their default.
	////////////////////////////////////////////////////////////////////////////

	waterSlice = append(waterSlice, "Enton")
	plantSlice = append(plantSlice, waterSlice...)

	var otherPokemons = make([]string, 3, 5)
	fmt.Println(otherPokemons) // [  ] -> length 3, capacity 5

	clear(plantSlice) // [    ]

	fmt.Println(waterSlice[0:2]) // ["Schiggy", "Schillok"]

	copyWaterPokemons := make([]string, 3)
	copy(copyWaterPokemons, waterSlice)

	plantPokemonSlice := plantArray[:]         // array to slice
	waterPokemonArray := [3]string(waterSlice) // slice to array
	fmt.Println(plantPokemonSlice, waterPokemonArray)

	////////////////////////////////////////////////////////////////////////////
	// Strings are made off Runes which is an alias for Bytes so you can use
	// the Array and Slice syntax on Strings as well.
	////////////////////////////////////////////////////////////////////////////

	waterPokemon := "Schiggy"
	fmt.Println(waterPokemon[:4])  // Schi
	fmt.Println(len(waterPokemon)) // 7

	////////////////////////////////////////////////////////////////////////////
	// When you want to associate one value type to another you can use [maps]
	// Maps are similar to slices. They grow as you add key:value pairs, you
	// can use [make], [len] and [maps.Equal].
	////////////////////////////////////////////////////////////////////////////

	bluePokedex := make(map[string][]string, 100)
	fmt.Println(len(bluePokedex)) // 100 key:value pairs

	redPokedex := map[string][]string{
		"Water Pokemons": {"Schiggy", "Schillok", "Turtok"},
		"Plant Pokemons": {"Bisasam", "Bisaknosp", "Bisaflor"},
		"Fire Pokemons":  {"Glumanda", "Glutexo", "Glurak"},
	}

	redPokedex["Other Pokemons"] = []string{"pikachu", "enton"}

	fmt.Println(redPokedex["Water Pokemons"])

	v, ok := redPokedex["Water Pokemons"]
	fmt.Println(v, ok) // ["Schiggy", "Schillok", "Turtok"], true

	delete(redPokedex, "Other Pokemons") // removes the key:value pair
	clear(redPokedex)                    // empties the map

	////////////////////////////////////////////////////////////////////////////
	// You can use Maps to simulate a Set - where values are not repeated.
	// However, you would need to create your own functions for union,
	// intersection and subtraction.
	////////////////////////////////////////////////////////////////////////////

	caughtPokemon := map[string]bool{
		"Pikachu":  true,
		"Schiggy":  true,
		"Glumanda": true,
	}

	fmt.Println(caughtPokemon["Pikachu"]) // true
	fmt.Println(caughtPokemon["Glurak"])  // false

	////////////////////////////////////////////////////////////////////////////
	// All values in a Map must be the same type and you can't constraint the
	// Map to certain keys. This is where Structs shine. Structs are the closest
	// thing Go has to classes. You can convert from one struct type to another
	// if the fields of both structs have the same names, order, and types.
	////////////////////////////////////////////////////////////////////////////

	type pokemon struct {
		name    string
		hp      int
		mp      int
		attacks []string
	}

	pikachu := pokemon{
		name:    "Pikachu",
		hp:      20,
		mp:      20,
		attacks: []string{"Donnerwelle", "Ruckzuckhieb", "Eisenschweif"},
	}

	fmt.Println(pikachu)

	// a variable implementing a struct type without name is an anonymous struct
	// this is useful for translating data to a struct (JSON)
	attack := struct {
		name     string
		strength int
	}{
		name:     "Donnerschock",
		strength: 15,
	}

	fmt.Println(attack)
}
