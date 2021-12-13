# pdqsort

The algorithm is mainly based on pattern-defeating quicksort by Orson Peters.

Compared to sort.Ints(Go1.18), it is **2x** faster in random slices, and **2x ~ 60x** faster in common patterns.

- Paper: https://arxiv.org/pdf/2106.05123.pdf
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
name                      time/op
Random/pdqsort-64         1.19µs ± 1%
Random/stdsort-64         2.38µs ± 0%
Random/pdqsort-256        6.35µs ± 1%
Random/stdsort-256        13.7µs ± 0%
Random/pdqsort-1024       33.1µs ± 0%
Random/stdsort-1024       62.1µs ± 0%
Random/pdqsort-4096        153µs ± 0%
Random/stdsort-4096        290µs ± 0%
Random/pdqsort-65536      3.20ms ± 0%
Random/stdsort-65536      6.12ms ± 0%
Sorted/pdqsort-64         90.1ns ± 1%
Sorted/stdsort-64          714ns ± 1%
Sorted/pdqsort-256         172ns ± 0%
Sorted/stdsort-256        3.37µs ± 0%
Sorted/pdqsort-1024        505ns ± 0%
Sorted/stdsort-1024       16.6µs ± 0%
Sorted/pdqsort-4096       1.81µs ± 0%
Sorted/stdsort-4096       79.0µs ± 0%
Sorted/pdqsort-65536      28.2µs ± 0%
Sorted/stdsort-65536      1.74ms ± 0%
Sorted90/pdqsort-64        147ns ± 2%
Sorted90/stdsort-64        750ns ± 1%
Sorted90/pdqsort-256      1.15µs ± 1%
Sorted90/stdsort-256      3.95µs ± 0%
Sorted90/pdqsort-1024     4.15µs ± 0%
Sorted90/stdsort-1024     19.7µs ± 0%
Sorted90/pdqsort-4096     17.8µs ± 0%
Sorted90/stdsort-4096     96.0µs ± 1%
Sorted90/pdqsort-65536     354µs ± 0%
Sorted90/stdsort-65536    2.12ms ± 0%
Reversed/pdqsort-64        116ns ± 1%
Reversed/stdsort-64        847ns ± 0%
Reversed/pdqsort-256       258ns ± 1%
Reversed/stdsort-256      3.72µs ± 1%
Reversed/pdqsort-1024      786ns ± 0%
Reversed/stdsort-1024     17.8µs ± 0%
Reversed/pdqsort-4096     2.90µs ± 0%
Reversed/stdsort-4096     84.1µs ± 0%
Reversed/pdqsort-65536    45.2µs ± 0%
Reversed/stdsort-65536    1.80ms ± 1%
Reversed90/pdqsort-64      449ns ± 2%
Reversed90/stdsort-64      938ns ± 1%
Reversed90/pdqsort-256    1.98µs ± 0%
Reversed90/stdsort-256    4.87µs ± 1%
Reversed90/pdqsort-1024   9.46µs ± 1%
Reversed90/stdsort-1024   23.7µs ± 1%
Reversed90/pdqsort-4096   36.6µs ± 0%
Reversed90/stdsort-4096    119µs ± 0%
Reversed90/pdqsort-65536   646µs ± 0%
Reversed90/stdsort-65536  2.52ms ± 0%
Mod8/pdqsort-64            338ns ± 1%
Mod8/stdsort-64            945ns ± 0%
Mod8/pdqsort-256          1.09µs ± 1%
Mod8/stdsort-256          3.79µs ± 1%
Mod8/pdqsort-1024         3.58µs ± 0%
Mod8/stdsort-1024         14.2µs ± 0%
Mod8/pdqsort-4096         12.0µs ± 0%
Mod8/stdsort-4096         59.1µs ± 0%
Mod8/pdqsort-65536         206µs ± 1%
Mod8/stdsort-65536        1.01ms ± 0%
```

