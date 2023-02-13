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

type InputP struct {
	a int64
	_ [56]byte
	b int64
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

func countInputsPedding(inputs []InputP) Result {
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

const counti = 2048

func main() {
	inputs := make([]Input, counti)
	inputsP := make([]InputP, counti)

	t := time.Now()
	countInputs(inputs)
	fmt.Println(time.Now().Sub(t))

	t = time.Now()
	countInputsPedding(inputsP) // sometimes it faster because of false sharding.
	fmt.Println(time.Now().Sub(t))
}
