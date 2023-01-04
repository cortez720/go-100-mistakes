package main

import "os"

func close() {
	f, err := os.OpenFile("filename", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil { // Don't leak, GC get it, but it's better to call Close.
			panic(err)
		}
	}()
}

func writeToFile(content []byte) (err error) {
	f, err := os.OpenFile("filename", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	defer func() { // Can return error if data still living in a buffer.
		if closeErr := f.Close(); err != nil { // Don't leak, GC get it, but it's better to call Close.
			err = closeErr
		}
	}()
	_, err = f.Write(content)
	if err != nil {
		return err
	}

	return nil
}

// But it dont guarantee that the file will be written on disk
// If durability critical factor, we can use the Sync() to commit change.
// In this case errors coming from Close can be safely ignored.

func writeToFile2(content []byte) (err error) {
	f, err := os.OpenFile("filename", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	return f.Sync()
}
