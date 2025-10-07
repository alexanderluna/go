# Bookworms

CLI tool that takes a list of bookworms and their book collections in the form
of a JSON file. We use the `encoding/json` package to decode JSON and the defer
keyword to make sure we close the file once we are done with it.

> `testdata` is ignored by the go tool and holds support files.

Next, we find all the Books that our Readers have in common by counting the
Books that overlap and returning a sorted slice. In order to sort the slice, we
use the sort package and implement the Interface (Len, Swap, Less).

Finally, we display the books that our readers have in common. The CLI tool has
a flag to pass in a custom path to a JSON file with a default JSON file if the
flag isn't used.

## Tests

First, create a map with struct for all the test cases. Each test case has a
two fields:

1. input: holds the input for testing our function
2. want: holds the desired output against which we can test

```go
testCases := map[string]struct {
  input []Bookworm
  want  []Book
 }{
    "test case description": {
        input: {},
        want: {},
    },
    // ... all other test cases follow the same pattern
 }
```

Now that we defined our test cases we can use a for loop to run each one.

```go
for name, tc := range testCases {
  t.Run(name, func(t *testing.T) {
   got := yourFunction(tc.input)

   if got != tc.want {
    t.Fatalf("our test failed: expected %v, got: %v", tc.want, got)
   }
  })
 }
```
