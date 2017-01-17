// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

// +build ignore

package main

import "github.com/tmthrgd/asm"

const header = `// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.
//
// This file is auto-generated - do not modify

// +build amd64,!gccgo,!appengine
`

func countBytesASM(a *asm.Asm) {
	a.NewFunction("countBytesASM")
	a.NoSplit()

	src := a.Argument("src", 8)
	length := a.Argument("len", 8)
	retArg := a.Argument("ret", 8)

	a.Start()

	bigloop := a.NewLabel("bigloop")
	loop := a.NewLabel("loop")
	tail := a.NewLabel("tail")
	ret := a.NewLabel("ret")

	si, cx, ax := asm.SI, asm.BX, asm.AX

	a.Movq(si, src)
	a.Movq(cx, length)

	/*
	 * Important performance information can be found at:
	 * http://stackoverflow.com/a/25089720
	 *
	 * POPCNT has a false-dependency bug that causes a
	 * performance hit. Thus, in bigloop four separate
	 * destination registers are used to allow
	 * intra-loop parallelization, also and in loop the
	 * the destination register is cleared (with no
	 * practical effect) before POPCNT to allow
	 * inter-loop parallelization.
	 */

	a.Xorq(ax, ax)

	a.Cmpq(asm.Constant(8), cx)
	a.Jb(tail)

	a.Cmpq(asm.Constant(32), cx)
	a.Jb(loop)

	a.Label(bigloop)

	for i, r := range []asm.Operand{asm.R11, asm.R10, asm.R9, asm.R8} {
		a.Popcntq(r, asm.Address(si, cx, asm.SX1, -8*(i+1)))
	}

	for _, r := range []asm.Operand{asm.R11, asm.R10, asm.R9, asm.R8} {
		a.Addq(ax, r)
	}

	a.Subq(cx, asm.Constant(32))
	a.Jz(ret)

	a.Cmpq(asm.Constant(32), cx)
	a.Jae(bigloop)

	a.Cmpq(asm.Constant(8), cx)
	a.Jb(tail)

	a.Label(loop)

	a.Xorq(asm.DX, asm.DX)
	a.Popcntq(asm.DX, asm.Address(asm.SI, asm.BX, asm.SX1, -8))
	a.Addq(ax, asm.DX)

	a.Subq(cx, asm.Constant(8))
	a.Jz(ret)

	a.Cmpq(asm.Constant(8), cx)
	a.Jae(loop)

	a.Label(tail)

	tail2 := tail.Suffix("2")
	tail3 := tail.Suffix("3")
	tail4 := tail.Suffix("4")

	a.Xorq(asm.DX, asm.DX)

	a.Cmpq(asm.Constant(4), cx)
	a.Jb(tail2)

	a.Movl(asm.DX, asm.Address(si, cx, asm.SX1, -4))

	a.Subq(cx, asm.Constant(4))
	a.Jz(tail4)

	a.Label(tail2)

	a.Cmpq(asm.Constant(2), cx)
	a.Jb(tail3)

	a.Shlq(asm.DX, asm.Constant(16))
	a.Orw(asm.DX, asm.Address(si, cx, asm.SX1, -2))

	a.Subq(cx, asm.Constant(2))
	a.Jz(tail4)

	a.Label(tail3)

	a.Shlq(asm.DX, asm.Constant(8))
	a.Orb(asm.DX, asm.Address(si, cx, asm.SX1, -1))

	a.Label(tail4)

	a.Popcntq(asm.DX, asm.DX)
	a.Addq(ax, asm.DX)

	a.Label(ret)

	a.Movq(retArg, ax)
	a.Ret()
}

func main() {
	if err := asm.Do("popcount_amd64.s", header, countBytesASM); err != nil {
		panic(err)
	}
}
