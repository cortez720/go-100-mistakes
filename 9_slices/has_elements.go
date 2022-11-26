package main

import "github.com/davecgh/go-spew/spew"

func main() {
	var sl []string
	sl2 := []string{}
	spew.Dump(len(sl) == len(sl2)) // True

	var m map[string]int
	m2 := map[string]int{}
	spew.Dump(len(m) == len(m2)) // True
}
