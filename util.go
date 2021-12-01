package pdqsort

import (
	"constraints"
	"math/bits"
	"strconv"
)

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func nextPowerOfTwo(length int) uint {
	shift := uint(strconv.IntSize - bits.LeadingZeros(uint(length)))
	return uint(1 << shift)
}
