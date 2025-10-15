package wordle

import "testing"

func TestPickWord(t *testing.T) {
	corpus := []string{"sickness", "secretary", "give", "country"}
	word := pickWord(corpus)

	if !inCorpus(corpus, word) {
		t.Errorf("expected a word to be in corpus, got %q", word)
	}
}

func inCorpus(corpus []string, word string) bool {
	for _, corpusWord := range corpus {
		if corpusWord == word {
			return true
		}
	}
	return false
}
