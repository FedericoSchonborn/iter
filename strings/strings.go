package strings

import (
	"github.com/FedericoSchonborn/iter"
	"github.com/FedericoSchonborn/iter/slices"
)

func NewBytes(s string) iter.Iterator[byte] {
	return slices.NewIterator([]byte(s))
}

func NewRunes(s string) iter.Iterator[rune] {
	return slices.NewIterator([]rune(s))
}
