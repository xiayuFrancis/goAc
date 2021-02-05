package main

import "fmt"

type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile
	father *Profile
}

func (person * Profile) increase_age() {
	person.age += 1
}
func main() {
	myself := Profile{
		name:   "xiaoming",
		age:    20,
		gender: "male",
		mother: nil,
		father: nil,
	}
	fmt.Printf("now age : %d\n", myself.age)
	myself.increase_age()
	fmt.Printf("now age : %d\n", myself.age)
}
