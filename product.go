package iter

import "golang.org/x/exp/constraints"

func Product[T constraints.Integer | constraints.Float | constraints.Complex](iter Iterator[T]) T {
	total, ok := iter.Next()
	if !ok {
		var zero T
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
