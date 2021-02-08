package main

import "fmt"

func main() {
	do_addNumber()
}
func do_addNumber()  {
	addNum := addNumber(5)
	addNum(1)
	addNum(1)
	addNum(1)
	fmt.Println("---------------------")
	addNum1 := addNumber(5)
	addNum1(1)
	addNum1(1)
	addNum1(1)
}

func addNumber(x int) func(int) {
	fmt.Printf("x: %d, addr of x:%p\n", x, &x)
	return func(y int) {
		k := x + y
		x = k
		y = k
		fmt.Printf("x: %d, addr of x:%p\n", x, &x)
		fmt.Printf("y: %d, addr of y:%p\n", y, &y)
	}
}