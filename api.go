package pdqsort

func SliceIsSorted[T ordered](v []T) bool {
	for i := len(v) - 1; i > 0; i-- {
		if v[i] < v[i-1] {
			return false
		}
	}
	return true
}
