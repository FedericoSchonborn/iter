package iter

type Adder interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~complex64 | ~complex128 |
		~string
}

func Sum[Item Adder](iter Iterator[Item]) Item {
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
