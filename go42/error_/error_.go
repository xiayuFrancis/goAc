package main

import (
	"errors"
	"fmt"
)

func main() {
	f, err := Sqrt(-8)
	if err != nil {
		fmt.Println(f)
		fmt.Println(err)

	} else {
		fmt.Println(f)
	}


}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New ("math - square root of negative number")
	}
	return  f*f*f, nil
}