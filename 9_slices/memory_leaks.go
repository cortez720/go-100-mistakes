package main

func consumeMessages() {
	for {
		msg := reciveMessage() // recieves 1MB message, 1_000_000 elements. Very often (1 per 100ms).
		// Do smth with msg
		storeMessageFirstBytes(getMessageFirstBytes(msg))
	}
}

func storeMessageFirstBytes(firstBytes []byte) {
	// Here we store firstBytes, slice with Len:5 and Cap:1_000_000
	// After 100 seconds we will have 1GB memory leak.
}

func reciveMessage() []byte {
	return nil
}

func getMessageFirstBytes(b []byte) []byte { // Bad getter
	return b[:5]
}

func getMessageFirstBytesV1(b []byte) []byte { // Also bad getter
	return b[:5:5] // GC just don't reclaim other elements
}

func getMessageFirstBytesV2(b []byte) []byte { // Best getter
	result := make([]byte, 5)
	copy(result, b)
	return result
}

// Rule of thumb: remember that slicing a large slice or array can lead to potential high-memory consumptoin.
// GC can't reclaim this memory
// Use a slice copy to prevent such a case

// Array under slice stay the same
