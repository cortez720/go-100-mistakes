package main

import (
	"fmt"
	"runtime"
)

type Fooz struct {
	v []byte
}

func main() {
	foos := make([]Fooz, 1_000)

	printAlloc()

	for i := 0; i < len(foos); i++ {
		foos[i] = Fooz{
			v: make([]byte, 1024*1024),
		}
	}
	printAlloc()

	two := keepFirstTwoElementsOnly(foos)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(two)
}

func keepFirstTwoElementsOnly(foos []Fooz) []Fooz {
	return foos[:2]
}

func keepFirstTwoElementsOnlyCopy(foos []Fooz) []Fooz {
	res := make([]Fooz, 2)
	copy(res, foos)
	return res
}

func keepFirstTwoElementsOnlyMarkNil(foos []Fooz) []Fooz { // We may use when n(2) closer to len(foos) then to 0.
	for i := 2; i < len(foos); i++ {
		foos[i].v = nil
	}
	return foos[:2]
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
