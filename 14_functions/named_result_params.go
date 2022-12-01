package main

type badDesignlocator interface {
	getCoordinates(address string) (float32, float32, error) // What is float32. Perhaps it's latitude and longitude,
	// but in which order?
}

type goodDesignlocator interface {
	getCoordinates(address string) (ltd, lng float32, err error) // We know what it is and it's order
}

// Use it when we need to increase readability. But it can decrease readability when we use it on simple return paraments
