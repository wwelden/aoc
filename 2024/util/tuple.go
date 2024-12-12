package util

type Tuple[T any] struct {
	a [2]T
}

func MkTuple[T any](a, b T) Tuple[T] {
	return Tuple[T]{[2]T{a, b}}
}
func Fst[T any](t Tuple[T]) T {
	return t.a[0]
}
func Snd[T any](t Tuple[T]) T {
	return t.a[1]
}
