package utils

// Converts the element of the provided `slice` using the provided `mapFunc`.
func Map[T any, R any](slice []T, mapFunc func(T) R) []R {
	result := make([]R, len(slice))
	for i, value := range slice {
		result[i] = mapFunc(value)
	}
	return result
}
