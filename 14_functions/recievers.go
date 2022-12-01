package main

type customer struct {
	balance float64
}

type data struct {
	balance float64
}

type customerV2 struct {
	Data *data // Highlight that customerV2 is mutable struct
}

func (c customer) add(v float64) { // Copy of object // Must use if we want  to enforce recievers immutability
	c.balance += v // mutate only local c object
}

func (c *customer) addV2(v float64) { // Copy of pointer
	c.balance += v // mutate original object by pointer
}

type clients []customer

func (sl *clients) add(c customer) { // Must use it if we need to mutate a reciever
	*sl = append(*sl, c) // Should be a pointer if reciver is a large pointer
}

func main() {
	c := customer{}
}

//Mixing recievers should be avoided by consensus
