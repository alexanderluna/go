// Types which are build into Go are called [predeclared types].
// There are 5 common [literals] which specify:
// 	- integer
// 	- floating-point
// 	- rune (characters)
// 	- string (interpreted rune literals)
// Literals are [untyped] meaning they will default to a type unless specified.

package main

func main() {
	// booleans are negative by default
	var flag bool = true

	// Go has 12 numeric types
	// 	- int8, int16, int32, int64
	// 	- uint8, uin16, uint32, uint64
	// 	- float32, float64 !!! DON'T represent decimals values exactly
	// 	- complex64, complex128
	var depthInMeters int16 = -2024
	var daysInAWeek int = 7
	var degrees float64 = 16.4
	var x complex128 = complex(3.1, 5.4)

	// The rune type is an alias for int32
	var myInitial rune = 'A'

	// [automatic type promotion] is not supported. Convert types manually
	var absolutDepth = uint16(depthInMeters)

	// Use the [:=] operator to declare variables when the type can be inferred
	// However, the := operator is not legal outside of functions.
	a, b := 6, 9

	// You can declare immutable variables with the [const] keyword. Constants
	// can only hold values which can be inferred at compile time. This means
	// you can't define a variable as a constant at runtime.
	const nameKey = "Alexander"
	const data = float64(daysInAWeek) + degrees // <- won't compile

	// In Go, every declared variable has to be read at least once. Failing to
	// do so will result in a [compile-time error].
}
