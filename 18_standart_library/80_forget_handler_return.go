package main

import "net/http"

func handler(w http.ResponseWriter, req *http.Request) {
	err := foo(req)
	if err != nil {
		http.Error(w, "foo", http.StatusInternalServerError)
		// Don't return the func.
	}

	_, _ = w.Write([]byte("all good")) // Response foo all good
	w.WriteHeader(http.StatusCreated)
}
