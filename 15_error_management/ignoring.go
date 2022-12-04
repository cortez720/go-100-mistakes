package main

func main() {
	// Ignoring the error, but the code reader can think that we just skipped it
	notify() // Bad ignoring

	// Good practice is to comment this.
	// Code reader sure that we ignore error on purpose with blank identifier.
	_ = notify() // Best ignoring.
}

func notify() error {
	return nil
}
