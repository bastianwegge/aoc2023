package util

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Filter[T any](ts []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, t := range ts {
		if fn(t) {
			result = append(result, t)
		}
	}
	return result
}
