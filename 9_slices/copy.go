package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2, 3, 4} // 4 Len
	dst := []int(nil)        // 0 Len

	copy(dst, src)          // nil, [1,2,3,4]
	fmt.Println(dst == nil) // true

	dst = make([]int, 0, 3)

	copy(dst, src)   // [1,2,3,4] []int{}
	fmt.Println(dst) // Empty slice with 3 Cap

	dst = make([]int, 3, 4)

	copy(dst, src)        // [1,2,3,4] [0,0,0]
	fmt.Println(dst)      // 1,2,3 slice.
	fmt.Println(dst[0:4]) // 4 is gone

	// In conclusion important only len

	dst = src
	dst[1] = 1
	src[1] = 0
	fmt.Println(dst, src) // Same array under SliceHeader

	dst = append([]int(nil), src...)
	dst[1] = 1
	src[1] = 0
	fmt.Println(dst, src) // Different array under SliceHeader
}
