package main

import (
	"fmt"
	"time"
)

func worker	(c chan int)  {
	for {
		n := <- c
		fmt.Printf("%d", n)
	}
}
func chanDemo() {
	  c := make(chan int)
	  go worker(c)
	  c <- 4
	  c <- 5
	  time.Sleep(time.Second)
}
func main() {
	chanDemo()
}
