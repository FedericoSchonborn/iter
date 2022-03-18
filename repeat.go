package iter

type repeatIter[T any] struct {
	item T
}

func Repeat[T any](item T) Iterator[T] {
	return &repeatIter[T]{
		item: item,
	}
}

func (ri *repeatIter[T]) Next() (_ T, ok bool) {
	return ri.item, true
}
