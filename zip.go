package iter

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

func (z *Zip[L, R, LI, RI]) Next() (item struct {
	Left  L
	Right R
}, ok bool) {
	left, ok := z.left.Next()
	if !ok {
		return item, false
	}

	right, ok := z.right.Next()
	if !ok {
		return item, false
	}

	item.Left = left
	item.Right = right
	return item, true
}
