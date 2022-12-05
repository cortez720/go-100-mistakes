package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func sequentialMergesort(s []int) {
	if len(s) == 1 {
		return
	}

	middle := len(s) / 2
	sequentialMergesort(s[:middle])
	sequentialMergesort(s[middle:])
	merge(s, middle)
}

func merge(s []int, middle int) {
	helper := make([]int, len(s))
	copy(helper, s)

	helperLeft := 0
	helperRight := middle
	current := 0
	high := len(s) - 1

	for helperLeft <= middle-1 && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			s[current] = helper[helperLeft]
			helperLeft++
		} else {
			s[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	for helperLeft <= middle-1 {
		s[current] = helper[helperLeft]
		current++
		helperLeft++
	}
}

// If workload is too slow, it will be work slower.
func parallelMergesortV1(s []int) { // 8 times slower
	if len(s) == 1 {
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)

	middle := len(s) / 2
	go func() {
		defer wg.Done()
		parallelMergesortV1(s[:middle])
	}()

	go func() {
		defer wg.Done()
		parallelMergesortV1(s[middle:])
	}()

	wg.Wait()
	merge(s, middle)
}

const max = 2048

func parallelMergesortV2(s []int) {
	if len(s) == 1 {
		return
	}

	var wg sync.WaitGroup

	middle := len(s) / 2
	if len(s) < max { // Optimize with constant. // Use the sequential if len < const.
		sequentialMergesort(s)
	} else {

		wg.Add(2)

		go func() {
			defer wg.Done()
			parallelMergesortV2(s[:middle])
		}()

		go func() {
			defer wg.Done()
			parallelMergesortV2(s[middle:])
		}()

	}

	wg.Wait()
	merge(s, middle)
}

const log = 4

func main() {
	value := int(math.Pow(2, log))

	slice := rand.Perm(value)
	slice2 := make([]int, 0, value)
	slice3 := make([]int, 0, value)

	copy(slice2, slice)
	copy(slice3, slice)

	start := time.Now()
	sequentialMergesort(slice)
	fmt.Print(slice[:10])
	fmt.Println(time.Now().Sub(start))

	start = time.Now()
	parallelMergesortV1(slice2)
	fmt.Print(slice2[:10])
	fmt.Println(time.Now().Sub(start))

	start = time.Now()
	parallelMergesortV2(slice3)
	fmt.Print(slice3[:10])
	fmt.Println(time.Now().Sub(start))

}
