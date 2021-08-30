package main

import (
	"fmt"
	"sync"
)

func main() {
	chanDemo()
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done: func() { wg.Done() },
	}
	go doWorker(id, w.in, w)
	return w
}

func doWorker(id int, c chan int, w worker) {
	for n := range c {
		fmt.Printf("doWorker %d received  %c\n", id, n)
		w.done()
	}
}

func chanDemo() {
	var chancellors [10]worker
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		chancellors[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		chancellors[i].in <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		chancellors[i].in <- 'A' + i
	}

}
