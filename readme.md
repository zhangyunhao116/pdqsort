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

