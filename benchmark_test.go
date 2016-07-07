// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package popcount

import (
	"math/rand"
	"testing"
)

func randRead64(s []uint64) {
	for i := range s {
		s[i] = uint64(rand.Int63())
	}
}

func benchmarkCountBytes(b *testing.B, l int) {
	if !usePOPCNT {
		b.Log("does not have POPCNT instruction")
	}

	s := make([]byte, l)
	rand.Read(s)

	b.SetBytes(int64(l))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CountBytes(s)
	}
}

func BenchmarkCountBytes_32(b *testing.B) {
	benchmarkCountBytes(b, 32)
}

func BenchmarkCountBytes_128(b *testing.B) {
	benchmarkCountBytes(b, 128)
}

func BenchmarkCountBytes_1k(b *testing.B) {
	benchmarkCountBytes(b, 1*1024)
}

func BenchmarkCountBytes_16k(b *testing.B) {
	benchmarkCountBytes(b, 16*1024)
}

func BenchmarkCountBytes_128k(b *testing.B) {
	benchmarkCountBytes(b, 128*1024)
}

func BenchmarkCountBytes_1M(b *testing.B) {
	benchmarkCountBytes(b, 1024*1024)
}

func BenchmarkCountBytes_16M(b *testing.B) {
	benchmarkCountBytes(b, 16*1024*1024)
}

func BenchmarkCountBytes_128M(b *testing.B) {
	benchmarkCountBytes(b, 128*1024*1024)
}

func BenchmarkCountBytes_512M(b *testing.B) {
	benchmarkCountBytes(b, 512*1024*1024)
}

func benchmarkCountSlice64(b *testing.B, l int) {
	if !usePOPCNT {
		b.Log("does not have POPCNT instruction")
	}

	s := make([]uint64, l/8)
	randRead64(s)

	b.SetBytes(int64(l))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CountSlice64(s)
	}
}

func BenchmarkCountSlice64_32(b *testing.B) {
	benchmarkCountSlice64(b, 32)
}

func BenchmarkCountSlice64_128(b *testing.B) {
	benchmarkCountSlice64(b, 128)
}

func BenchmarkCountSlice64_1k(b *testing.B) {
	benchmarkCountSlice64(b, 1*1024)
}

func BenchmarkCountSlice64_16k(b *testing.B) {
	benchmarkCountSlice64(b, 16*1024)
}

func BenchmarkCountSlice64_128k(b *testing.B) {
	benchmarkCountSlice64(b, 128*1024)
}

func BenchmarkCountSlice64_1M(b *testing.B) {
	benchmarkCountSlice64(b, 1024*1024)
}

func BenchmarkCountSlice64_16M(b *testing.B) {
	benchmarkCountSlice64(b, 16*1024*1024)
}

func BenchmarkCountSlice64_128M(b *testing.B) {
	benchmarkCountSlice64(b, 128*1024*1024)
}

func BenchmarkCountSlice64_512M(b *testing.B) {
	benchmarkCountSlice64(b, 512*1024*1024)
}

func benchmarkCountBytesGo(b *testing.B, l int) {
	s := make([]byte, l)
	rand.Read(s)

	b.SetBytes(int64(l))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		countBytesGo(s)
	}
}

func BenchmarkCountBytesGo_32(b *testing.B) {
	benchmarkCountBytesGo(b, 32)
}

func BenchmarkCountBytesGo_128(b *testing.B) {
	benchmarkCountBytesGo(b, 128)
}

func BenchmarkCountBytesGo_1k(b *testing.B) {
	benchmarkCountBytesGo(b, 1*1024)
}

func BenchmarkCountBytesGo_16k(b *testing.B) {
	benchmarkCountBytesGo(b, 16*1024)
}

func BenchmarkCountBytesGo_128k(b *testing.B) {
	benchmarkCountBytesGo(b, 128*1024)
}

func BenchmarkCountBytesGo_1M(b *testing.B) {
	benchmarkCountBytesGo(b, 1024*1024)
}

func BenchmarkCountBytesGo_16M(b *testing.B) {
	benchmarkCountBytesGo(b, 16*1024*1024)
}

func BenchmarkCountBytesGo_128M(b *testing.B) {
	benchmarkCountBytesGo(b, 128*1024*1024)
}

func BenchmarkCountBytesGo_512M(b *testing.B) {
	benchmarkCountBytesGo(b, 512*1024*1024)
}

func benchmarkCountSlice64Go(b *testing.B, l int) {
	s := make([]uint64, l/8)
	randRead64(s)

	b.SetBytes(int64(l))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		countSlice64Go(s)
	}
}

func BenchmarkCountSlice64Go_32(b *testing.B) {
	benchmarkCountSlice64Go(b, 32)
}

func BenchmarkCountSlice64Go_128(b *testing.B) {
	benchmarkCountSlice64Go(b, 128)
}

func BenchmarkCountSlice64Go_1k(b *testing.B) {
	benchmarkCountSlice64Go(b, 1*1024)
}

func BenchmarkCountSlice64Go_16k(b *testing.B) {
	benchmarkCountSlice64Go(b, 16*1024)
}

func BenchmarkCountSlice64Go_128k(b *testing.B) {
	benchmarkCountSlice64Go(b, 128*1024)
}

func BenchmarkCountSlice64Go_1M(b *testing.B) {
	benchmarkCountSlice64Go(b, 1024*1024)
}

func BenchmarkCountSlice64Go_16M(b *testing.B) {
	benchmarkCountSlice64Go(b, 16*1024*1024)
}

func BenchmarkCountSlice64Go_128M(b *testing.B) {
	benchmarkCountSlice64Go(b, 128*1024*1024)
}

func BenchmarkCountSlice64Go_512M(b *testing.B) {
	benchmarkCountSlice64Go(b, 512*1024*1024)
}
