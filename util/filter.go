package util

func Filter[T any](xs []T, keep func(T) bool) []T {
	res := make([]T, 0, len(xs))
	for _, x := range xs {
		if keep(x) {
			res = append(res, x)
		}
	}
	return res
}
