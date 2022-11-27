package main

import "fmt"

func main() {
	m := map[int]bool{
		1: true,
		2: false,
		3: true}

	for k, v := range m {
		if v {
			m[k+10] = true // Map entry created during iteration.
			// It can be produced or skipped during iteration
		}
	}

	fmt.Println(m) // Unpredictable result

	// Possible solution with map copy
	m = map[int]bool{
		1: true,
		2: false,
		3: true}

	m2 := make(map[int]bool, len(m)) // Just copy the map
	for k, v := range m {
		m2[k] = v
	}

	for k, v := range m2 { // Iterate over copy, add to original map
		if v {
			m[k+10] = true
		}
	}

	fmt.Println(m) // Result that we needed
}
