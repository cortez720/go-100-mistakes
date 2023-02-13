package main

import (
	"fmt"
	"time"
)

func main() {
	sl := make([]int64, 10000000)

	start := time.Now()
	sum2(sl)
	fmt.Println(2, time.Now().Sub(start))

	sl = make([]int64, 10000000)
	start = time.Now()
	sum8(sl)
	fmt.Println(8, time.Now().Sub(start)) // about 2 times faster then sum8, but expected 4. The reason is related to cache lines
}

func sum2(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 2 {
		total += s[i]
	}
	return total
}

func sum8(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 8 {
		total += s[i]
	}
	return total
}
