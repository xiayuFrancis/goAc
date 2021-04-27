package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	http.HandleFunc("/", sayHelloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sayHelloHandler(w http.ResponseWriter, r *http.Request)  {
	helloworld(10000)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Fprintf(w, "Hello world!\n")
}

func helloworld(times int) {
	time.Sleep(time.Second)
	var counter int
	for i:=0; i < times; i++ {
		for j := 0; j < times; j++ {
			counter++
		}
	}
}