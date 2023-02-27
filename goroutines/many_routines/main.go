package main

import (
	"fmt"
)

func main() {
	// create a notification channel
	ch := make(chan struct{})
	var counter = 0
	// run the sample go routine
	go sample(&counter, ch)
	go sample(&counter, ch)
	go sample(&counter, ch)
	<-ch
	<-ch
	<-ch
}

func sample(counter *int, ch chan struct{}) {
	*counter++
	fmt.Println(*counter)
	// send a notification to the channel
	ch <- struct{}{}
}
