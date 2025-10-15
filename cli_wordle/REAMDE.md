# Wordle

Wordle CLI is copy of the famous Wordle game. First, I load a list of words from
a text file through the `ReadCorpus` function which makes use of sentinel errors
to check if a file is empty.

```go
corpus, err := wordle.ReadCorpus("corpus/english.txt")
```

Then I initialize the game with a reader (`bufio.Reader`) to handle user input,
a file with a list of words and a maxAttempt constant. In both the file reader
and game initializer I propagate the errors to the caller.

```go
g, err := wordle.New(bufio.NewReader(os.Stdin), corpus, maxAttempts)
```

Within the initializer, I check that the list of words isn't empty and pick a
random word from the list, otherwise I return a sentinel error. Once the game is
initialized and configured, the game starts. The play method prompts the player
for input and returns a feedback until he/she has run out of attempts.

```zsh
Welcome to Wordle
Enter a 12-character guess:
denouuuuauou    
✅✅🟧✅⬜️⬜️⬜️⬜️✅⬜️✅⬜️
You lost 👾
The solution was: DEMONSTRATOR
```

In order to provide hints, I make use of an Enum and a switch statement to
decide which UI element to assign. The feedback and hint types implement the
Stringer interface. Inside the `computeFeedback` function I loop over each
character `[]rune` and assign it the most useful hint.

The player wins by guessing the word with the help of the hints within the
allowed attempts.

## Testing

```zsh
go test ./wordle
```
