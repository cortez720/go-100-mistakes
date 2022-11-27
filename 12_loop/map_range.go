package main

import "fmt"

type Customer struct {
	ID      string
	Balance int
}

type Store struct {
	m map[string]*Customer
}

func main() {
	store := Store{m: make(map[string]*Customer)}
	customers := []Customer{{"1", 10}, {"2", 20}, {"3", 0}}

	for _, customer := range customers { // Create copy of object and redeclate it each time
		store.m[customer.ID] = &customer // Same pointer every time
	}

	fmt.Println(store) // Hence same pointer, balance: 0 and ID: 3 in each obj

	for _, customer := range customers { // V1 possible solution
		current := customer // Create local obj copy // Initialize during each iteration
		// that's why it has unique address
		store.m[customer.ID] = &current
	}

	fmt.Println(store) // Different poiners, different values

	for i := range customers { // Best possible solution with not relying on slice
		customer := &customers[i] // Initialize during each iteration, unique address
		store.m[customer.ID] = &customers[i]
	}

	fmt.Println(store) // Different poiners, different values
}
