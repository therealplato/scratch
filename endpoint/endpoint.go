package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	e := endpoint{a: A{}}
	fmt.Println(e.start("localhost:8080"))
}

// dependency
type A struct{}

func (a *A) b(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

type endpoint struct {
	a A
}

// TODO: Learn how to declare ServeHTTP on endpoint
// As of now, you can test each route independently with httptest, but not the routing logic
func (e *endpoint) start(addr string) error {
	if addr == "" {
		log.Fatal("endpoint: requires listening address")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/a/b", e.a.b)
	mux.HandleFunc("/c", e.c)
	mux.HandleFunc("/d", e.d)
	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	fmt.Printf("Listening %s...\n", addr)
	return server.ListenAndServe()
}

func (e *endpoint) c(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
	}
	return
}

func (e *endpoint) d(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("hello")
	responsePayload, err := json.MarshalIndent(map[string]interface{}{"hello": 1}, "", "  ")
	if err != nil {
		log.Println("marshal failed", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(responsePayload)
}

// invocation:
// curl -I localhost:8080/a/b
// curl -I localhost:8080/c
// curl -I localhost:8080/d
