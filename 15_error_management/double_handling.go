package main

import (
	"fmt"
	"log"
)

type Route struct{}

func GetRoute1(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates1(srcLat, srcLng)
	if err != nil {
		log.Println("failed to validate source coordinates") // Error 1
		return Route{}, err                                  // Error 2. >> Two errors in logs. Hence, harder to debug.
	}

	err = validateCoordinates1(dstLat, dstLng)
	if err != nil {
		log.Println("failed to validate target coordinates") // Here the same
		return Route{}, err
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func validateCoordinates1(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		log.Printf("invalid latitude: %f", lat)        // Log - first
		return fmt.Errorf("invalid latitude: %f", lat) // Pass error - second
	}
	if lng > 180.0 || lng < -180.0 {
		log.Printf("invalid longitude: %f", lng)        // Log - first
		return fmt.Errorf("invalid longitude: %f", lng) // Pass error - second
	}
	return nil
}

func GetRoute2(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates2(srcLat, srcLng)
	if err != nil {
		return Route{}, err // Lose context here.
	}

	err = validateCoordinates2(dstLat, dstLng) // And here.
	if err != nil {
		return Route{}, err
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func validateCoordinates2(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		return fmt.Errorf("invalid latitude: %f", lat) // 1 errors in log, but we lost context.
		// Source or destination coords?
	}
	if lng > 180.0 || lng < -180.0 {
		return fmt.Errorf("invalid longitude: %f", lng) // The same.
	}
	return nil
}

func GetRoute3(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates2(srcLat, srcLng)
	if err != nil {
		return Route{},
			fmt.Errorf("failed to validate source coordinates: %w", err) // Wrapping - best solution.
		// Not losing context, pass the core error through.
	}

	err = validateCoordinates2(dstLat, dstLng)
	if err != nil {
		return Route{},
			fmt.Errorf("failed to validate target coordinates: %w", err) // Not losing context.
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func getRoute(lat, lng, lat2, lng2 float32) (Route, error) {
	return Route{}, nil
}

// Handling an error should be done only twice.
// In conclusion: we should either log or return error.
