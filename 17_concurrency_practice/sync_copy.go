package main

import "sync"

type Counter struct {
	mu       sync.Mutex // 1st Solution: change mu field to pointer.
	counters map[string]int
}

func NewCounter() Counter {
	return Counter{counters: map[string]int{}} // Init Mutex straightforward here if mu is pointer, overwise mu will be nil
}

func (c Counter) Increment(name string) { // 2nd Solution: change the reciever
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++ // Here we have data race because of the copy of Counter.
}

// Face the issue

// 1. Calling a method with a value reciver
// 2. Calling a function with a sync argument
// 3. Calling a function with an argument that contains a sync field (struct Counter)

// Sync type should never be copied:

// sync.Cond
// sync.Map
// sync.Mutex
// sync.RWMutex
// sync.Once
// sync.Pool
// sync.WaitGroup
