package iter

type EnumerateItem[Item any] struct {
	Index int
	Value Item
}

type EnumerateIter[Item any] struct {
	iter  Iterator[Item]
	index int
}

func Enumerate[Item any](iter Iterator[Item]) Iterator[EnumerateItem[Item]] {
	return &EnumerateIter[Item]{
		iter:  iter,
		index: 0,
	}
}

func (ei *EnumerateIter[Item]) Next() (_ EnumerateItem[Item], ok bool) {
	item, ok := ei.iter.Next()
	if !ok {
		var zero EnumerateItem[Item]
		return zero, false
	}

	value := EnumerateItem[Item]{
		Index: ei.index,
		Value: item,
	}
	ei.index++
	return value, true
}
