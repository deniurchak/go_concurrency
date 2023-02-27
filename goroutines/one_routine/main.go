package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	// print stack of the main goroutine
	debug.PrintStack()
	// create a notification channel
	ch := make(chan struct{})
	// run the sample go routine
	go sample(ch)
	<-ch
}

func sample(ch chan struct{}) {
	// print the stack
	debug.PrintStack()
	n := runtime.NumGoroutine()
	fmt.Println("Number of goroutines:", n)
	// send a notification to the channel
	ch <- struct{}{}
}