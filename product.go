package iter

func Product[Item Number](iter Iterator[Item]) Item {
	total, ok := iter.Next()
	if !ok {
		var zero Item
		return zero
	}

	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		total *= item
	}

	return total
}
