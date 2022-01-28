package iter

func Sum[Item Number | ~string](iter Iterator[Item]) Item {
	var total Item
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		total += item
	}

	return total
}
