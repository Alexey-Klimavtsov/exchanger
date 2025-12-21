package util

func Filter[T any](xs []T, keep func(T) bool) []T {
	res := make([]T, 0)
	for _, x := range xs {
		if keep(x) {
			res = append(res, x)
		}
	}
	return res
}
