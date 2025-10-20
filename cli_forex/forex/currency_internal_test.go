package forex

import (
	"errors"
	"testing"
)

func TestParseCurrency_Success(t *testing.T) {
	testCases := map[string]struct {
		code string
		want Currency
	}{
		"EUR, precision 2": {
			code: "EUR",
			want: Currency{code: "EUR", precision: 2},
		},
		"BHD, precision 3": {
			code: "BHD",
			want: Currency{code: "BHD", precision: 3},
		},
		"CNY, precision 1": {
			code: "CNY",
			want: Currency{code: "CNY", precision: 1},
		},
		"IRR, precision 0": {
			code: "IRR",
			want: Currency{code: "IRR", precision: 0},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := ParseCurrency(tc.code)

			if err != nil {
				t.Errorf("expected no error, got %s", err.Error())
			}

			if got != tc.want {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestParseCurrency_Unknown(t *testing.T) {
	_, err := ParseCurrency("INVALID")

	if !errors.Is(err, ErrInvalidCurrencyCode) {
		t.Errorf("expected error %s, got %v", ErrInvalidCurrencyCode, err)
	}
}
