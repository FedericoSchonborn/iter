package iter

type EnumerateIter[Item any] struct {
	iter  Iterator[Item]
	index int
}

func Enumerate[Item any](iter Iterator[Item]) Iterator[struct {
	Index int
	Item  Item
}] {
	return &EnumerateIter[Item]{
		iter:  iter,
		index: 0,
	}
}

func (ei *EnumerateIter[Item]) Next() (_ struct {
	Index int
	Item  Item
}, ok bool) {
	item, ok := ei.iter.Next()
	if !ok {
		var zero struct {
			Index int
			Item  Item
		}
		return zero, false
	}

	value := struct {
		Index int
		Item  Item
	}{
		Index: ei.index,
		Item:  item,
	}
	ei.index++
	return value, true
}
