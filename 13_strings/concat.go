package main

import (
	"fmt"
	"strings"
)

func main() {
	concatV3([]string{"10", "20", "30"})
}

func concat(strs []string) { // ~40 times slower then builder with Grow
	res := ""
	for _, s := range strs {
		res += s // New allocation each time
	}

	fmt.Println(res)
}

func concatV2(strs []string) { // Optimize solution with builder without Grow // 78% slower
	res := strings.Builder{}
	for _, s := range strs {
		res.WriteString(s) // No unneccessary allocations
	}

	fmt.Println(res.String())
}

func concatV3(strs []string) {
	total := 0
	for _, s := range strs {
		total += len(s)
	}

	res := strings.Builder{}
	res.Grow(total) // Allocate neccessary memory
	for _, s := range strs {
		res.WriteString(s) // No unneccessary allocations
	}

	fmt.Println(res.String())
}

// Rule of thumb: use builder if we want to concat more then 5 strings
