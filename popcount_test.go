// Copyright 2015 Hideaki Ohno. All rights reserved.
// Use of this source code is governed by an MIT License
// that can be found in the LICENSE file.

package popcount

import (
	"encoding/binary"
	"testing"
	"testing/quick"
)

type testVector struct {
	n uint64
	s []uint64
	b []byte
}

func newTestVector(n uint64, s []uint64) testVector {
	b := make([]byte, len(s)*8)

	for i, x := range s {
		binary.LittleEndian.PutUint64(b[i*8:], x)
	}

	return testVector{n, s, b}
}

var testVectors = []testVector{
	newTestVector(120, []uint64{0xFF, 0xFFFF, 0xFFFFFFFF, 0xFFFFFFFFFFFFFFFF}),
	newTestVector(17, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
}

func testCountBytes(t *testing.T, count func(s []byte) uint64) {
	for _, tc := range testVectors {
		if n := count(tc.b); n != tc.n {
			t.Errorf("Expected %d, got %d", tc.n, n)
		}
	}
}

func TestCountBytes(t *testing.T) {
	testCountBytes(t, CountBytes)
}

func TestCountBytesGo(t *testing.T) {
	testCountBytes(t, countBytesGo)
}

func testCountSlice64(t *testing.T, count func(s []uint64) uint64) {
	for _, tc := range testVectors {
		if n := count(tc.s); n != tc.n {
			t.Errorf("Expected %d, got %d", tc.n, n)
		}
	}
}

func TestCountSlice64(t *testing.T) {
	testCountSlice64(t, CountSlice64)
}

func TestCountSlice64Go(t *testing.T) {
	testCountSlice64(t, countSlice64Go)
}

func TestCountBytesCompare(t *testing.T) {
	if !usePOPCNT {
		t.Skip("does not have POPCNT instruction")
	}

	for _, tc := range testVectors {
		if a, b := CountBytes(tc.b), countBytesGo(tc.b); a != b {
			t.Errorf("CountBytes(%[1]v) = %[2]d; countBytesGo(%[1]v) = %[3]d", tc.b, a, b)
		}
	}

	if err := quick.CheckEqual(countBytesGo, CountBytes, &quick.Config{
		MaxCountScale: 1000,
	}); err != nil {
		t.Error(err)
	}
}

func TestCountSlice64Compare(t *testing.T) {
	if !usePOPCNT {
		t.Skip("does not have POPCNT instruction")
	}

	for _, tc := range testVectors {
		if a, b := CountSlice64(tc.s), countSlice64Go(tc.s); a != b {
			t.Errorf("CountSlice64(%[1]v) = %[2]d; countSlice64Go(%[1]v) = %[3]d", tc.s, a, b)
		}
	}

	if err := quick.CheckEqual(countSlice64Go, CountSlice64, &quick.Config{
		MaxCountScale: 1000,
	}); err != nil {
		t.Error(err)
	}
}
