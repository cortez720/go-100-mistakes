package main

import "fmt"

func main() {
	s := "hêllo"
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s)) // Len 6, but 5 chars. Ã instead ê

	for i, r := range s { // first approach
		fmt.Printf("position %d: %c\n", i, r) // Use the copy of value, to print real rune.
	}

	print("\n")

	rs := []rune(s)
	for i := range rs { // second approach
		fmt.Printf("position %d: %c\n", i, rs[i]) //Print slice of runes
	}

	fmt.Println(len("Ъ")) // Два байта
}
