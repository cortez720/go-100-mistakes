package main

import (
	"github.com/davecgh/go-spew/spew"
	"math"
)

func main() {
	var n = math.MaxInt
	spew.Dump(n)
	n++
	spew.Dump(n)
}
