// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

package popcount

import (
	"math/rand"
	"testing"
)

type size struct {
	name string
	l    int
}

var sizes = []size{
	{"32", 32},
	{"128", 128},
	{"1K", 1 * 1024},
	{"16K", 16 * 1024},
	{"128K", 128 * 1024},
	{"1M", 1024 * 1024},
	{"16M", 16 * 1024 * 1024},
	{"128M", 128 * 1024 * 1024},
	{"512M", 512 * 1024 * 1024},
}

func randRead64(s []uint64) {
	for i := range s {
		s[i] = uint64(rand.Int63())
	}
}

func BenchmarkCountBytes(b *testing.B) {
	if !usePOPCNT {
		b.Log("does not have POPCNT instruction")
	}

	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			s := make([]byte, size.l)
			rand.Read(s)

			b.SetBytes(int64(size.l))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				CountBytes(s)
			}
		})
	}
}

func BenchmarkCountSlice64(b *testing.B) {
	if !usePOPCNT {
		b.Log("does not have POPCNT instruction")
	}

	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			s := make([]uint64, size.l/8)
			randRead64(s)

			b.SetBytes(int64(size.l))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				CountSlice64(s)
			}
		})
	}
}

func BenchmarkCountBytesGo(b *testing.B) {
	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			s := make([]byte, size.l)
			rand.Read(s)

			b.SetBytes(int64(size.l))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				countBytesGo(s)
			}
		})
	}
}

func BenchmarkCountSlice64Go(b *testing.B) {
	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			s := make([]uint64, size.l/8)
			randRead64(s)

			b.SetBytes(int64(size.l))
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				countSlice64Go(s)
			}
		})
	}
}

func BenchmarkCount64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Count64(^uint64(0))
	}
}
