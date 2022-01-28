package iter

type SliceIter[Item any] struct {
	slice []Item
	index int
}

func New[Item any](items ...Item) Iterator[Item] {
	return From(items)
}

func From[Item any](slice []Item) Iterator[Item] {
	return &SliceIter[Item]{
		slice: slice,
		index: 0,
	}
}

func (si *SliceIter[Item]) Next() (_ Item, ok bool) {
	if si.index >= len(si.slice) {
		var zero Item
		return zero, false
	}

	item := si.slice[si.index]
	si.index++
	return item, true
}
