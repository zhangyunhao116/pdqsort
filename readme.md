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
Random/pdqsort-64         1.19µs ± 0%
Random/stdsort-64         2.37µs ± 0%
Random/pdqsort-256        6.36µs ± 1%
Random/stdsort-256        12.2µs ± 0%
Random/pdqsort-1024       31.7µs ± 1%
Random/stdsort-1024       60.0µs ± 0%
Random/pdqsort-4096        151µs ± 0%
Random/stdsort-4096        287µs ± 1%
Random/pdqsort-65536      3.21ms ± 0%
Random/stdsort-65536      6.09ms ± 0%
Sorted/pdqsort-64         90.9ns ± 1%
Sorted/stdsort-64          738ns ± 0%
Sorted/pdqsort-256         171ns ± 1%
Sorted/stdsort-256        3.43µs ± 1%
Sorted/pdqsort-1024        515ns ± 1%
Sorted/stdsort-1024       16.7µs ± 1%
Sorted/pdqsort-4096       1.82µs ± 0%
Sorted/stdsort-4096       80.1µs ± 2%
Sorted/pdqsort-65536      28.3µs ± 0%
Sorted/stdsort-65536      1.76ms ± 0%
Sorted90/pdqsort-64        148ns ± 2%
Sorted90/stdsort-64        760ns ± 0%
Sorted90/pdqsort-256      1.18µs ± 1%
Sorted90/stdsort-256      3.99µs ± 1%
Sorted90/pdqsort-1024     4.24µs ± 0%
Sorted90/stdsort-1024     19.8µs ± 1%
Sorted90/pdqsort-4096     18.3µs ± 0%
Sorted90/stdsort-4096     96.7µs ± 1%
Sorted90/pdqsort-65536     364µs ± 0%
Sorted90/stdsort-65536    2.14ms ± 0%
Reversed/pdqsort-64        117ns ± 2%
Reversed/stdsort-64        956ns ± 2%
Reversed/pdqsort-256       256ns ± 0%
Reversed/stdsort-256      3.73µs ± 2%
Reversed/pdqsort-1024      813ns ± 2%
Reversed/stdsort-1024     17.9µs ± 2%
Reversed/pdqsort-4096     2.99µs ± 0%
Reversed/stdsort-4096     84.4µs ± 1%
Reversed/pdqsort-65536    45.4µs ± 0%
Reversed/stdsort-65536    1.82ms ± 0%
Reversed90/pdqsort-64      497ns ± 1%
Reversed90/stdsort-64      955ns ± 1%
Reversed90/pdqsort-256    2.02µs ± 1%
Reversed90/stdsort-256    4.89µs ± 1%
Reversed90/pdqsort-1024   9.40µs ± 1%
Reversed90/stdsort-1024   23.8µs ± 1%
Reversed90/pdqsort-4096   37.0µs ± 1%
Reversed90/stdsort-4096    121µs ± 1%
Reversed90/pdqsort-65536   658µs ± 0%
Reversed90/stdsort-65536  2.54ms ± 0%
Mod8/pdqsort-64            345ns ± 2%
Mod8/stdsort-64           1.02µs ± 1%
Mod8/pdqsort-256          1.06µs ± 2%
Mod8/stdsort-256          3.81µs ± 1%
Mod8/pdqsort-1024         3.57µs ± 0%
Mod8/stdsort-1024         14.4µs ± 1%
Mod8/pdqsort-4096         12.3µs ± 1%
Mod8/stdsort-4096         60.4µs ± 1%
Mod8/pdqsort-65536         221µs ± 0%
Mod8/stdsort-65536        1.02ms ± 0%
```

