package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestFooFatals(t *testing.T) {
	fmt.Println("TestFooFatals")
	outer := os.Getenv("FATAL_TESTING") == ""
	if outer {
		fmt.Println("Outer process: Spawning inner `go test` process, looking for failure from fatal")
		cmd := exec.Command(os.Args[0], "-test.run=TestFooFatals")
		cmd.Env = append(os.Environ(), "FATAL_TESTING=1")
		// cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		err := cmd.Run()
		fmt.Printf("Outer process: Inner process returned %v\n", err)
		if e, ok := err.(*exec.ExitError); ok && !e.Success() {
			// fmt.Println("Success: inner process returned 1, passing test")
			return
		}
		t.Fatalf("Failure: inner function returned %v, want exit status 1", err)
	} else {
		// We're in the spawned process. Do something that should fatal, failing the test.
		foo()
	}
}

// should fatal every time
func foo() {
	log.Printf("oh my goodness, i see %q\n", os.Getenv("FATAL_TESTING"))
	// log.Fatal("oh my gosh")
}
