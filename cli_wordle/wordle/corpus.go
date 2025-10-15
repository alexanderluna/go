package wordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const ErrCorpusIsEmpty = corpusError("Corpus is empty")

// ReadCorpus reads a file at a given path and returns a list of words
func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q due to an error: %w", path, err)
	}

	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	// expect file to be line or space separated
	words := strings.Fields(string(data))
	return words, nil
}

// pickWord returns a random word from the corpus
func pickWord(corpus []string) string {
	index := rand.Intn(len(corpus))
	return corpus[index]
}
