package pdqsort

import "constraints"

// choosePivot chooses a pivot in `v` and returns the index and `true` if the slice is likely already sorted.
//
// Elements in `v` might be reordered in the process.
//
// [0,8): choose a static pivot.
// [8,ShortestNinther): use the simple median-of-three method.
// [ShortestNinther,∞): use the Tukey’s ninther method.
func choosePivot[T constraints.Ordered](v []T) (pivotidx int, likelySorted bool) {
	const (
		// ShortestNinther is the minimum length to choose the Tukey’s ninther method.
		// Shorter slices use the simple median-of-three method.
		ShortestNinther = 50
		// MaxSwaps is the maximum number of swaps that can be performed in this function.
		MaxSwaps = 4 * 3
	)

	l := len(v)

	var (
		// Counts the total number of swaps we are about to perform while sorting indices.
		swaps int
		// Three indices near which we are going to choose a pivot.
		a = l / 4 * 1
		b = l / 4 * 2
		c = l / 4 * 3
	)

	if l >= 8 {
		if l >= ShortestNinther {
			// Find medians in the neighborhoods of `a`, `b`, and `c`.
			sortAdjacent(v, &a, &swaps)
			sortAdjacent(v, &b, &swaps)
			sortAdjacent(v, &c, &swaps)
		}
		// Find the median among `a`, `b`, and `c`.
		sort3(v, &a, &b, &c, &swaps)
	}

	if swaps < MaxSwaps {
		return b, (swaps == 0)
	} else {
		// The maximum number of swaps was performed. Chances are the slice is descending or mostly
		// descending, so reversing will probably help sort it faster.
		reverseRange(v)
		return (l - 1 - b), true
	}
}

// sort2 swaps `a` `b` so that `v[a] <= v[b]`.
func sort2[T constraints.Ordered](v []T, a, b, swaps *int) {
	if v[*b] < v[*a] {
		*a, *b = *b, *a
		*swaps++
	}
}

// sort3 swaps `a` `b` `c` so that `v[a] <= v[b] <= v[c]`.
func sort3[T constraints.Ordered](v []T, a, b, c, swaps *int) {
	sort2(v, a, b, swaps)
	sort2(v, b, c, swaps)
	sort2(v, a, b, swaps)
}

// sortAdjacent finds the median of `v[a - 1], v[a], v[a + 1]` and stores the index into `a`.
func sortAdjacent[T constraints.Ordered](v []T, a, swaps *int) {
	t1 := *a - 1
	t2 := *a + 1
	sort3(v, &t1, a, &t2, swaps)
}

func reverseRange[T constraints.Ordered](v []T) {
	i := 0
	j := len(v) - 1
	for i < j {
		v[i], v[j] = v[j], v[i]
		i++
		j--
	}
}
