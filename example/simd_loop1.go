// +build !simd

//go:generate gensimd

package main

import "github.com/bjwbell/gensimd/simd"

func simd_loop1(v4 *[4]int) int {
	var tmp int
	var ret int
	var tmp2 simd.Int4
	tmp = 4
	x := v4[0]
	y := v4[1]
	v := simd.Int4(*v4)
	for i := 0; i < 10; i++ {
		tmp2 = simd.Int4Add(v, v)
		if i == 5 {
			break
		}
		tmp2 = simd.Int4(*v4)
		tmp = x*3 + 6*y
	}
	if tmp == 4 {
		ret = 1
	}
	if tmp2[0] == 0 {
		ret = 0
	}
	return ret
}

func simd_loop2(v4 []int) int {
	var tmp int
	var ret int
	var tmp2 simd.Int4
	tmp = 4
	x := v4[0]
	y := v4[1]
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		tmp = x*3 + 6*y
	}
	if tmp == 4 {
		ret = 1
	}
	if tmp2[0] == 0 {
		ret = 0
	}
	return ret
}
