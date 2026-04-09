package main

import (
	"os"
	"os/signal"
	"syscall"
)

func shutdown() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
	killBinary()
	os.Exit(0)
}
