package iter

import "constraints"

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}
