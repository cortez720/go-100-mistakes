package main

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
)

type MultiErrorSl struct {
	errors []error
}

func (m MultiErrorSl) Add(err error) {
	m.errors = append(m.errors, err)
}

func main() {
	me := MultiErrorSl{errors: make([]error, 0, 1)} // Allocate memory, thats why we append element in existing array under slice
	// If there are no memory, new array will be allocated in a copy of struct and we will not have error there.
	err := errors.New("error 1")
	me.Add(err) // Copy with same array under slice, but not samae len value in slice struct
	spew.Dump(me.errors[:1])
}
