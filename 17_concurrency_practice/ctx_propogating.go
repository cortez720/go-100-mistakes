package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

type detachCtx struct {
	ctx context.Context
}

func (detachCtx) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

func (detachCtx) Done() <-chan struct{} {
	return nil
}

func (detachCtx) Err() error {
	return nil
}

func (d detachCtx) Value(a any) any {
	return d.ctx.Value(a)
}

func hanlder(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(r.Context(), time.Millisecond*100)
	ctx = context.WithValue(ctx, "ID", uuid.Must(uuid.NewV7()))

	response, err := doSomeTask(ctx, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		err := publishKafka(ctx, response) // Pass the original context.
		if err != nil {
			fmt.Printf("Warning, error: %v\n", err.Error())
		}
	}()

	w.Write([]byte("OK"))
}

func hanlderV2(w http.ResponseWriter, r *http.Request) {
	ctx, _ := context.WithTimeout(r.Context(), time.Millisecond*100)
	ctx = context.WithValue(ctx, "ID", uuid.Must(uuid.NewV7()))

	response, err := doSomeTask(ctx, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		// Pass the detached context. We cant get error from it.
		err := publishKafka(detachCtx{ctx: ctx}, response) // Send message to kafka anyway. Even context deadline exceeded.
		if err != nil {
			fmt.Println("Warning, error: %w\n", err)
		}
	}()

	w.Write([]byte("OK"))
}

func publishKafka(ctx context.Context, m string) error { // Context-aware function.
	ch := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 110) // Too long execution.

		fmt.Printf("Published to Kafka: %v", <-ch)
	}()

	select { // But we can use Value from context here. That's why we can't pass the Background() context.
	case ch <- fmt.Sprintf("ID: %v, message: %v", ctx.Value("ID"), m):
		return nil
	case <-ctx.Done(): // We cannot send the message if the context deadline exceeded
		return ctx.Err() // But we must send the message.
	}
}

func doSomeTask(context.Context, *http.Request) (string, error) {
	return "very important message", nil
}

func main() {
	w := response{}
	r := http.Request{}

	hanlder(w, &r)
	hanlderV2(w, &r)
	time.Sleep(time.Millisecond * 200)
}

type response struct {
}

func (response) Header() http.Header {
	return nil
}

func (response) Write([]byte) (int, error) {
	return 0, nil
}
func (response) WriteHeader(statusCode int) {}
