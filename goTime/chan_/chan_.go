package main

import (
	"fmt"
	"time"
)

func main() {
	pipline := make(chan int)

	go func() {
		fmt.Println("ready send data : 100")
		pipline <- 100
	}()

	go func() {
		num := <- pipline
		fmt.Printf("recv data : %d", num)
	}()

	time.Sleep(time.Second)
}
