package main

import "fmt"

type LargeStruct struct {
	foo string
}

func main() {
	var m = make(map[int]LargeStruct, 1)
	m[0] = LargeStruct{"0"}
	fmt.Println(m)
	MutateMapWithCopy(m)
	fmt.Println(m)
}

func MutateMapWithCopy(m map[int]LargeStruct) {
	value := LargeStruct{"1"}
	m[0] = value
	// m[0].foo = "1" // UnaddressableFieldAssign occurs when trying to assign to a struct field
}

func MutateMap(m map[int]*LargeStruct) {
	m[0].foo = "1" // Allow to mutate with pointers
}
