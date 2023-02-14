package main

import (
	"testing"
)

var globalValue int
var globalPtr *int

func BenchmarkSumValue(b *testing.B) {
	b.ReportAllocs()
	var local int
	for i := 0; i < b.N; i++ {
		local = sumValue(i, i) // MUCH faster because of the stack.
	}
	globalValue = local
}

func BenchmarkSumPtr(b *testing.B) {
	b.ReportAllocs()
	var local *int
	for i := 0; i < b.N; i++ {
		local = sumPtr(i, i) // Heap MUCH slower.
	}
	globalPtr = local
}

// Therefore, use the pointer then the object has to be shared.
// CPU are extemely efficient with copy data.

func sumValue(x, y int) int {
	z := x + y
	return z
}

func sumPtr(x, y int) *int {
	z := x + y
	return &z
}
