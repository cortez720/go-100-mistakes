package main

import (
	"bytes"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

func main() {
	string1 := "sourasdfasdfasdffasdce"
	reader := strings.NewReader(string1)
	dest := bytes.NewBuffer(make([]byte, len(string1)))

	t, _ := reader.Read(dest.Bytes())
	spew.Dump(dest, t)
}
