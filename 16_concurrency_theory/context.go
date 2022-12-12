package main

import (
	"context"
	"fmt"
	"time"
)

type impl int

type Position struct {
	Lat float32
	Lng float32
}

type publisher interface {
	Publish(ctx context.Context, position Position) error
}

type publishHandler struct {
	pub publisher
}

func (h publishHandler) publishPosition(position Position) error {
	ctx, cancel := context.WithTimeout(context.Background(), 9*time.Millisecond)
	defer cancel() // Safeguard to don't leave retained objects in memory
	return h.pub.Publish(ctx, position)
}

func (impl) Publish(ctx context.Context, position Position) error {
	deadline, _ := ctx.Deadline()

	time.Sleep(time.Millisecond * 10)

	if deadline.Before(time.Now()) {
		return context.DeadlineExceeded
	}

	fmt.Println(position)

	return nil
}

func main() {
	var p impl
	ph := publishHandler{p}
	err := ph.publishPosition(Position{13.5542, 2123.123213})
	if err != nil {
		fmt.Println(err)
	}
}

// Use context.TODO instead of context.Background if we in doubt to use context
// All contexts in standart library are safe gorutine-safe.