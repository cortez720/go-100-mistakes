package main

import "fmt"

func listing1() {
	a := [3]int{0, 1, 2}
	for i, v := range a {
		a[2] = 10
		if i == 2 {
			fmt.Println(v)
		}
	}
}

func listing2() {
	a := [3]int{0, 1, 2}
	for i := range a {
		a[2] = 10
		if i == 2 {
			fmt.Println(a[2])
		}
	}
}

func listing3() {
	a := [3]int{0, 1, 2}
	for i, v := range &a { //We assign a copy of array pointer. 
		// Both pointers reference the same arrray, accesing v also return 10.
		a[2] = 10
		if i == 2 {
			fmt.Println(v)
		}
	}
}

func main(){
	listing1()
	listing2()
	listing3()
}

