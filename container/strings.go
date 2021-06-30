package main

import "fmt"

func main() {
	s:="Yes我爱你"
	for _, b := range []byte(s){
		fmt.Printf("%x ", b)
	}
}
