package main

import (
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		ReadHeaderTimeout: time.Second,                                            // maximum amount of time to read the request headers
		ReadTimeout:       time.Second,                                            // maximum amount of time to read the entire request
		Handler:           http.TimeoutHandler(handler, time.Second*5, "message"), // Wrapper func to specify the maximum amount of time for a handler to complete.
		IdleTimeout:       time.Second * 5,                                        // Maxixmum amount of time fot the next request.
	}

}

// For production-grade applications we need to make sure not to use default HTTP clients and servers.
// Otherwise, requests may be stuck forever die to an absence  of timeouts or ever malicious clients exploit it.
