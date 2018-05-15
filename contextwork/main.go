package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Things can still be read from closed, buffered channels:
	messages := make(chan string, 100)
	for i := 0; i < 20; i++ {
		messages <- fmt.Sprintf("message_%d", i+1)
	}
	close(messages)

	if err := doWork(messages); err != nil {
		fmt.Printf("doWork returned error: %v", err)
	}

}

func doWork(messages chan string) (err error) {
	errs := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(ctx, i, messages, errs, &wg)
	}

	workersFinished := make(chan struct{})
	go func(wg *sync.WaitGroup, workersFinished chan<- struct{}) {
		wg.Wait()
		close(workersFinished)
		return
	}(&wg, workersFinished)

	for {
		select {
		case e := <-errs:
			fmt.Println("encountered error, canceling the doWork context:", e)
			err = e
			_ = cancel
			// cancel()
		case <-workersFinished:
			fmt.Println("all workers have finished")
			return
		}
	}
}

func worker(ctx context.Context, id int, messages <-chan string, errors chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Starting worker [%v]\n", id)

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker [%v] exiting\n", id)
			return
		case st, ok := <-messages:
			if !ok {
				// messages is closed and empty
				return
			}

			fmt.Printf("Worker [%v], msg [%s] \n", id, st)

			if id == 4 {
				errors <- fmt.Errorf("Worker 4 is broken")
				return
			}

			time.Sleep(1 * time.Second)
		}
	}
}
