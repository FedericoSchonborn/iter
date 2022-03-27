package iter

type Once[T any] struct {
	item T
	done bool
}

func NewOnce[T any](item T) *Once[T] {
	return &Once[T]{
		item: item,
		done: false,
	}
}

func (o *Once[T]) Next() (_ T, ok bool) {
	if o.done {
		var zero T
		return zero, false
	}

	o.done = true
	return o.item, true
}
