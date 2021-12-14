package iter

type RepeatIter[Item any] struct {
	item Item
}

func Repeat[Item any](item Item) Iterator[Item] {
	return &RepeatIter[Item]{
		item: item,
	}
}

func (ri *RepeatIter[Item]) Next() (_ Item, ok bool) {
	return ri.item, true
}
