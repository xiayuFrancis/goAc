package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
}

func main() {
	var s = make([]int, 0, 50)
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s2 := make([]int, 16)
	copy(s2, s)
	printSlice(s2)
	fmt.Println(s2)
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println(s2)
}
