package cli

import (
	"context"
	"testing"
	"time"
)

// create the timeout goroutine and return the cancelFunc
func createTimeout(t *testing.T, duration time.Duration, name string) (cancel func()) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	go checkTimeout(t, ctx, name)
	return cancel
}

// repeatedly call the done() method for the context, fail if the deadline exceeds
func checkTimeout(t *testing.T, ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				t.Fatal("Deadline exceeded:", name)
			}
			return
		}
	}
}
