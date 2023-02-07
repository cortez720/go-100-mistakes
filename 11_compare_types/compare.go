package main

import (
	"fmt"
	"reflect"
)

type customer struct{
	id int
	name string
}

type customer2 struct{
	id int
	name string
	sl []float64 // cannot be compared
}

func main(){
	c1 := customer{1, "Vanya"}
	c2 := customer{1, "Vanya"}

	fmt.Println(c1==c2)

	// var a any = []float64{}
	// var b any = []float64{}

	// fmt.Println(a==b) // panic: runtime error: comparing uncomparable type []float64

	c3 := customer2{1, "Vanya", []float64{1.} }
	c4 := customer2{1, "Vanya", []float64{1.}}

	fmt.Println(reflect.DeepEqual(c3, c4))
	
	fmt.Println(ourOwnComprasion(c3, c4))
}

func ourOwnComprasion(c, c2 customer2) bool{ // about 100 faster then reflect.DeepEqual
	if c.id != c2.id{
		return false
	}

	if c.name != c2.name{
		return false
	}

	if len(c.sl) != len(c2.sl){
		return false
	}

	for i := 0; i < len(c.sl); i++{
		if c.sl[i] != c2.sl[i]{
			return false
		}
	}

	return true
}