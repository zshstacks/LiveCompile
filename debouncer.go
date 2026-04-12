package main

import "time"

func debouncer(ch chan string) {

	var timer *time.Timer

	for range ch {
		if timer == nil {
			timer = time.AfterFunc(time.Duration(*flagDelay)*time.Millisecond, builder)
		} else {
			timer.Reset(time.Duration(*flagDelay) * time.Millisecond)
		}
	}
}
