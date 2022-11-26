package main

import (
	"fmt"
	"sort"
)

type SliceFn[T any] struct {
	S      []T
	LessFn func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.S)
}
func (s SliceFn[T]) Less(i, j int) bool {
	return s.LessFn(s.S[i], s.S[j])
}
func (s SliceFn[T]) Swap(i, j int) {
	s.S[i], s.S[j] = s.S[j], s.S[i]
}

func main() {
	s := SliceFn[float64]{
		S: []float64{3.4654654, 2.465465465, 1.646545465},
		LessFn: func(a, b float64) bool {
			return a < b
		},
	}
	sort.Sort(s)
	fmt.Println(s.S)
}
