package concurrency

import (
	"go-shared/utils"
	"sync"
)

// Collect

func GoCollect[T, O any](inputs []T, transform func(T) O) chan O {
	workerPool := NewWorkerPool(len(inputs))

	defer func() {
		go workerPool.Stop()
	}()

	return GoCollectVia(workerPool, inputs, transform)
}

func GoCollectVia[I, O any](workerPool *WorkerPool, inputs []I, transform func(I) O) chan O {
	var wg sync.WaitGroup
	outputsChan := make(chan O, len(inputs))

	for _, input := range inputs {
		wg.Add(1)
		workerPool.Submit(func() {
			defer wg.Done()
			output := transform(input)
			outputsChan <- output
		})
	}

	go func() {
		wg.Wait()
		close(outputsChan)
	}()

	return outputsChan
}

// Map

func GoMap[T, U any](inputs []T, transform func(T) U) []U {
	workerPool := NewWorkerPool(len(inputs))
	defer workerPool.Stop()

	return GoMapVia(workerPool, inputs, transform)
}

func GoMapVia[I, O any](workerPool *WorkerPool, inputs []I, transform func(I) O) []O {
	indexedInputs := make([]utils.IndexedValue[I], len(inputs))
	for i, input := range inputs {
		indexedInputs[i] = utils.NewIndexedValue(i, input)
	}

	indexedOutputsChan := GoCollectVia(workerPool, indexedInputs, func(input utils.IndexedValue[I]) utils.IndexedValue[O] {
		output := transform(input.Value)
		return utils.NewIndexedValue(input.Index, output)
	})

	outputs := make([]O, len(inputs))
	for indexedInput := range indexedOutputsChan {
		outputs[indexedInput.Index] = indexedInput.Value
	}
	return outputs
}

// ForEach

func GoForEach[T any](inputs []T, body func(T)) {
	workerPool := NewWorkerPool(len(inputs))
	defer workerPool.Stop()

	GoForEachVia(workerPool, inputs, body)
}

func GoForEachVia[I any](workerPool *WorkerPool, inputs []I, body func(I)) {
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)
		workerPool.Submit(func() {
			defer wg.Done()
			body(input)
		})
	}

	wg.Wait()
}
