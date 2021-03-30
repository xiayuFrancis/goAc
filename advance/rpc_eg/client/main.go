package main

import (
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

func main() {
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
	var result Result

	if err := client.Call("Cal.Square", 12, &result); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}
	log.Printf("%d^2 = %d", result.Num, result.Ans)

	var result2 Result

	asyncCall := client.Go("Cal.Square", 12, &result2, nil)
	log.Printf("%d^2 = %d", result2.Num, result2.Ans)

	<-asyncCall.Done
	log.Printf("%d^2 = %d", result2.Num, result2.Ans)
}
