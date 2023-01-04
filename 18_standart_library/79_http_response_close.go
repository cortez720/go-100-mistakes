package main

import (
	"io"
	"net/http"
)

func getStatusCode(body io.Reader) (int, error) {
	resp, err := http.Post("url", "application/json", body)
	if err != nil {
		return 0, err
	}

	// We can ignore the response if there is an error.
	// HTTP connection automatically close the body on any error.
	defer func() {
		err := resp.Body.Close() // We must close http body here
		// Overwise we keep some memory allocated, GC can't get resourses
		// And may prevent clients from reusing the TCP connection in the worst cases.
		// If we dont close the body, it may cause memory leak.
		if err != nil {
			panic(err)
		}
	}()

	return resp.StatusCode, nil
}

func getStatusCode2(body io.Reader) (int, error) {
	resp, err := http.Post("url", "application/json", body)
	if err != nil {
		return 0, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	// If we close body after reading, connection keeps living. If dont read, close the connection.
	_, _ = io.Copy(io.Discard, resp.Body) // Read the respone without copyng.

	return resp.StatusCode, nil // Reuse same connection
}

// Should close resourses with closer interface.
