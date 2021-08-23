package main

import (
	"fmt"
)

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	fmt.Println(s)
	//updateSlice(s)
	s2 :=s[3:5]
	fmt.Println(s2)
	fmt.Println(len(s), cap(s))
	fmt.Println(len(s2), cap(s2))
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s3, s4, s5)
	fmt	.Println(arr)
}
