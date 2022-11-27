package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		switch i {
		default:
		case 2:
			break // Here we brake only switch
		}
	}

	fmt.Println()

loop:
	for i := 0; i < 5; i++ {
		switch i {
		default:
		case 3:
			break loop // Here we brake the loop with label
		case 1:
			continue loop // Skip 1 printing with continue
		}
		fmt.Println(i)
	}
}

// Lables used in standart library
// We can aslo use continue with labels
// Same thing work with select statement
