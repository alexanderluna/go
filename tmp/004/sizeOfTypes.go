package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("%T,\t size of: %d\n", int(10), unsafe.Sizeof(int(10)))
	fmt.Printf("%T,\t size of: %d\n", int8(10), unsafe.Sizeof(int8(10)))
	fmt.Printf("%T,\t size of: %d\n", int16(10), unsafe.Sizeof(int16(10)))
	fmt.Printf("%T,\t size of: %d\n", int32(10), unsafe.Sizeof(int32(10)))
	fmt.Printf("%T,\t size of: %d\n", int64(10), unsafe.Sizeof(int64(10)))
	fmt.Printf("%T,\t size of: %d\n", string("a"), unsafe.Sizeof(string("a")))
	fmt.Printf("%T,\t size of: %d\n", rune(97), unsafe.Sizeof(rune(97)))
}
