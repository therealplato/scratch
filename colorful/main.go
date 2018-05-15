package main

import (
	"fmt"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func main() {
	pal, _ := colorful.SoftPalette(50)

	for _, c := range pal {
		fmt.Println(c.Hex())
	}
}
