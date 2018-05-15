package main

import (
	"testing"
	"time"
)

var x = 1

func f1(y *int) {
	for {
		_ = *y
	}
}
func f2(y *int) {
	for {
		_ = *y
	}
}
func TestConcurrentRead(t *testing.T) {
	go f1(&x)
	go f2(&x)
	// x = 1
	time.Sleep(time.Millisecond)
}
