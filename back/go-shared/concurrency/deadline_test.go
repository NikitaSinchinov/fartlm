package concurrency

import (
	"context"
	"testing"
	"time"
)

const (
	epsilon = 500 * time.Millisecond
)

func TestRunWithDeadline_Success(t *testing.T) {
	fn := func(context.Context) {
		time.Sleep(epsilon)
	}

	err := SyncWithDeadline(2*epsilon, fn)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
}

func TestRunWithDeadline_Timeout(t *testing.T) {
	fn := func(context.Context) {
		time.Sleep(2 * epsilon)
	}

	err := SyncWithDeadline(1*epsilon, fn)
	if err == nil {
		t.Errorf("expected error, but got %v", err)
	}
}

func TestRunWithDeadline_Context(t *testing.T) {
	signals := make(chan struct{}, 2)

	fn := func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				close(signals)
				return
			default:
				signals <- struct{}{}
				time.Sleep(2 * epsilon)
			}
		}
	}

	err := SyncWithDeadline(epsilon, fn)
	if err == nil {
		t.Errorf("expected error, but got %v", err)
	}

	signalsCount := 0
	for range signals {
		signalsCount++
	}

	if signalsCount != 1 {
		t.Errorf("expected 1 signal, but got %d", signalsCount)
	}
}
