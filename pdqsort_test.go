package pdqsort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestBreakPatterns(t *testing.T) {
	// Special case used to trigger breakPattern.
	x := make([]int, 30)
	for i := range x {
		x[i] = 10
	}
	x[(len(x) / 4)] = 0
	x[(len(x)/4)*2] = 1
	x[(len(x)/4)*3] = 2
	Slice(x)
}

func TestPDQSort(t *testing.T) {
	fuzzTestSort(t, func(data []int) {
		Slice(data)
	})
}

func TestPartialInsertionSort(t *testing.T) {
	randomTestTimes := rand.Intn(1000)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := rand.Intn(100)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		v2 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			randomValue := rand.Intn(randomLenth)
			v1[j] = randomValue
			v2[j] = randomValue
		}
		sort.Ints(v1)
		if partialInsertionSort(v2) {
			for idx := range v1 {
				if v1[idx] != v2[idx] {
					t.Fatal("invalid sort:", idx, v1[idx], v2[idx])
				}
			}
		}
	}
}

func TestPartitionEqual(t *testing.T) {
	randomTestTimes := rand.Intn(1000)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := rand.Intn(100)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			randomValue := rand.Intn(randomLenth/2 + 1)
			v1[j] = randomValue
		}
		minvalue := v1[0]
		minvalueidx := 0
		mincount := 0
		for i, v := range v1 {
			if v < minvalue {
				minvalue = v
				minvalueidx = i
			}
		}
		for _, v := range v1 {
			if v == minvalue {
				mincount++
			}
		}
		if mincount != partitionEqual(v1, minvalueidx) {
			t.Fatal()
		}
	}
}

func TestPartition(t *testing.T) {
	fuzzTestPartition(t, func(data []int, pivotidx int) int {
		idx, _ := partition(data, pivotidx)
		return idx
	})
}

func TestBreakPatternsFuzz(t *testing.T) {
	return
	randomTestTimes := rand.Intn(1000)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := rand.Intn(1000)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			v1[j] = rand.Intn(randomLenth)
		}
		breakPatterns(v1)
	}
}

func fuzzTestPartition(t *testing.T, f func(data []int, pivotidx int) int) {
	const times = 2048
	randomTestTimes := rand.Intn(times)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := rand.Intn(times)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			randomValue := rand.Intn(randomLenth)
			v1[j] = randomValue
		}
		pivotidx := rand.Intn(len(v1))
		newpivotidx := f(v1, pivotidx)
		pivot := v1[newpivotidx]
		for i, v := range v1 {
			if i < newpivotidx && v > pivot {
				t.Fatal(i, v, pivotidx, pivot)
			}
			if i > newpivotidx && v < pivot {
				t.Fatal(i, v, pivotidx, pivot)
			}
		}
	}
}

func fuzzTestSort(t *testing.T, f func(data []int)) {
	const times = 2048
	randomTestTimes := rand.Intn(times)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := rand.Intn(times)
		v1 := make([]int, randomLenth)
		v2 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			randomValue := rand.Intn(randomLenth)
			v1[j] = randomValue
			v2[j] = randomValue
		}
		sort.Ints(v1)
		f(v2)
		for idx := range v1 {
			if v1[idx] != v2[idx] {
				t.Fatal("invalid sort:", idx, v1[idx], v2[idx])
			}
		}
	}
}
