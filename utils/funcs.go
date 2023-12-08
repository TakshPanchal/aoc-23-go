package utils

func MapSlice[T, U any](data []T, f func(T) U) []U {
	mapped := make([]U, len(data))

	for i := range data {
		mapped[i] = f(data[i])
	}

	return mapped
}
