package iter

type OnceIter[Item any] struct {
	item Item
	done bool
}

func Once[Item any](item Item) Iterator[Item] {
	return &OnceIter[Item]{
		item: item,
		done: false,
	}
}

func (oi *OnceIter[Item]) Next() (_ Item, ok bool) {
	if oi.done {
		var zero Item
		return zero, false
	}

	oi.done = true
	return oi.item, true
}
