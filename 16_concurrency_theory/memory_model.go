package main

import "fmt"

func listing1() {
	i := 0
	go func() {
		i++
	}()
}

func listing2() {
	i := 0
	go func() {
		i++
	}()
	fmt.Println(i)
}

// This patterns prtoect from data race, but not from race condition.

func listing3() {
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch
		fmt.Println(i)
	}()
	i++
	ch <- struct{}{}
}

func listing4() {
	i := 0
	ch := make(chan struct{})
	go func() {
		<-ch // Race condition, but not data race
		fmt.Println(i)
	}()
	i++
	close(ch)
}

func listing5() {
	i := 0
	ch := make(chan struct{}, 1) // Buffered. Data race
	go func() {
		i = 1
		<-ch
	}()
	ch <- struct{}{}
	fmt.Println(i)
}

func listing6() {
	i := 0
	ch := make(chan struct{}) // No buf, no data race
	go func() {
		i = 1
		<-ch
	}()
	ch <- struct{}{}
	fmt.Println(i)
}

func main() {
	listing4()
}
