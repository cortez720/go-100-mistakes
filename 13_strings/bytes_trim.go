package main

import (
	"bytes"
	"io"
	"strings"
)

func getBytes(reader io.Reader) ([]byte, error) {
	res, err := io.ReadAll(reader)
	if err != nil {

		return nil, err
	}
	// Big price of additional conversions
	return []byte(sanitize(string(res))), nil // A lot of unnceccesery allocations, slower code.
}

func sanitize(s string) string {
	return strings.TrimSpace(s)
}

func sanitizeBytes(s []byte) []byte {
	return bytes.TrimSpace(s) // No extra allocations, better speed.
}
