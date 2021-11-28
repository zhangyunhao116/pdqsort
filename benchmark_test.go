package pdqsort

import (
	"fmt"
	"sort"
	"testing"

	"github.com/zhangyunhao116/fastrand"
)

func BenchmarkRandom(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("pdqsort-%d", size), func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, size)
				for i := 0; i < len(data); i++ {
					data[i] = fastrand.Int()
				}
				b.StartTimer()
				Slice(data)
				b.StopTimer()
			}
		})
		b.Run(fmt.Sprintf("stdsort-%d", size), func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, size)
				for i := 0; i < len(data); i++ {
					data[i] = fastrand.Int()
				}
				b.StartTimer()
				sort.Ints(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkSorted(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("pdqsort-%d", size), func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, size)
				for i := 0; i < len(data); i++ {
					data[i] = i
				}
				b.StartTimer()
				Slice(data)
				b.StopTimer()
			}
		})
		b.Run(fmt.Sprintf("stdsort-%d", size), func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, size)
				for i := 0; i < len(data); i++ {
					data[i] = i
				}
				b.StartTimer()
				sort.Ints(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("pdqsort-%d", size), func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, size)
				for i := 0; i < len(data); i++ {
					data[i] = len(data) - i
				}
				b.StartTimer()
				Slice(data)
				b.StopTimer()
			}
		})
		b.Run(fmt.Sprintf("stdsort-%d", size), func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, size)
				for i := 0; i < len(data); i++ {
					data[i] = len(data) - i
				}
				b.StartTimer()
				sort.Ints(data)
				b.StopTimer()
			}
		})
	}
}

func BenchmarkAlmostDuplicate(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("pdqsort-%d", size), func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, size)
				for i := 0; i < len(data); i++ {
					data[i] = fastrand.Intn(10)
				}
				b.StartTimer()
				Slice(data)
				b.StopTimer()
			}
		})
		b.Run(fmt.Sprintf("stdsort-%d", size), func(b *testing.B) {
			b.StopTimer()
			for i := 0; i < b.N; i++ {
				data := make([]int, size)
				for i := 0; i < len(data); i++ {
					data[i] = fastrand.Intn(10)
				}
				b.StartTimer()
				sort.Ints(data)
				b.StopTimer()
			}
		})
	}
}
