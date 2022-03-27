package iter

type Zipped[L, R any] struct {
	Left  L
	Right R
}

type Zip[L, R any, LI Iterator[L], RI Iterator[R]] struct {
	left  LI
	right RI
}

func New[L, R any, LI Iterator[L], RI Iterator[R]](left LI, right RI) *Zip[L, R, LI, RI] {
	return &Zip[L, R, LI, RI]{
		left:  left,
		right: right,
	}
}

func (z *Zip[L, R, LI, RI]) Next() (_ Zipped[L, R], ok bool) {
	lv, ok := z.left.Next()
	if !ok {
		var zero Zipped[L, R]
		return zero, false
	}

	rv, ok := z.right.Next()
	if !ok {
		var zero Zipped[L, R]
		return zero, false
	}

	return Zipped[L, R]{
		Left:  lv,
		Right: rv,
	}, true
}
