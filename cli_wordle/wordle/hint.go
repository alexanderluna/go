package wordle

import (
	"fmt"
	"os"
	"strings"
)

// hint describes the validity of a character in a word
type hint int

// feedback is a list of hints, one for each character
type feedback []hint

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

// String implements the Stringer interface.
func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜️"
	case wrongPosition:
		return "🟧"
	case correctPosition:
		return "✅"
	default:
		return "❌"
	}
}

// String implements the String interface for a slice of hints.
func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}

// computeFeedback verifies each character of the guess against the solution
func computeFeedback(guess, solution []rune) feedback {
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))

	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error! Guess and solution"+
			" have different lengths: %d vs %d", len(guess), len(solution))
		return result
	}

	// check for correctly positioned letters
	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctPosition
			used[posInGuess] = true
		}
	}

	// check for wrongly positioned letters
	for posInGuess, guessChar := range guess {
		// ignore already marked characters
		if result[posInGuess] != absentCharacter {
			continue
		}

		for posInSolution, solutionChar := range solution {
			// ignore already assigned characters of the solution
			if used[posInSolution] {
				continue
			}

			// mark character, assign as used and skip to the next letter
			if guessChar == solutionChar {
				result[posInGuess] = wrongPosition
				used[posInSolution] = true
				break
			}
		}
	}
	return result
}

// Equal checks for equality of two feedbacks.
func (fb feedback) Equal(other feedback) bool {
	if len(fb) != len(other) {
		return false
	}

	for index := range fb {
		if other[index] != fb[index] {
			return false
		}
	}

	return true
}
