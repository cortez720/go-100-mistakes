package main

import "fmt"

func fn() {
	disconnectCh := make(chan bool) // Bad pattern

	// If false, connection continues

	go func() {
		disconnectCh <- true
	}()

	if <-disconnectCh {
		fmt.Println("disconnected, return")
		return
	}
	// Perhaps, we should only expect true for continue connection
}

// We need a channel wihtout data if we dont need specific informaiton.

func emptyStruct() {
	disconnectCh := make(chan struct{}) // Good pattern

	go func() {
		// Is a de facto standard to convey an absence of meaning
		disconnectCh <- struct{}{} // Size of empty struct is 0 bytes. // Size of empty interface is 8 or 16 bytes.
	}()

	select {
	case <-disconnectCh:
		fmt.Println("disconnected, return")
		return
	}

}

// With Go standarts, if channel without data, we should use empty struct{}
// In go these channels called Notification channels.

func main() {
	emptyStruct()
}
