package main

import (
	"fmt"
	"sync"
)

func fn0() {
	messageCh := make(chan int, 10)
	disconnectCh := make(chan struct{}, 1)

	for i := 0; i < 10; i++ {
		messageCh <- i
	}

	disconnectCh <- struct{}{}

	for {
		select {
		case m := <-messageCh: // Random selection between cases.
			fmt.Printf("message: %v\n", m)
		case <-disconnectCh: // Not determenistic order.
			fmt.Println("disconnection, return")
			return
		}
	}
}

// Prevent selection from only one channel.

// We can just make channels unbuffered, and will this solve the problem.

func fn1() { // Solution with buffered channels
	messageCh := make(chan int, 10)

	for i := 0; i < 10; i++ {
		messageCh <- i // There is only written values before closing channnel.
	}

	close(messageCh) // Close the channel.

	for {
		select {
		case m, open := <-messageCh: // Random selection between cases.
			if !open {
				fmt.Println("disconnection, return")
				return
			}
			fmt.Printf("message: %v\n", m)
		}
	}

}

type Message struct {
	message    int
	disconnect bool
}

func fn2() { // Solution with buffered channels 2
	messageCh := make(chan Message, 10)

	for i := 0; i < 10; i++ {
		if i == 9 {
			messageCh <- Message{message: i, disconnect: true}
			continue
		}
		messageCh <- Message{message: i} // There is only written values before closing channnel.
	}

	for {
		select {
		case m := <-messageCh:
			fmt.Println(m.message)
			if m.disconnect {
				return
			}
		}
	}

}

func fn3() { // Solutions with multiple goroutines.
	messageCh := make(chan int, 30)
	disconnectCh := make(chan struct{}, 1)

	var wg sync.WaitGroup

	wg.Add(3)

	go writter(0, 10, messageCh, &wg)
	go writter(10, 20, messageCh, &wg)
	go writter(20, 30, messageCh, &wg)

	wg.Wait()

	disconnectCh <- struct{}{}

	for {
		select {
		case m := <-messageCh:
			fmt.Printf("message: %v\n", m)
		case <-disconnectCh: // Random select if both message and disocnnect channels are ready.
			for {
				select {
				case m := <-messageCh: // select Chose that if message is ready. // We will miss message if sent to chan after goroutine is returned.
					fmt.Printf("message: %v\n", m)
				default: // Disconnect after recieving all messages.
					fmt.Println("disconnection, return")
					return
				}
			}
		}
	}
}

func writter(low, high int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := low; i < high; i++ {
		ch <- i
	}
}

func main() {
	fn3()
}
