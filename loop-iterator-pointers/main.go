package main

import (
	"fmt"
)

func main() {
	var a [10]int
	for i, j := range a {
		fmt.Printf("%p | %p\n", &i, &j)
		go func(i int) {
			fmt.Printf("%p\n", i)
		}(i)
	}
}
