package iter

type SliceIter[S ~[]Item, Item any] struct {
	slice S
	index int
}

func FromSlice[S ~[]Item, Item any](slice S) Iterator[Item] {
	return &SliceIter[S, Item]{
		slice: slice,
		index: 0,
	}
}

func (si *SliceIter[S, Item]) Next() (_ Item, ok bool) {
	if si.index >= len(si.slice) {
		var zero Item
		return zero, false
	}

	value := si.slice[si.index]
	si.index++
	return value, true
}
