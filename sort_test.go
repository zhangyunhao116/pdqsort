package pdqsort

import "testing"

func TestSorts(t *testing.T) {
	fuzzTestSort(t, func(data []int) {
		insertionSort(data)
	})
	fuzzTestSort(t, func(data []int) {
		heapSort(data)
	})
	fuzzTestSort(t, func(data []int) {
		simpleQS(data)
	})
}
