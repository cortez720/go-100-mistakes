package main

import "fmt"

func main() {
	s := Struct{id: "foo"}
	defer s.print() 
	defer s.printPrt()
	s.id = "bar"
}

type Struct struct {
	id string
}

func (s Struct) print() { // Prints foo. Func immidiately evaluates with copy. Same as print(s Struct)
	fmt.Println(s.id)
}

func (s *Struct) printPrt() {
	fmt.Println(s.id) // Prints bar. Func immidiately evaluates with copy of pointer. Same as print(s *Struct)
}