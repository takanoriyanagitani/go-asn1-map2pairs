package map2pairs

import (
	"encoding/asn1"
	"iter"
	"slices"
)

type RawMap map[string]string

type RawPair struct {
	Key string `asn1:"utf8"`
	Val string `asn1:"utf8"`
}

func (r RawMap) ToPairs() iter.Seq[RawPair] {
	return func(yield func(RawPair) bool) {
		for key, val := range r {
			pair := RawPair{
				Key: key,
				Val: val,
			}
			if !yield(pair) {
				return
			}
		}
	}
}

func (r RawMap) ToDerBytes() ([]byte, error) {
	var pairs []RawPair = slices.Collect(r.ToPairs())
	return asn1.Marshal(pairs)
}

func RawMapNew(m map[string]string) (RawMap, error) { return RawMap(m), nil }
