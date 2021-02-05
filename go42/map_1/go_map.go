package main

import (
	. "fmt"
)

func main() {
	map2 := make(map[string]float32, 100)
	map2["ala"]=200.44
	map2["blb"]=200.55
	Println(map2)

	if _, ok := map2["two"]; !ok {
		Println("no entry")
	}

	map_map := make(map[string]map[string]float32)

	map_map["1"] = map2

	Println(map_map)
}
