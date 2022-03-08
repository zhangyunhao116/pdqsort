package pdqsort

func Slice[T ordered](v []T) {
	if len(v) <= 1 {
		return
	}
	var tmp T // meaningless variable
	limit := usize - bitsLeadingZeros(uint(len(v)))
	recurse(v, tmp, false, limit)
}

// recurse sorts `v` recursively.
//
// If the slice had a predecessor in the original array, it is specified as `pred`(must be the minimum value if exist).
//
// `limit` is the number of allowed imbalanced partitions before switching to `heapsort`. If zero,
// this function will immediately switch to heapsort.
func recurse[T ordered](v []T, pred T, predExist bool, limit int) {
	const MaxInsertion = 24 // slices of up to this length get sorted using insertion sort.

	var (
		// True if the last partitioning was reasonably balanced.
		wasBalanced = true
		// True if the last partitioning didn't shuffle elements (the slice was already partitioned).
		wasPartitioned = true
	)

	for {
		length := len(v)

		// Very short slices get sorted using insertion sort.
		if length <= MaxInsertion {
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
		wasPartitioned = wasP

		left, right := v[:mid], v[mid+1:]
		pivot := v[mid]
		pivotExist := true

		if len(left) > len(right) {
			wasBalanced = len(right) >= len(v)/8
			recurse(left, pred, predExist, limit)
			v = right
			pred = pivot
			predExist = pivotExist
		} else {
			wasBalanced = len(left) >= len(v)/8
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
func partition[T ordered](v []T, pivotidx int) (int, bool) {
	pivot := v[pivotidx]
	v[0], v[pivotidx] = v[pivotidx], v[0]
	i, j := 1, len(v)-1

	for i <= j && v[i] < pivot {
		i++
	}
	for i <= j && v[j] >= pivot {
		j--
	}
	if i > j {
		v[j], v[0] = v[0], v[j]
		return j, true
	}

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
	return j, false
}

type xorshift uint64

func (r *xorshift) Next() uint64 {
	*r ^= *r << 13
	*r ^= *r >> 17
	*r ^= *r << 5
	return uint64(*r)
}

// breakPatterns scatters some elements around in an attempt to break patterns that might cause imbalanced
// partitions in quicksort.
// Warning: this function will panic if `len(v) < 4`,
// it's the caller's responsibility to make sure the `len(v) >= 4`.
func breakPatterns[T ordered](v []T) {
	length := len(v)

	r := xorshift(length)

	modulus := nextPowerOfTwo(length)

	var idxs [3]uint
	idxs[0] = uint(length/4)*2 - 1
	idxs[1] = uint(length/4) * 2
	idxs[2] = uint(length/4)*2 + 1

	for _, idx := range idxs {
		other := int(uint(r.Next()) & (modulus - 1))
		if other >= length {
			other -= length
		}
		v[idx], v[other] = v[other], v[idx]
	}
}

// partitionEqual partitions `v` into elements equal to `v[pivotidx]` followed by elements greater than `v[pivotidx]`.
//
// Returns the number of elements equal to the pivot. It is assumed that `v` does not contain
// elements smaller than the pivot.
func partitionEqual[T ordered](v []T, pivotidx int) int {
	v[0], v[pivotidx] = v[pivotidx], v[0]
	pivot := v[0] // minimum value

	L := 1
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
	return L
}

// partialInsertionSort partially sorts a slice by shifting several out-of-order elements around.
// Returns `true` if the slice is sorted at the end. This function is `O(n)` worst-case.
func partialInsertionSort[T ordered](v []T) bool {
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

func shiftTail[T ordered](v []T, a, b int) {
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

func shiftHead[T ordered](v []T, a, b int) {
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

func nextPowerOfTwo(length int) uint {
	shift := uint(usize - bitsLeadingZeros(uint(length)))
	return uint(1 << shift)
}
