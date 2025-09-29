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
