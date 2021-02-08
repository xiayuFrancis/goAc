package main

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool	"An import answer"
	field2 string "The same of the thing"
	field3 int "How much there are"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixFiled := ttType.Field(ix)
	fmt.Printf("%v\n", ixFiled.Tag)
}

func CallRefTag() {
	tt := TagType{true, "barak obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}