package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://slow:8080")
	if err != nil {
		log.Fatalf("get config request failed: %q", err)
	}
	cfg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("got unreadable config: %q", err)
	}
	fmt.Println(cfg)
	os.Exit(0)
}
