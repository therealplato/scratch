package main

import "fmt"

func main() {
	left, right := foo()
	fmt.Println(left)
	fmt.Println(right)
}

func foo() (a int, b int) {
	a = 1
	b = 2
	return b, a
}
