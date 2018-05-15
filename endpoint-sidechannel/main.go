package main

import (
	"fmt"
	"net/http"
)

type endpoint struct {
	state chan bool
}

func (e *endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		e.state <- false
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	case "GET":
		e.state <- true
	}
}

func main() {
	c := make(chan bool)
	e := &endpoint{c}
	go func() {
		for {
			success := <-c
			fmt.Printf("main received channel communication from endpoint: %v\n", success)
		}
	}()
	s := &http.Server{Addr: "127.0.0.1:8080", Handler: e}
	fmt.Println(s.ListenAndServe())
}
