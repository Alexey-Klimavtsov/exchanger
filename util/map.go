package util

func Map[T any, R any](xs []T, f func(T) R) []R {
	res := make([]R, 0, len(xs))
	for _, x := range xs {
		res = append(res, f(x))
	}
	return res
}
