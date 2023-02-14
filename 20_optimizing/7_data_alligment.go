package main

import (
	"fmt"
	"time"
)

// Because of data alliment we will have that:

type Foo1 struct {
	b1 byte // 1 byte
	//	_ [7] byte added 7 bytes
	i int64 // 8 byte
	// 	_ [7] byte added 7 bytes
	b2 byte // 1 byte
}

// Overall we have 24 bytes instead of 10

type Foo2 struct {
	i  int64 // 8 byte
	b1 byte  // 1 byte
	b2 byte  // 1 byte
	// _ [6] byte added 6 bytes.
}

// Overall we have 16 bytes struct.

// All structs must be multiple of biggest field size. Data alligment

const foosCount = 1_000_000

func main() {
	foos1 := make([]Foo1, foosCount)
	t := time.Now()
	sumFoos1(foos1)
	fmt.Println(time.Now().Sub(t))

	foos2 := make([]Foo2, foosCount)
	t = time.Now()
	sumFoos2(foos2) // Must be faster. // reduces the frequency of GC and spatial locality.
	fmt.Println(time.Now().Sub(t))
}

func sumFoos1(foos []Foo1) int64 {
	var sum int64
	for i := range foos {
		sum += foos[i].i
	}

	return sum
}

func sumFoos2(foos []Foo2) int64 {
	var sum int64
	for i := range foos {
		sum += foos[i].i
	}
	return sum
}
