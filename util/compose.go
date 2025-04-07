package util

import (
	mi "github.com/takanoriyanagitani/go-asn1-map2pairs"
)

func ComposeErr[T, U, V any](
	f func(T) (U, error),
	g func(U) (V, error),
) func(T) (V, error) {
	return mi.ComposeErr(f, g)
}
