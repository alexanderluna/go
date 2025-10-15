package wordle

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	testCases := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in english": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in arabic": {
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		"5 characters in japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			g, _ := New(strings.NewReader(tc.input), []string{string(tc.want)}, 0)

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuess(t *testing.T) {
	testCases := map[string]struct {
		guess []rune
		err   error
	}{
		"correct length": {
			guess: []rune("guess"),
			err:   nil,
		},
		"too long": {
			guess: []rune("palindrome"),
			err:   errInvalidWordLength,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			g, _ := New(nil, []string{"HELLO"}, 0)

			err := g.validateGuess(tc.guess)
			if !errors.Is(err, tc.err) {
				// %c = character, %q = single quoted character
				t.Errorf("%c, expected %q, got %q", tc.guess, tc.err, err)
			}
		})
	}
}

func Test_splitToUppercaseCharacters(t *testing.T) {
	testCases := map[string]struct {
		guess string
		want  []rune
	}{
		"Capitalized String": {
			guess: "Hello",
			want:  []rune("HELLO"),
		},
		"Mixed case": {
			guess: "HeLLo",
			want:  []rune("HELLO"),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			output := splitToUppercaseCharacters(tc.guess)
			if !slices.Equal(output, tc.want) {
				t.Errorf("Expected %v, got %v", string(tc.want), string(output))
			}
		})
	}
}

func Test_pluralizeGuess(t *testing.T) {
	testCases := map[string]struct {
		guesses int
		want    string
	}{
		"1 guess": {
			guesses: 1,
			want:    "1 guess",
		},
		"2 or more guesses": {
			guesses: 4,
			want:    "4 guesses",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			guesses := pluralizeGuess(tc.guesses)
			if guesses != tc.want {
				t.Errorf("Expected: %s, got: %s", tc.want, guesses)
			}
		})
	}
}

func Test_computeFeedback(t *testing.T) {
	testCases := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"all characters in correct position": {
			guess:            "ABCDE",
			solution:         "ABCDE",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double characters in correct position": {
			guess:            "AABCD",
			solution:         "AABCD",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double characters with wrong answer": {
			guess:            "ABCCC",
			solution:         "ABCCD",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, absentCharacter},
		},
		"5 times the same character but only 2 are there": {
			guess:            "CCCCC",
			solution:         "ABCCD",
			expectedFeedback: feedback{absentCharacter, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
		"double character but in the wrong position": {
			guess:            "ACCBD",
			solution:         "ABCCD",
			expectedFeedback: feedback{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
		"triple character but in the wrong position": {
			guess:            "ACCCD",
			solution:         "ABCCD",
			expectedFeedback: feedback{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"1 correct, 1 incorrect, 1 absent": {
			guess:            "CCCZZ",
			solution:         "ABCCD",
			expectedFeedback: feedback{wrongPosition, absentCharacter, correctPosition, absentCharacter, absentCharacter},
		},
		"2 characters are in a swapped position": {
			guess:            "ADCCB",
			solution:         "ABCCD",
			expectedFeedback: feedback{correctPosition, wrongPosition, correctPosition, correctPosition, wrongPosition},
		},
		"only 1 character is absent": {
			guess:            "AZCDE",
			solution:         "ABCDE",
			expectedFeedback: feedback{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"1 absent character and 1 incorrect character": {
			guess:            "AXCDD",
			solution:         "ABCDE",
			expectedFeedback: feedback{correctPosition, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !tc.expectedFeedback.Equal(fb) {
				t.Errorf("guess: %q, got the wrong feedback, expected %v, got: %v", tc.guess, tc.expectedFeedback, fb)
			}
		})
	}
}
