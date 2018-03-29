package main

import (
	"fmt"
	"unsafe"
)

func main() {
	normalInt := int(10)
	small8 := int8(10)
	small16 := int16(10)
	small32 := int32(10)
	large64 := int64(10)

	fmt.Printf("normal: %T, size of: %d\n", normalInt, unsafe.Sizeof(normalInt))
	fmt.Printf("small: %T, size of: %d\n", small8, unsafe.Sizeof(small8))
	fmt.Printf("small: %T, size of: %d\n", small16, unsafe.Sizeof(small16))
	fmt.Printf("small: %T, size of: %d\n", small32, unsafe.Sizeof(small32))
	fmt.Printf("large: %T, size of %d\n", large64, unsafe.Sizeof(large64))
}
