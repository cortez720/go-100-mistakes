package main

import "fmt"

func nil_0() {
	var ch chan int

	go func() {
		ch <- 1 // panic. write to nil
	}()

	for {
		fmt.Println(<-ch) // panic. read from nil
	}
}

func merge(ch1, ch2 chan int) chan int { // Merge without pay attention to closing.
	ch := make(chan int)

	go func() {
		for {
			select {
			case v := <-ch1:
				ch <- v
			case v := <-ch2: // Read from closed channel, not from nil channel.
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}

func nil_1() {
	var ch1 chan int // Nil channel.
	ch2 := make(chan int, 5)

	for i := 0; i < 5; i++ {
		ch2 <- i
	}
	close(ch2)

	ch := merge(ch1, ch2)

	for {
		select {
		case v, open := <-ch:
			if !open {
				return
			}
			fmt.Printf("value: %d\n", v)
		}
	}
}

func merge1(ch1, ch2 chan int) chan int {
	ch := make(chan int)

	ch1Closed, ch2Closed := false, false

	go func() {
		for {
			select {
			// We will serve both channels, even one of them closed. We waste CPU cycles here.
			case v, open := <-ch1: // Infinity loop for closed channel.
				if !open {
					ch1Closed = true
					break
				}

				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2Closed = true
					break
				}
				ch <- v
			}

			if ch1Closed && ch2Closed {
				close(ch)
				return
			}
		}
	}()

	return ch
}

func nil_2() {
	ch1 := make(chan int, 5)
	close(ch1)

	ch2 := make(chan int, 5)

	for i := 0; i < 5; i++ {
		ch2 <- i
	}
	close(ch2)

	ch := merge1(ch1, ch2)

	for {
		select {
		case v, open := <-ch:
			if !open {
				return
			}
			fmt.Printf("value: %d\n", v)
		}
	}
}

func merge2(ch1, ch2 chan int) chan int {
	ch := make(chan int)

	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case v, open := <-ch1: // Select don't read from nil channels. That's why we save a CPU time here.
				if !open {
					ch1 = nil // We remove one case from serving here.
					break
				}

				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2 = nil
					break
				}
				ch <- v
			}
		}
		close(ch)

	}()

	return ch
}

func nil_3() {
	var ch1 chan int // Nil channel.
	ch2 := make(chan int, 5)

	for i := 0; i < 5; i++ {
		ch2 <- i
	}
	close(ch2)

	ch := merge2(ch1, ch2)

	for {
		select {
		case v, open := <-ch:
			if !open {
				return
			}
			fmt.Printf("value: %d\n", v)
		}
	}
}

func main() {
	nil_3()
}

// Using nil channels can help to save CPU time.
// Nil channels should be part of Go developers toolset when dealing with concurrent code.
