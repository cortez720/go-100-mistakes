package main

import (
	"github.com/davecgh/go-spew/spew"
	"sync"
)

type Getter interface {
	Get(int) string
}

type Setter interface {
	Set(int, string)
}

type GetSetter interface {
	Getter
	Setter
}

func main() {
	var obj GetSetter
	obj = myType{}
	spew.Dump(obj)
}

type myType struct {
	m map[int]string
	sync.Mutex
}

func (t myType) Get(x int) string {
	return t.m[x]
}

func (t myType) Set(x int, s string) {
	t.Lock()
	t.m[x] = s
	t.Unlock()
}
