package pdqsort

import (
	"constraints"
	"math/bits"
	"strconv"
)

func Slice[T constraints.Ordered](v []T) {
	var tmp T // meaningless variable
	limit := strconv.IntSize - bits.LeadingZeros(uint(len(v)))
	recurse(v, tmp, false, limit)
}

// recurse sorts `v` recursively.
//
// If the slice had a predecessor in the original array, it is specified as `pred`(must be the minimum value if exist).
//
// `limit` is the number of allowed imbalanced partitions before switching to `heapsort`. If zero,
// this function will immediately switch to heapsort.
func recurse[T constraints.Ordered](v []T, pred T, predExist bool, limit int) {
	// Slices of up to this length get sorted using insertion sort.
	const MAX_INSERTION = 24

	var (
		// True if the last partitioning was reasonably balanced.
		wasBalanced = true
		// True if the last partitioning didn't shuffle elements (the slice was already partitioned).
		wasPartitioned = true
	)

	for {
		length := len(v)

		// Very short slices get sorted using insertion sort.
		if length <= MAX_INSERTION {
			insertionSort(v)
			return
		}

		// If too many bad pivot choices were made, simply fall back to heapsort in order to
		// guarantee `O(n log n)` worst-case.
		if limit == 0 {
			heapSort(v)
			return
		}

		// If the last partitioning was imbalanced, try breaking patterns in the slice by shuffling
		// some elements around. Hopefully we'll choose a better pivot this time.
		if !wasBalanced {
			breakPatterns(v)
			limit--
		}

		// Choose a pivot and try guessing whether the slice is already sorted.
		pivotidx, likelySorted := choosePivot(v)

		// If the last partitioning was decently balanced and didn't shuffle elements, and if pivot
		// selection predicts the slice is likely already sorted...
		if wasBalanced && wasPartitioned && likelySorted {
			// Try identifying several out-of-order elements and shifting them to correct
			// positions. If the slice ends up being completely sorted, we're done.
			if partialInsertionSort(v) {
				return
			}
		}

		// If the chosen pivot is equal to the predecessor, then it's the smallest element in the
		// slice. Partition the slice into elements equal to and elements greater than the pivot.
		// This case is usually hit when the slice contains many duplicate elements.
		if predExist && pred == v[pivotidx] {
			mid := partitionEqual(v, pivotidx)
			v = v[mid:]
			continue
		}

		// Partition the slice.
		mid, wasP := partition(v, pivotidx)
		wasBalanced = min(mid, len(v)-mid) >= len(v)/8
		wasPartitioned = wasP

		left, right := v[:mid], v[mid:]
		pivot := right[0]
		pivotExist := true
		right = right[1:]

		if len(left) > len(right) {
			recurse(left, pred, predExist, limit)
			v = right
			pred = pivot
			predExist = pivotExist
		} else {
			recurse(right, pivot, pivotExist, limit)
			v = left
		}
	}
}

// Partitions `v` into elements smaller than `v[pivotidx]`, followed by elements greater than or
// equal to `v[pivotidx]`.
//
// Returns a tuple of:
//
// 1. New pivot index.
// 2. True if `v` was already partitioned.
func partition[T constraints.Ordered](v []T, pivotidx int) (int, bool) {
	pivot := v[pivotidx]
	v[0], v[pivotidx] = v[pivotidx], v[0]
	i, j := 1, len(v)-1
	for {
		for i <= j && v[i] < pivot {
			i++
		}
		for i <= j && v[j] >= pivot {
			j--
		}
		if i > j {
			break
		}
		v[i], v[j] = v[j], v[i]
		i++
		j--
	}
	v[j], v[0] = v[0], v[j]
	return j, j == pivotidx
}

// breakPatterns scatters some elements around in an attempt to break patterns that might cause imbalanced
// partitions in quicksort.
func breakPatterns[T constraints.Ordered](v []T) {
	length := len(v)
	if length >= 8 {
		// Xorshift paper: https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
		random := uint(length)
		random ^= random << 13
		random ^= random >> 17
		random ^= random << 5
		modulus := nextPowerOfTwo(length)
		pos := length / 8

		for i := 0; i < 3; i++ {
			other := int(random & (modulus - 1))
			if other >= length {
				other -= length
			}
			v[pos-1+i], v[other] = v[other], v[pos-1+i]
		}
	}
}

// partitionEqual partitions `v` into elements equal to `v[pivotidx]` followed by elements greater than `v[pivotidx]`.
//
// Returns the number of elements equal to the pivot. It is assumed that `v` does not contain
// elements smaller than the pivot.
func partitionEqual[T constraints.Ordered](v []T, pivotidx int) int {
	v[0], v[pivotidx] = v[pivotidx], v[0]
	pivot := v[0] // minimum value
	v = v[1:]

	L := 0
	R := len(v)
	for {
		for L < R && pivot >= v[L] {
			L++
		}
		for L < R && pivot < v[R-1] {
			R--
		}
		if L >= R {
			break
		}
		R--
		v[L], v[R] = v[R], v[L]
		L++
	}
	return L + 1
}

// partialInsertionSort partially sorts a slice by shifting several out-of-order elements around.
// Returns `true` if the slice is sorted at the end. This function is `O(n)` worst-case.
func partialInsertionSort[T constraints.Ordered](v []T) bool {
	const (
		MaxSteps         = 5  // maximum number of adjacent out-of-order pairs that will get shifted
		ShortestShifting = 50 // if the slice is shorter than this, don't shift any elements
	)
	length := len(v)
	i := 1
	for k := 0; k < MaxSteps; k++ {
		// Find the next pair of adjacent out-of-order elements.
		for i < length && v[i] >= v[i-1] {
			i++
		}

		// Are we done?
		if i == length {
			return true
		}

		// Don't shift elements on short arrays, that has a performance cost.
		if length < ShortestShifting {
			return false
		}

		// Swap the found pair of elements. This puts them in correct order.
		v[i-1], v[i] = v[i], v[i-1]

		// Shift the smaller element to the left.
		shiftTail(v, 0, i)
		// Shift the greater element to the right.
		shiftHead(v, i, len(v))
	}

	return false
}

func shiftTail[T constraints.Ordered](v []T, a, b int) {
	l := b - a
	if l >= 2 {
		for i := l - 1; i >= 1; i-- {
			if v[i] >= v[i-1] {
				break
			}
			v[i], v[i-1] = v[i-1], v[i]
		}
	}
}

func shiftHead[T constraints.Ordered](v []T, a, b int) {
	l := b - a
	if l >= 2 {
		for i := 1; i < l; i++ {
			if v[i] >= v[i-1] {
				break
			}
			v[i], v[i-1] = v[i-1], v[i]
		}
	}
}
