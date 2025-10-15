package wordle

// page 192

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Game holds all the required information needed to play Wordle.
type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

// New returns a Game variable which can be used to play Wordle.
func New(playerInput io.Reader, corpus []string, maxAttempts int) (*Game, error) {
	if len(corpus) == 0 {
		return nil, ErrCorpusIsEmpty
	}
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacters(pickWord(corpus)),
		maxAttempts: maxAttempts,
	}

	return g, nil
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Wordle")

	for currentAttempt := 1; currentAttempt < g.maxAttempts; currentAttempt++ {
		guess := g.ask()

		fb := computeFeedback(guess, g.solution)

		fmt.Println(fb.String())

		if slices.Equal(guess, g.solution) {
			fmt.Printf("You won 🏆\nYou found the right word in %v.\nThe word was: %s\n", pluralizeGuess(currentAttempt), string(g.solution))
			return
		}
	}

	fmt.Printf("You lost 👾\nThe solution was: %s\n", string(g.solution))
}

// ask reads the input until a valid suggestion is made (and returned).
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Wordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := splitToUppercaseCharacters(string(playerInput))

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid: %s.\n", err.Error())
		} else {
			return guess
		}
	}
}

// errInvalidWordLength is returned when the word has the wrong number of characters
var errInvalidWordLength = fmt.Errorf("invalid guess, the number of characters doesn't match solution")

// validateGuess makes sure the guess is valid
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}
	return nil
}

// splitToUppercaseCharacters is a naive implementation to turn a string into a list of characters
func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// pluralizeGuess returns pluralizes the guess if guesses is greater than 1
func pluralizeGuess(guesses int) string {
	if guesses > 1 {
		return fmt.Sprintf("%d guesses", guesses)
	}

	return "1 guess"
}
