package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func read(r io.Reader) (int64, error) {
	var count int64
	wg := sync.WaitGroup{}
	var n = 10

	ch := make(chan []byte, n)
	wg.Add(n)

	for i := 0; i < n; i++ { // Async task processing from channel.
		go func() {
			defer wg.Done()
			for b := range ch {
				v := task(b)
				atomic.AddInt64(&count, int64(v))
			}
		}()
	}

	for { // Write to 10 gorutines.
		b := make([]byte, 1024)
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return count, err
			}
		}
		ch <- b
	}

	close(ch)
	wg.Wait()

	return count, nil
}

func task(b []byte) int {
	f, err := os.OpenFile("data1", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	count, err := f.Write(b)
	if err != nil {
		fmt.Println(err)
	}

	return count
}

func main() {
	t := time.Now()
	r := strings.NewReader(strings.Repeat("abcdefghABCDEFGH12345678mnghcbrt\n", 64000000))
	count, err := read(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(count, time.Since(t))
}
