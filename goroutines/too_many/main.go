package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "")
	})
	go http.ListenAndServe(":80", nil)
	fmt.Println("Debugging info is available on http://localhost:80/debug/pprof/")
	for {
		go increment(interrupt)
	}
}

func increment(interrupt chan os.Signal) {
	for range interrupt {
		debug.PrintStack()
		os.Exit(0)
		return
	}
}
