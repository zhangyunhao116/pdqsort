package pdqsort

import (
	"testing"

	"github.com/zhangyunhao116/fastrand"
)

func TestChoosePivotFuzz(t *testing.T) {
	randomTestTimes := fastrand.Intn(1000)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := fastrand.Intn(1000)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			v1[j] = fastrand.Intn(randomLenth)
		}
		pivotidx, _ := choosePivot(v1)
		_ = v1[pivotidx]
	}
}
