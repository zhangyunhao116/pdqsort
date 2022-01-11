package pdqsort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSearch(t *testing.T) {
	const times = 2048
	randomTestTimes := rand.Intn(times)
	for i := 0; i < randomTestTimes; i++ {
		randomLenth := rand.Intn(times)
		if randomLenth == 0 {
			continue
		}
		v1 := make([]int, randomLenth)
		for j := 0; j < randomLenth; j++ {
			v1[j] = j
		}
		randomItem := v1[rand.Intn(len(v1))]
		if sort.SearchInts(v1, randomItem) != Search(v1, randomItem) {
			t.Fatal(randomItem)
		}
	}
}
