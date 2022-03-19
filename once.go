package iter

type OnceIterator[T any] struct {
	item T
	done bool
}

func Once[T any](item T) *OnceIterator[T] {
	return &OnceIterator[T]{
		item: item,
		done: false,
	}
}

func (oi *OnceIterator[T]) Next() (_ T, ok bool) {
	if oi.done {
		var zero T
		return zero, false
	}

	oi.done = true
	return oi.item, true
}
