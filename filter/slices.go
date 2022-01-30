package filter

import (
	"github.com/sonirico/container/types"
	"github.com/sonirico/container/utils"
)

// SliceWithInPlaceMutation filters the given array by applying in-place state mutation, which mutates
// the underlying array. Accepts a flag to indicate whether to copy the result array as the result to
// prevent keeping references to the original array which could be way bigger than the filtered result
// Time complexity: O(n)
// Space complexity: O(1) | O(n)
func SliceWithInPlaceMutation[T types.Any](slice []T, shouldCopy bool, fn func(T) bool) []T {
	i := 0
	for _, x := range slice {
		if ok := fn(x); ok {
			slice[i] = x
			i++
		}
	}
	if !shouldCopy {
		return slice[:i]
	}

	return utils.SliceCopy(slice[:i])
}

// SliceWithAppend filters the given array by creating a new one while filling it with `append`. Beware
// of the fact that the longer the filtered result may be, the more empty slots into the array will be
// created.
// Time complexity: O(n)
// Space complexity: O(n)
func SliceWithAppend[T types.Any](slice []T, fn func(T) bool) []T {
	i := 0
	var r []T
	for _, x := range slice {
		if ok := fn(x); ok {
			r = append(r, x)
			i++
		}
	}
	return r
}

// SliceWithAppendCopy behaves as SliceWithAppend but copies the filtered result in order to remove empty
// slots of the result.
// Time complexity: O(n)
// Space complexity: O(n)
func SliceWithAppendCopy[T types.Any](slice []T, fn func(T) bool) []T {
	i := 0
	var r []T
	for _, x := range slice {
		if ok := fn(x); ok {
			r = append(r, x)
			i++
		}
	}
	return utils.SliceCopy(r)
}

// SliceWithDoubleFor filters the given slice without mutating it with as much less memory footprint as possible.
// However, it does so by iterating twice the slice. The first round is to know how many slots to allocate.
// Time complexity: O(2n)
// Space complexity: O(n)
func SliceWithDoubleFor[T types.Any](slice []T, fn func(T) bool) []T {
	i := 0
	for _, x := range slice {
		if ok := fn(x); ok {
			i++
		}
	}
	r := make([]T, i)
	i = 0
	for _, x := range slice {
		if ok := fn(x); ok {
			r[i] = x
			i++
		}
	}
	return r
}

// Slice defaults to filtering the given array by applying in-place mutation as SliceWithInPlaceMutation.
func Slice[T types.Any](data []T, fn func(T) bool) []T {
	return SliceWithInPlaceMutation(data, false, fn)
}
