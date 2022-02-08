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
func siftDown[T ordered](v []T, node int) {
	for {
		child := 2*node + 1
		if child >= len(v) {
			break
		}
		if child+1 < len(v) && v[child] < v[child+1] {
			child++
		}
		if v[node] >= v[child] {
			return
		}
		v[node], v[child] = v[child], v[node]
		node = child
	}
}

func heapSort[T ordered](v []T) {
	// Build heap with greatest element at top.
	for i := (len(v) - 1) / 2; i >= 0; i-- {
		siftDown(v, i)
	}

	// Pop elements into end of v.
	for i := len(v) - 1; i >= 1; i-- {
		v[0], v[i] = v[i], v[0]
		siftDown(v[:i], 0)
	}
}
