package main

import (
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
)

func main() {
	random := browser.Chrome()
	fmt.Printf("random", random)
}
