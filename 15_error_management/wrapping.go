package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func bar() error {
	return barError{}
}

type barError struct{}

func (b barError) Error() string {
	return "bar error"
}

func listing1() error {
	err := bar()
	if err != nil {
		return err // Without wrapping
	}
	// ...
	return nil
}

type BarError struct {
	Err error
}

func (b BarError) Error() string {
	return "bar failed:" + b.Err.Error()
}

func listing2() error {
	err := bar()
	if err != nil {
		return BarError{Err: err} // New type of error. Possible to mark.
	}
	// ...
	return nil
}

func listing3() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("bar failed: %w", err) // Wrapping with fmt %w. With Go 1.13.
	}
	// ...
	return nil
}

func listing4() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("bar failed: %v", err) // Wrapping with fmt %v. Not visible source.
	}
	return nil
}

func main() {
	spew.Dump(listing1())
	spew.Dump(listing2())
	spew.Dump(listing3())
	spew.Dump(listing4())
}
