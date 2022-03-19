package iter

var _ Iterator[Numerated[any]] = (*NumerateIterator[any, *EmptyIterator[any]])(nil)

type Numerated[T any] struct {
	Index int
	Value T
}

type NumerateIterator[T any, I Iterator[T]] struct {
	iter  I
	index int
}

func Numerate[T any, I Iterator[T]](iter I) *NumerateIterator[T, I] {
	return &NumerateIterator[T, I]{
		iter:  iter,
		index: 0,
	}
}

func (ni *NumerateIterator[T, I]) Next() (_ Numerated[T], ok bool) {
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
