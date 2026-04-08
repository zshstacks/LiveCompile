package main



func main() {

	ch := make(chan string)

	go watcher(ch)
	debouncer(ch)
}