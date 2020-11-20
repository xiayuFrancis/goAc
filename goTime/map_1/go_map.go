package main

import (
	"fmt"
)

func main() {
	scores := make(map[string]int)
	scores["english"] = 70

	fmt.Println(scores)

	if math, ok := scores["math"]; ok {
		fmt.Printf("math :%d", math)
	} else {
		fmt.Println("math not exits")
	}

	for key, value := range scores {
		fmt.Printf("key %s, value %d", key, value)
	}
}
