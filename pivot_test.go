package pdqsort

import (
	"math/rand"
	"testing"
)

func TestChoosePivotFuzz(t *testing.T) {
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
		pivotidx, _ := choosePivot(v1)
		_ = v1[pivotidx]
	}
}

func TestReverseRange(t *testing.T) {
	data := []int{1, 2, 3, 4, 4, 5, 7, 8}
	reverseRange(data)
	for i := len(data) - 1; i > 0; i-- {
		if data[i] > data[i-1] {
			t.Fatalf("reverRange didn't work")
		}
	}

	data1 := []int{1, 2, 3, 4, 5, 7, 8}
	data2 := []int{1, 2, 5, 4, 3, 7, 8}
	reverseRange(data1[2:5])
	for i, v := range data1 {
		if v != data2[i] {
			t.Fatalf("reverRange didn't work")
		}
	}
}
