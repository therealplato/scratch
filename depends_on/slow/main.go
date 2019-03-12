package main

import (
	"net/http"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	http.ListenAndServe(":8080", http.HandlerFunc(f))
}

func f(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"foo":"bar"}`))
}
