package maps

import "github.com/sonirico/container/types"

// Slice maps a slice of type `T` to another slice of type `R` without mutating the original
func Slice[T types.Any, R types.Any](source []T, predicate func(T) R) []R {
	r := make([]R, len(source))
	for i, x := range source {
		r[i] = predicate(x)
	}
	return r
}

// Map maps the source maps to a new one by receiving a predicate which takes as parameter the key-value pairs from
// the original
func Map[
	K1 types.Comparable, V1 types.Any,
	K2 types.Comparable, V2 types.Any](
	source map[K1]V1,
	predicate func(K1, V1) (K2, V2),
) map[K2]V2 {
	r := make(map[K2]V2, len(source))
	for k1, v1 := range source {
		k2, v2 := predicate(k1, v1)
		r[k2] = v2
	}
	return r
}
