package main

import (
	"fmt"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	var sl []string
	sl2 := []string{}
	spew.Dump(len(sl) == len(sl2)) // True

	var m map[string]int
	m2 := map[string]int{}
	spew.Dump(len(m) == len(m2)) // True

	fmt.Println("sl", unsafe.Sizeof(sl), unsafe.Sizeof(sl2))
	spew.Dump(sl, sl2)
	fmt.Println("m", unsafe.Sizeof(m), unsafe.Sizeof(m2))
	spew.Dump(m, m2)
}
