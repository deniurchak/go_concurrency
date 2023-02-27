package main

import (
	"fmt"
)

func main() {
	// create a notification channel
	ch := make(chan struct{})
	// run the sample go routine
	go sample(ch)
	go sample(ch)
	go sample(ch)
	<-ch
	<-ch
	<-ch
}

func sample(ch chan struct{}) {
	fmt.Println("new routine")
	// send a notification to the channel
	ch <- struct{}{}
}
