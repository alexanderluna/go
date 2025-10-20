package forex

import (
	"fmt"
	"strconv"
	"strings"
)

// Decimal is capable of storing a float with a fixed precision.
// example: 1.52 = 152 * 10^(-2) will be stored as {152, 2}
type Decimal struct {
	// multiply subunits by the precision to get the real value
	subunits int64
	// precision is the number of subunits in a unit, expressed as a power of 10
	precision byte
}

const (
	// ErrInvalidDecimal is returned if the decimal is malformed.
	ErrInvalidDecimal = Error("unable to convert the decimal")

	// ErrTooLarge is returned if the quantity is too large.
	ErrTooLarge = Error("quantity over 10^12 is too large")
)

// maxDecimal is 1 trillion, using the short scale -- 10^12
const maxDecimal = 1e12

// ParseDecimal converts a string into a Decimal.
// It assumes there is at most one decimal (.) separator.
func ParseDecimal(value string) (Decimal, error) {
	// cut into integer and decimal parts
	intPart, fracPart, _ := strings.Cut(value, ".")

	// parse the value and validate
	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	if subunits > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	// calculate precision, simply and return Decimal
	precision := byte(len(fracPart))
	d := Decimal{subunits: subunits, precision: precision}
	d.simplify()

	return d, nil
}

// simplify removes the trailing 0 if it doesn't affect the value of the Decimal
func (d *Decimal) simplify() {
	// 1550%10 = 0, 1550 /= 10 = 155
	for d.subunits%10 == 0 && d.precision > 0 {
		d.precision--
		d.subunits /= 10
	}
}
