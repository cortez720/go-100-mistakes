package main

import (
	"fmt"
	"time"
)

const n = 1_000_000

func add(s [2]int64) [2]int64 {
	for i := 0; i < n; i++ {
		s[0]++
		if s[0]%2 == 0 {
			s[1]++
		}
	}

	return s
}

func add2(s [2]int64) [2]int64 {
	for i := 0; i < n; i++ {
		v := s[0]
		s[0] = v + 1
		if v%2 != 0 {
			s[1]++
		}
	}

	return s
}

func main() {
	t := time.Now()
	add([2]int64{0, 0})
	fmt.Println(time.Now().Sub(t))

	t = time.Now()
	add2([2]int64{0, 0}) // Faster mainly because of ILP
	fmt.Println(time.Now().Sub(t))
}
