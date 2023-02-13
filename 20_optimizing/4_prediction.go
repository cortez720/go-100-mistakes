package main

import (
	"fmt"
	"time"
)

type node struct {
	value int64
	next  *node
}

func linkedList(n *node) int64 {
	var total int64
	for n != nil {
		total += n.value
		n = n.next
	}

	return total
}

func slice(sl []int64) int64 {
	var total int64
	for i := 0; i < len(sl); i += 2 {
		total += sl[i]
	}

	return total
}

const count = 512

func main() {
	var first node
	first.fillNode(count)

	t := time.Now()
	linkedList(&first) // non-unit stride. Unpredictable for CPU
	fmt.Println(time.Now().Sub(t))

	sl := make([]int64, count)
	t = time.Now()
	slice(sl) // constant stride, every 2 elements. Predictable for CPU
	fmt.Println(time.Now().Sub(t))

	// Similar spatial location but not similar strides.

}

func (first *node) fillNode(count int) {
	var n *node
	for i := 0; i < count+1; i += 1 {
		oldNode := n

		if i == count {
			n = first
		} else {
			n = &node{}
		}

		n.value = 0
		n.next = oldNode
	}
}
