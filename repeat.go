package iter

type RepeatIterator[T any] struct {
	item T
}

func Repeat[T any](item T) *RepeatIterator[T] {
	return &RepeatIterator[T]{
		item: item,
	}
}

func (ri *RepeatIterator[T]) Next() (_ T, ok bool) {
	return ri.item, true
}
