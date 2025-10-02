package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

// print out Title and Author of a list of Books
func displayBooks(books []Book) {
	fmt.Println("Here are the common books:")
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
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

// returns Books are are on more than one bookworm's bookshelf
func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := countBooks(bookworms)

	var commonBooks []Book

	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}

	return sortBooks(commonBooks)
}

// registers all the books and their occurrences
func countBooks(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)

	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}

	return count
}

// custom type to implement sort.Interface
type byAuthor []Book

// implement sort.Interface length of BookByAuthor
func (b byAuthor) Len() int { return len(b) }

// implement sort.Interface swapping of Books
func (b byAuthor) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// implement sort.Interface Less to sort by Author then Title
func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}
	return b[i].Title < b[j].Title
}

// sort Books first by Author then Title
func sortBooks(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}
