package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "Go语言四十二章经"
	for k, v := range s{
		fmt.Printf("k: %d, v:%c == %d\n", k, v, v)
	}

	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(", ")
	buffer.WriteString("world")

	var bu strings.Builder
	bu.WriteString("world")

	fmt.Println(buffer.String())
	fmt.Println(bu.String())

}
