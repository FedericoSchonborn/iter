package iter

import "golang.org/x/exp/constraints"

func Sum[T constraints.Integer | constraints.Float | constraints.Complex | ~string](iter Iterator[T]) T {
	var total T
	for {
		item, ok := iter.Next()
		if !ok {
			break
		}

		total += item
	}

	return total
}
