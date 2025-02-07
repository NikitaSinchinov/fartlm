package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func StartAndListenToSigTerm(body func(), shutdown func()) {
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, os.Interrupt, syscall.SIGTERM)

	done := make(chan struct{})

	go func() {
		body()
		close(done)
	}()

	select {
	case <-sigTerm:
	case <-done:
	}

	shutdown()
}
