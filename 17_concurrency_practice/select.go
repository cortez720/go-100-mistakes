package main

import "fmt"

func fn() {
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

func fn2() { // Solution with buffered channels
	messageCh := make(chan int, 10)

	for i := 0; i < 10; i++ {
		messageCh <- i
	}

	close(messageCh) // Close the channel.

	for {
		select {
		case m, open := <-messageCh: // Random selection between cases.
			if !open{
				fmt.Println("disconnection, return")
				return
			}
			fmt.Printf("message: %v\n", m)
		}
	}

}

func main(){
	fn2()
}