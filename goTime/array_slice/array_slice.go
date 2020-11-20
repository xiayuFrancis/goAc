package main

import "fmt"

func main() {
	arr01 := [...]int{1, 2, 3}
	arr02 := [...]int{1, 2, 3, 4}
	fmt.Printf("%d 的类型是: %T\n", arr01, arr01)
	fmt.Printf("%d 的类型是: %T\n", arr02, arr02)

	myarr := []int{1}
	myarr = append(myarr, 2)
	myarr = append(myarr, 3, 4)
	myarr = append(myarr, []int{7,8}...)
	myarr = append([]int{0}, myarr...)

	fmt.Println(myarr)

}
