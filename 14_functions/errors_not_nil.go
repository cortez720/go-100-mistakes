package main

import (
	"errors"
	"log"
	"strings"
)

type MultiError struct {
	errors []string
}

func (m *MultiError) Add(err error) { // It just a syntetick sugar for method. It's equal to Add(m *MultiError, err error).
	//  Hence, it's valid to pass nil like a pointer
	m.errors = append(m.errors, err.Error())
}

func (m *MultiError) Error() string {
	return strings.Join(m.errors, ";")
}

type Customer struct {
	Age  int
	Name string
}

func (c Customer) Validate1() error {
	var m *MultiError

	if c.Age < 0 {
		m = &MultiError{}
		m.Add(errors.New("age is negative"))
	}
	if c.Name == "" {
		if m == nil {
			m = &MultiError{}
		}
		m.Add(errors.New("name is nil"))
	}

	// Nil have the m variable, but we wrap it into error interface. It's error interface, and in it we have nil MultiError pointer.
	return m
	// Best solution: return nil explicitly if we dont have errors: 'return nil'
}

func main() {
	c := Customer{30, "John"} // Valid Customer
	if err := c.Validate1(); err != nil {
		log.Print("invalid validation: %w", err) // Printed nil MultiError, why?
	}
}
