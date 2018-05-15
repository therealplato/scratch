package main

import "fmt"

func main() {
	f()
}

func f() {
	defer g()
	panic("asdf")
}

func g() {
	fmt.Println("hi from g")
}
