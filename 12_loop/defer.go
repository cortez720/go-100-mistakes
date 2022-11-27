package main

import "os"

func readFiles1(ch <-chan string) error {
	for path := range ch {
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close() // Not closed on each iteration
		// Produced then readFiles1 returns

		// Do something with file
	}
	return nil
}

func readFiles2(ch <-chan string) error {
	for path := range ch {
		// Call the func, that contains file closure
		if err := readFile(path); err != nil { // Possible solution V1.
			return err
		}
	}
	return nil
}

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close() // Produced each time, therefore, on each iteration

	// Do something with file
	return nil
}

func readFiles3(ch <-chan string) error {
	for path := range ch {
		err := func() error { // Possible solution V2. // Anonym function
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			defer file.Close() // Produces then anonym func returns

			// Do something with file
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}
