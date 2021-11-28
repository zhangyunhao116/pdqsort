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
Random/pdqsort-16               225ns ± 1%
Random/stdsort-16               528ns ± 1%
Random/pdqsort-256             7.17µs ± 0%
Random/stdsort-256             13.1µs ± 0%
Random/pdqsort-1024            35.1µs ± 0%
Random/stdsort-1024            64.1µs ± 0%
Random/pdqsort-4096             167µs ± 0%
Random/stdsort-4096             308µs ± 1%
Random/pdqsort-65536           3.50ms ± 0%
Random/stdsort-65536           6.60ms ± 0%
Sorted/pdqsort-16              42.5ns ± 1%
Sorted/stdsort-16               228ns ± 1%
Sorted/pdqsort-256              251ns ± 1%
Sorted/stdsort-256             4.59µs ± 0%
Sorted/pdqsort-1024             777ns ± 0%
Sorted/stdsort-1024            22.4µs ± 0%
Sorted/pdqsort-4096            2.92µs ± 0%
Sorted/stdsort-4096             106µs ± 1%
Sorted/pdqsort-65536           45.1µs ± 0%
Sorted/stdsort-65536           2.32ms ± 0%
Reverse/pdqsort-16              121ns ± 1%
Reverse/stdsort-16              246ns ± 2%
Reverse/pdqsort-256             335ns ± 1%
Reverse/stdsort-256            4.92µs ± 0%
Reverse/pdqsort-1024           1.06µs ± 0%
Reverse/stdsort-1024           23.7µs ± 0%
Reverse/pdqsort-4096           3.95µs ± 0%
Reverse/stdsort-4096            112µs ± 1%
Reverse/pdqsort-65536          61.5µs ± 0%
Reverse/stdsort-65536          2.36ms ± 0%
AlmostDuplicate/pdqsort-16      218ns ± 1%
AlmostDuplicate/stdsort-16      494ns ± 1%
AlmostDuplicate/pdqsort-256    3.81µs ± 0%
AlmostDuplicate/stdsort-256    7.44µs ± 0%
AlmostDuplicate/pdqsort-1024   13.0µs ± 0%
AlmostDuplicate/stdsort-1024   27.0µs ± 0%
AlmostDuplicate/pdqsort-4096   48.2µs ± 0%
AlmostDuplicate/stdsort-4096    104µs ± 0%
AlmostDuplicate/pdqsort-65536   727µs ± 0%
AlmostDuplicate/stdsort-65536  1.62ms ± 0%
```

