package iter

type OnceWithIter[Item any] struct {
	fn   func() Item
	done bool
}

func OnceWith[Item any](fn func() Item) Iterator[Item] {
	return &OnceWithIter[Item]{
		fn:   fn,
		done: false,
	}
}

func (owi *OnceWithIter[Item]) Next() (_ Item, ok bool) {
	if owi.done {
		var zero Item
		return zero, false
	}

	owi.done = true
	return owi.fn(), true
}
