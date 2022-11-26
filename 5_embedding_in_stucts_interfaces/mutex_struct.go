package main

import (
	"io"
	"os"
	"sync"
)

func main() {
	obj := myType1{}
	obj.Lock() //Не правильно, доступно извне

	//obj3 := myType3{}
	//	obj3.Lock() //Правильно, не доступно извне, т.к. инкапсуляция

	logger := Logger{WriteCloser: os.Stdout}                     //Композиция
	logger.Write([]byte{75, 55, 84, 80, 14, 80, 10, 36, 89, 42}) //Правильное использование вложенныйх структур
	logger.Close()
}

type myType1 struct {
	m map[int]string
	sync.Mutex
}

func (t myType1) Get(x int) string {
	return t.m[x]
}

func (t myType1) Set(x int, s string) {
	t.Lock()
	t.m[x] = s
	t.Unlock()
}

type myType3 struct {
	m  map[int]string
	mu sync.Mutex
}

func (t myType3) Get(x int) string {
	return t.m[x]
}

func (t myType3) Set(x int, s string) {
	t.mu.Lock()
	t.m[x] = s
	t.mu.Unlock()
}

//-------------------------------

type Logger struct {
	io.WriteCloser
}
