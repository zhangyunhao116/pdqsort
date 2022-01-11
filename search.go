package pdqsort

import (
	"constraints"
)

// Search searches for x in a sorted slice of Ordereds and returns the index
// as specified by Search. The return value is the index to insert x if x is
// not present (it could be len(a)).
// The slice must be sorted in ascending order.
func Search[T constraints.Ordered](v []T, x T) int {
	i, j := 0, len(v)
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if v[h] < x {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}
