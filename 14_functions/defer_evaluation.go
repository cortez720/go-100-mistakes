package main

import "fmt"

const (
	StatusSuccess  = "success"
	StatusErrorFoo = "error_foo"
	StatusErrorBar = "error_bar"
)

func main() {
	test()
	_ = f1()
	_ = f2()
	_ = f3()
}

func test(){
	i, j := 0, 0
	defer func (i int){ // Pass i right away, thats why we have 0 on i.
		fmt.Println(i, j)	// Evaluate j when closure is executed.
	}(i)
	i++
	j++
}

func f1() error {
	var status string
	defer notify(status) // Pass args right away, we have empty string here
	defer incrementCounter(status)

	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}

	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}

	status = StatusSuccess
	return nil
}

func f2() error {
	var status string
	defer notifyPtr(&status)  // Pass pointers. It works, but we changed funcs' signatures.
	defer incrementCounterPtr(&status)

	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}

	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}

	status = StatusSuccess
	return nil
}

func f3() error {
	var status string
	defer func() { // Evaluates only anonym func. We pass fresh args to funcs.
		notify(status) // Perfectly works. No need to change the signature.
		incrementCounter(status)
	}()

	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}

	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}

	status = StatusSuccess
	return nil
}

func notify(status string) {
	fmt.Println("notify:", status)
}

func incrementCounter(status string) {
	fmt.Println("increment:", status)
}

func notifyPtr(status *string) {
	fmt.Println("notify:", *status)
}

func incrementCounterPtr(status *string) {
	fmt.Println("increment:", *status)
}

func foo() error {
	return nil
}

func bar() error {
	return nil
}


