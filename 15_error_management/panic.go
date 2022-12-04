package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r) // Panic was caught // Doesn't stop the goroutine.
		}
	}()

	f()
}

func f() {
	fmt.Println("A")
	panic("foo")
	fmt.Println("B")
}

// Two cases to use panic
// 1. To signal programmer error: HTTP code < 100 || > 999
// 2. Application fails to make mandatory dependecy.
