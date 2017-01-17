// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

// +build appengine

package popcount

import "encoding/binary"

func countBytesGo(s []byte) (count uint64) {
	for i := 0; i+8 <= len(s); i += 8 {
		x := binary.LittleEndian.Uint64(s[i:])
		count += Count64(x)
	}

	s = s[len(s)&^7:]

	var left uint64

	if len(s) >= 4 {
		left = uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
		s = s[4:]
	}

	if len(s) >= 2 {
		left = left<<16 | uint64(s[1])<<8 | uint64(s[0])
		s = s[2:]
	}

	if len(s) == 1 {
		left = left<<8 | uint64(s[0])
	}

	count += Count64(left)
	return
}
