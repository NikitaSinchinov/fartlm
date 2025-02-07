package concurrency

import (
	"bytes"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"bou.ke/monkey"
)

const (
	waitTimeToFinish = 100 * time.Millisecond
)

func TestWorkerPool_Basic(t *testing.T) {
	wp := NewDefaultWorkerPool()

	var counter int32

	// Add tasks that increment the counter
	for i := 0; i < 100; i++ {
		wp.Submit(func() {
			atomic.AddInt32(&counter, 1)
		})
	}

	wp.Stop()
	if !wp.isStopped {
		t.Error("WorkerPool was not completed")
	}

	if counter != 100 {
		t.Errorf("Expected counter to be 100, got %d", counter)
	}
}

func TestWorkerPool_Basic_WithGoroutineLeak(t *testing.T) {
	_TestGoroutineLeak(t, TestWorkerPool_Basic)
}

func TestWorkerPool_Concurrency(t *testing.T) {
	wp := NewDefaultWorkerPool()

	// Track concurrent executions
	var maxConcurrent int32
	var currentConcurrent int32

	// Add tasks that sleep briefly to test concurrency
	for i := 0; i < 32; i++ {
		wp.Submit(func() {
			current := atomic.AddInt32(&currentConcurrent, 1)
			if current > atomic.LoadInt32(&maxConcurrent) {
				atomic.StoreInt32(&maxConcurrent, current)
			}
			time.Sleep(10 * time.Millisecond)
			atomic.AddInt32(&currentConcurrent, -1)
		})
	}

	wp.Stop()
	if !wp.isStopped {
		t.Error("WorkerPool was not completed")
	}

	if int(maxConcurrent) > GlobalConfig.MaxConcurrentGoroutines {
		t.Errorf("Max concurrent executions %d exceeded limit of %d", maxConcurrent, GlobalConfig.MaxConcurrentGoroutines)
	}
}

func TestWorkerPool_OrderIndependence(t *testing.T) {
	wp := NewDefaultWorkerPool()

	var mu sync.Mutex
	results := make([]int32, 100)

	// Add tasks that set values in different orders
	for i := 0; i < 100; i++ {
		index := i
		wp.Submit(func() {
			time.Sleep(time.Duration(100-index) * time.Millisecond)
			mu.Lock()
			defer mu.Unlock()
			results[index] = int32(index)
		})
	}

	wp.Stop()
	if !wp.isStopped {
		t.Error("WorkerPool was not completed")
	}

	// Verify all values were set
	for i := 0; i < 100; i++ {
		if results[i] != int32(i) {
			t.Errorf("Expected results[%d] to be %d, got %d", i, i, results[i])
		}
	}
}

func TestWorkerPool_StressTest(t *testing.T) {
	wp := NewDefaultWorkerPool()

	var successCounter int32
	var errorCounter int32

	// Add many quick tasks
	for i := 0; i < 10000; i++ {
		wp.Submit(func() {
			if time.Now().UnixNano()%2 == 0 {
				atomic.AddInt32(&successCounter, 1)
			} else {
				atomic.AddInt32(&errorCounter, 1)
			}
		})
	}

	wp.Stop()
	if !wp.isStopped {
		t.Error("WorkerPool was not completed")
	}

	total := successCounter + errorCounter
	if total != 10000 {
		t.Errorf("Expected 10000 total executions, got %d", total)
	}
}

func TestWorkerPool_EmptyPool(t *testing.T) {
	wp := NewDefaultWorkerPool()
	wp.Stop() // Should not block or panic
	if !wp.isStopped {
		t.Error("WorkerPool was not completed")
	}
}

func TestWorkerPool_SingleTask(t *testing.T) {
	wp := NewDefaultWorkerPool()

	var executed bool
	wp.Submit(func() {
		executed = true
	})

	wp.Stop()
	if !wp.isStopped {
		t.Error("WorkerPool was not completed")
	}

	if !executed {
		t.Error("Single task was not executed")
	}
}

func TestWorkerPool_LogFatalOnGC(t *testing.T) {
	var wp *WorkerPool

	createPool := func() {
		wp = NewDefaultWorkerPool()
	}

	triggerGC := func() {
		wp = nil
		runtime.GC()
		time.Sleep(waitTimeToFinish)
	}

	// Test case 1: WorkerPool is properly stopped before GC
	t.Run("No log.Fatal when stopped", func(t *testing.T) {
		createPool()
		wp.Stop()

		// Capture log output
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer log.SetOutput(os.Stderr) // Restore default output

		triggerGC()

		// Check if any unexpected log message was produced
		if buf.Len() > 0 {
			t.Errorf("Unexpected log output: %s", buf.String())
		}
	})

	// Test case 2: WorkerPool is not stopped before GC
	t.Run("log.Fatal when not stopped", func(t *testing.T) {
		createPool()

		// Capture log output
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer log.SetOutput(os.Stderr) // Restore default output

		// Use a channel to signal if os.Exit was called
		exitCalled := make(chan bool)
		patch := monkey.Patch(os.Exit, func(int) {
			exitCalled <- true
		})
		defer patch.Unpatch()

		go func() {
			triggerGC()
			exitCalled <- false
		}()

		select {
		case wasExitCalled := <-exitCalled:
			if !wasExitCalled {
				t.Error("Expected os.Exit to be called, but it wasn't")
			}
		case <-time.After(waitTimeToFinish):
			t.Error("Test timed out")
		}

		// Check if the expected message was logged
		if !strings.Contains(buf.String(), "WorkerPool was not stopped before being garbage collected") {
			t.Error("Expected log message not found")
		}
	})
}

func _TestGoroutineLeak(t *testing.T, body func(t *testing.T)) {
	initialCount := runtime.NumGoroutine()

	body(t)

	time.Sleep(waitTimeToFinish)

	finalCount := runtime.NumGoroutine()

	if finalCount > initialCount {
		t.Errorf("Goroutine leak detected: %d goroutines before, %d after",
			initialCount, finalCount)
	}
}

func BenchmarkWorkerPool(b *testing.B) {
	wp := NewDefaultWorkerPool()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wp.Submit(func() {
			time.Sleep(time.Microsecond)
		})
	}
	wp.Stop()
}
