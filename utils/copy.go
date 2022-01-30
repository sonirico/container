package utils

import "github.com/sonirico/container/types"

func SliceCopy[T types.Any](source []T) []T {
	r := make([]T, len(source))
	copy(r, source)
	return r
}
