package main

import "time"

func debouncer(ch chan string) {

	var timer *time.Timer

	for range ch {
		if timer == nil {
			timer = time.AfterFunc(300 * time.Millisecond, builder)
		} else {
				timer.Reset(300 * time.Millisecond)
		}
	}
}