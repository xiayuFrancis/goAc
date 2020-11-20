package main

import "fmt"

func set_data(x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var arr [10]int
	arr[x] = 88
}
func main() {
	set_data(20)
	fmt.Println("it's ok")
}
