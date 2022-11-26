package main

import (
	"fmt"
)

// type SliceHeader struct {
//	Data uintptr
//	Len  int
//	Cap  int
//}

func main() {
	sl := make([]int, 3, 6) // Len: 3, Capacity: 6. len(arr) = 6, but Len = 3, Cap = 6 in runtime.
	// Last elements are allocated but not used

	sl[0], sl[1], sl[2] = 1, 1, 1 // Array is [1,1,1,0,0,0], but slice is [1,1,1]

	sl2 := sl[2:6] //Data uintptr to third element, array stay the same, Cap: 4, Len:4

	fmt.Println(sl, sl2) // [1 1 1] [1 1 1 0 0 0]

	sl = append(sl, 2, 2, 2, 2) //Creates another, doubling the capacity, copying all the elements. Now different array to sl1

	fmt.Println(sl, sl[0:12]) // [1 1 1 2 2 2 2] [1 1 1 2 2 2 2 0 0 0 0 0] // Len: 7, Cap: 12

}

// To summarize, the slice length is the number of available elements in the slice,
// whereas the slice capacity is the number of elements in the backing array.