package iter

type TakeIter[Item any] struct {
	iter  Iterator[Item]
	count int
}

func Take[Item any](iter Iterator[Item], n int) Iterator[Item] {
	return &TakeIter[Item]{
		iter:  iter,
		count: n,
	}
}

func (ti *TakeIter[Item]) Next() (_ Item, ok bool) {
	if ti.count <= 0 {
		var zero Item
		return zero, false
	}

	item, ok := ti.iter.Next()
	if !ok {
		var zero Item
		return zero, false
	}

	ti.count--
	return item, true
}
