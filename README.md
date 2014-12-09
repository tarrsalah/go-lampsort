## Lampsort
Leslie Lamport's non-recursive quicksort described [here](http://bertrandmeyer.com/2014/12/07/lampsort/).

![](http://upload.wikimedia.org/wikipedia/commons/6/6a/Sorting_quicksort_anim.gif)


```
go test -bench=.
PASS
BenchmarkQuick	 1000000	      2216 ns/op
BenchmarkLamp	  500000	      4559 ns/op
ok  	github.com/tarrsalah/go-lampsort	4.573s
```
