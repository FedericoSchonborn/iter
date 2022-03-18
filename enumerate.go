package iter

type Enumerated[T any] struct {
	Index int
	Value T
}

type enumerateIter[T any] struct {
	iter  Iterator[T]
	index int
}

func Enumerate[T any](iter Iterator[T]) Iterator[Enumerated[T]] {
	return &enumerateIter[T]{
		iter:  iter,
		index: 0,
	}
}

func (ei *enumerateIter[T]) Next() (_ Enumerated[T], ok bool) {
	item, ok := ei.iter.Next()
	if !ok {
		var zero Enumerated[T]
		return zero, false
	}

	value := Enumerated[T]{
		Index: ei.index,
		Value: item,
	}
	ei.index++
	return value, true
}
