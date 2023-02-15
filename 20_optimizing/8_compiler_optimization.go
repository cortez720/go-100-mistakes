package main

type cache struct {
	m map[string]int
}

func (c *cache) get(bytes []byte) (v int, contains bool) {
	key := string(bytes)
	v, contains = c.m[key]
	return
}

func (c *cache) get1(bytes []byte) (v int, contains bool) {
	v, contains = c.m[string(bytes)] // Faster because of Go compiler optimization. Compiler avoid doing  bytes-to-string conversation.
	return
}
