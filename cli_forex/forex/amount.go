package forex

// Amount defines a quantity of money in a given Currency.
type Amount struct {
	quantity Decimal
	Currency Currency
}
