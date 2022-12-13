package main

import (
	"fmt"
	"time"
)

func listing_1() {
	s := []int{1, 2, 3}

	for _, i := range s {
		go func() {
			fmt.Print(i) // Got the value out of closure // 333, 221 ...
		}()
	}
}

func listing_2() {
	s := []int{1, 2, 3}

	for _, i := range s {
		val := i // In each iteration we create the new variable. That's why it's not rewrites.
		go func() {
			fmt.Print(val)
		}()
	}
}

func listing_3() {
	s := []int{1, 2, 3}

	for _, i := range s {
		go func(i int) {
			fmt.Print(i)
		}(i)
	}
}

func main(){
	listing_1()

	time.Sleep(time.Millisecond*1)
}