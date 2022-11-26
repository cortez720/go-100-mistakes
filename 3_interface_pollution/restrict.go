package main

type IntConfig struct {
	configValue int
}

func (c *IntConfig) Get() int {
	return c.configValue
}

func (c *IntConfig) Set(x int) {
	c.configValue = x
}

type ReadOnly interface {
	Get() int
}

type ReadOnlyObj struct {
	ReadOnly
}

func NewReadOnlyObj(impl ReadOnly) ReadOnlyObj {
	return ReadOnlyObj{impl}
}

func main() {
	obj := NewReadOnlyObj(&IntConfig{})
	obj.Get()

	raw := IntConfig{configValue: 0}
	obj2 := ReadOnly(&raw)
	obj2.Get()
}