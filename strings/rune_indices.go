package strings

import "github.com/FedericoSchonborn/go-iter"

type RuneIndices struct {
	runes []rune
	len   int
	index int
}

func NewRuneIndices(s string) *RuneIndices {
	r := []rune(s)
	return &RuneIndices{
		runes: r,
		len:   len(r),
		index: 0,
	}
}

func (ri *RuneIndices) Next() (_ struct {
	Index int
	Rune  rune
}, ok bool) {
	if ri.index >= ri.len {
		return iter.Zero[struct {
			Index int
			Rune  rune
		}](), false
	}

	next := struct {
		Index int
		Rune  rune
	}{
		ri.index,
		ri.runes[ri.index],
	}
	ri.index++
	return next, true
}
