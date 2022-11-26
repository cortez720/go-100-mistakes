package main

import "github.com/davecgh/go-spew/spew"

func main() {
	var s []string //nil, no allocation
	spew.Dump(s)

	s = []string(nil) //nil, no allocation
	spew.Dump(s)

	s = []string{} // not nil, zero elements, allocation
	spew.Dump(s)

	s = make([]string, 0) // not nil, zero elements, allocation
	spew.Dump(s)

	s2 := append([]string(nil), []string{}...) //nil
	spew.Dump(s2)
}
