package iter

func ForEach[I Iterator[Item], Item any](iter I, fn func(Item)) {
	for {
		value, ok := iter.Next()
		if !ok {
			break
		}

		fn(value)
	}
}
