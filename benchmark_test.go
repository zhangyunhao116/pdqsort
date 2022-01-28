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
		f:    Slice[int],
	},
	{
		name: "stdsort",
		f:    sort.Ints,
	},
}

func benchmarkBase(b *testing.B, dataset func(x []int)) {
	for _, size := range sizes {
		for _, task := range benchTasks {
			b.Run(fmt.Sprintf(task.name+"_%d", size), func(b *testing.B) {
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

func BenchmarkNearlySorted(b *testing.B) {
	benchmarkBase(b, func(x []int) {
		for i := range x {
			x[i] = i
		}
		for i := 0; i < len(x)/20; i++ {
			a, b := rand.Intn(len(x)), rand.Intn(len(x))
			x[a], x[b] = x[b], x[a]
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

func BenchmarkNearlyReversed(b *testing.B) {
	benchmarkBase(b, func(x []int) {
		for i := range x {
			x[i] = len(x) - i
		}
		for i := 0; i < len(x)/20; i++ {
			a, b := rand.Intn(len(x)), rand.Intn(len(x))
			x[a], x[b] = x[b], x[a]
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

func BenchmarkAllEqual(b *testing.B) {
	benchmarkBase(b, func(x []int) {
		for i := range x {
			x[i] = 1
		}
	})
}
