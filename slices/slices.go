package slices

import "github.com/FedericoSchonborn/go-iter"

var _ iter.SizedBilateralIterator[any] = (*Iterator[any, []any])(nil)

type Iterator[T any, S ~[]T] struct {
	slice S
	front int
	back  int
}

func NewIterator[T any, S ~[]T](slice S) *Iterator[T, S] {
	return &Iterator[T, S]{
		slice: slice,
		front: 0,
		back:  len(slice) - 1,
	}
}

func (i *Iterator[T, S]) Next() (_ T, ok bool) {
	if i.front >= len(i.slice) {
		return iter.Zero[T](), false
	}

	next := i.slice[i.front]
	i.front++
	return next, true
}

func (i *Iterator[T, S]) NextBack() (_ T, ok bool) {
	if i.back < 0 {
		return iter.Zero[T](), false
	}

	next := i.slice[i.back]
	i.back--
	return next, true
}

func (i *Iterator[T, S]) Len() int {
	return len(i.slice)
}
