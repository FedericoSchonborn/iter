package iter

type Zipped[L, R any] struct {
	Left  L
	Right R
}

type zipIter[L, R any] struct {
	left  Iterator[L]
	right Iterator[R]
}

func Zip[L, R any](left Iterator[L], right Iterator[R]) Iterator[Zipped[L, R]] {
	return &zipIter[L, R]{
		left:  left,
		right: right,
	}
}

func (zi *zipIter[L, R]) Next() (_ Zipped[L, R], ok bool) {
	lv, ok := zi.left.Next()
	if !ok {
		var zero Zipped[L, R]
		return zero, false
	}

	rv, ok := zi.right.Next()
	if !ok {
		var zero Zipped[L, R]
		return zero, false
	}

	return Zipped[L, R]{
		Left:  lv,
		Right: rv,
	}, true
}
