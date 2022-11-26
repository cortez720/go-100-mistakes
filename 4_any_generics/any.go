package main

import "fmt"

func main() {
	var i any

	i = 42
	i = "foo"
	i = struct {
		string
	}{string: "bar"}
	i = f

	fmt.Println(i)
}

func f() {}
