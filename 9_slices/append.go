package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:2]
	_ = append(s2, 10) // Modifid original array from s1 slice

	fmt.Println(s1) // [1 2 10]

	s2 = []int{1, 2, 3}
	s3 := s2[:2:2]      // Point out max (capacity) // [low:high:max]
	s3 = append(s3, 10) // Create another slice with new array.
	fmt.Println(s2, s3) // [1 2 3] [1 2 10]
}

// While append if len == capcaity, go creates new array and copy all elements.
