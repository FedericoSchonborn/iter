package iter

var _ Iterator[Numerated[any]] = (*NumerateIterator[any])(nil)

type Numerated[T any] struct {
	Index int
	Value T
}

type NumerateIterator[T any] struct {
	iter  Iterator[T]
	index int
}

func Numerate[T any](iter Iterator[T]) *NumerateIterator[T] {
	return &NumerateIterator[T]{
		iter:  iter,
		index: 0,
	}
}

func (ni *NumerateIterator[T]) Next() (_ Numerated[T], ok bool) {
	item, ok := ni.iter.Next()
	if !ok {
		var zero Numerated[T]
		return zero, false
	}

	value := Numerated[T]{
		Index: ni.index,
		Value: item,
	}
	ni.index++
	return value, true
}
