package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	c1, c2 := generator(), generator()

	after := time.After(5 * time.Second)
	for {
		select {
		case n := <-c1:
			fmt.Println("recieved from c1: ", n)
		case n := <-c2:
			fmt.Println("recieved from c2: ", n)
		case <-time.After(800 * time.Millisecond):
			fmt.Println("time out")
			return
		case <-after:
			fmt.Println("time after")
			return
		}
	}

}
