package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// やっと来たよ。
}

func TestGreeting(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello World",
		},
		"Japanese": {
			lang: "jp",
			want: "やっと来たよ。",
		},
		"Spanish": {
			lang: "es",
			want: "Hola Mundo",
		},
		"Chinese, not supported": {
			lang: "ch",
			want: `unsupported language: "ch"`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)

			if got != tc.want {
				t.Errorf("expected: %q, got: %q", tc.want, got)
			}
		})
	}
}
