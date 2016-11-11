# go-popcount

[![GoDoc](https://godoc.org/github.com/tmthrgd/go-popcount?status.svg)](https://godoc.org/github.com/tmthrgd/go-popcount)
[![Build Status](https://travis-ci.org/tmthrgd/go-popcount.svg?branch=master)](https://travis-ci.org/tmthrgd/go-popcount)

A population count implementation for Golang.

An x86-64 implementation is provided that uses the POPCNT instruction.

## Download

```
go get github.com/tmthrgd/go-popcount
```

## Benchmark

```
BenchmarkCountBytes/32-8      	300000000	         5.83 ns/op	5488.90 MB/s
BenchmarkCountBytes/128-8     	100000000	        10.0 ns/op	12747.10 MB/s
BenchmarkCountBytes/1k-8      	30000000	        47.3 ns/op	21637.04 MB/s
BenchmarkCountBytes/16k-8     	 3000000	       577 ns/op	28374.53 MB/s
BenchmarkCountBytes/128k-8    	  300000	      4987 ns/op	26281.76 MB/s
BenchmarkCountBytes/1M-8      	   30000	     50773 ns/op	20652.13 MB/s
BenchmarkCountBytes/16M-8     	    1000	   1304461 ns/op	12861.41 MB/s
BenchmarkCountBytes/128M-8    	     100	  10249707 ns/op	13094.79 MB/s
BenchmarkCountBytes/512M-8    	      30	  41865546 ns/op	12823.69 MB/s
BenchmarkCountSlice64/32-8    	200000000	         6.30 ns/op	5082.89 MB/s
BenchmarkCountSlice64/128-8   	100000000	        11.2 ns/op	11459.70 MB/s
BenchmarkCountSlice64/1k-8    	30000000	        50.4 ns/op	20315.12 MB/s
BenchmarkCountSlice64/16k-8   	 3000000	       607 ns/op	26968.84 MB/s
BenchmarkCountSlice64/128k-8  	  300000	      5319 ns/op	24639.93 MB/s
BenchmarkCountSlice64/1M-8    	   30000	     52579 ns/op	19942.75 MB/s
BenchmarkCountSlice64/16M-8   	    1000	   1284949 ns/op	13056.72 MB/s
BenchmarkCountSlice64/128M-8  	     100	  10340755 ns/op	12979.49 MB/s
BenchmarkCountSlice64/512M-8  	      30	  41242836 ns/op	13017.31 MB/s
BenchmarkCountBytesGo/32-8    	50000000	        23.8 ns/op	1346.65 MB/s
BenchmarkCountBytesGo/128-8   	20000000	        66.9 ns/op	1914.63 MB/s
BenchmarkCountBytesGo/1k-8    	 3000000	       470 ns/op	2174.79 MB/s
BenchmarkCountBytesGo/16k-8   	  200000	      7284 ns/op	2249.15 MB/s
BenchmarkCountBytesGo/128k-8  	   30000	     58083 ns/op	2256.61 MB/s
BenchmarkCountBytesGo/1M-8    	    3000	    456864 ns/op	2295.16 MB/s
BenchmarkCountBytesGo/16M-8   	     200	   7491107 ns/op	2239.62 MB/s
BenchmarkCountBytesGo/128M-8  	      20	  63084066 ns/op	2127.60 MB/s
BenchmarkCountBytesGo/512M-8  	       5	 239421622 ns/op	2242.37 MB/s
BenchmarkCountSlice64Go/32-8  	100000000	        15.8 ns/op	2024.47 MB/s
BenchmarkCountSlice64Go/128-8 	30000000	        57.0 ns/op	2244.07 MB/s
BenchmarkCountSlice64Go/1k-8  	 3000000	       448 ns/op	2283.13 MB/s
BenchmarkCountSlice64Go/16k-8 	  200000	      7048 ns/op	2324.58 MB/s
BenchmarkCountSlice64Go/128k-8	   30000	     56173 ns/op	2333.34 MB/s
BenchmarkCountSlice64Go/1M-8  	    3000	    450511 ns/op	2327.53 MB/s
BenchmarkCountSlice64Go/16M-8 	     200	   7534657 ns/op	2226.67 MB/s
BenchmarkCountSlice64Go/128M-8	      20	  62670803 ns/op	2141.63 MB/s
BenchmarkCountSlice64Go/512M-8	       5	 251465567 ns/op	2134.97 MB/s
```

## License

Unless otherwise noted, the go-popcount source files are distributed under the Modified BSD License
found in the LICENSE file.
