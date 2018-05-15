package main

import "fmt"

var junk = "hi​​"

func main() {
	fmt.Println("string: ", junk)
	fmt.Printf("quoted: %q\n", junk)
	fmt.Println("len:", len(junk))
	for i, c := range junk {
		fmt.Printf("index: %v value: %v\n", i, c)
	}
}
