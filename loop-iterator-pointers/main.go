package main

import (
	"fmt"
)

func main() {
	var a [10]int
	for i, j := range a {
		fmt.Printf("%p | %p\n", &i, &j)
	}
}
