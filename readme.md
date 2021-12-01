# pdqsort

Pattern-defeating quicksort (pdqsort) is a novel sorting algorithm that combines the fast average case of randomized quicksort with the fast worst case of heapsort, while achieving linear time on inputs with certain patterns. The algorithm is based on pattern-defeating quicksort by Orson Peters, the original C++ version  implementation: https://github.com/orlp/pdqsort

```
Best        Average     Worst       Memory      Stable      Deterministic
n           n log n     n log n     log n       No          Yes
```



## Benchmark

Go version: go1.18 linux/amd64

CPU: AMD 3700x(8C16T), running at 3.6GHz

OS: ubuntu 18.04

MEMORY: 16G x 2 (3200MHz)

```text
name                           time/op
Random/pdqsort-32               478ns ± 1%
Random/stdsort-32              1.06µs ± 1%
Random/pdqsort-256             6.32µs ± 0%
Random/stdsort-256             12.7µs ± 0%
Random/pdqsort-1024            31.4µs ± 0%
Random/stdsort-1024            62.3µs ± 0%
Random/pdqsort-4096             150µs ± 0%
Random/stdsort-4096             295µs ± 0%
Random/pdqsort-65536           3.18ms ± 0%
Random/stdsort-65536           6.21ms ± 0%
Sorted/pdqsort-32              57.4ns ± 1%
Sorted/stdsort-32               171ns ± 0%
Sorted/pdqsort-256              172ns ± 1%
Sorted/stdsort-256              597ns ± 0%
Sorted/pdqsort-1024             506ns ± 1%
Sorted/stdsort-1024            2.05µs ± 1%
Sorted/pdqsort-4096            1.82µs ± 0%
Sorted/stdsort-4096            7.48µs ± 0%
Sorted/pdqsort-65536           28.3µs ± 0%
Sorted/stdsort-65536            116µs ± 0%
Reverse/pdqsort-32              108ns ± 1%
Reverse/stdsort-32              505ns ± 1%
Reverse/pdqsort-256             256ns ± 0%
Reverse/stdsort-256             856ns ± 1%
Reverse/pdqsort-1024            786ns ± 0%
Reverse/stdsort-1024           2.94µs ± 1%
Reverse/pdqsort-4096           2.90µs ± 0%
Reverse/stdsort-4096           10.9µs ± 1%
Reverse/pdqsort-65536          45.2µs ± 0%
Reverse/stdsort-65536           173µs ± 0%
AlmostDuplicate/pdqsort-32      455ns ± 1%
AlmostDuplicate/stdsort-32      959ns ± 1%
AlmostDuplicate/pdqsort-256    3.47µs ± 0%
AlmostDuplicate/stdsort-256    6.71µs ± 0%
AlmostDuplicate/pdqsort-1024   11.7µs ± 0%
AlmostDuplicate/stdsort-1024   23.4µs ± 1%
AlmostDuplicate/pdqsort-4096   43.4µs ± 0%
AlmostDuplicate/stdsort-4096   87.9µs ± 0%
AlmostDuplicate/pdqsort-65536   672µs ± 0%
AlmostDuplicate/stdsort-65536  1.36ms ± 0%
```

