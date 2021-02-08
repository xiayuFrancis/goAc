package main

import (
	"fmt"
	"time"
)

func Do() {
	c := make(chan int, 10)
	defer close(c)

	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send ready ", i)
		c <- i
		fmt.Println("send suc ", i)
	}
}

func recv(c <-chan int) {
	for i := range c {
		fmt.Println("received ", i)
	}
}
