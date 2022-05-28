package main

import (
	"fmt"
	"unsafe"
)

var justString string

type String struct {
    str *byte
    len int
}

func someFunc() {
	v := createHugeString(1 << 10)

	//bytes are copied to a new location and the memory for original bytes is released
	justString = string([]byte(v[:100]))

	//underlying byte slice is kept in memory
	// justString = v[:100] 

	fmt.Println("old", *(*String)(unsafe.Pointer(&v)), "new", *(*String)(unsafe.Pointer(&justString)))
	// fmt.Println(&v, &justString)
}

func main() {
	someFunc()
	// println(justString)
}

func createHugeString(size int) string {
	result := make([]byte, size)
	for i := range result {
		result[i] = 'a'
	}
	// resultStr := string(result)
	// fmt.Println("result", *(*String)(unsafe.Pointer(&result)))
	// fmt.Println("resultStr", *(*String)(unsafe.Pointer(&resultStr)))
	return string(result)
}
