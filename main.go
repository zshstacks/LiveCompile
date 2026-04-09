package main

import "flag"

func main() {
	flag.Parse()
	ch := make(chan string)

	go shutdown()
	go watcher(ch)
	debouncer(ch)
}
