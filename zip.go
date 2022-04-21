package iter

type Zipped[L, R any] struct {
	Left  L
	Right R
}

type ZipIterator[L, R any, LI Iterator[L], RI Iterator[R]] struct {
	left  LI
	right RI
}

func Zip[L, R any, LI Iterator[L], RI Iterator[R]](left LI, right RI) *ZipIterator[L, R, LI, RI] {
	return &ZipIterator[L, R, LI, RI]{
		left:  left,
		right: right,
	}
}

func (zi *ZipIterator[L, R, LI, RI]) Next() (_ Zipped[L, R], ok bool) {
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
