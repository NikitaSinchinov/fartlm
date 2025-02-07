package concurrency

import (
	"go-shared/logger"
	"runtime"
	"sync"
)

type Config struct {
	MaxConcurrentGoroutines int
}

var GlobalConfig Config

func init() {
	GlobalConfig = Config{runtime.NumCPU()}
}

type Task func()

// Worker

type Worker struct {
	tasks <-chan Task
	wg    *sync.WaitGroup
}

func NewWorker(jobs <-chan Task, wg *sync.WaitGroup) *Worker {
	return &Worker{jobs, wg}
}

func (w *Worker) Start() {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		w.dispatch()
	}()
}

func (w *Worker) dispatch() {
	for task := range w.tasks {
		task()
	}
}

// WorkerPool

type WorkerPool struct {
	tasks chan Task

	waitFunc  func()
	isStopped bool
}

func NewWorkerPool(maxWorkers int) *WorkerPool {
	workersCount := min(maxWorkers, GlobalConfig.MaxConcurrentGoroutines)
	tasks := make(chan Task, GlobalConfig.MaxConcurrentGoroutines)

	var wg sync.WaitGroup
	for w := 0; w < workersCount; w++ {
		worker := NewWorker(tasks, &wg)
		worker.Start()
	}

	wp := &WorkerPool{
		tasks:    tasks,
		waitFunc: wg.Wait,
	}

	runtime.SetFinalizer(wp, func(wp *WorkerPool) {
		if !wp.isStopped {
			logger.Fatal("WorkerPool was not stopped before being garbage collected")
		}
	})

	return wp
}

func NewDefaultWorkerPool() *WorkerPool {
	return NewWorkerPool(GlobalConfig.MaxConcurrentGoroutines)
}

func (wp *WorkerPool) Submit(task Task) {
	wp.tasks <- task
}

func (wp *WorkerPool) SyncSubmit(task Task) {
	done := make(chan struct{}, 1)
	wp.tasks <- func() {
		task()
		done <- struct{}{}
		close(done)
	}
	<-done
}

func (wp *WorkerPool) Stop() {
	close(wp.tasks)

	wp.waitFunc()
	wp.isStopped = true
}
