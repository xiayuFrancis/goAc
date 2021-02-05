package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r:= recover(); r != nil {
			var ok bool
			err , ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(fmt.Sprintf("%v: field :%v err: %v", idx, field, err))
		}
		numbers = append(numbers, num)
	}
	return
}

func main() {
	number, err := Parse("")
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(number)
		fmt.Println(err.Error())
	}
}
