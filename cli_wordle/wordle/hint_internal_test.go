package wordle

import "testing"

func Test_feedback_String(t *testing.T) {
	testCases := map[string]struct {
		fb   feedback
		want string
	}{
		"3 characters in the correct position": {
			fb:   feedback{correctPosition, correctPosition, correctPosition},
			want: "✅✅✅",
		},
		"1 in correct position, 1 in wrong position, 1 absent": {
			fb:   feedback{correctPosition, wrongPosition, absentCharacter},
			want: "✅🟧⬜️",
		},
		"1 in wrong position, 1 absent, 1 in correct position": {
			fb:   feedback{wrongPosition, absentCharacter, correctPosition},
			want: "🟧⬜️✅",
		},
		"unknown position should return default case": {
			fb:   feedback{404},
			want: "❌",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := tc.fb.String()
			if got != tc.want {
				t.Errorf("String() = %v, want %v", got, tc.want)
			}
		})
	}
}
