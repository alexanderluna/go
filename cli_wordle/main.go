package main

import (
	"bufio"
	"cli_wordle/wordle"
	"fmt"
	"os"
)

const maxAttempts = 6

func main() {
	corpus, err := wordle.ReadCorpus("corpus/english.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to read corpus: %s", err)
		return
	}

	g, err := wordle.New(bufio.NewReader(os.Stdin), corpus, maxAttempts)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to start the game: %s", err)
		return
	}

	g.Play()
}
