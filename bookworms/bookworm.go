package main

import (
	"encoding/json"
	"os"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// reads the file and returns the list of bookworms and their books
func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// initialize the type in which the file will be decoded
	var bookworms []Bookworm

	// decode the json file and store the content in bookworms
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil
}
