package utils

func MillisecondsToSeconds(ms int64) int64 {
	return ms / 1000
}

func PairsToMap(pairs []Pair[string, *float64]) map[string]*float64 {
	m := make(map[string]*float64, len(pairs))
	for _, pair := range pairs {
		m[pair.First] = pair.Second
	}
	return m
}
