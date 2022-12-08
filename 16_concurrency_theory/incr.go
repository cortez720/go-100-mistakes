package main

import (
	"fmt"
	"sync"
)

// Atomics

// Mutex

// Data race (access to 1 memory address at the same time)
func f1() {
	var i int
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	go func() {
		ch <- 1
	}()

	i += <-ch
	i += <-ch

	fmt.Println(i)
}

// Not determine result
// Either 1 or 2

// Race condition (not solid access order)

func f2() {
	var m sync.Mutex
	var i int

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		m.Lock()
		defer func(){
			m.Unlock()
			wg.Done()
		}()
		i = 1
	}()

	go func() {
		m.Lock()
		defer func(){
			m.Unlock()
			wg.Done()
		}()
		i = 2
	}()
	
	wg.Wait()

	fmt.Print(i)
}

func main() {
	f2()
}
