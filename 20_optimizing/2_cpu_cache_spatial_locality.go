package main

import (
	"fmt"
	"time"
)

type Foo struct {
	a int64
	b int64
}

func sumFoo(foos []Foo) int64 {
	var total int64
	for i := 0; i < len(foos); i++ {
		total += foos[i].a
	}
	return total
}

type Bar struct {
	a []int64
	b []int64
}

func sumBar(bar Bar) int64 {
	var total int64
	for i := 0; i < len(bar.a); i++ {
		total += bar.a[i]
	}
	return total
}

func main() {
	sl := make([]Foo, 4096)

	start := time.Now()
	sumFoo(sl)
	fmt.Println(2, time.Now().Sub(start))

	var b Bar
	b.a = make([]int64, 4096)
	b.b = make([]int64, 4096)

	start = time.Now()
	sumBar(b) // About 2 times faster. the reason is better spatial locality. 
	// CPU fetch fewer cache lines from memory.
	fmt.Println(8, time.Now().Sub(start))
}
