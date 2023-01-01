package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "1", "1")
	ctx = context.WithValue(ctx, "2", "2")
	ctx = context.WithValue(ctx, "3", "3")
	ctx = context.WithValue(ctx, "4", "4")
	ctx = context.WithValue(ctx, "5", "5")

	fmt.Printf("ctx: %v\n", ctx)

}
