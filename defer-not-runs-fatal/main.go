package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer g()
	f()
	time.Sleep(1 * time.Second)
}

func f() {
	log.Fatal("asdf")
}

func g() {
	fmt.Println("hi from g")
}
