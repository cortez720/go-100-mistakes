package main

import "fmt"

type customConstraint interface {
	~int | ~string
	String() string
}

func getKeys[K customConstraint, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {

	mapa := map[int]string{1: "1", 2: "2", 3: "3"}
	mapa[5] = "5"
	fmt.Println(getKeys(mapa)[0])

}
