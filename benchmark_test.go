package pdqsort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

var sizes = []int{1 << 6, 1 << 8, 1 << 10, 1 << 12, 1 << 16}

type benchTask struct {
	name string
	f    func([]int)
}

var benchTasks = []benchTask{
	{
		name: "pdqsort",
		f: func(i []int) {
			Slice(i)
		},
	},
	{
		name: "stdsort",
		f:    sort.Ints,
	},
}

func benchmarkBase(b *testing.B, dataset func(x []int)) {
	for _, size := range sizes {
		for _, task := range benchTasks {
			b.Run(fmt.Sprintf(task.name+"-%d", size), func(b *testing.B) {
				b.StopTimer()
				for i := 0; i < b.N; i++ {
					data := make([]int, size)
					dataset(data)
					b.StartTimer()
					task.f(data)
					b.StopTimer()
				}
			})
		}
	}
}

func BenchmarkRandom(b *testing.B) {
	benchmarkBase(b, func(x []int) {
		for i := range x {
			x[i] = rand.Int()
		}
	})
}

func BenchmarkSorted(b *testing.B) {
	benchmarkBase(b, func(x []int) {
		for i := range x {
			x[i] = i
		}
	})
}

func BenchmarkSorted90(b *testing.B) {
	benchmarkBase(b, func(x []int) {
		for i := range x {
			if i < len(x)-(len(x)/10) {
				x[i] = i
			} else {
				x[i] = rand.Int()
			}
		}
	})
}

func BenchmarkReversed(b *testing.B) {
	benchmarkBase(b, func(x []int) {
		for i := range x {
			x[i] = len(x) - i
		}
	})
}

func BenchmarkReversed90(b *testing.B) {
	benchmarkBase(b, func(x []int) {
		for i := range x {
			if i < len(x)-(len(x)/10) {
				x[i] = len(x) - i
			} else {
				x[i] = rand.Int()
			}
		}
	})
}

func BenchmarkMod8(b *testing.B) {
	benchmarkBase(b, func(x []int) {
		for i := range x {
			x[i] = i % 8
		}
	})
}
