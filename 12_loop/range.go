package main

import "github.com/davecgh/go-spew/spew"

type account struct {
	balance float64
}

func main() {
	accounts := []account{
		{100.},
		{200.},
		{300.},
	}

	for _, a := range accounts { // Everything (a) here is copy
		a.balance += 1000 // Присваение копии, // Mutates only a in this scope
	}
	
	spew.Dump(accounts)

	for i := range accounts {
		accounts[i].balance += 1000 // Присваение сущ. объекту
	}

	accountsPointers := []*account{
		{100.},
		{200.},
		{300.},
	}

	for _, a := range accountsPointers { // Copy of memory address
		a.balance += 1000 // Mutates in memory addres
	}

	spew.Dump(accounts, accountsPointers)
}

// In general, we should remember that the value element is a range loop is a copy.
// Therefore, if the value is a struct we need to mutate, we will onlyu update the copy, not the element itseld,
// unless the value or filed we modify is a pointer.
