package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var name *string
	name = flag.String("name", "covfefe", "a name")
	name = flag.String("n", "covfefe", "a name")
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(os.Args)
}

// invocation:
// go run main.go -name asdf
// go run main.go -n asdf
