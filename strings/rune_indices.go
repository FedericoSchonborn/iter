package strings

type RuneIndices struct {
	inner []rune
	len   int
	index int
}

func NewRuneIndices(s string) *RuneIndices {
	r := []rune(s)
	return &RuneIndices{
		inner: r,
		len:   len(r),
		index: 0,
	}
}

func (ri *RuneIndices) Next() (item struct {
	Index int
	Rune  rune
}, ok bool) {
	if ri.index >= ri.len {
		return item, false
	}

	item.Index = ri.index
	item.Rune = ri.inner[ri.index]
	ri.index++
	return item, true
}
