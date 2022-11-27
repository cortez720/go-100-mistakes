package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}

	for range s { // provided expression is evaluated only once
		// s copy to tmp varualable and iterates over it
		s = append(s, 10) // original slice also updates
	}

//	for i:= 0; i< len(s); i++{ // Loop never ends, because len evaluate each time
//		s = append(s, 20)
//	}

	fmt.Println(s) // accountsPointers
}
