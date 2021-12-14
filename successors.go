package iter

type SuccessorsIter[Item any] struct {
	first   bool
	current Item
	fn      func(Item) (Item, bool)
}

func Successors[Item any](first Item, fn func(Item) (Item, bool)) Iterator[Item] {
	return &SuccessorsIter[Item]{
		first:   true,
		current: first,
		fn:      fn,
	}
}

func (si *SuccessorsIter[Item]) Next() (_ Item, ok bool) {
	if si.first {
		si.first = false
		return si.current, true
	}

	si.current, ok = si.fn(si.current)
	if !ok {
		var zero Item
		return zero, false
	}

	return si.current, true
}
