package main

import (
	"fmt"
	"os"
)

func main() {
	echo2()
}

func echo1()  {
	var s, sep string
	for i:= 0; i< len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}