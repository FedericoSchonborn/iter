package iter

type Once[T any] struct {
	value T
	done  bool
}

func NewOnce[T any](value T) *Once[T] {
	return &Once[T]{
		value: value,
		done:  false,
	}
}

func (o *Once[T]) Next() (_ T, ok bool) {
	if o.done {
		return Zero[T](), false
	}

	o.done = true
	return o.value, true
}
