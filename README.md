# Binary Search

Find the position of an item within a sorted array.

Worst Case: O(log n)

Best Case: O(1)

Avg Case: O(log n)

Space Complexity: O(1)

Benchmark:
```sh
goos: linux
goarch: amd64
pkg: github.com/theorangutan/binarysearch
BenchmarkIterative10-8     	30000000	        51.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkIterative100-8    	20000000	        83.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkIterative1000-8   	20000000	       109 ns/op	       0 B/op	       0 allocs/op
BenchmarkRecursive10-8     	20000000	        63.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkRecursive100-8    	20000000	       106 ns/op	       0 B/op	       0 allocs/op
BenchmarkRecursive1000-8   	10000000	       150 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 88.9% of statements
ok  	github.com/theorangutan/binarysearch	10.937s
Success: Benchmarks passed.
```