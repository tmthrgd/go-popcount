// Copyright 2015 Hideaki Ohno. All rights reserved.
// Use of this source code is governed by an MIT License
// that can be found in the LICENSE file.

// Package popcount is a population count implementation for Golang.
package popcount

// Count64 function counts the number of non-zero bits in a 64bit unsigned integer.
func Count64(x uint64) uint64 {
	x = (x & 0x5555555555555555) + ((x & 0xAAAAAAAAAAAAAAAA) >> 1)
	x = (x & 0x3333333333333333) + ((x & 0xCCCCCCCCCCCCCCCC) >> 2)
	x = (x & 0x0F0F0F0F0F0F0F0F) + ((x & 0xF0F0F0F0F0F0F0F0) >> 4)
	x *= 0x0101010101010101
	return ((x >> 56) & 0xFF)
}

func countSlice64Go(s []uint64) (count uint64) {
	for _, x := range s {
		count += Count64(x)
	}

	return
}
