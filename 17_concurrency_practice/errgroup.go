package main

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"
)

type (
	Circle struct{} //
	Result struct{}
)

// First implemetation

func handler(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	wg := sync.WaitGroup{}
	wg.Add(len(results))

	for i, circle := range circles {
		i := i // Create new var each iteration.
		circle := circle

		go func() {
			defer wg.Done()

			result, err := foo(ctx, circle)
			if err != nil {
				// ?
			}
			results[i] = result
		}()
	}

	wg.Wait()

	return results, nil
}

// Impl with errgroup
func handlerErrGroup(ctx context.Context, circles []Circle) ([]Result, error) {
	results := make([]Result, len(circles))
	g, ctx := errgroup.WithContext(ctx)

	for i, circle := range circles {
		i := i
		circle := circle
		g.Go(func() error { // Func in g.Go must be context-aware, overwise cancelling the context won't have any effect.
			result, err := foo(ctx, circle)
			if err != nil {
				return nil
			}
			results[i] = result
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}

func foo(ctx context.Context, c Circle) (Result, error) {
	// Calculations
	return Result{}, nil
}

func main() {

}
