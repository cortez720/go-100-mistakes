package main

import (
	"errors"
	"fmt"
	"net/http"
)

type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprint("transient error: %w", t.err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transaction")

	amount, err := getTransactionAmount1(transactionID)
	if err != nil {
		switch err := err.(type) { // Compare types to find the error. // Before Go 1.13.
		case transientError:
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// Write response
	_ = amount
}

func getTransactionAmount1(transactionID string) (float32, error) {
	if len(transactionID) != 5 {
		return 0, fmt.Errorf("id is invalid: %s", transactionID)
	}

	amount, err := getTransactionAmountFromDB1(transactionID)
	if err != nil {
		return 0, transientError{err: err} // Wrap directly to type to define in later.
	}
	return amount, nil
}

func getTransactionAmountFromDB1(id string) (float32, error) {
	return 0, nil
}

func handler2(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transaction")

	amount, err := getTransactionAmount2(transactionID)
	if err != nil { // Go 1.13
		if errors.As(err, &transientError{}) { // errors.As unwraps all error's chain and will succeed if there err with transientErr type
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			// We can use errors.Is if we want to compare expected error's value and the error's value we got
			// It will succeed if there expected value in errors' chain.
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// Write response
	_ = amount
}

func getTransactionAmount2(transactionID string) (float32, error) {
	// Check transaction ID validity

	amount, err := getTransactionAmountFromDB2(transactionID)
	if err != nil {
		return 0, fmt.Errorf("failed to get transaction %s: %w", // We can wrap it as we wish.
			transactionID, err)
	}
	return amount, nil
}

func getTransactionAmountFromDB2(transactionID string) (float32, error) {
	// ...
	var err error
	if err != nil {
		return 0, transientError{err: err} // Wrap error to transient type. No matter where it is.
	}
	// ...
	return 0, nil
}

func main() {}
