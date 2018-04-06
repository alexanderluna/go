package main

import (
	. "fmt"
	. "unsafe"
)

func main() {
	Printf("%T,\t size of: %d\n", int(10), Sizeof(int(10)))
	Printf("%T,\t size of: %d\n", int8(10), Sizeof(int8(10)))
	Printf("%T,\t size of: %d\n", int16(10), Sizeof(int16(10)))
	Printf("%T,\t size of: %d\n", int32(10), Sizeof(int32(10)))
	Printf("%T,\t size of: %d\n", int64(10), Sizeof(int64(10)))
	Printf("%T,\t size of: %d\n", string("a"), Sizeof(string("a")))
	Printf("%T,\t size of: %d\n", rune(97), Sizeof(rune(97)))
}
