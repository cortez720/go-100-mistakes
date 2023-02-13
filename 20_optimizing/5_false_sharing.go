package main

import (
	"fmt"
	"sync"
	"time"
)

type Input struct {
	a int64
	b int64
}

type ResultP struct {
	sumA int64
	_ [56]byte
	sumB int64
}

type Result struct {
	sumA int64
	sumB int64
}

func countInputs(inputs []Input) Result {
	wg := sync.WaitGroup{}
	wg.Add(2)

	result := Result{}

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}

		wg.Done()
	}()

	wg.Wait()
	return result
}

func countInputsPedding(inputs []Input) ResultP {
	wg := sync.WaitGroup{}
	wg.Add(2)

	result := ResultP{}

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}

		wg.Done()
	}()

	wg.Wait()
	return result
}

const counti = 2048

func main() {
	inputs := make([]Input, counti)

	t := time.Now()
	countInputs(inputs)
	fmt.Println(time.Now().Sub(t))

	t = time.Now()
	countInputsPedding(inputs) // Faster about 50% because of false cache sharding.
	fmt.Println(time.Now().Sub(t))
}
