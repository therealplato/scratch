package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	bb := []byte("hi")
	fmt.Println(base64.URLEncoding.EncodeToString(bb))
}
