package iter

func ForEach[I Iterator[T], T any](iter I, fn func(T)) {
	for {
		value, ok := iter.Next()
		if !ok {
			break
		}

		fn(value)
	}
}
