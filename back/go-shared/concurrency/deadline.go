package concurrency

import (
	"context"
	"time"
)

func SyncWithDeadline(deadline time.Duration, fn func(context.Context)) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(deadline))
	defer cancel()

	ch := make(chan struct{}, 1)

	go func() {
		fn(ctx)
		ch <- struct{}{}
		close(ch)
	}()

	select {
	case <-ch:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
