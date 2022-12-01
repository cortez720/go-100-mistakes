package main

import (
	"context"
	"errors"
)

type loc struct{}

func (l loc) getCoordinates1(ctx context.Context, address string) (
	lat, lng float32, err error) {
	isValid := l.validateAddress(address)
	if !isValid {
		return 0, 0, errors.New("invalid address")
	}

	if ctx.Err() != nil {
		return 0, 0, err // We could just make mistake and return nil error as a nil param.
		// But without named param we get compilation error
	}

	// Get and return coordinates
	return 0, 0, nil
}
