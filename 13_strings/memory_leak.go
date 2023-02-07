package main

import (
	"errors"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type store struct {
	data   map[string]string
	lastID int
}

func main() {
	s := store{data: make(map[string]string)}
	s.handleLogV2(strings.Repeat("A", 50))
	spew.Dump(s)
}

func (s *store) handleLog(log string) error {
	if len(log) < 36 {
		return errors.New("invalid format")
	}

	s.store(log[:36]) // The log message can consists thousand of bytes, but we keep whole backing array in memory

	return nil
}

func (s *store) handleLogV2(log string) error {
	if len(log) < 36 {
		return errors.New("invalid format")
	}

	s.store(string([]byte(log[:36]))) // we keep only 36 len string with new byte slice allocation
	s.store(strings.Clone(log[:36]))  // 36 len string copy

	return nil
}

func (s *store) store(str string) {
	s.lastID++
	s.data[string(s.lastID)] = str
}
