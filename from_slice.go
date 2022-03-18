package iter

type fromSliceIter[T any] struct {
	slice []T
	index int
}

func FromSlice[T any](slice []T) Iterator[T] {
	return &fromSliceIter[T]{
		slice: slice,
		index: 0,
	}
}

func (si *fromSliceIter[T]) Next() (_ T, ok bool) {
	if si.index >= len(si.slice) {
		var zero T
		return zero, false
	}

	item := si.slice[si.index]
	si.index++
	return item, true
}
