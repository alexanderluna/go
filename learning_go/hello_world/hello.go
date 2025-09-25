// modules are organized into one or more packages. The "main" package is where
// your Go program starts.
package main

// import the required packages in a list. You can only import whole packages.
import (
	"flag"
	"fmt"
)

// the main function is always the starting point
func main() {
	var lang string
	flag.StringVar(&lang,
		"lang",
		"jp",
		"The required language, e.g en, jp, es...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

type language string

var phrasebook = map[language]string{
	"en": "Hello World",
	"jp": "やっと来たよ。",
	"es": "Hola Mundo",
}

func greet(l language) string {
	greeting, ok := phrasebook[l]

	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
