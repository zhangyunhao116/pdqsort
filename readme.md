# pdqsort

The algorithm is mainly based on pattern-defeating quicksort by Orson Peters.

Compared to sort.Ints(Go1.18), it is **2x** faster in random slices, and **3x ~ 60x** faster in common patterns.

- C++  implementation: https://github.com/orlp/pdqsort
- Rust implementation: https://docs.rs/pdqsort/latest/pdqsort/

```
Best        Average     Worst       Memory      Stable      Deterministic
n           n log n     n log n     log n       No          Yes
```



## Features

- **Unstable sort**, may reorder equal elements.
- Disable the optimization from [BlockQuickSort](https://dl.acm.org/doi/10.1145/3274660), since its poor performance in Go.



## QuickStart

```go
package main

import (
	"fmt"

	"github.com/zhangyunhao116/pdqsort"
)

func main() {
	x := []int{3, 1, 2, 4, 5, 9, 8, 7}
	pdqsort.Slice(x)
	fmt.Printf("%v\n", x)
}

```



## Benchmark

Go version: go1.18-a412b5f0d8 linux/amd64

CPU: Intel 11700k(8C16T)

OS: ubuntu 20.04

MEMORY: 16G x 2 (3200MHz)

```text
name                           time/op
Random/pdqsort-32               476ns ± 3%
Random/stdsort-32              1.02µs ± 0%
Random/pdqsort-256             6.33µs ± 0%
Random/stdsort-256             12.2µs ± 0%
Random/pdqsort-1024            31.4µs ± 0%
Random/stdsort-1024            60.7µs ± 2%
Random/pdqsort-4096             150µs ± 0%
Random/stdsort-4096             291µs ± 1%
Random/pdqsort-65536           3.18ms ± 0%
Random/stdsort-65536           6.09ms ± 0%
Sorted/pdqsort-32              57.0ns ± 2%
Sorted/stdsort-32               331ns ± 1%
Sorted/pdqsort-256              171ns ± 1%
Sorted/stdsort-256             3.43µs ± 1%
Sorted/pdqsort-1024             508ns ± 1%
Sorted/stdsort-1024            16.7µs ± 1%
Sorted/pdqsort-4096            1.81µs ± 0%
Sorted/stdsort-4096            82.0µs ± 1%
Sorted/pdqsort-65536           28.3µs ± 0%
Sorted/stdsort-65536           1.75ms ± 0%
Reverse/pdqsort-32              110ns ± 1%
Reverse/stdsort-32              367ns ± 1%
Reverse/pdqsort-256             255ns ± 0%
Reverse/stdsort-256            3.72µs ± 1%
Reverse/pdqsort-1024            785ns ± 0%
Reverse/stdsort-1024           17.8µs ± 0%
Reverse/pdqsort-4096           2.90µs ± 0%
Reverse/stdsort-4096           84.1µs ± 0%
Reverse/pdqsort-65536          45.2µs ± 0%
Reverse/stdsort-65536          1.81ms ± 0%
AlmostDuplicate/pdqsort-32      454ns ± 2%
AlmostDuplicate/stdsort-32      900ns ± 1%
AlmostDuplicate/pdqsort-256    3.46µs ± 0%
AlmostDuplicate/stdsort-256    6.76µs ± 0%
AlmostDuplicate/pdqsort-1024   11.7µs ± 1%
AlmostDuplicate/stdsort-1024   24.4µs ± 0%
AlmostDuplicate/pdqsort-4096   43.4µs ± 0%
AlmostDuplicate/stdsort-4096   94.0µs ± 0%
AlmostDuplicate/pdqsort-65536   670µs ± 0%
AlmostDuplicate/stdsort-65536  1.48ms ± 0%
```

