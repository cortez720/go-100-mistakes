package main

import (
	"fmt"
	"sync"
)

func main() {
	c := Cache{
		balances: make(map[string]float64),
	}
	c.AddBalance("1", 1.0)
	c.AddBalance("2", 3.0)
	fmt.Println(c.AverageBalance1())
}

type Cache struct {
	mu       sync.RWMutex
	balances map[string]float64
}

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

func (c *Cache) AverageBalance1() float64 {
	c.mu.RLock()
	balances := c.balances
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range balances { // Data race here. Same data in map. Pointer of runtime.hmap struct
		sum += balance
	}
	return sum / float64(len(balances))
}

func (c *Cache) AverageBalance2() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock() // Solution, Just mute all func

	sum := 0.
	for _, balance := range c.balances {
		sum += balance
	}
	return sum / float64(len(c.balances))
}

func (c *Cache) AverageBalance3() float64 {
	c.mu.RLock()
	m := make(map[string]float64, len(c.balances)) // Solution 2. Deep copy of all map.
	for k, v := range c.balances {
		m[k] = v
	}
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range m {
		sum += balance
	}
	return sum / float64(len(m))
}
