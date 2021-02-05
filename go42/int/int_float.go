package main

import "fmt"

func main() {
	var num01 int = 0b1100
	var num02 int = 0o14
	var num03 int = 0xc

	fmt.Printf("2进制数 %b 表示的是: %d \n", num01, num01)
	fmt.Printf("8进制数 %o 表示的是: %d \n", num02, num02)
	fmt.Printf("16进制数 %X 表示的是: %d \n", num03, num03)

	var myfloat float32 = 10000018
	fmt.Println("myfloat: ", myfloat)
	fmt.Println("myfloat: ", myfloat+1)
}
