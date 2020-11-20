package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a byte = 65
	var b uint8 = 66

	var c rune = 'B'
	fmt.Printf("a 的值: %c \nb 的值: %c\n", a, b)
	fmt.Printf("a 占用 %d 个字节数\nc 占用 %d 个字节数", unsafe.Sizeof(a), unsafe.Sizeof(c))

	mystr02 := `\r\n`
	fmt.Println(mystr02)
}
