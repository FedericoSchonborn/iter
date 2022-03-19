package iter

var _ SizedBilateralIterator[any] = (*fromSliceIter[any])(nil)

type fromSliceIter[T any] struct {
	inner []T
	front int
	back  int
}

func FromSlice[T any, S ~[]T](slice S) Iterator[T] {
	return &fromSliceIter[T]{
		inner: slice,
		front: 0,
	}
}

func (fsi *fromSliceIter[T]) Next() (_ T, ok bool) {
	if fsi.front >= len(fsi.inner) {
		var zero T
		return zero, false
	}

	item := fsi.inner[fsi.front]
	fsi.front++
	return item, true
}

func (fsi *fromSliceIter[T]) NextBack() (_ T, ok bool) {
	if fsi.back < 0 {
		var zero T
		return zero, false
	}

	item := fsi.inner[fsi.back]
	fsi.back--
	return item, true
}

func (fsi *fromSliceIter[T]) Len() int {
	return len(fsi.inner)
}
