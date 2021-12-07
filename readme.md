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
Random/pdqsort-64         1.18µs ± 1%
Random/stdsort-64         2.37µs ± 0%
Random/pdqsort-256        6.28µs ± 0%
Random/stdsort-256        12.2µs ± 0%
Random/pdqsort-1024       31.0µs ± 0%
Random/stdsort-1024       60.0µs ± 0%
Random/pdqsort-4096        148µs ± 0%
Random/stdsort-4096        291µs ± 1%
Random/pdqsort-65536      3.16ms ± 0%
Random/stdsort-65536      6.10ms ± 0%
Sorted/pdqsort-64         87.1ns ± 1%
Sorted/stdsort-64          738ns ± 0%
Sorted/pdqsort-256         170ns ± 0%
Sorted/stdsort-256        3.42µs ± 1%
Sorted/pdqsort-1024        501ns ± 0%
Sorted/stdsort-1024       16.7µs ± 1%
Sorted/pdqsort-4096       1.82µs ± 0%
Sorted/stdsort-4096       80.1µs ± 2%
Sorted/pdqsort-65536      28.3µs ± 1%
Sorted/stdsort-65536      1.76ms ± 0%
Sorted90/pdqsort-64        146ns ± 1%
Sorted90/stdsort-64        767ns ± 1%
Sorted90/pdqsort-256      1.18µs ± 0%
Sorted90/stdsort-256      3.99µs ± 0%
Sorted90/pdqsort-1024     4.09µs ± 0%
Sorted90/stdsort-1024     19.7µs ± 0%
Sorted90/pdqsort-4096     17.5µs ± 0%
Sorted90/stdsort-4096     96.6µs ± 1%
Sorted90/pdqsort-65536     350µs ± 0%
Sorted90/stdsort-65536    2.14ms ± 0%
Reversed/pdqsort-64        117ns ± 1%
Reversed/stdsort-64        850ns ± 1%
Reversed/pdqsort-256       256ns ± 0%
Reversed/stdsort-256      3.71µs ± 1%
Reversed/pdqsort-1024      785ns ± 0%
Reversed/stdsort-1024     18.0µs ± 1%
Reversed/pdqsort-4096     2.96µs ± 0%
Reversed/stdsort-4096     84.8µs ± 1%
Reversed/pdqsort-65536    45.2µs ± 0%
Reversed/stdsort-65536    1.82ms ± 0%
Reversed90/pdqsort-64      488ns ± 1%
Reversed90/stdsort-64      953ns ± 1%
Reversed90/pdqsort-256    1.99µs ± 1%
Reversed90/stdsort-256    4.89µs ± 1%
Reversed90/pdqsort-1024   9.43µs ± 1%
Reversed90/stdsort-1024   23.8µs ± 2%
Reversed90/pdqsort-4096   36.0µs ± 0%
Reversed90/stdsort-4096    121µs ± 2%
Reversed90/pdqsort-65536   642µs ± 0%
Reversed90/stdsort-65536  2.54ms ± 0%
Mod8/pdqsort-64            336ns ± 1%
Mod8/stdsort-64            975ns ± 0%
Mod8/pdqsort-256          1.09µs ± 2%
Mod8/stdsort-256          3.80µs ± 1%
Mod8/pdqsort-1024         3.10µs ± 2%
Mod8/stdsort-1024         14.4µs ± 2%
Mod8/pdqsort-4096         10.6µs ± 1%
Mod8/stdsort-4096         60.7µs ± 2%
Mod8/pdqsort-65536         202µs ± 1%
Mod8/stdsort-65536        1.02ms ± 0%
```

