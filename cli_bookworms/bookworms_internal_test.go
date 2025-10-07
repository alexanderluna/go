package main

import "testing"

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookworms_Success(t *testing.T) {
	tests := map[string]struct {
		bookwormsFile string
		want          []Bookworm
		wantErr       bool
	}{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nothing")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if !equalBookworms(t, got, tc.want) {
				t.Fatalf("different result: got %v, expected %v", got, tc.want)
			}
		})
	}
}

func TestBooksCount(t *testing.T) {
	testCases := map[string]struct {
		input []Bookworm
		want  map[Book]uint
	}{
		"nominal use case": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{
					handmaidsTale, theBellJar,
				}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				theBellJar:    1,
				oryxAndCrake:  1,
				janeEyre:      1},
		},
		"no bookworms": {
			input: []Bookworm{},
			want:  map[Book]uint{},
		},
		"bookworm without books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale}},
				{Name: "Peggy", Books: []Book{}},
			},
			want: map[Book]uint{handmaidsTale: 1},
		},
		"bookworm with twice the same book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, handmaidsTale}},
				{Name: "Peggy", Books: []Book{janeEyre, janeEyre}},
			},
			want: map[Book]uint{
				handmaidsTale: 2,
				janeEyre:      2,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := countBooks(tc.input)
			if !equalBooksCount(t, tc.want, got) {
				t.Fatalf("got a wrong list of books: %v, expected: %v", got, tc.want)
			}
		})
	}
}

func TestFindCommonBooks(t *testing.T) {
	testCases := map[string]struct {
		input []Bookworm
		want  []Book
	}{
		"no common books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, oryxAndCrake}},
				{Name: "Peggy", Books: []Book{janeEyre}},
			},
			want: nil,
		},
		"one common book": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, oryxAndCrake}},
				{Name: "Peggy", Books: []Book{janeEyre, handmaidsTale}},
			},
			want: []Book{handmaidsTale},
		},
		"multiple bookworms have common books": {
			input: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, oryxAndCrake}},
				{Name: "Peggy", Books: []Book{janeEyre, handmaidsTale, oryxAndCrake}},
				{Name: "John", Books: []Book{handmaidsTale}},
			},
			want: []Book{oryxAndCrake, handmaidsTale},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := findCommonBooks(tc.input)

			if !equalBooks(t, got, tc.want) {
				t.Fatalf("got a different list of books: %v, expected %v", got, tc.want)
			}
		})
	}
}

// helper to test equality of two lists of bookworms
func equalBookworms(t *testing.T, bookworms, target []Bookworm) bool {
	t.Helper()

	// return early if length of bookworms don't match
	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		// return early if bookworm names don't match
		if bookworms[i].Name != target[i].Name {
			return false
		}
		// return early if books don't match
		if !equalBooks(t, bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}

// helper to test equality of two book arrays
func equalBooks(t *testing.T, books, target []Book) bool {
	t.Helper()

	// return early if the books length don't match
	if len(books) != len(target) {
		return false
	}

	for i := range books {
		// return early if books don't match
		if books[i] != target[i] {
			return false
		}
	}

	return true
}

// helper to test equality of two maps of books count
func equalBooksCount(t *testing.T, got, want map[Book]uint) bool {
	t.Helper()

	// return early if lengths don't match
	if len(got) != len(want) {
		return false
	}

	// return early if book not found or the count doesn't match
	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || count != targetCount {
			return false
		}
	}

	return true
}
