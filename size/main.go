package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(int(0)))   // 8
	fmt.Println(unsafe.Sizeof(int8(0)))  // 1
	fmt.Println(unsafe.Sizeof(int16(0))) // 2
	fmt.Println(unsafe.Sizeof(int32(0))) // 4
	fmt.Println(unsafe.Sizeof(int64(0))) // 8

	fmt.Println(unsafe.Sizeof(uint(0)))   // 8
	fmt.Println(unsafe.Sizeof(uint8(0)))  // 1
	fmt.Println(unsafe.Sizeof(uint16(0))) // 2
	fmt.Println(unsafe.Sizeof(uint32(0))) // 4
	fmt.Println(unsafe.Sizeof(uint64(0))) // 8

	fmt.Println(unsafe.Sizeof(byte(0)))       // 1
	fmt.Println(unsafe.Sizeof(rune(0)))       // 4
	fmt.Println(unsafe.Sizeof(uintptr(0)))    // 8
	fmt.Println(unsafe.Sizeof(float32(0)))    // 4
	fmt.Println(unsafe.Sizeof(float64(0)))    // 8
	fmt.Println(unsafe.Sizeof(complex64(0)))  // 8
	fmt.Println(unsafe.Sizeof(complex128(0))) // 16

	fmt.Println(unsafe.Sizeof(false))    // 1
	fmt.Println(unsafe.Sizeof("string")) // 16
}
