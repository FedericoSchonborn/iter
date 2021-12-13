package iter

func Fold[I Iterator[Item], Item, B any](iter I, init B, fn func(B, Item) B) B {
	acc := init
	for {
		value, ok := iter.Next()
		if !ok {
			break
		}

		acc = fn(acc, value)
	}

	return acc
}
