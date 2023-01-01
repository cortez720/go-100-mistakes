package main

import (
	"fmt"
	"sync"
	"time"
)

func slicerace1() {
	sl := make([]int, 1) // No ploblem

	var sl1, sl2 []int

	go func() {
		sl1 = append(sl, 1)
	}()

	go func() {
		sl2 = append(sl, 1)
	}()

	time.Sleep(time.Millisecond)

	fmt.Println(sl1, sl2)
}

func slicerace2() {
	sl := make([]int, 0, 1) // Problem. Data race

	var sl1, sl2 []int

	go func() {
		sl1 = append(sl, 1)
	}()

	go func() {
		sl2 = append(sl, 1)
	}()

	time.Sleep(time.Millisecond)

	fmt.Println(sl1, sl2)
}

func slicerace3() { // Solution
	sl := make([]int, 0, 1)

	var sl1, sl2 []int

	var wg sync.WaitGroup // Wg for not
	wg.Add(2)

	go func() {
		defer wg.Done()
		sCopy := make([]int, len(sl), len(sl))
		copy(sCopy, sl)

		sl1 = append(sCopy, 1)
	}()

	go func() {
		defer wg.Done()
		sCopy := make([]int, len(sl), len(sl))
		copy(sCopy, sl)

		sl2 = append(sCopy, 1)
	}()

	wg.Wait()

	fmt.Println(sl1, sl2) // Wg for not data race here
}

func main() {
	slicerace3()
}
