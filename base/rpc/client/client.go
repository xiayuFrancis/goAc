package main

import (
	"fmt"
	"goAc/base/rpc/rpc_objects"
	"log"
	"net/rpc"
)

func main() {
	client , err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	args := &rpc_objects.Args{7, 8}
	var reply int
	errr := client.Call("Args.Multiply", args, &reply)
	if errr != nil {
		log.Fatal("args error", errr)
	}
	fmt.Printf("Args : %d * %d = %d", args.N, args.M, reply)
}
