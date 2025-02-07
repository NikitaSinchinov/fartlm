package utils

import (
	"reflect"
	"testing"
)

func TestMapSlice(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{2, 4, 6, 8}
	mapFunc := func(x int) int { return x * 2 }

	result := Map(input, mapFunc)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestForEach(t *testing.T) {
	input := []int{1, 2, 3, 4}
	sum := 0
	forEachFunc := func(x int) {
		sum += x
	}

	ForEach(input, forEachFunc)

	expectedSum := 10
	if sum != expectedSum {
		t.Errorf("expected sum %v, got %v", expectedSum, sum)
	}
}
