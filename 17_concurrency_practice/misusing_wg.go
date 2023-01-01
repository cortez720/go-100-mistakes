package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func misusing1() {
	var v uint64
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		go func() { // Func won't start until main goroutine ends. wg.Add(1) don't executed.
			wg.Add(1)
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v) // 1-3 random.
}

func misusing2() { // Solution
	var v uint64
	wg := sync.WaitGroup{}

	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}

func misusing3() { // Solution 2
	var v uint64
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint64(&v, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(v)
}
