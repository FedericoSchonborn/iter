package strings

import (
	"github.com/FedericoSchonborn/go-iter"
	"github.com/FedericoSchonborn/go-iter/slices"
)

func Bytes(s string) iter.Iterator[byte] {
	return slices.IntoIterator([]byte(s))
}

func Runes(s string) iter.Iterator[rune] {
	return slices.IntoIterator([]rune(s))
}
