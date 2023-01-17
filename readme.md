# pdqsort

> The pdqsort has been merged into the Go standard library since Go 1.19, please use `sort` or `slices` directly instead of this package.
>
> issue: https://github.com/golang/go/issues/50154
>
> commit: https://github.com/golang/go/commit/72e77a7f41bbf45d466119444307fd3ae996e257

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
name                          time/op
Random/pdqsort_64             1.18µs ± 0%
Random/stdsort_64             2.38µs ± 0%
Random/pdqsort_256            6.24µs ± 3%
Random/stdsort_256            13.2µs ± 7%
Random/pdqsort_1024           32.4µs ± 0%
Random/stdsort_1024           62.2µs ± 0%
Random/pdqsort_4096            149µs ± 0%
Random/stdsort_4096            291µs ± 0%
Random/pdqsort_65536          3.14ms ± 0%
Random/stdsort_65536          6.11ms ± 0%
Sorted/pdqsort_64             94.4ns ± 4%
Sorted/stdsort_64              711ns ± 0%
Sorted/pdqsort_256             171ns ± 1%
Sorted/stdsort_256            3.37µs ± 0%
Sorted/pdqsort_1024            507ns ± 0%
Sorted/stdsort_1024           16.6µs ± 0%
Sorted/pdqsort_4096           1.82µs ± 0%
Sorted/stdsort_4096           78.9µs ± 0%
Sorted/pdqsort_65536          28.2µs ± 0%
Sorted/stdsort_65536          1.73ms ± 0%
NearlySorted/pdqsort_64        353ns ± 2%
NearlySorted/stdsort_64        931ns ± 1%
NearlySorted/pdqsort_256      1.78µs ± 1%
NearlySorted/stdsort_256      4.99µs ± 1%
NearlySorted/pdqsort_1024     8.28µs ± 0%
NearlySorted/stdsort_1024     26.1µs ± 1%
NearlySorted/pdqsort_4096     38.2µs ± 0%
NearlySorted/stdsort_4096      117µs ± 1%
NearlySorted/pdqsort_65536     792µs ± 0%
NearlySorted/stdsort_65536    2.46ms ± 0%
Reversed/pdqsort_64            113ns ± 1%
Reversed/stdsort_64            845ns ± 0%
Reversed/pdqsort_256           253ns ± 1%
Reversed/stdsort_256          3.69µs ± 0%
Reversed/pdqsort_1024          785ns ± 0%
Reversed/stdsort_1024         17.8µs ± 1%
Reversed/pdqsort_4096         2.89µs ± 0%
Reversed/stdsort_4096         83.7µs ± 0%
Reversed/pdqsort_65536        45.1µs ± 0%
Reversed/stdsort_65536        1.80ms ± 0%
NearlyReversed/pdqsort_64      435ns ± 2%
NearlyReversed/stdsort_64     1.33µs ± 1%
NearlyReversed/pdqsort_256    2.13µs ± 1%
NearlyReversed/stdsort_256    7.23µs ± 1%
NearlyReversed/pdqsort_1024   10.6µs ± 1%
NearlyReversed/stdsort_1024   38.4µs ± 1%
NearlyReversed/pdqsort_4096   50.2µs ± 1%
NearlyReversed/stdsort_4096    176µs ± 0%
NearlyReversed/pdqsort_65536  1.04ms ± 1%
NearlyReversed/stdsort_65536  3.57ms ± 0%
Mod8/pdqsort_64                345ns ± 2%
Mod8/stdsort_64                949ns ± 1%
Mod8/pdqsort_256              1.02µs ± 1%
Mod8/stdsort_256              3.78µs ± 0%
Mod8/pdqsort_1024             3.13µs ± 2%
Mod8/stdsort_1024             14.3µs ± 0%
Mod8/pdqsort_4096             10.2µs ± 1%
Mod8/stdsort_4096             59.3µs ± 0%
Mod8/pdqsort_65536             190µs ± 2%
Mod8/stdsort_65536            1.00ms ± 0%
AllEqual/pdqsort_64           89.4ns ± 3%
AllEqual/stdsort_64            381ns ± 1%
AllEqual/pdqsort_256           170ns ± 1%
AllEqual/stdsort_256          1.03µs ± 1%
AllEqual/pdqsort_1024          507ns ± 0%
AllEqual/stdsort_1024         3.66µs ± 0%
AllEqual/pdqsort_4096         1.82µs ± 0%
AllEqual/stdsort_4096         14.0µs ± 0%
AllEqual/pdqsort_65536        28.2µs ± 0%
AllEqual/stdsort_65536         224µs ± 0%
```

