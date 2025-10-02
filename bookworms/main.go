package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var path string
	flag.StringVar(
		&path,
		"path",
		"testdata/bookworms.json",
		"path to your JSON file",
	)
	flag.Parse()

	bookworms, err := loadBookworms(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	commonBooks := findCommonBooks(bookworms)
	displayBooks(commonBooks)
}
