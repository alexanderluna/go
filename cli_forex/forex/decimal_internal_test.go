package forex

import (
	"errors"
	"testing"
)

func TestParseDecimal(t *testing.T) {
	testCases := map[string]struct {
		value string
		want  Decimal
		err   error
	}{
		"2 decimal digits": {
			value: "2.54",
			want:  Decimal{subunits: 254, precision: 2},
			err:   nil,
		},
		"no decimal digits": {
			value: "5",
			want:  Decimal{subunits: 5, precision: 0},
			err:   nil,
		},
		"decimal digit ends in 0": {
			value: "4.40",
			want:  Decimal{subunits: 44, precision: 1},
			err:   nil,
		},
		"decimal digit start with 0": {
			value: "1.04",
			want:  Decimal{subunits: 104, precision: 2},
			err:   nil,
		},
		"multiple of 10": {
			value: "880",
			want:  Decimal{subunits: 880, precision: 0},
			err:   nil,
		},
		"invalid decimal part": {
			value: "gibberish.99",
			err:   ErrInvalidDecimal,
		},
		"not a number": {
			value: "NaN",
			err:   ErrInvalidDecimal,
		},
		"empty string": {
			value: "",
			err:   ErrInvalidDecimal,
		},
		"too large": {
			value: "10000000000000",
			err:   ErrTooLarge,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := ParseDecimal(tc.value)

			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}

			if got != tc.want {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}
