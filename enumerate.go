package iter

var _ Iterator[Enumerated[any]] = (*Enumerate[any, *Empty[any]])(nil)

type Enumerated[T any] struct {
	Index int
	Value T
}

type Enumerate[T any, I Iterator[T]] struct {
	iter  I
	index int
}

func NewEnumerate[T any, I Iterator[T]](iter I) *Enumerate[T, I] {
	return &Enumerate[T, I]{
		iter:  iter,
		index: 0,
	}
}

func (e *Enumerate[T, I]) Next() (_ Enumerated[T], ok bool) {
	item, ok := e.iter.Next()
	if !ok {
		var zero Enumerated[T]
		return zero, false
	}

	value := Enumerated[T]{
		Index: e.index,
		Value: item,
	}
	e.index++
	return value, true
}
