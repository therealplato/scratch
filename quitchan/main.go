package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("vim-go")
	wg := &sync.WaitGroup{}
	q := make(chan struct{})
	wg.Add(2)
	go fast(wg, q)
	go slow(wg, q)
	go func(q chan struct{}) {
		<-time.After(100 * time.Millisecond)
		close(q)
	}(q)
	wg.Wait()
	fmt.Println("main: bye")
}

func fast(wg *sync.WaitGroup, q chan struct{}) {
	defer wg.Done()
	t := time.NewTicker(3 * time.Millisecond)
	for {
		select {
		case <-t.C:
			fmt.Println("fast: hi")
		case <-q:
			fmt.Println("fast: bye")
			return
			// default:
			// 	fmt.Println("fast: no reads available")
		}
	}
}

func slow(wg *sync.WaitGroup, q chan struct{}) {
	defer wg.Done()
	t := time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case <-t.C:
			fmt.Println("slow: hi")
		case <-q:
			fmt.Println("slow: bye")
			return
		}
	}
}
