package main

import "fmt"

func main() {
	m := map[string]string{
		"name": "xy",
		"age":  "20",
		"time": "22:52",
		"year": "2020",
	}
	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	if name, ok := m["name"]; ok {
		fmt.Println(name)
	}

	delete(m, "name")
	fmt.Println(m)
}
