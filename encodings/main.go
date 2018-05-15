package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	bb := []byte("hi")

	fmt.Println(base64.URLEncoding.EncodeToString(bb))

	fmt.Println(hex.EncodeToString(bb))
}
