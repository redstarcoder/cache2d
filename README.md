cache2d
======

Package cache2d is for experimenting and demonstrating potential improvements caching can bring to [draw2d](https://github.com/llgcode/draw2d). Currently this relys on a [fork of draw2d](github.com/redstarcoder/draw2d) for `PathBuilder.CopyPath`.

Benchmarks
---------------

```
$ go test -bench . -benchtime 10s
testing: warning: no tests to run
BenchmarkFillStringAt-4         	   10000	   2239271 ns/op
BenchmarkFillStringAtCached-4   	   10000	   1108676 ns/op
PASS
ok  	github.com/redstarcoder/cache2d	40.433s
```

Acknowledgments
---------------

[redstarcoder](https://github.com/redstarcoder) wrote this library.
[Laurent Le Goff](https://github.com/llgcode) wrote draw2d and assisted in development via reviews and advice.

