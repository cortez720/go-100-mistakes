package main

import "fmt"

func main() {
	str := "Hello world"
	fmt.Println(str[:5]) // Print first 5 BYTES, not RUNES

	str = "Привет, мир!"
	fmt.Println(str[:5]) // Print first 5 BYTES, not RUNES

	fmt.Println(string([]rune(str)[:5])) // Print first 5 RUNES
}
