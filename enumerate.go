package iter

var _ Iterator[Enumerated[any]] = (*EnumerateIterator[any, *EmptyIterator[any]])(nil)

type Enumerated[T any] struct {
	Index int
	Value T
}

type EnumerateIterator[T any, I Iterator[T]] struct {
	iter  I
	index int
}

func Enumerate[T any, I Iterator[T]](iter I) *EnumerateIterator[T, I] {
	return &EnumerateIterator[T, I]{
		iter:  iter,
		index: 0,
	}
}

func (ei *EnumerateIterator[T, I]) Next() (_ Enumerated[T], ok bool) {
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
