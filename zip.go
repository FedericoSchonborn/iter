package iter

type ZipItem[L, R any] struct {
	Left  L
	Right R
}

type ZipIter[L, R any] struct {
	left  Iterator[L]
	right Iterator[R]
}

func Zip[L, R any](left Iterator[L], right Iterator[R]) Iterator[ZipItem[L, R]] {
	return &ZipIter[L, R]{
		left:  left,
		right: right,
	}
}

func (zi *ZipIter[L, R]) Next() (_ ZipItem[L, R], ok bool) {
	lv, ok := zi.left.Next()
	if !ok {
		var zero ZipItem[L, R]
		return zero, false
	}

	rv, ok := zi.right.Next()
	if !ok {
		var zero ZipItem[L, R]
		return zero, false
	}

	return ZipItem[L, R]{
		Left:  lv,
		Right: rv,
	}, true
}
