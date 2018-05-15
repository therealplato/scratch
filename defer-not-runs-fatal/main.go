package main

import (
	"fmt"
	"log"
)

func main() {
	f()
}

func f() {
	defer g()
	log.Fatal("asdf")
}

func g() {
	fmt.Println("hi from g")
}
