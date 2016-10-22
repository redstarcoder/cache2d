cache2d
======
[![GoDoc](https://godoc.org/github.com/redstarcoder/cache2d?status.svg)](https://godoc.org/github.com/redstarcoder/cache2d)

Package cache2d is for experimenting and demonstrating potential improvements caching can bring to [draw2d](https://github.com/llgcode/draw2d). Currently this relys on a [fork of draw2d](https://github.com/redstarcoder/draw2d) for `GraphicContext.GetPath`.

Benchmarks
---------------

```
$ go test -bench=. -benchtime=5s
testing: warning: no tests to run
BenchmarkFillStringAt-4         	    3000	   2152957 ns/op
BenchmarkFillStringAtCached-4   	   10000	    997937 ns/op
PASS
ok  	github.com/redstarcoder/cache2d	18.417s
```

```
$ go test -bench=. -benchtime=10s -cpu 1
testing: warning: no tests to run
BenchmarkFillStringAt       	   10000	   2108196 ns/op
BenchmarkFillStringAtCached 	   10000	   1203926 ns/op
PASS
ok  	github.com/redstarcoder/cache2d	34.654s
```

Acknowledgments
---------------

[redstarcoder](https://github.com/redstarcoder) wrote this library.
[Laurent Le Goff](https://github.com/llgcode) wrote draw2d and assisted in development via reviews and advice.

