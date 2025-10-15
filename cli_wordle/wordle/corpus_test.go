package wordle_test

import (
	"cli_wordle/wordle"
	"errors"
	"testing"
)

func TestReadCoprus(t *testing.T) {
	testCases := map[string]struct {
		file   string
		length int
		err    error
	}{
		"English corpus": {
			file:   "../corpus/english.txt",
			length: 34,
			err:    nil,
		},
		"Empty corpus": {
			file:   "../corpus/empty.txt",
			length: 0,
			err:    wordle.ErrCorpusIsEmpty,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			words, err := wordle.ReadCorpus(tc.file)
			if !errors.Is(tc.err, err) {
				t.Errorf("expected err %v, got %v", tc.err, err)
			}

			if tc.length != len(words) {
				t.Errorf("expected %d, got %d", tc.length, len(words))
			}
		})
	}
}
