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

func (z *Zip[L, R, LI, RI]) Next() (_ struct {
	Left  L
	Right R
}, ok bool) {
	left, lok := z.left.Next()
	right, rok := z.right.Next()
	if !lok || !rok {
		return Zero[struct {
			Left  L
			Right R
		}](), false
	}

	next := struct {
		Left  L
		Right R
	}{
		left,
		right,
	}
	return next, true
}
