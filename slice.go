package iter

type SliceIter[S ~[]T, T any] struct {
	slice S
	index int
}

func FromSlice[S ~[]T, T any](slice S) Iterator[T] {
	return &SliceIter[S, T]{
		slice: slice,
		index: 0,
	}
}

func (si *SliceIter[S, T]) Next() (_ T, ok bool) {
	if si.index >= len(si.slice) {
		var zero T
		return zero, false
	}

	value := si.slice[si.index]
	si.index++
	return value, true
}
