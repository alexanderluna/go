package forex

// Currency defines the ISO-4217 code of a currency and its decimal precision
type Currency struct {
	code      string
	precision byte
}

const ErrInvalidCurrencyCode = Error("invalid currency code")

// ParseCurrency returns the associated Currency
// and may return ErrInvalidCurrencyCode
func ParseCurrency(code string) (Currency, error) {
	//  ISO-4217 standard
	if len(code) != 3 {
		return Currency{}, ErrInvalidCurrencyCode
	}

	switch code {
	case "IRR":
		return Currency{code: code, precision: 0}, nil
	case "CNY", "VND":
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "KWD", "LYD", "OMR", "TND":
		return Currency{code: code, precision: 3}, nil
	default:
		return Currency{code: code, precision: 2}, nil
	}
}
