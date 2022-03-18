package iter

type onceIter[T any] struct {
	item T
	done bool
}

func Once[T any](item T) Iterator[T] {
	return &onceIter[T]{
		item: item,
		done: false,
	}
}

func (oi *onceIter[T]) Next() (_ T, ok bool) {
	if oi.done {
		var zero T
		return zero, false
	}

	oi.done = true
	return oi.item, true
}
