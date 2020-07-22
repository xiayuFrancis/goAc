package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type ColorGroup struct {
		ID int
		Name string
		Colors []string
		Type map[string]int
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
		Type: map[string]int{"key":1,},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println(string(b))
	//os.Stdout.Write(b)
}
