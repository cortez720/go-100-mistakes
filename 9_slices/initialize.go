package main

type Foo int
type Bar string

func main() {

}

func convert(foos []Foo) []Bar {
	bars := make([]Bar, 0) // Unifficent init

	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}

	return bars
}

func convertV2(foos []Foo) []Bar {
	bars := make([]Bar, 0, len(foos)) // Norm init

	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}

	return bars
}

func convertV3(foos []Foo) []Bar {
	bars := make([]Bar, len(foos)) // Norm init 2

	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}

	return bars
}

func fooToBar(Foo) Bar {
	return "converted"
}
