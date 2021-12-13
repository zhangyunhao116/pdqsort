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
