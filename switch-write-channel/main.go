package main

import (
	"fmt"
)

func main() {
	choke := make(chan string, 150000)
	var i int
	for {
		select {
		case choke <- string(i):
			i++
		default:
			fmt.Println("Full")
			return
		}
	}
}

/*
func main() {
	choke := make(chan string, 150000)

	go func() {
		for i := 0; i < 10000000; i++ {
			choke <- string(i)
			fmt.Println("i=", i)
		}
	}()

	for {
		fmt.Println(len(choke))
		if len(choke) >= 150000 {
			fmt.Println("Full")
		}
	}
}

*/
