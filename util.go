package pdqsort

import (
	"constraints"
	"fmt"
	"math/bits"
	"strconv"
)

func debugPrintln(args ...interface{}) {
	return
	fmt.Println(args...)
}

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
