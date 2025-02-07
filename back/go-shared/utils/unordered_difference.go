package utils

func UnorderedDifference[T comparable](a, b []T) []T {
	// Convert arrays to sets (maps in Go)
	set1 := make(map[T]bool)
	set2 := make(map[T]bool)

	for _, v := range a {
		set1[v] = true
	}
	for _, v := range b {
		set2[v] = true
	}

	difference := []T{}
	for v := range set1 {
		if !set2[v] {
			difference = append(difference, v)
		}
	}

	return difference
}
