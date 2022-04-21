package iter

type SuccessorsIterator[T any] struct {
	first   bool
	current T
	fn      func(T) (T, bool)
}

func Successors[T any](first T, fn func(T) (T, bool)) *SuccessorsIterator[T] {
	return &SuccessorsIterator[T]{
		first:   true,
		current: first,
		fn:      fn,
	}
}

func (si *SuccessorsIterator[T]) Next() (_ T, ok bool) {
	if si.first {
		si.first = false
		return si.current, true
	}

	si.current, ok = si.fn(si.current)
	if !ok {
		var zero T
		return zero, false
	}

	return si.current, true
}
