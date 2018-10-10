package main

//Hi, could anyone give me a sense of how complicated of a project building a tool that makes a bunch of http requests and saves the json to a file would be?I need to make about 60k requests and from what I can tell I'll get significant time saving if it's done in go

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 6, "number of requests")
	flag.Parse()
	f, err := ioutil.TempFile("", "*.out")
	if err != nil {
		log.Fatalf("failure making temp file: %q", err)
	}
	fmt.Printf("logging to %s\n", f.Name())
	for i := 0; i < n; i++ {
		request(f)
	}
}

func request(w io.Writer) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://httpstat.us/202", nil)
	if err != nil {
		log.Fatalf("failure constructing request: %q", err)
	}
	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("failure requesting: %q", err)
	}
	n, err := io.Copy(w, res.Body)
	if err != nil {
		log.Fatalf("failure copying response to file: %q", err)
	}
	if n == 0 {
		log.Println("received empty response")
	}
}
