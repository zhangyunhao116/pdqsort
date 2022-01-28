package pdqsort

// insertionSort sorts v[begin:end) using insertion sort.
func insertionSort[T ordered](v []T) {
	for cur := 1; cur < len(v); cur++ {
		for j := cur; j > 0 && v[j] < v[j-1]; j-- {
			v[j], v[j-1] = v[j-1], v[j]
		}
	}
}

// siftDown implements the heap property on v[lo:hi].
func siftDown[T ordered](v []T, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && v[child] < v[child+1] {
			child++
		}
		if v[root] >= v[child] {
			return
		}
		v[root], v[child] = v[child], v[root]
		root = child
	}
}

func heapSort[T ordered](v []T) {
	lo := 0
	hi := len(v)

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(v, i, hi)
	}

	// Pop elements into end of v.
	for i := hi - 1; i >= 0; i-- {
		v[0], v[i] = v[i], v[0]
		siftDown(v, lo, i)
	}
}

func simpleQS[T ordered](v []T) {
	if len(v) > 1 {
		p := simplePartition(v, 0)
		simpleQS(v[:p])
		simpleQS(v[p+1:])
	}
}

func simplePartition[T ordered](v []T, pivotidx int) int {
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
	return j
}
