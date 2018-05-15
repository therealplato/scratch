package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(release())
}

func release() (err error) {
	defer func() {
		err = errors.New("error")
	}()

	return nil
}
