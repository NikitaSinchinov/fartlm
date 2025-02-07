package concurrency

import (
	"reflect"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestGoCollect(t *testing.T) {
	inputs := []int{1, 2, 3, 4, 5}
	transform := func(i int) int { return i * 2 }

	outputChan := GoCollect(inputs, transform)

	var results []int
	for result := range outputChan {
		results = append(results, result)
	}

	expected := []int{2, 4, 6, 8, 10}
	sort.Ints(results)

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %v, but got %v", expected, results)
	}
}

func TestGoCollect_WithGoroutineLeak(t *testing.T) {
	_TestGoroutineLeak(t, TestGoCollect)
}

func TestGoCollectVia(t *testing.T) {
	inputs := []string{"a", "b", "c"}
	transform := func(s string) string { return s + s }

	workerPool := NewDefaultWorkerPool()
	defer workerPool.Stop()

	outputChan := GoCollectVia(workerPool, inputs, transform)

	var results []string
	for result := range outputChan {
		results = append(results, result)
	}

	expected := []string{"aa", "bb", "cc"}
	sort.Strings(results)

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %v, but got %v", expected, results)
	}
}

func TestGoMap(t *testing.T) {
	inputs := []float64{1.1, 2.2, 3.3}
	transform := func(f float64) int { return int(f * 2) }

	results := GoMap(inputs, transform)

	expected := []int{2, 4, 6}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %v, but got %v", expected, results)
	}
}

func TestGoMap_WithGoroutineLeak(t *testing.T) {
	_TestGoroutineLeak(t, TestGoMap)
}

func TestGoMapVia(t *testing.T) {
	inputs := []int{1, 2, 3}
	transform := func(i int) string { return string(rune(i + 64)) }

	workerPool := NewDefaultWorkerPool()
	defer workerPool.Stop()

	results := GoMapVia(workerPool, inputs, transform)

	expected := []string{"A", "B", "C"}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %v, but got %v", expected, results)
	}
}

func TestGoForEach(t *testing.T) {
	inputs := []int{1, 2, 3}

	var mu sync.Mutex
	sum := 0

	body := func(i int) {
		mu.Lock()
		defer mu.Unlock()
		sum += i
	}

	GoForEach(inputs, body)

	expected := 6
	if sum != expected {
		t.Errorf("Expected sum to be %d, but got %d", expected, sum)
	}
}

func TestGoForEach_WithGoroutineLeak(t *testing.T) {
	_TestGoroutineLeak(t, TestGoForEach)
}

func TestGoForEachVia(t *testing.T) {
	inputs := []string{"a", "b", "c"}

	var mu sync.Mutex
	result := ""

	body := func(s string) {
		mu.Lock()
		defer mu.Unlock()
		result += s
	}

	workerPool := NewDefaultWorkerPool()
	defer workerPool.Stop()

	GoForEachVia(workerPool, inputs, body)

	if len(result) != 3 {
		t.Errorf("Expected result length to be 3, but got %d", len(result))
	}

	for _, char := range []string{"a", "b", "c"} {
		if !strings.Contains(result, char) {
			t.Errorf("Expected result to contain %s, but it doesn't", char)
		}
	}
}

func TestActualConcurrency(t *testing.T) {
	inputs := make([]int, 1000)
	for i := range inputs {
		inputs[i] = i
	}

	transform := func(i int) int {
		time.Sleep(10 * time.Millisecond)
		return i * 2
	}

	start := time.Now()
	results := GoMap(inputs, transform)
	duration := time.Since(start)

	if duration > 2*time.Second {
		t.Errorf("Expected concurrent execution to take less than 2 seconds, but it took %v", duration)
	}

	for i, result := range results {
		if result != i*2 {
			t.Errorf("Expected result at index %d to be %d, but got %d", i, i*2, result)
		}
	}
}
