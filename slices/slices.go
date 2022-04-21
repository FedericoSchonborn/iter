package slices

import "github.com/FedericoSchonborn/go-iter"

var _ iter.SizedBilateralIterator[any] = (*Iterator[any, []any])(nil)

type Iterator[T any, S ~[]T] struct {
	inner S
	front int
	back  int
}

func IntoIterator[T any, S ~[]T](slice S) *Iterator[T, S] {
	return &Iterator[T, S]{
		inner: slice,
		front: 0,
		back:  len(slice) - 1,
	}
}

func (i *Iterator[T, S]) Next() (_ T, ok bool) {
	if i.front >= len(i.inner) {
		var zero T
		return zero, false
	}

	item := i.inner[i.front]
	i.front++
	return item, true
}

func (i *Iterator[T, S]) NextBack() (_ T, ok bool) {
	if i.back < 0 {
		var zero T
		return zero, false
	}

	item := i.inner[i.back]
	i.back--
	return item, true
}

func (i *Iterator[T, S]) Len() int {
	return len(i.inner)
}
