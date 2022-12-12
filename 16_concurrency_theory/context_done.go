package main

import (
	"context"
	"fmt"
	"time"
)

func handler(ctx context.Context, ch chan string) {
	for {
		select {
		case m := <-ch:
			fmt.Println(m)
		case <-ctx.Done():
			panic(ctx.Err())
		}
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100)

	ch := make(chan string)
	m := "message"

	go func() {
		handler(ctx, ch)
	}()

	time.Sleep(80 * time.Millisecond)

	ch <- m

	time.Sleep(10 * time.Millisecond)

	ch <- m
	ch <- m

	time.Sleep(15 * time.Millisecond) // Sleep here to not exit too early. Wait to print all messages.
}
