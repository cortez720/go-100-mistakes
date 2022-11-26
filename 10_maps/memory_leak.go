package main

import (
	"fmt"
	"runtime"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// Init
	n := 1_000_000
	m := make(map[int]*[128]byte)
	printAlloc()

	// Add elements
	for i := 0; i < n; i++ {
		m[i] = randBytes()

	}
	printAlloc()

	// Remove elements
	for i := 0; i < n; i++ { //Map never shrinks, just delete the elements
		delete(m, i)
	}
	spew.Dump(m[857412]) // Nil or zero values
	// End
	runtime.GC() // With map[int]*[128]byte, we will have nil
	// With map[int][128]byte, we will have empty 128 len array
	printAlloc()
	runtime.KeepAlive(m)
}

func randBytes() *[128]byte {
	return &[128]byte{}
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d MB\n", m.Alloc/1024/1024)
}
