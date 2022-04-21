package iter

var _ Iterator[struct {
	Index int
	Value any
}] = (*Enumerate[any, *Empty[any]])(nil)

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

func (e *Enumerate[T, I]) Next() (item struct {
	Index int
	Value T
}, ok bool) {
	value, ok := e.iter.Next()
	if !ok {
		return item, false
	}

	item.Index = e.index
	item.Value = value
	e.index++
	return item, true
}
