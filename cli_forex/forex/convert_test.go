package forex_test

import (
	"cli_forex/forex"
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	testCases := map[string]struct {
		amount   forex.Amount
		to       forex.Currency
		validate func(t *testing.T, got forex.Amount, err error)
	}{
		"34.98 USD to EUR": {
			amount: forex.Amount{},
			to:     forex.Currency{},
			validate: func(t *testing.T, got forex.Amount, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err.Error())
				}

				expected := forex.Amount{}

				if !reflect.DeepEqual(got, expected) {
					t.Errorf("expected %v, got %v", expected, got)
				}
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := forex.Convert(tc.amount, tc.to)
			tc.validate(t, got, err)
		})
	}
}
