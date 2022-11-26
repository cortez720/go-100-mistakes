package main

import (
	"github.com/davecgh/go-spew/spew"
	"sort"
)

type myNumbers []myNumber

type myNumber float64

func main() {
	numbers := myNumbers{1.0, 0.1, 0.165465, 0.85564, 0.9889798, 500.0, 100.635486486, 165.65432131, 88.98462}
	spew.Dump(numbers, sort.IsSorted(numbers))
	sort.Sort(numbers)
	spew.Dump(numbers, sort.IsSorted(numbers))
}
func (slice myNumbers) Len() int {
	return len(slice)
}

func (slice myNumbers) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice myNumbers) Swap(i, j int) {
	tmp := slice[j]
	slice[j] = slice[i]
	slice[i] = tmp
}
