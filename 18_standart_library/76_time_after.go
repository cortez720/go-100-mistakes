package main

import (
	"context"
	"log"
	"time"
)

type Event struct{}

func consumer(ch <-chan Event) {
	for {
		select {
		case event := <-ch:
			handle(event)
		case <-time.After(time.Hour): // Memory leak here. The resuources created here released once the timeout expires.
			log.Println("warning: no message received")
		}
	}
}

func consumer2(ch <-chan Event) { // Solution with context with cancelation.
	for {
		ctx, cancel := context.WithCancel(context.Background()) // Creating context everytime. Expensive operation because of the creating channel.
		select {
		case event := <-ch:
			cancel()
			handle(event)
		case <-ctx.Done():
			log.Println("warning: no message received")
		}
	}
}

func consumer3(ch <-chan Event) { // Solution with timer. // Best option.
	timeDuration := time.Hour
	timer := time.NewTimer(timeDuration)
	for {
		timer.Reset(timeDuration) // Less cumberstone operation then creatin new context every time.
		select {
		case event := <-ch:
			handle(event)
		case <-timer.C:
			log.Println("warning: no message received")
		}
	}
	// defer timer.Stop() after timer creation.
}

func handle(Event) {

}
