package main

import "strings"

type store map[string]string

func main{

}

func (s store) handleLog(log string) error{
	if len(log) < 36{
		return errors.New("invalid format")
	}

	s.store(log[:36]) // The log message can consists thousand of bytes, but we keep whole backing array in memory
}

func (s store) handleLogV2(log string) error{
	if len(log) < 36{
		return errors.New("invalid format")
	}

	s.store(string(byte(log[:36]))) // we keep only 36 len string with new byte slice allocation
	s.store(strings.Clone(log[:36])) // 36 len string copy
}

func (s store) store(str string){

}