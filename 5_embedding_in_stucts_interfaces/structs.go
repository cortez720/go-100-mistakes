package main

import "github.com/davecgh/go-spew/spew"

type Foo struct {
	Bar
}

type Bar struct {
	Value int
}

func main() {
	obj := Foo{}
	obj.Value = 5
	spew.Dump(obj)
}
